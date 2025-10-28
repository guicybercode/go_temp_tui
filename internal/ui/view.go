package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if m.width == 0 {
		return ""
	}
	if m.err != nil {
		return m.styles.Container.Render("Error: " + m.err.Error())
	}

	head := m.styles.Header.Render(strings.TrimSpace(m.TitleFiglet()))
	center := m.styles.Bordered.Render(
		lipgloss.JoinVertical(lipgloss.Center,
			m.Icon()+"  "+m.MetricLine(),
			m.styles.Secondary.Render(fmt.Sprintf("%s • %s", m.city, m.conditionText())),
			m.styles.Secondary.Render(m.forecastLine()),
		),
	)

	footer := m.footerSignature()

	content := lipgloss.JoinVertical(lipgloss.Center, head, center, footer)
	return m.styles.Container.Width(m.width).Render(content)
}

func (m Model) conditionText() string {
	if len(m.cur.Data) == 0 {
		return "—"
	}
	return m.cur.Data[0].Weather.Description
}

func (m Model) footerSignature() string {
	text := "made by gui기กีギ"
	pad := 0
	if m.width > len(text)+2 {
		pad = m.width - len(text) - 4
	}
	return m.styles.Footer.Render(strings.Repeat(" ", pad) + text)
}

func (m Model) forecastLine() string {
	if len(m.forecast.Data) == 0 {
		return ""
	}
	max := 4
	if len(m.forecast.Data) < max {
		max = len(m.forecast.Data)
	}
	unit := "°C"
	if m.units == "imperial" {
		unit = "°F"
	}
	parts := make([]string, 0, max)
	for i := 0; i < max; i++ {
		it := m.forecast.Data[i]
		parts = append(parts, fmt.Sprintf("%.0f%s", it.Temp, unit))
	}
	return "next: " + strings.Join(parts, "  ")
}
