package ascii

func IconFor(condition string, isNight bool) string {
	switch condition {
	case "Thunderstorm":
		return "⚡"
	case "Drizzle":
		return "☔"
	case "Rain":
		return "🌧"
	case "Snow":
		return "❄"
	case "Mist", "Smoke", "Haze", "Dust", "Fog", "Sand", "Ash", "Squall", "Tornado":
		return "🌫"
	case "Clear":
		if isNight {
			return "🌙"
		}
		return "☀"
	case "Clouds":
		if isNight {
			return "☁"
		}
		return "⛅"
	default:
		return "✦"
	}
}
