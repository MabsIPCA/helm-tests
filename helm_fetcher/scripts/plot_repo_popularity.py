#!/usr/bin/env python3
"""Generate a 2D bar plot (SVG) with repository popularity (stars)."""

from __future__ import annotations

import argparse
import json
from pathlib import Path

from svg_bar_plot import render_horizontal_bar_svg


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Create repository popularity (stars) bar plot as SVG."
    )
    parser.add_argument(
        "--search-in",
        default="github_search.json",
        help="Path to github_search.json",
    )
    parser.add_argument(
        "--output",
        default="plots/repo_popularity_stars.svg",
        help="Output SVG path",
    )
    parser.add_argument(
        "--top",
        type=int,
        default=50,
        help="Number of repositories to include",
    )
    return parser.parse_args()


def _repo_label(repo_url: str) -> str:
    value = (repo_url or "").strip()
    if value.startswith("https://github.com/"):
        return value.removeprefix("https://github.com/")
    if value.startswith("http://github.com/"):
        return value.removeprefix("http://github.com/")
    return value or "unknown"


def load_repo_stars(path: Path) -> list[tuple[str, int]]:
    payload = json.loads(path.read_text(encoding="utf-8"))
    repos = payload.get("repos") or []

    rows: list[tuple[str, int]] = []
    for item in repos:
        label = _repo_label(str(item.get("repo_url") or ""))
        stars = int(item.get("stars") or 0)
        rows.append((label, stars))

    rows.sort(key=lambda item: item[1], reverse=True)
    return rows


def main() -> None:
    args = parse_args()
    input_path = Path(args.search_in)
    output_path = Path(args.output)

    rows = load_repo_stars(input_path)
    top = max(int(args.top), 1)
    selected = rows[:top]

    render_horizontal_bar_svg(
        title=f"Repository Popularity by Stars (Top {len(selected)})",
        x_label="GitHub stars",
        rows=selected,
        output_path=output_path,
        bar_color="#2563eb",
    )

    print(f"Wrote plot: {output_path}")


if __name__ == "__main__":
    main()

