package main

import (
	"flag"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/guicybercode/go_temp_tui/internal/ui"
)

func main() {
	units := flag.String("units", "metric", "units: metric or imperial")
	theme := flag.String("theme", "tokyo", "theme: tokyo|nord|dracula")
	lang := flag.String("lang", "en", "language code")
	flag.Parse()

	var th ui.ThemeName
	switch *theme {
	case "nord":
		th = ui.ThemeNord
	case "dracula":
		th = ui.ThemeDrac
	default:
		th = ui.ThemeTokyo
	}

	m := ui.NewModel(ui.ThemeName(th), *units, *lang)
	if err := tea.NewProgram(m, tea.WithAltScreen()).Start(); err != nil {
		log.Fatal(err)
	}
}
