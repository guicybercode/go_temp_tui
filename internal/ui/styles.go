package ui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/guicybercode/go_temp_tui/internal/theme"
)

type Styles struct {
	Container lipgloss.Style
	Header    lipgloss.Style
	Metric    lipgloss.Style
	Secondary lipgloss.Style
	Bordered  lipgloss.Style
	Footer    lipgloss.Style
}

func NewStyles(p theme.Palette, width int) Styles {
	container := lipgloss.NewStyle().
		Foreground(p.Foreground).
		Background(p.Background).
		Padding(1, 2)

	header := lipgloss.NewStyle().
		Foreground(p.Accent1).
		Bold(true)

	metric := lipgloss.NewStyle().
		Foreground(p.Accent2).
		Bold(true)

	secondary := lipgloss.NewStyle().
		Foreground(p.Accent3)

	bordered := lipgloss.NewStyle().
		BorderForeground(p.Border).
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2)

	footer := lipgloss.NewStyle().
		Foreground(p.Accent1)

	_ = width
	return Styles{Container: container, Header: header, Metric: metric, Secondary: secondary, Bordered: bordered, Footer: footer}
}
