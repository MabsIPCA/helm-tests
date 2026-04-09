"""Small helper to render scatter plots as standalone SVG files."""

from __future__ import annotations

import math
from pathlib import Path
from typing import Iterable


Point = tuple[str, int, int]


def _svg_escape(value: str) -> str:
    return (
        value.replace("&", "&amp;")
        .replace("<", "&lt;")
        .replace(">", "&gt;")
        .replace('"', "&quot;")
    )


def _format_star_tick(value: int) -> str:
    if value >= 1_000_000:
        return f"{value // 1_000_000}M"
    if value >= 1_000:
        return f"{value // 1_000}k"
    return str(value)


def _build_log_star_ticks(max_x: int) -> list[int]:
    ticks = [0]
    if max_x <= 0:
        return [0, 1]

    value = 1
    while value <= max_x:
        for factor in (1, 2, 5):
            tick = factor * value
            if tick <= max_x:
                ticks.append(tick)
        value *= 10

    if max_x not in ticks:
        ticks.append(max_x)

    ticks = sorted(set(ticks))
    # Keep labels readable when the range is very large.
    if len(ticks) > 12:
        step = max(1, len(ticks) // 10)
        compact = [ticks[0]] + ticks[1:-1:step] + [ticks[-1]]
        ticks = sorted(set(compact))
    return ticks


def render_scatter_svg(
    title: str,
    x_label: str,
    y_label: str,
    points: Iterable[Point],
    output_path: Path,
    point_color: str,
    x_scale: str = "linear",
) -> None:
    data = list(points)
    if not data:
        raise ValueError("No points to plot.")

    max_x = max(x for _, x, _ in data)
    max_y = max(y for _, _, y in data)
    if max_x <= 0:
        max_x = 1
    if max_y <= 0:
        max_y = 1

    width = 1300
    height = 900
    margin_left = 90
    margin_right = 40
    margin_top = 70
    margin_bottom = 90
    chart_width = width - margin_left - margin_right
    chart_height = height - margin_top - margin_bottom

    parts = [
        '<?xml version="1.0" encoding="UTF-8"?>',
        f'<svg xmlns="http://www.w3.org/2000/svg" width="{width}" height="{height}" viewBox="0 0 {width} {height}">',
        '<style>',
        '.title { font: 700 22px sans-serif; fill: #111827; }',
        '.axis { font: 12px sans-serif; fill: #374151; }',
        '.grid { stroke: #e5e7eb; stroke-width: 1; }',
        '.dot { fill-opacity: 0.75; }',
        '</style>',
        f'<text class="title" x="24" y="36">{_svg_escape(title)}</text>',
    ]

    tick_count = 6

    if x_scale not in {"linear", "log1p"}:
        raise ValueError("x_scale must be 'linear' or 'log1p'.")

    if x_scale == "log1p":
        x_denominator = math.log10(max_x + 1)
        if x_denominator <= 0:
            x_denominator = 1.0

        def _x_to_svg(value: int) -> float:
            transformed = math.log10(max(value, 0) + 1)
            return margin_left + (transformed / x_denominator) * chart_width

        for tick_value in _build_log_star_ticks(max_x):
            x = _x_to_svg(tick_value)
            parts.append(
                f'<line class="grid" x1="{x:.2f}" y1="{margin_top}" x2="{x:.2f}" y2="{margin_top + chart_height}" />'
            )
            parts.append(
                f'<text class="axis" x="{x:.2f}" y="{margin_top + chart_height + 20}" text-anchor="middle">{_format_star_tick(tick_value)}</text>'
            )
    else:
        def _x_to_svg(value: int) -> float:
            return margin_left + (value / max_x) * chart_width

        # Vertical grid and X ticks.
        for i in range(tick_count + 1):
            tick_value = round(max_x * i / tick_count)
            x = margin_left + (chart_width * i / tick_count)
            parts.append(
                f'<line class="grid" x1="{x:.2f}" y1="{margin_top}" x2="{x:.2f}" y2="{margin_top + chart_height}" />'
            )
            parts.append(
                f'<text class="axis" x="{x:.2f}" y="{margin_top + chart_height + 20}" text-anchor="middle">{tick_value}</text>'
            )

    # Horizontal grid and Y ticks.
    for i in range(tick_count + 1):
        tick_value = round(max_y * i / tick_count)
        y = margin_top + chart_height - (chart_height * i / tick_count)
        parts.append(
            f'<line class="grid" x1="{margin_left}" y1="{y:.2f}" x2="{margin_left + chart_width}" y2="{y:.2f}" />'
        )
        parts.append(
            f'<text class="axis" x="{margin_left - 10}" y="{y + 4:.2f}" text-anchor="end">{tick_value}</text>'
        )

    # Axis labels.
    parts.append(
        f'<text class="axis" x="{margin_left + chart_width / 2:.2f}" y="{height - 28}" text-anchor="middle">{_svg_escape(x_label)}</text>'
    )
    parts.append(
        f'<text class="axis" transform="translate(24 {margin_top + chart_height / 2:.2f}) rotate(-90)" text-anchor="middle">{_svg_escape(y_label)}</text>'
    )

    # Scatter points with tooltip labels.
    for label, x_value, y_value in data:
        cx = _x_to_svg(x_value)
        cy = margin_top + chart_height - (y_value / max_y) * chart_height
        tooltip = f"{label} | stars={x_value}, problems={y_value}"
        parts.append(
            f'<circle class="dot" cx="{cx:.2f}" cy="{cy:.2f}" r="4" fill="{point_color}"><title>{_svg_escape(tooltip)}</title></circle>'
        )

    parts.append("</svg>")

    output_path.parent.mkdir(parents=True, exist_ok=True)
    output_path.write_text("\n".join(parts), encoding="utf-8")


