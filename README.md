# go_temp_tui – Clean Terminal Weather Widget

<div align="center">

A clean, animated, phone-style weather widget for the terminal.

<img src="clima.png" width="500"/>
</div>

- ASCII icons and figlet titles
- Smooth animations powered by Bubble Tea ticks
- Lip Gloss borders, alignment, and colors
- IP-based geolocation and OpenWeatherMap integration
- Themes: Tokyo Night, Nord, Dracula
- Footer signature: "made by gui기กีギ"

## Features

- Location detection via IP (ipapi.co)
- Current weather and short forecast via Weatherbit
- Metric or Imperial units
- Minimal, modern TUI layout

## Requirements

- Go 1.22+
- Weatherbit API key in `WEATHERBIT_API_KEY`

## Install

```bash
git clone https://github.com/guicybercode/go_temp_tui
cd go_temp_tui
go run ./cmd/weather --theme tokyo --units metric --lang en
```

## Usage

```bash
# Theme: tokyo|nord|dracula
# Units: metric|imperial
# Language: ISO code like en, pt, es
WEATHERBIT_API_KEY=your_key go run ./cmd/weather \
  --theme tokyo \
  --units metric \
  --lang en
```

The app uses IP geolocation to infer city and coordinates. It then fetches the current conditions and a near-term forecast. The temperature text gently pulses and data refreshes periodically.

## Configuration

- Env: `WEATHERBIT_API_KEY`
- Flags: `--theme`, `--units`, `--lang`

## Theming

- Tokyo Night: deep navy background, violet/blue accents
- Nord: cool polar palette with cyan and green accents
- Dracula: dark purple base with neon accents

## Credits

- Charmbracelet Bubble Tea and Lip Gloss
- ipapi.co for geolocation
- Weatherbit for weather data

## License

MIT

## About

Built by `guicybercode` · GitHub: `https://github.com/guicybercode`

References: [GitHub profile](https://github.com/guicybercode), [Target repository](https://github.com/guicybercode/go_temp_tui)

## 성경 구절 (시편 4편)

평안히 내가 눕고 자기도 하리니 나를 안전히 살게 하시는 이는 오직 여호와이시니이다.

