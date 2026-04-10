package git

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func useNoSleep(t *testing.T) {
	t.Helper()
	original := artifactHubSleep
	artifactHubSleep = func(_ time.Duration) {}
	t.Cleanup(func() { artifactHubSleep = original })
}

func TestExtractRepoURLFromString_Variants(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "schemeless github URL",
			in:   "github.com/OrgOne/RepoOne",
			want: "https://github.com/OrgOne/RepoOne",
		},
		{
			name: "ssh github URL",
			in:   "git@github.com:OrgTwo/RepoTwo.git",
			want: "https://github.com/OrgTwo/RepoTwo",
		},
		{
			name: "embedded markdown link",
			in:   "Source: [repo](https://github.com/OrgThree/RepoThree/tree/main)",
			want: "https://github.com/OrgThree/RepoThree",
		},
		{
			name: "json escaped URL",
			in:   "https:\\/\\/github.com\\/OrgFour\\/RepoFour",
			want: "https://github.com/OrgFour/RepoFour",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := extractRepoURLFromString(tc.in)
			if got != tc.want {
				t.Fatalf("extractRepoURLFromString() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestExtractRepoURLFromRawArtifactPackage_FindsNestedSource(t *testing.T) {
	raw := []byte(`{
		"name": "example",
		"metadata": {
			"notes": "git@github.com:OrgNested/RepoNested.git"
		},
		"repository": {
			"url": "https://charts.example.com"
		}
	}`)

	got := extractRepoURLFromRawArtifactPackage(raw)
	want := "https://github.com/OrgNested/RepoNested"
	if got != want {
		t.Fatalf("extractRepoURLFromRawArtifactPackage() = %q, want %q", got, want)
	}
}

func TestSearchTopArtifactHubRepos_UsesFallbackAndDeduplicates(t *testing.T) {
	useNoSleep(t)
	original := artifactHubFetchAPI
	defer func() { artifactHubFetchAPI = original }()

	responses := map[int]string{
		0: `{
			"packages": [
				{
					"name": "one",
					"links": [{"name": "source", "url": "github.com/OrgOne/RepoOne"}],
					"repository": {"url": "https://charts.example.com"}
				},
				{
					"name": "two",
					"links": [{"name": "docs", "url": "https://example.com"}],
					"extra": "git@github.com:OrgTwo/RepoTwo.git"
				},
				{
					"name": "dup",
					"links": [{"name": "source", "url": "https://github.com/OrgOne/RepoOne"}]
				}
			]
		}`,
		60: `{
			"packages": [
				{
					"name": "three",
					"repository": {"url": "ssh://git@github.com/OrgThree/RepoThree.git"}
				},
				{
					"name": "four",
					"summary": "see https://github.com/OrgFour/RepoFour/releases"
				}
			]
		}`,
		120: `{"packages": []}`,
	}

	artifactHubFetchAPI = func(url string, _ ...string) ([]byte, error) {
		var offset int
		if _, err := fmt.Sscanf(url, "https://artifacthub.io/api/v1/packages/search?kind=0&sort=stars&limit=60&offset=%d", &offset); err != nil {
			return nil, fmt.Errorf("unexpected url: %s", url)
		}
		body, ok := responses[offset]
		if !ok {
			return nil, fmt.Errorf("no stubbed response for offset %d", offset)
		}
		return []byte(body), nil
	}

	repos, err := SearchTopArtifactHubRepos(4, 60)
	if err != nil {
		t.Fatalf("SearchTopArtifactHubRepos() returned error: %v", err)
	}

	if len(repos) != 4 {
		t.Fatalf("expected 4 repos, got %d (%v)", len(repos), repos)
	}

	joined := strings.Join(repos, ",")
	for _, want := range []string{
		"https://github.com/OrgOne/RepoOne",
		"https://github.com/OrgTwo/RepoTwo",
		"https://github.com/OrgThree/RepoThree",
		"https://github.com/OrgFour/RepoFour",
	} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing repo %q in %v", want, repos)
		}
	}
}

