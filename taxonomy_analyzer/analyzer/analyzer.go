package analyzer

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/classifier"
	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/model"
)

const maxExamplesPerBucket = 5

// Analyzer keeps running aggregation state while catalog entries are streamed.
type Analyzer struct {
	report      model.TaxonomyReport
	fixedIndex  map[string]*model.FixedResult
	transitions map[string]int // "start.sub -> final.sub" -> count
	// seen dedupes occurrences by (repo, error message). The fetcher renders up
	// to 100 value-file combinations per chart, so one broken chart yields many
	// near-identical failures (e.g. translator-devops' control-character value
	// files inflate malformed_yaml). Collapsing identical errors within a project
	// keeps the taxonomy summary one-row-per-distinct-error.
	seen map[string]struct{}
}

// New creates a new analyzer report collector.
// fixedCatalog and fixedIndex may be empty/nil when --fixed is not provided.
func New(sourceCatalog, fixedCatalog string, fixedIndex map[string]*model.FixedResult) *Analyzer {
	r := model.TaxonomyReport{
		GeneratedAt:   time.Now().UTC(),
		SourceCatalog: sourceCatalog,
		ByKind:        map[string]model.TaxonomyBucket{},
		BySubKind:     map[string]model.TaxonomyBucket{},
		ByRepo:        map[string]map[string]int{},
	}
	if fixedCatalog != "" {
		r.FixedCatalog = fixedCatalog
	}
	return &Analyzer{report: r, fixedIndex: fixedIndex, transitions: map[string]int{}, seen: map[string]struct{}{}}
}

// ConsumeRepo feeds one repo result into the taxonomy analyzer.
func (a *Analyzer) ConsumeRepo(repo model.RepoResult) {
	a.report.Totals.Repos++
	a.report.Totals.Runs += repo.TotalRuns
	a.report.Totals.TemplateFailures += repo.TotalFailures
	a.report.Totals.DependencyErrors += repo.TotalDepFailures

	for _, chart := range repo.Charts {
		if chart.DepBuildFailure && chart.DepBuildError != "" {
			a.consumeOccurrence(model.ErrorOccurrence{
				RepoURL:      repo.RepoURL,
				RepoName:     repo.RepoName,
				ChartPath:    chart.ChartPath,
				ErrorSource:  "dependency",
				ErrorMessage: chart.DepBuildError,
			})
		}

		for _, run := range chart.Runs {
			if run.Success || run.ErrorMessage == "" {
				continue
			}
			occ := model.ErrorOccurrence{
				RepoURL:      repo.RepoURL,
				RepoName:     repo.RepoName,
				ChartPath:    run.ChartPath,
				ValuesFiles:  run.ValuesFiles,
				HelmCommand:  run.HelmCommand,
				ErrorSource:  "template",
				ErrorMessage: run.ErrorMessage,
			}
			if a.fixedIndex != nil {
				occ.Fixed = a.fixedIndex[run.HelmCommand]
			}
			a.consumeOccurrence(occ)
		}
	}
}

func (a *Analyzer) consumeOccurrence(occ model.ErrorOccurrence) {
	// Collapse duplicate errors within the same project. The same error message
	// repeated across value-file combinations (or charts) of one repo counts once,
	// so combinatorial explosion does not skew the summary. Different projects with
	// an identical message still count separately (the repo is part of the key).
	dedupKey := occ.RepoName + "\x00" + occ.ErrorMessage
	if _, dup := a.seen[dedupKey]; dup {
		a.report.Totals.DuplicatesCollapsed++
		return
	}
	a.seen[dedupKey] = struct{}{}

	res := classifier.Classify(occ.ErrorMessage, occ.ErrorSource)
	occ.ErrorKind = res.Kind
	occ.ErrorSubKind = res.SubKind

	if !res.Classified {
		a.report.Totals.UnclassifiedError++
		a.report.Unclassified = append(a.report.Unclassified, occ)
	} else {
		a.report.Totals.ClassifiedErrors++
	}

	kindKey := occ.ErrorKind
	subKindKey := fmt.Sprintf("%s.%s", occ.ErrorKind, occ.ErrorSubKind)

	// Classify the FINAL error cumulatively: the error the chart is STILL failing
	// with at the end. This covers every run that is not resolved — unresolved
	// fixes (use the recorded blocker, or the original error if none was
	// captured) AND runs that were never attempted (the original error stands as
	// the final). Resolved runs have no final error. Done before the buckets are
	// bumped so examples carry the final-error fields too.
	resolved := occ.Fixed != nil && occ.Fixed.Resolved
	if !resolved {
		// The recorded FinalError is only a genuine, different blocker when the
		// fixer actually injected something (non-empty fix chain). With an empty
		// chain nothing changed, so the chart still has its ORIGINAL error — any
		// captured FinalError is a fixer-environment artifact (e.g. deps not built
		// producing a dependency error). Unfixable kinds we never attempt
		// (values_schema_validation, custom_validation, ...) thus keep
		// final == start, and not-attempted runs (Fixed == nil) likewise.
		injected := occ.Fixed != nil && len(occ.Fixed.FixChain) > 0
		finalText := occ.ErrorMessage
		if injected && occ.Fixed.FinalError != "" {
			finalText = occ.Fixed.FinalError
		}
		fres := classifier.Classify(finalText, occ.ErrorSource)
		occ.FinalErrorKind = fres.Kind
		occ.FinalErrorSubKind = fres.SubKind
		finalSubKey := fres.Kind + "." + fres.SubKind
		a.bumpFinalCount(a.report.ByKind, fres.Kind)
		a.bumpFinalCount(a.report.BySubKind, finalSubKey)
		// Transitions track the journey of an actual fix attempt — only record
		// when injection happened, so the panel stays about post-injection movement.
		if injected {
			a.transitions[subKindKey+" -> "+finalSubKey]++
		}
	}

	a.bumpBucket(a.report.ByKind, kindKey, occ)
	a.bumpBucket(a.report.BySubKind, subKindKey, occ)

	// Flag non-fixable subkinds (chart-author validation / structural blockers)
	// so the viewer can separate them from the fixable surface. Only the subkind
	// bucket is flagged — a whole kind ("template") mixes fixable and not.
	if classifier.IsNonFixable(occ.ErrorKind, occ.ErrorSubKind) {
		a.report.Totals.NonFixableErrors++
		b := a.report.BySubKind[subKindKey]
		b.NonFixable = true
		a.report.BySubKind[subKindKey] = b
	}

	if occ.Fixed != nil {
		a.report.Totals.FixAttempted++
		if occ.Fixed.Resolved {
			a.report.Totals.FixResolved++
		} else {
			a.report.Totals.FixUnresolved++
		}
		a.bumpFixOutcome(a.report.ByKind, kindKey, occ.Fixed)
		a.bumpFixOutcome(a.report.BySubKind, subKindKey, occ.Fixed)
	}

	if _, ok := a.report.ByRepo[occ.RepoName]; !ok {
		a.report.ByRepo[occ.RepoName] = map[string]int{}
	}
	a.report.ByRepo[occ.RepoName][subKindKey]++
	a.report.AllClassified = append(a.report.AllClassified, occ)
}

