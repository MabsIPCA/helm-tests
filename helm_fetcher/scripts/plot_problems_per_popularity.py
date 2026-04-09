#!/usr/bin/env python3
"""Generate a 2D scatter plot (SVG) for problems per repository popularity (stars)."""

from __future__ import annotations

import argparse
import json
from pathlib import Path

from svg_scatter_plot import Point, render_scatter_svg


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Create problems-vs-popularity scatter plot as SVG."
    )
    parser.add_argument(
        "--catalog-in",
        default="catalog_by_project.json",
        help="Path to catalog_by_project.json",
    )
    parser.add_argument(
        "--search-in",
        default="github_search.json",
        help="Path to github_search.json",
    )
    parser.add_argument(
        "--output",
        default="plots/problems_per_popularity.svg",
        help="Output SVG path",
    )
    parser.add_argument(
        "--top",
        type=int,
        default=200,
        help="Use top N repositories by stars",
    )
    parser.add_argument(
        "--x-scale",
        choices=("log1p", "linear"),
        default="log1p",
        help="X-axis scale for stars: log1p spreads low-star repos, linear keeps raw spacing",
    )
    return parser.parse_args()


def _normalize_repo_key(repo_url: str) -> str:
    value = (repo_url or "").strip().rstrip("/")
    if value.startswith("https://github.com/"):
        value = value.removeprefix("https://github.com/")
    elif value.startswith("http://github.com/"):
        value = value.removeprefix("http://github.com/")
    return value.lower()


def _label_from_key(repo_key: str) -> str:
    return repo_key or "unknown"


def load_points(catalog_path: Path, search_path: Path, top: int) -> list[Point]:
    search_data = json.loads(search_path.read_text(encoding="utf-8"))
    repos = search_data.get("repos") or []

    stars_by_repo: dict[str, int] = {}
    for item in repos:
        key = _normalize_repo_key(str(item.get("repo_url") or ""))
        if not key:
            continue
        stars_by_repo[key] = int(item.get("stars") or 0)

    catalog_data = json.loads(catalog_path.read_text(encoding="utf-8"))
    points: list[Point] = []
    for entry in catalog_data:
        key = _normalize_repo_key(str(entry.get("repo_url") or ""))
        if not key or key not in stars_by_repo:
            continue

        failures = int(entry.get("total_failures") or 0)
        dep_failures = int(entry.get("total_dep_failures") or 0)
        problems = failures + dep_failures
        if problems <= 0:
            continue
        points.append((_label_from_key(key), stars_by_repo[key], problems))

    points.sort(key=lambda item: item[1], reverse=True)
    return points[: max(top, 1)]


def main() -> None:
    args = parse_args()
    points = load_points(Path(args.catalog_in), Path(args.search_in), int(args.top))

    render_scatter_svg(
        title=f"Problems per Popularity (non-zero only, Top {len(points)} by stars)",
        x_label="GitHub stars (log10(stars+1))" if args.x_scale == "log1p" else "GitHub stars",
        y_label="Problems (total_failures + total_dep_failures)",
        points=points,
        output_path=Path(args.output),
        point_color="#7c3aed",
        x_scale=args.x_scale,
    )

    print(f"Wrote plot: {args.output}")


if __name__ == "__main__":
    main()



