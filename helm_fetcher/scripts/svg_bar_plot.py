"""Small helper to render horizontal bar charts as standalone SVG files."""

from __future__ import annotations

from pathlib import Path
from typing import Iterable


def _svg_escape(value: str) -> str:
    return (
        value.replace("&", "&amp;")
        .replace("<", "&lt;")
        .replace(">", "&gt;")
        .replace('"', "&quot;")
    )


def render_horizontal_bar_svg(
    title: str,
    x_label: str,
    rows: Iterable[tuple[str, int]],
    output_path: Path,
    bar_color: str,
) -> None:
    data = list(rows)
    if not data:
        raise ValueError("No rows to plot.")

    max_value = max(value for _, value in data)
    if max_value <= 0:
        max_value = 1

    width = 1400
    top_margin = 70
    right_margin = 40
    bottom_margin = 60
    left_margin = 480
    row_height = 22
    row_gap = 8
    bar_height = row_height
    chart_height = len(data) * (row_height + row_gap)
    height = top_margin + chart_height + bottom_margin
    bar_area_width = width - left_margin - right_margin

    parts = [
        '<?xml version="1.0" encoding="UTF-8"?>',
        f'<svg xmlns="http://www.w3.org/2000/svg" width="{width}" height="{height}" viewBox="0 0 {width} {height}">',
        '<style>',
        '.title { font: 700 22px sans-serif; fill: #111827; }',
        '.axis { font: 12px sans-serif; fill: #374151; }',
        '.label { font: 12px sans-serif; fill: #111827; }',
        '.value { font: 11px sans-serif; fill: #111827; }',
        '.grid { stroke: #e5e7eb; stroke-width: 1; }',
        '</style>',
        f'<text class="title" x="24" y="36">{_svg_escape(title)}</text>',
    ]

    tick_count = 6
    for i in range(tick_count + 1):
        tick_value = round(max_value * i / tick_count)
        x = left_margin + (bar_area_width * i / tick_count)
        parts.append(
            f'<line class="grid" x1="{x:.2f}" y1="{top_margin}" x2="{x:.2f}" y2="{top_margin + chart_height}" />'
        )
        parts.append(
            f'<text class="axis" x="{x:.2f}" y="{top_margin + chart_height + 18}" text-anchor="middle">{tick_value}</text>'
        )

    parts.append(
        f'<text class="axis" x="{left_margin + bar_area_width / 2:.2f}" y="{height - 18}" text-anchor="middle">{_svg_escape(x_label)}</text>'
    )

    for idx, (label, value) in enumerate(data):
        y = top_margin + idx * (row_height + row_gap)
        bar_width = (value / max_value) * bar_area_width
        label_y = y + bar_height * 0.72

        parts.append(
            f'<text class="label" x="{left_margin - 12}" y="{label_y:.2f}" text-anchor="end">{_svg_escape(label)}</text>'
        )
        parts.append(
            f'<rect x="{left_margin}" y="{y}" width="{bar_width:.2f}" height="{bar_height}" rx="2" ry="2" fill="{bar_color}" />'
        )
        parts.append(
            f'<text class="value" x="{left_margin + bar_width + 6:.2f}" y="{label_y:.2f}">{value}</text>'
        )

    parts.append("</svg>")

    output_path.parent.mkdir(parents=True, exist_ok=True)
    output_path.write_text("\n".join(parts), encoding="utf-8")

