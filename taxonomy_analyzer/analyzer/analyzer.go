package analyzer

import (
	"fmt"
	"sort"
	"time"

	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/classifier"
	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/model"
)

const maxExamplesPerBucket = 5

// Analyzer keeps running aggregation state while catalog entries are streamed.
type Analyzer struct {
	report model.TaxonomyReport
}

// New creates a new analyzer report collector.
func New(sourceCatalog string) *Analyzer {
	return &Analyzer{
		report: model.TaxonomyReport{
			GeneratedAt:   time.Now().UTC(),
			SourceCatalog: sourceCatalog,
			ByKind:        map[string]model.TaxonomyBucket{},
			BySubKind:     map[string]model.TaxonomyBucket{},
			ByRepo:        map[string]map[string]int{},
		},
	}
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
			a.consumeOccurrence(model.ErrorOccurrence{
				RepoURL:      repo.RepoURL,
				RepoName:     repo.RepoName,
				ChartPath:    run.ChartPath,
				ValuesFiles:  run.ValuesFiles,
				HelmCommand:  run.HelmCommand,
				ErrorSource:  "template",
				ErrorMessage: run.ErrorMessage,
			})
		}
	}
}

func (a *Analyzer) consumeOccurrence(occ model.ErrorOccurrence) {
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
	a.bumpBucket(a.report.ByKind, kindKey, occ)
	a.bumpBucket(a.report.BySubKind, subKindKey, occ)

	if _, ok := a.report.ByRepo[occ.RepoName]; !ok {
		a.report.ByRepo[occ.RepoName] = map[string]int{}
	}
	a.report.ByRepo[occ.RepoName][subKindKey]++
	a.report.AllClassified = append(a.report.AllClassified, occ)
}

func (a *Analyzer) bumpBucket(buckets map[string]model.TaxonomyBucket, key string, occ model.ErrorOccurrence) {
	bucket := buckets[key]
	bucket.Count++
	if len(bucket.Examples) < maxExamplesPerBucket {
		bucket.Examples = append(bucket.Examples, occ)
	}
	buckets[key] = bucket
}

// Report returns a normalized copy of the report for exporting.
func (a *Analyzer) Report() model.TaxonomyReport {
	report := a.report
	report.ByKind = sortBucketsByCount(report.ByKind)
	report.BySubKind = sortBucketsByCount(report.BySubKind)
	report.ByRepo = sortRepoKinds(report.ByRepo)
	return report
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