// bumpFinalCount increments the destination counter for a taxonomy bucket. The
// bucket may not yet exist (a taxonomy that only ever appears as a final blocker,
// never as a start) — that is the point, and it surfaces with Count 0.
func (a *Analyzer) bumpFinalCount(buckets map[string]model.TaxonomyBucket, key string) {
	bucket := buckets[key]
	bucket.FinalCount++
	buckets[key] = bucket
}

func (a *Analyzer) bumpBucket(buckets map[string]model.TaxonomyBucket, key string, occ model.ErrorOccurrence) {
	bucket := buckets[key]
	bucket.Count++
	if len(bucket.Examples) < maxExamplesPerBucket {
		bucket.Examples = append(bucket.Examples, occ)
	}
	buckets[key] = bucket
}

func (a *Analyzer) bumpFixOutcome(buckets map[string]model.TaxonomyBucket, key string, fixed *model.FixedResult) {
	bucket := buckets[key]
	fo := bucket.FixOutcome
	if fo == nil {
		fo = &model.FixOutcome{}
	}
	fo.Attempted++
	if fixed.Resolved {
		fo.Resolved++
	} else {
		fo.Unresolved++
		if fixed.StopReason != "" {
			if fo.ByStopReason == nil {
				fo.ByStopReason = map[string]int{}
			}
			fo.ByStopReason[fixed.StopReason]++
		}
	}
	bucket.FixOutcome = fo
	buckets[key] = bucket
}

// Report returns a normalized copy of the report for exporting.
func (a *Analyzer) Report() model.TaxonomyReport {
	report := a.report
	report.ByKind = sortBucketsByCount(report.ByKind)
	report.BySubKind = sortBucketsByCount(report.BySubKind)
	report.ByRepo = sortRepoKinds(report.ByRepo)
	report.Transitions = sortTransitions(a.transitions)
	return report
}

// sortTransitions flattens the start->final transition counter into a slice
// sorted by count desc (then key) for deterministic, readable output.
func sortTransitions(src map[string]int) []model.TransitionStat {
	if len(src) == 0 {
		return nil
	}
	out := make([]model.TransitionStat, 0, len(src))
	for k, c := range src {
		from, to, _ := strings.Cut(k, " -> ")
		out = append(out, model.TransitionStat{From: from, To: to, Count: c})
	}
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Count == out[j].Count {
			if out[i].From == out[j].From {
				return out[i].To < out[j].To
			}
			return out[i].From < out[j].From
		}
		return out[i].Count > out[j].Count
	})
	return out
}

// map order is not guaranteed in JSON output, so we recreate maps in sorted order
// to get deterministic marshal output for tests and diffs.
func sortBucketsByCount(src map[string]model.TaxonomyBucket) map[string]model.TaxonomyBucket {
	type pair struct {
		key   string
		count int
	}
	pairs := make([]pair, 0, len(src))
	for k, v := range src {
		pairs = append(pairs, pair{key: k, count: v.Count})
	}
	sort.SliceStable(pairs, func(i, j int) bool {
		if pairs[i].count == pairs[j].count {
			return pairs[i].key < pairs[j].key
		}
		return pairs[i].count > pairs[j].count
	})

	out := make(map[string]model.TaxonomyBucket, len(src))
	for _, p := range pairs {
		out[p.key] = src[p.key]
	}
	return out
}

func sortRepoKinds(src map[string]map[string]int) map[string]map[string]int {
	repoNames := make([]string, 0, len(src))
	for repo := range src {
		repoNames = append(repoNames, repo)
	}
	sort.Strings(repoNames)

	out := make(map[string]map[string]int, len(src))
	for _, repo := range repoNames {
		kinds := src[repo]
		keys := make([]string, 0, len(kinds))
		for k := range kinds {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		normalized := make(map[string]int, len(kinds))
		for _, k := range keys {
			normalized[k] = kinds[k]
		}
		out[repo] = normalized
	}
	return out
}
