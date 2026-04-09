#!/usr/bin/env python3
"""Generate a 2D bar plot (SVG) with problems per project from catalog_by_project.json."""

from __future__ import annotations

import argparse
import json
from pathlib import Path

from svg_bar_plot import render_horizontal_bar_svg


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Create problems-per-project bar plot as SVG."
    )
    parser.add_argument(
        "--catalog-in",
        default="catalog_by_project.json",
        help="Path to catalog_by_project.json",
    )
    parser.add_argument(
        "--output",
        default="plots/problems_per_project.svg",
        help="Output SVG path",
    )
    parser.add_argument(
        "--top",
        type=int,
        default=40,
        help="Number of projects to include",
    )
    return parser.parse_args()


def _repo_label(entry: dict) -> str:
    repo_name = str(entry.get("repo_name") or "").strip()
    if repo_name:
        return repo_name

    repo_url = str(entry.get("repo_url") or "").strip()
    if repo_url.startswith("https://github.com/"):
        return repo_url.removeprefix("https://github.com/")
    if repo_url.startswith("http://github.com/"):
        return repo_url.removeprefix("http://github.com/")
    return repo_url or "unknown"


def load_project_problems(path: Path) -> list[tuple[str, int]]:
    data = json.loads(path.read_text(encoding="utf-8"))

    rows: list[tuple[str, int]] = []
    for entry in data:
        failures = int(entry.get("total_failures") or 0)
        dep_failures = int(entry.get("total_dep_failures") or 0)
        problems = failures + dep_failures
        rows.append((_repo_label(entry), problems))

    rows.sort(key=lambda item: item[1], reverse=True)
    return rows


def main() -> None:
    args = parse_args()
    input_path = Path(args.catalog_in)
    output_path = Path(args.output)

    rows = load_project_problems(input_path)
    top = max(int(args.top), 1)
    selected = rows[:top]

    render_horizontal_bar_svg(
        title=f"Helm Template Problems per Project (Top {len(selected)})",
        x_label="Problem count (total_failures + total_dep_failures)",
        rows=selected,
        output_path=output_path,
        bar_color="#dc2626",
    )

    print(f"Wrote plot: {output_path}")


if __name__ == "__main__":
    main()


