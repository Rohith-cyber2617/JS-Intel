package colors

const (
	Reset  = "\033[0m"

	Green  = "\033[32m"
	Yellow = "\033[93m"
	Blue   = "\033[96m"
	Red    = "\033[91m"
	Cyan   = "\033[36m"
	White  = "\033[97m"
)

func GreenText(text string) string {
	return Green + text + Reset
}

func YellowText(text string) string {
	return Yellow + text + Reset
}

func BlueText(text string) string {
	return Blue + text + Reset
}

func RedText(text string) string {
	return Red + text + Reset
}

func CyanText(text string) string {
	return Cyan + text + Reset
}

func WhiteText(text string) string {
	return White + text + Reset
}
