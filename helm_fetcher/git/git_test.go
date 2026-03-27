package git

import "testing"

func TestRankAndTrimReposByStarsDesc(t *testing.T) {
	candidates := []scoredRepo{
		{url: "https://github.com/org/c", stars: 2},
		{url: "https://github.com/org/a", stars: 10},
		{url: "https://github.com/org/b", stars: 10},
	}

	ranked := rankAndTrimReposByStars(candidates, 2, OrderDesc)
	if len(ranked) != 2 {
		t.Fatalf("expected 2 repos, got %d", len(ranked))
	}
	if ranked[0].url != "https://github.com/org/a" || ranked[1].url != "https://github.com/org/b" {
		t.Fatalf("unexpected rank order: %#v", ranked)
	}
}

func TestRankAndTrimReposByStarsAsc(t *testing.T) {
	candidates := []scoredRepo{
		{url: "https://github.com/org/c", stars: 2},
		{url: "https://github.com/org/a", stars: 10},
		{url: "https://github.com/org/b", stars: 5},
	}

	ranked := rankAndTrimReposByStars(candidates, 2, OrderAsc)
	if len(ranked) != 2 {
		t.Fatalf("expected 2 repos, got %d", len(ranked))
	}
	if ranked[0].url != "https://github.com/org/c" || ranked[1].url != "https://github.com/org/b" {
		t.Fatalf("unexpected rank order: %#v", ranked)
	}
}
