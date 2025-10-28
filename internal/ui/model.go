package ui

import (
	"context"
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/guicybercode/go_temp_tui/internal/ascii"
	"github.com/guicybercode/go_temp_tui/internal/geo"
	"github.com/guicybercode/go_temp_tui/internal/theme"
	"github.com/guicybercode/go_temp_tui/internal/weather"
)

type ThemeName string

const (
	ThemeTokyo ThemeName = "tokyo"
	ThemeNord  ThemeName = "nord"
	ThemeDrac  ThemeName = "dracula"
)

type Model struct {
	width  int
	height int

	styles Styles
	pal    theme.Palette

	apiKey string
	units  string
	lang   string
	th     ThemeName

	loading   bool
	phaseTick bool

	city     string
	cur      weather.Current
	forecast weather.Forecast
	err      error
}

type fetchMsg struct{}
type weatherMsg struct {
	cur  weather.Current
	fc   weather.Forecast
	city string
	err  error
}
type tickMsg time.Time

func NewModel(th ThemeName, units, lang string) Model {
	var pal theme.Palette
	switch th {
	case ThemeNord:
		pal = theme.Nord()
	case ThemeDrac:
		pal = theme.Dracula()
	default:
		pal = theme.TokyoNight()
	}
	return Model{pal: pal, th: th, units: units, lang: lang, apiKey: os.Getenv("WEATHERBIT_API_KEY"), loading: true}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.fetchCmd(), m.tickCmd())
}

func (m Model) fetchCmd() tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = ctx
		ip, err := geo.GetIPLocation()
		if err != nil {
			return weatherMsg{err: err}
		}
		cli := weather.NewClient(m.apiKey, m.units)
		cur, err := cli.CurrentByCoords(ip.Lat, ip.Lon, m.lang)
		if err != nil {
			return weatherMsg{err: err}
		}
		fc, err := cli.ForecastByCoords(ip.Lat, ip.Lon, m.lang)
		if err != nil {
			return weatherMsg{err: err}
		}
		city := ip.City
		if city == "" {
			city = cur.Name
		}
		return weatherMsg{cur: cur, fc: fc, city: city}
	}
}

func (m Model) tickCmd() tea.Cmd {
	return tea.Tick(750*time.Millisecond, func(t time.Time) tea.Msg { return tickMsg(t) })
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch v := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = v.Width, v.Height
		m.styles = NewStyles(m.pal, m.width)
		return m, nil
	case weatherMsg:
		m.loading = false
		m.cur = v.cur
		m.forecast = v.fc
		m.city = v.city
		m.err = v.err
		return m, tea.Tick(10*time.Minute, func(t time.Time) tea.Msg { return fetchMsg{} })
	case fetchMsg:
		m.loading = true
		return m, m.fetchCmd()
	case tickMsg:
		m.phaseTick = !m.phaseTick
		return m, m.tickCmd()
	}
	return m, nil
}

func (m Model) isNight() bool {
	if len(m.cur.Data) == 0 {
		return false
	}
	return len(m.cur.Data[0].Weather.Icon) > 0 && m.cur.Data[0].Weather.Icon[len(m.cur.Data[0].Weather.Icon)-1] == 'n'
}

func (m Model) TitleFiglet() string {
	return `
██╗    ██╗███████╗ █████╗ ████████╗██╗  ██╗███████╗██████╗ 
██║    ██║██╔════╝██╔══██╗╚══██╔══╝██║  ██║██╔════╝██╔══██╗
██║ █╗ ██║█████╗  ███████║   ██║   ███████║█████╗  ██████╔╝
██║███╗██║██╔══╝  ██╔══██║   ██║   ██╔══██║██╔══╝  ██╔══██╗
╚███╔███╔╝███████╗██║  ██║   ██║   ██║  ██║███████╗██║  ██║
 ╚══╝╚══╝ ╚══════╝╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝
`
}

func (m Model) Icon() string {
	if len(m.cur.Data) == 0 {
		return "✦"
	}
	return ascii.IconFor(m.cur.Data[0].Weather.Main, m.isNight())
}

func (m Model) MetricLine() string {
	unit := "°C"
	if m.units == "imperial" {
		unit = "°F"
	}
	if len(m.cur.Data) == 0 {
		return "—"
	}
	val := fmt.Sprintf("%.0f%s", m.cur.Data[0].Temp, unit)
	if m.phaseTick {
		return m.styles.Metric.Render(val)
	}
	return m.styles.Metric.Copy().Faint(true).Render(val)
}
