package todo_cli

import "fmt"

// ColorDefault represents the default color for text.
const (
	ColorDefault = "\033[0m"

	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorWhite   = "\033[37m"
)

// Colorize applies the specified color to the given text and returns the colorized text.
func Colorize(color, text string) string {
	//return color + text + ColorDefault
	return fmt.Sprintf("%s%s%s", color, text, ColorDefault)
}

func red(text string) string {
	return Colorize(ColorRed, text)
}

func green(text string) string {
	return Colorize(ColorGreen, text)
}

func yellow(text string) string {
	return Colorize(ColorYellow, text)
}

func blue(text string) string {
	return Colorize(ColorBlue, text)
}

func magenta(text string) string {
	return Colorize(ColorMagenta, text)
}

func cyan(text string) string {
	return Colorize(ColorCyan, text)
}

func white(text string) string {
	return Colorize(ColorWhite, text)
}