func TestSearchTopArtifactHubRepos_UsesPackageDetailsFallback(t *testing.T) {
	useNoSleep(t)
	original := artifactHubFetchAPI
	defer func() { artifactHubFetchAPI = original }()

	searchURL := "https://artifacthub.io/api/v1/packages/search?kind=0&sort=stars&limit=60&offset=0"
	detailURL := "https://artifacthub.io/api/v1/packages/helm/prometheus-community/kube-prometheus-stack"

	artifactHubFetchAPI = func(url string, _ ...string) ([]byte, error) {
		switch url {
		case searchURL:
			return []byte(`{
				"packages": [
					{
						"package_id": "p1",
						"name": "kube-prometheus-stack",
						"normalized_name": "kube-prometheus-stack",
						"repository": {"name": "prometheus-community", "url": "https://prometheus-community.github.io/helm-charts"}
					}
				]
			}`), nil
		case detailURL:
			return []byte(`{
				"name": "kube-prometheus-stack",
				"links": [
					{"name": "source", "url": "https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack"}
				]
			}`), nil
		default:
			return []byte(`{"packages": []}`), nil
		}
	}

	repos, err := SearchTopArtifactHubRepos(1, 60)
	if err != nil {
		t.Fatalf("SearchTopArtifactHubRepos() returned error: %v", err)
	}
	if len(repos) != 1 {
		t.Fatalf("expected 1 repo, got %d (%v)", len(repos), repos)
	}
	if repos[0] != "https://github.com/prometheus-community/helm-charts" {
		t.Fatalf("unexpected repo %q", repos[0])
	}
}

func TestSearchTopArtifactHubRepos_SkipsInvalidBatchAfterRetries(t *testing.T) {
	useNoSleep(t)
	original := artifactHubFetchAPI
	defer func() { artifactHubFetchAPI = original }()

	attemptsByOffset := map[int]int{}

	artifactHubFetchAPI = func(url string, _ ...string) ([]byte, error) {
		var offset int
		if _, err := fmt.Sscanf(url, "https://artifacthub.io/api/v1/packages/search?kind=0&sort=stars&limit=60&offset=%d", &offset); err == nil {
			attemptsByOffset[offset]++
			switch offset {
			case 0:
				// Always invalid JSON to trigger retries and batch skip.
				return []byte(""), nil
			case 60:
				return []byte(`{
					"packages": [
						{
							"name": "ok",
							"links": [{"name": "source", "url": "https://github.com/OrgOk/RepoOk"}]
						}
					]
				}`), nil
			default:
				return []byte(`{"packages": []}`), nil
			}
		}
		return nil, fmt.Errorf("unexpected url: %s", url)
	}

	repos, err := SearchTopArtifactHubRepos(1, 60)
	if err != nil {
		t.Fatalf("SearchTopArtifactHubRepos() returned error: %v", err)
	}
	if len(repos) != 1 {
		t.Fatalf("expected 1 repo, got %d (%v)", len(repos), repos)
	}
	if repos[0] != "https://github.com/OrgOk/RepoOk" {
		t.Fatalf("unexpected repo %q", repos[0])
	}
	if attemptsByOffset[0] != artifactHubBatchRetryAttempts {
		t.Fatalf("expected %d attempts for failing batch, got %d", artifactHubBatchRetryAttempts, attemptsByOffset[0])
	}
}

func TestSearchTopArtifactHubRepos_PartialUnder500DoesNotFail(t *testing.T) {
	useNoSleep(t)
	original := artifactHubFetchAPI
	defer func() { artifactHubFetchAPI = original }()

	artifactHubFetchAPI = func(url string, _ ...string) ([]byte, error) {
		var offset int
		if _, err := fmt.Sscanf(url, "https://artifacthub.io/api/v1/packages/search?kind=0&sort=stars&limit=60&offset=%d", &offset); err != nil {
			return nil, fmt.Errorf("unexpected url: %s", url)
		}
		switch offset {
		case 0:
			return []byte(`{
				"packages": [
					{"name":"a","links":[{"name":"source","url":"https://github.com/A/one"}]},
					{"name":"b","links":[{"name":"source","url":"https://github.com/B/two"}]},
					{"name":"c","links":[{"name":"source","url":"https://github.com/C/three"}]}
				]
			}`), nil
		default:
			return []byte(`{"packages": []}`), nil
		}
	}

	repos, err := SearchTopArtifactHubRepos(500, 60)
	if err != nil {
		t.Fatalf("SearchTopArtifactHubRepos() returned error: %v", err)
	}
	if len(repos) != 3 {
		t.Fatalf("expected 3 repos, got %d (%v)", len(repos), repos)
	}
}
