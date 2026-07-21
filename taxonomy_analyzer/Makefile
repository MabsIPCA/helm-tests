run:
	go run .

build:
	go build ./...

test:
	go test ./...

# catalog_sources_v2.json = github + artifacthub input catalogs concatenated;
# catalog_fixed_cumulative_v2.json = their two fixed catalogs concatenated. Pairing
# the matched input+fixed keeps Fix attempts == Template failures (every failure
# correlates by helm_command). These *_v2 files are the CURRENT basis (2026-07-17)
# and reproduce out/cumulative_v2 — the numbers used in the thesis. The June
# catalog_sources_merged.json / catalog_fixed_cumulative.json are superseded.
analyze-cumulative:
	go run . -input ../helm_fetcher/catalog_sources_v2.json --fixed ../helm_fetcher/catalog_fixed_cumulative_v2.json -out out/cumulative_v2

# analyze-cumulative-june: the previous (superseded) basis, kept for provenance.
analyze-cumulative-june:
	go run . -input ../helm_fetcher/catalog_sources_merged.json --fixed ../helm_fetcher/catalog_fixed_cumulative.json -out out/cumulative_final

analyze-github:
	go run . -input ../helm_fetcher/results/github/catalog_by_project.json --fixed ../helm_fetcher/catalog_fixed_github_final.json -out out/github_final

analyze-artifacthub:
	go run . -input ../helm_fetcher/runs/20260529_135859_artifacthub/catalog_by_project.json --fixed ../helm_fetcher/catalog_fixed_ah_final.json -out out/artifacthub_final


