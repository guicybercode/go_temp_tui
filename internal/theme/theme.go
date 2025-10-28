package theme

import "github.com/charmbracelet/lipgloss"

type Palette struct {
	Background lipgloss.Color
	Foreground lipgloss.Color
	Accent1    lipgloss.Color
	Accent2    lipgloss.Color
	Accent3    lipgloss.Color
	Border     lipgloss.Color
}

func TokyoNight() Palette {
	return Palette{
		Background: lipgloss.Color("#1a1b26"),
		Foreground: lipgloss.Color("#c0caf5"),
		Accent1:    lipgloss.Color("#7aa2f7"),
		Accent2:    lipgloss.Color("#bb9af7"),
		Accent3:    lipgloss.Color("#9ece6a"),
		Border:     lipgloss.Color("#2f334d"),
	}
}

func Nord() Palette {
	return Palette{
		Background: lipgloss.Color("#2e3440"),
		Foreground: lipgloss.Color("#e5e9f0"),
		Accent1:    lipgloss.Color("#88c0d0"),
		Accent2:    lipgloss.Color("#b48ead"),
		Accent3:    lipgloss.Color("#a3be8c"),
		Border:     lipgloss.Color("#3b4252"),
	}
}

func Dracula() Palette {
	return Palette{
		Background: lipgloss.Color("#282a36"),
		Foreground: lipgloss.Color("#f8f8f2"),
		Accent1:    lipgloss.Color("#bd93f9"),
		Accent2:    lipgloss.Color("#50fa7b"),
		Accent3:    lipgloss.Color("#8be9fd"),
		Border:     lipgloss.Color("#44475a"),
	}
}
