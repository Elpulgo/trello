package commandColors

import (
	"github.com/gookit/color"
)

func Red(value string) string {
	return color.FgRed.Render(value)
}

func RedBold(value string) string {
	return color.Style{color.FgRed, color.OpBold}.Render(value)
}

func Green(value string) string {
	return color.FgGreen.Render(value)
}

func GreenBold(value string) string {
	return color.Style{color.FgGreen, color.OpBold}.Render(value)
}

func Yellow(value string) string {
	return color.FgYellow.Render(value)

}
func YellowBold(value string) string {
	return color.Style{color.FgYellow, color.OpBold}.Render(value)
}

func Cyan(value string) string {
	return color.FgCyan.Render(value)

}
func CyanBold(value string) string {
	return color.Style{color.FgCyan, color.OpBold}.Render(value)
}

func DarkGrey(value string) string {
	return color.FgDarkGray.Render(value)

}
func DarkGreyBold(value string) string {
	return color.Style{color.FgDarkGray, color.OpBold}.Render(value)
}
