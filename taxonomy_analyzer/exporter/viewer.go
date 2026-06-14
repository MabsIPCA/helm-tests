package exporter

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/MabsIPCA/helm-tests/taxonomy_analyzer/model"
)

// viewerSubKind is a slim per-subkind summary row for the viewer sidebar.
type viewerSubKind struct {
	Key        string `json:"key"`
	Count      int    `json:"count"`
	Resolved   int    `json:"resolved"`
	Unresolved int    `json:"unresolved"`
	// Final is how many unresolved fixes ended on this taxonomy as the final
	// blocker (the destination counter, distinct from Count which is start).
	Final int `json:"final"`
	// NonFixable marks a subkind the fixer cannot resolve by value injection.
	NonFixable bool `json:"non_fixable,omitempty"`
}

// viewerPayload is the JSON blob embedded into viewer.html.
type viewerPayload struct {
	GeneratedAt   string                  `json:"generated_at"`
	SourceCatalog string                  `json:"source_catalog"`
	FixedCatalog  string                  `json:"fixed_catalog,omitempty"`
	Totals        model.ReportTotals      `json:"totals"`
	SubKinds      []viewerSubKind         `json:"sub_kinds"`
	Transitions   []model.TransitionStat  `json:"transitions,omitempty"`
	Occurrences   []model.ErrorOccurrence `json:"occurrences"`
}

// writeViewerHTML emits a self-contained, single-file HTML viewer for the
// classified failures. It embeds the occurrence data (including each fix
// chain) so the file can be opened directly in a browser with no server.
func writeViewerHTML(path string, report model.TaxonomyReport) error {
	payload := viewerPayload{
		GeneratedAt:   report.GeneratedAt.Format("2006-01-02 15:04:05 MST"),
		SourceCatalog: report.SourceCatalog,
		FixedCatalog:  report.FixedCatalog,
		Totals:        report.Totals,
		Transitions:   report.Transitions,
		Occurrences:   report.AllClassified,
	}

	for key, b := range report.BySubKind {
		row := viewerSubKind{Key: key, Count: b.Count, Final: b.FinalCount, NonFixable: b.NonFixable}
		if b.FixOutcome != nil {
			row.Resolved = b.FixOutcome.Resolved
			row.Unresolved = b.FixOutcome.Unresolved
		}
		payload.SubKinds = append(payload.SubKinds, row)
	}
	sortViewerSubKinds(payload.SubKinds)

	// MarshalIndent escapes <, > and & to < etc by default, so the blob
	// is safe to embed inside a <script> tag (no </script> break-out).
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal viewer payload: %w", err)
	}

	html := strings.Replace(viewerTemplate, "/*__DATA__*/", string(data), 1)
	if err := os.WriteFile(path, []byte(html), 0o644); err != nil {
		return fmt.Errorf("write viewer html: %w", err)
	}
	return nil
}

func sortViewerSubKinds(rows []viewerSubKind) {
	// Unresolved desc, then count desc, then key — surfaces the biggest
	// still-to-fix buckets first.
	for i := 1; i < len(rows); i++ {
		for j := i; j > 0 && lessViewerSubKind(rows[j], rows[j-1]); j-- {
			rows[j], rows[j-1] = rows[j-1], rows[j]
		}
	}
}

func lessViewerSubKind(a, b viewerSubKind) bool {
	if a.Unresolved != b.Unresolved {
		return a.Unresolved > b.Unresolved
	}
	if a.Count != b.Count {
		return a.Count > b.Count
	}
	return a.Key < b.Key
}

const viewerTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>Helm Failure Viewer</title>
<style>
  :root {
    --bg: #0f1115; --panel: #171a21; --panel2: #1e222b; --border: #2a2f3a;
    --text: #e6e9ef; --muted: #9aa3b2; --accent: #4f9cf9; --green: #3fb950;
    --red: #f85149; --amber: #d29922; --nfx: #a371f7; --mono: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  }
  * { box-sizing: border-box; }
  body { margin: 0; background: var(--bg); color: var(--text);
    font: 14px/1.5 system-ui, -apple-system, Segoe UI, Roboto, sans-serif; }
  header { padding: 14px 20px; border-bottom: 1px solid var(--border);
    background: var(--panel); position: sticky; top: 0; z-index: 5; }
  header h1 { margin: 0 0 4px; font-size: 16px; }
  header .meta { color: var(--muted); font-size: 12px; }
  .totals { display: flex; gap: 18px; margin-top: 10px; flex-wrap: wrap; }
  .totals .stat { font-size: 12px; color: var(--muted); }
  .totals .stat b { display: block; font-size: 18px; color: var(--text); }
  .totals .stat.red b { color: var(--red); }
  .totals .stat.green b { color: var(--green); }
  .totals .stat.muted b { color: var(--muted); }
  .totals .stat.nfx b { color: var(--nfx); }
  .layout { display: flex; align-items: flex-start; }
  aside { width: 484px; flex: none; border-right: 1px solid var(--border);
    height: calc(100vh - 92px); overflow-y: auto; position: sticky; top: 92px;
    background: var(--panel); }
  aside h2 { font-size: 11px; text-transform: uppercase; letter-spacing: .05em;
    color: var(--muted); padding: 12px 16px 6px; margin: 0; }
  .skrow { display: flex; align-items: center; gap: 8px; padding: 7px 16px;
    cursor: pointer; border-left: 3px solid transparent; font-size: 13px; }
  .skrow:hover { background: var(--panel2); }
  .skrow.active { background: var(--panel2); border-left-color: var(--accent); }
  .skrow .name { flex: 1; min-width: 0; font-family: var(--mono); font-size: 12px;
    white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }
  .skrow .badge { flex: none; font-size: 11px; color: var(--muted); width: 34px;
    text-align: right; font-family: var(--mono); }
  .skrow .badge.red { color: var(--red); }
  .skrow .badge.fin { color: var(--amber); }
  .skrow .badge.nfx { color: var(--nfx); }
  .skrow .badge.zero { color: var(--border); }
  aside h2 { display: flex; align-items: baseline; }
  aside h2 .leg { margin-left: auto; display: flex; gap: 8px; text-transform: none;
    letter-spacing: 0; font-family: var(--mono); font-size: 10px; }
  aside h2 .leg span { width: 34px; text-align: right; }
  aside h2 .leg .u { color: var(--red); }
  aside h2 .leg .f { color: var(--amber); }
  aside h2 .leg .x { color: var(--nfx); }
  /* highlight the active filter column: start dims the ->fin col, final dims tot/unfix */
  aside:not(.dim-final) .skrow .badge.fin,
  aside:not(.dim-final) .leg .f { opacity: .45; }
  aside.dim-final .skrow .badge:not(.fin):not(.nfx),
  aside.dim-final .leg span:not(.f):not(.x) { opacity: .45; }
  aside.dim-final h2 .leg .f { text-decoration: underline; }
  main { flex: 1; padding: 16px 20px; min-width: 0; }
  .controls { display: flex; gap: 10px; align-items: center; margin-bottom: 14px;
    flex-wrap: wrap; }
  .controls input[type=text] { flex: 1; min-width: 220px; padding: 8px 10px;
    background: var(--panel2); border: 1px solid var(--border); color: var(--text);
    border-radius: 6px; font-size: 13px; }
  .controls select { padding: 8px 10px; background: var(--panel2);
    border: 1px solid var(--border); color: var(--text); border-radius: 6px; }
  .controls label { color: var(--muted); font-size: 12px; }
  .count { color: var(--muted); font-size: 12px; margin-left: auto; }
  .project { margin-bottom: 18px; }
  .project > .phead { display: flex; gap: 10px; align-items: baseline; cursor: pointer;
    padding: 9px 12px; background: var(--panel2); border: 1px solid var(--border);
    border-radius: 8px; position: sticky; top: 92px; z-index: 2; }
  .project > .phead:hover { border-color: var(--accent); }
  .project > .phead .tw { color: var(--muted); width: 12px; flex: none; }
  .project > .phead .pname { font-weight: 600; font-size: 15px; }
  .project > .phead .purl { color: var(--muted); font-size: 12px; font-family: var(--mono);
    word-break: break-all; }
  .project > .phead .pcount { margin-left: auto; font-size: 12px; color: var(--muted);
    flex: none; }
  .project > .phead .pcount b.red { color: var(--red); }
  .project > .pbody { padding: 10px 0 0 14px; display: none; }
  .project.open > .pbody { display: block; }
  .project.open > .phead .tw::before { content: "\25be"; }
  .project:not(.open) > .phead .tw::before { content: "\25b8"; }
  .card { background: var(--panel); border: 1px solid var(--border);
    border-radius: 8px; margin-bottom: 12px; overflow: hidden; }
  .card .head { padding: 11px 14px; cursor: pointer; display: flex; gap: 10px;
    align-items: baseline; }
  .card .head:hover { background: var(--panel2); }
  .card .repo { font-weight: 600; }
  .card .chart { color: var(--muted); font-family: var(--mono); font-size: 12px;
    word-break: break-all; }
  .pill { font-size: 11px; padding: 2px 8px; border-radius: 999px; flex: none;
    font-family: var(--mono); }
  .pill.sk { background: #1f2937; color: var(--accent); }
  .pill.sk.shifted { background: rgba(210,153,34,.15); color: var(--amber); }
  .transitions { margin-bottom: 16px; }
  .transitions table { border-collapse: collapse; font-size: 12px; }
  .transitions th, .transitions td { text-align: left; padding: 5px 10px;
    border: 1px solid var(--border); }
  .transitions th { background: var(--panel2); color: var(--muted); font-weight: 600; }
  .transitions td.mono { font-family: var(--mono); }
  .transitions td.n { text-align: right; color: var(--text); }
  .transitions .same { color: var(--muted); }
  .transitions .shift { color: var(--amber); }
  .transitions summary { cursor: pointer; color: var(--muted); font-size: 12px;
    padding: 6px 0; user-select: none; }
  .pill.resolved { background: rgba(63,185,80,.15); color: var(--green); }
  .pill.unresolved { background: rgba(248,81,73,.15); color: var(--red); }
  .pill.noattempt { background: rgba(210,153,34,.15); color: var(--amber); }
  .card .body { padding: 0 14px 14px; display: none; }
  .card.open .body { display: block; }
  .field { margin-top: 12px; }
  .field .lbl { font-size: 11px; text-transform: uppercase; letter-spacing: .05em;
    color: var(--muted); margin-bottom: 4px; }
  pre { margin: 0; background: var(--bg); border: 1px solid var(--border);
    border-radius: 6px; padding: 10px 12px; overflow-x: auto; font-family: var(--mono);
    font-size: 12px; white-space: pre-wrap; word-break: break-word; }
  .chain { border-collapse: collapse; width: 100%; font-size: 12px; }
  .chain th, .chain td { text-align: left; padding: 6px 8px; border: 1px solid var(--border);
    vertical-align: top; }
  .chain th { background: var(--panel2); color: var(--muted); font-weight: 600;
    white-space: nowrap; }
  .chain td.mono { font-family: var(--mono); }
  .chain td .err { color: var(--muted); white-space: pre-wrap; word-break: break-word;
    max-width: 520px; }
  .stop { margin-top: 6px; font-size: 12px; }
  .stop b { font-family: var(--mono); color: var(--amber); }
  pre.blocker { border-color: var(--red); border-left-width: 3px;
    background: rgba(248,81,73,.06); }
  .ovr { display: flex; flex-direction: column; gap: 4px; }
  .ovr .kv { font-family: var(--mono); font-size: 12px; background: var(--bg);
    border: 1px solid var(--border); border-radius: 6px; padding: 5px 9px; }
  .ovr .kv .k { color: var(--accent); }
  .ovr .kv .v { color: var(--green); }
  .ovr .kv .v.empty { color: var(--muted); font-style: italic; }
  .empty { color: var(--muted); padding: 40px; text-align: center; }
  code { font-family: var(--mono); }
  .dimtoggle { display: inline-flex; align-items: center; gap: 6px; }
  .seg { display: inline-flex; border: 1px solid var(--border); border-radius: 6px;
    overflow: hidden; }
  .seg button { background: var(--panel2); color: var(--muted); border: none;
    padding: 7px 11px; font: inherit; font-size: 12px; cursor: pointer; }
  .seg button + button { border-left: 1px solid var(--border); }
  .seg button.on { background: var(--accent); color: #fff; }
  .seg button[data-dim="final"].on { background: var(--amber); color: #1a1206; }
</style>
</head>
<body>
<header>
  <h1>Helm Failure Viewer</h1>
  <div class="meta" id="meta"></div>
  <div class="totals" id="totals"></div>
</header>
<div class="layout">
  <aside id="sidebar"></aside>
  <main>
    <div class="controls">
      <input type="text" id="search" placeholder="Search repo, chart, command or error message...">
      <label>Fix status
        <select id="status">
          <option value="unresolved">Still to fix</option>
          <option value="all">All</option>
          <option value="resolved">Resolved</option>
          <option value="noattempt">No fix attempt</option>
        </select>
      </label>
      <label class="dimtoggle" title="Filter the selected subkind by the original error, or by the final blocker the fixer ended on">
        Match on
        <span class="seg" id="dim">
          <button data-dim="start" class="on">original</button><button data-dim="final">final&nbsp;error</button>
        </span>
      </label>
      <span class="count" id="count"></span>
    </div>
    <div id="transitions"></div>
    <div id="list"></div>
  </main>
</div>
<script id="data" type="application/json">/*__DATA__*/</script>
<script>
  const DATA = JSON.parse(document.getElementById('data').textContent);
  const state = { subkind: null, status: 'unresolved', q: '', dim: 'start' };

  function statusOf(o) {
    if (!o.fixed) return 'noattempt';
    return o.fixed.resolved ? 'resolved' : 'unresolved';
  }

  function esc(s) {
    return (s || '').replace(/[&<>]/g, c => ({'&':'&amp;','<':'&lt;','>':'&gt;'}[c]));
  }

  function renderHeader() {
    document.getElementById('meta').innerHTML =
      'Generated <code>' + esc(DATA.generated_at) + '</code> &middot; source <code>' +
      esc(DATA.source_catalog) + '</code>';
    const t = DATA.totals;
    const rows = [
      ['Template failures', t.template_failures || 0, ''],
      ['Fix attempts', t.fix_attempted || 0, ''],
      ['Resolved', t.fix_resolved || 0, 'green'],
      ['Still to fix', t.fix_unresolved || 0, 'red'],
    ];
    if (t.non_fixable_errors) {
      rows.push(['Non-fixable', t.non_fixable_errors, 'nfx']);
    }
    if (t.duplicates_collapsed) {
      rows.push(['Dupes collapsed', t.duplicates_collapsed, 'muted']);
    }
    document.getElementById('totals').innerHTML = rows
      .map(([l, v, c]) => '<div class="stat ' + c + '"><b>' + v + '</b>' + l + '</div>').join('');
  }

  function renderSidebar() {
    const total = DATA.occurrences.length;
    const finalTotal = DATA.sub_kinds.reduce((s, sk) => s + (sk.final || 0), 0);
    let html = '<h2>Subkinds <span class="leg">' +
      '<span title="occurrences as start error">tot</span>' +
      '<span class="u" title="still to fix">unfix</span>' +
      '<span class="f" title="still failing here at the end (cumulative: unresolved + not-attempted)">&rarr;fin</span>' +
      '<span class="x" title="non-fixable by design (chart-author validation / structural blockers)">nfx</span></span></h2>';
    html += skRow({ key: '(all)', count: total, unresolved: DATA.totals.fix_unresolved || 0, final: finalTotal, nfx: DATA.totals.non_fixable_errors || 0 }, state.subkind === null);
    for (const sk of DATA.sub_kinds) {
      // nfx column: a non-fixable subkind contributes its whole count, else 0.
      html += skRow(Object.assign({}, sk, { nfx: sk.non_fixable ? sk.count : 0 }), state.subkind === sk.key);
    }
    const aside = document.getElementById('sidebar');
    // dim-final emphasizes the column the selected subkind currently filters on.
    aside.classList.toggle('dim-final', state.dim === 'final');
    aside.innerHTML = html;
    aside.querySelectorAll('.skrow').forEach(el => {
      el.onclick = () => { state.subkind = el.dataset.key === '(all)' ? null : el.dataset.key; render(); };
    });
  }

  function skRow(sk, active) {
    // Three separate columns: total occurrences as the START error, how many of
    // those are still unresolved, and how many fixes end HERE as the final blocker.
    const tot = '<span class="badge" title="occurrences as start error">' + sk.count + '</span>';
    const unr = '<span class="badge ' + (sk.unresolved ? 'red' : 'zero') +
      '" title="still to fix">' + (sk.unresolved || 0) + '</span>';
    const fin = '<span class="badge fin' + (sk.final ? '' : ' zero') +
      '" title="still failing here at the end (cumulative: unresolved + not-attempted)">&rarr;' + (sk.final || 0) + '</span>';
    const nfx = '<span class="badge nfx' + (sk.nfx ? '' : ' zero') +
      '" title="non-fixable by design (chart-author validation / structural blockers)">' + (sk.nfx || 0) + '</span>';
    return '<div class="skrow ' + (active ? 'active' : '') + '" data-key="' + esc(sk.key) + '">' +
      '<span class="name">' + esc(sk.key) + '</span>' + tot + unr + fin + nfx + '</div>';
  }

  function matches(o) {
    if (state.subkind) {
      // 'start' matches the original error's taxonomy; 'final' matches the
      // blocker the fixer ended on (only present on unresolved fixes).
      const key = state.dim === 'final'
        ? (o.final_error_kind ? o.final_error_kind + '.' + o.final_error_sub_kind : '')
        : o.error_kind + '.' + o.error_sub_kind;
      if (key !== state.subkind) return false;
    }
    if (state.status !== 'all' && statusOf(o) !== state.status) return false;
    if (state.q) {
      const hay = (o.repo_name + ' ' + o.chart_path + ' ' + (o.helm_command || '') + ' ' +
        o.error_message).toLowerCase();
      if (!hay.includes(state.q)) return false;
    }
    return true;
  }

  function appliedOverrides(fixed) {
    // The fixer applies --set value_path=value_injected cumulatively across
    // iterations; collapse the chain to the final set of overrides actually
    // provided (last value wins per path), preserving first-seen order.
    if (!fixed || !fixed.fix_chain) return [];
    const order = [], byPath = {};
    for (const a of fixed.fix_chain) {
      if (!a.value_path) continue;
      if (!(a.value_path in byPath)) order.push(a.value_path);
      byPath[a.value_path] = a.value_injected || '';
    }
    return order.map(p => ({ path: p, value: byPath[p] }));
  }

  function overridesBlock(fixed) {
    const ovr = appliedOverrides(fixed);
    if (!ovr.length) return '';
    const items = ovr.map(o =>
      '<div class="kv"><span class="k">' + esc(o.path) + '</span> = ' +
      (o.value === '' ? '<span class="v empty">(empty string)</span>'
                      : '<span class="v">' + esc(o.value) + '</span>') + '</div>'
    ).join('');
    return '<div class="field"><div class="lbl">Values provided to fix (--set overrides)</div>' +
      '<div class="ovr">' + items + '</div></div>';
  }

  function blockerBlock(o) {
    // The error the chart was STILL failing with when the fixer gave up. This
    // is the real blocker and often differs from the original error_message —
    // including landing in a different taxonomy, shown as the final-error pill.
    const fixed = o.fixed;
    if (!fixed || fixed.resolved || !fixed.final_error) return '';
    let tag = '';
    if (o.final_error_kind) {
      const fk = o.final_error_kind + '.' + o.final_error_sub_kind;
      const start = o.error_kind + '.' + o.error_sub_kind;
      const shifted = fk !== start;
      tag = ' <span class="pill sk' + (shifted ? ' shifted' : '') + '" title="' +
        (shifted ? 'final taxonomy differs from start ' + esc(start) : 'same taxonomy as start') +
        '">' + esc(fk) + (shifted ? ' ≠ start' : '') + '</span>';
    }
    return '<div class="field"><div class="lbl">Blocking error when fixer gave up' + tag + '</div>' +
      '<pre class="blocker">' + esc(fixed.final_error) + '</pre></div>';
  }

  function chainTable(o) {
    const fixed = o.fixed;
    const blocker = blockerBlock(o);
    if (!fixed || !fixed.fix_chain || !fixed.fix_chain.length) {
      // No injections were applied, but a blocker may still exist (immediately
      // unfixable). The stop reason is still worth surfacing.
      if (!blocker && !(fixed && fixed.stop_reason)) return '';
      return blocker +
        (fixed && fixed.stop_reason ? '<div class="stop">Stopped: <b>' + esc(fixed.stop_reason) + '</b></div>' : '');
    }
    let rows = fixed.fix_chain.map(a =>
      '<tr><td>' + a.attempt + '</td>' +
      '<td class="mono">' + esc(a.kind || '-') + '</td>' +
      '<td class="mono">' + esc(a.value_path || '-') + '</td>' +
      '<td class="mono">' + esc(a.value_injected || '-') + '</td>' +
      '<td><div class="err">' + esc(a.error_seen || '') + '</div></td></tr>'
    ).join('');
    return '<div class="field"><div class="lbl">Fix chain (' + fixed.fix_chain.length + ' attempts)</div>' +
      '<table class="chain"><tr><th>#</th><th>kind</th><th>value_path</th><th>injected</th><th>error seen</th></tr>' +
      rows + '</table>' +
      (fixed.stop_reason ? '<div class="stop">Stopped: <b>' + esc(fixed.stop_reason) + '</b></div>' : '') +
      '</div>' + blocker;
  }

  function card(o, idx) {
    const st = statusOf(o);
    const pill = { resolved: 'resolved', unresolved: 'unresolved', noattempt: 'noattempt' }[st];
    const pillLabel = { resolved: 'resolved', unresolved: 'still to fix', noattempt: 'no attempt' }[st];
    const sk = o.error_kind + '.' + o.error_sub_kind;
    let body = '';
    if (o.helm_command) body += '<div class="field"><div class="lbl">Command</div><pre>' + esc(o.helm_command) + '</pre></div>';
    if (o.values_files && o.values_files.length) body += '<div class="field"><div class="lbl">Values files</div><pre>' + esc(o.values_files.join('\n')) + '</pre></div>';
    body += '<div class="field"><div class="lbl">Error message</div><pre>' + esc(o.error_message) + '</pre></div>';
    body += overridesBlock(o.fixed);
    body += chainTable(o);
    return '<div class="card" data-idx="' + idx + '">' +
      '<div class="head">' +
        '<span class="pill sk">' + esc(sk) + '</span>' +
        '<span class="pill ' + pill + '">' + pillLabel + '</span>' +
        '<span class="chart">' + esc(o.chart_path) + '</span>' +
      '</div>' +
      '<div class="body">' + body + '</div>' +
    '</div>';
  }

  function groupByProject(occs) {
    const map = new Map();
    for (const o of occs) {
      let g = map.get(o.repo_name);
      if (!g) { g = { name: o.repo_name, url: o.repo_url, items: [], unresolved: 0 }; map.set(o.repo_name, g); }
      g.items.push(o);
      if (statusOf(o) === 'unresolved') g.unresolved++;
    }
    const groups = Array.from(map.values());
    // Most still-to-fix first, then most failures, then name.
    groups.sort((a, b) => b.unresolved - a.unresolved || b.items.length - a.items.length || a.name.localeCompare(b.name));
    return groups;
  }

  function renderTransitions() {
    const el = document.getElementById('transitions');
    const tr = DATA.transitions || [];
    if (!tr.length) { el.innerHTML = ''; return; }
    const rows = tr.map(t => {
      const shift = t.from !== t.to;
      return '<tr><td class="mono">' + esc(t.from) + '</td>' +
        '<td class="mono ' + (shift ? 'shift' : 'same') + '">' +
          (shift ? '&rarr; ' : '= ') + esc(t.to) + '</td>' +
        '<td class="n">' + t.count + '</td></tr>';
    }).join('');
    const shifts = tr.filter(t => t.from !== t.to).reduce((s, t) => s + t.count, 0);
    el.innerHTML = '<details class="transitions"><summary>Start &rarr; final taxonomy shifts ' +
      '(' + shifts + ' of ' + tr.reduce((s, t) => s + t.count, 0) +
      ' unresolved fixes changed taxonomy after injection)</summary>' +
      '<table><tr><th>start</th><th>final blocker</th><th>count</th></tr>' + rows + '</table></details>';
  }

  function render() {
    renderSidebar();
    renderTransitions();
    const filtered = DATA.occurrences.filter(matches);
    document.getElementById('count').textContent = filtered.length + ' failures in ' +
      new Set(filtered.map(o => o.repo_name)).size + ' projects';
    const list = document.getElementById('list');
    if (!filtered.length) { list.innerHTML = '<div class="empty">No matching failures.</div>'; return; }

    const groups = groupByProject(filtered);
    const cap = 600;
    let shown = 0, truncated = false, html = '';
    for (const g of groups) {
      if (shown >= cap) { truncated = true; break; }
      const cards = g.items.slice(0, cap - shown).map((o, i) => card(o, i)).join('');
      shown += Math.min(g.items.length, cap - shown);
      const purl = g.url ? '<span class="purl">' + esc(g.url) + '</span>' : '';
      const ucount = g.unresolved ? '<b class="red">' + g.unresolved + '</b> to fix / ' : '';
      html += '<div class="project open">' +
        '<div class="phead"><span class="tw"></span>' +
          '<span class="pname">' + esc(g.name) + '</span>' + purl +
          '<span class="pcount">' + ucount + g.items.length + ' failures</span>' +
        '</div>' +
        '<div class="pbody">' + cards + '</div>' +
      '</div>';
    }
    if (truncated) html += '<div class="empty">Showing first ' + cap + ' failures. Narrow with search or a subkind.</div>';
    list.innerHTML = html;

    list.querySelectorAll('.project > .phead').forEach(h => {
      h.onclick = () => h.parentElement.classList.toggle('open');
    });
    list.querySelectorAll('.card .head').forEach(h => {
      h.onclick = e => { e.stopPropagation(); h.parentElement.classList.toggle('open'); };
    });
  }

  document.getElementById('search').addEventListener('input', e => { state.q = e.target.value.toLowerCase().trim(); render(); });
  document.getElementById('status').addEventListener('change', e => { state.status = e.target.value; render(); });
  document.getElementById('dim').addEventListener('click', e => {
    const btn = e.target.closest('button[data-dim]');
    if (!btn || btn.dataset.dim === state.dim) return;
    state.dim = btn.dataset.dim;
    document.querySelectorAll('#dim button').forEach(b => b.classList.toggle('on', b.dataset.dim === state.dim));
    // Final error is cumulative over every still-broken run (unresolved AND
    // not-attempted). The default 'unresolved' status would hide not-attempted
    // runs, so widen to 'all' when entering final mode; restore on the way back.
    if (state.dim === 'final' && state.status === 'unresolved') {
      state.status = 'all';
      document.getElementById('status').value = 'all';
    } else if (state.dim === 'start' && state.status === 'all') {
      state.status = 'unresolved';
      document.getElementById('status').value = 'unresolved';
    }
    render();
  });

  renderHeader();
  render();
</script>
</body>
</html>`
