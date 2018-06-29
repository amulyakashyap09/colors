package chalk

import (
	"log"
	"strings"
)

type Message interface{}

func Custom(fgColor string, bgColor string, message ...Message) {

	var FgColorCode string
	var BgColorCode string

	fgColor = strings.ToLower(fgColor)
	bgColor = strings.ToLower(bgColor)

	switch fgColor {
	case "red":
		FgColorCode = FgRed
	case "black":
		FgColorCode = FgBlack
	case "blue":
		FgColorCode = FgBlue
	case "cyan":
		FgColorCode = FgCyan
	case "magenta":
		FgColorCode = FgMagenta
	case "green":
		FgColorCode = FgGreen
	case "white":
		FgColorCode = FgWhite
	case "yellow":
		FgColorCode = FgYellow
	default:
		FgColorCode = FgWhite
	}

	switch bgColor {
	case "red":
		BgColorCode = BgRed
	case "black":
		BgColorCode = BgBlack
	case "blue":
		BgColorCode = BgBlue
	case "cyan":
		BgColorCode = BgCyan
	case "magenta":
		BgColorCode = BgMagenta
	case "green":
		BgColorCode = BgGreen
	case "white":
		BgColorCode = BgWhite
	case "yellow":
		BgColorCode = BgYellow
	default:
		BgColorCode = BgBlack
	}

	log.Println(BgColorCode, FgColorCode, message, Reset)
}

func Underline(message ...Message) {
	log.Println(Underscore, message, Reset)
}

func Bold(message ...Message) {
	log.Println(Bright, message, Reset)
}

func Flash(message ...Message) {
	log.Println(Blink, message, Reset)
}

func Inverse(message ...Message) {
	log.Println(Reverse, message, Reset)
}

func Panic(message ...Message) {
	log.Panic(FgRed, message, Reset)
}

func Highlight(message ...Message) {
	log.Println(BgYellow, FgRed, message, Reset)
}

func Important(message ...Message) {
	log.Println(BgRed, FgWhite, message, Reset)
}

func Success(message ...Message) {
	log.Println(FgGreen, message, Reset)
}

func Info(message ...Message) {
	log.Println(FgBlue, message, Reset)
}

func Warn(message ...Message) {
	log.Println(FgYellow, message, Reset)
}

func Error(message ...Message) {
	log.Println(FgMagenta, Bright, message, Reset)
}

func Green(message ...Message) {
	log.Println(FgGreen, message, Reset)
}

func Blue(message ...Message) {
	log.Println(FgBlue, message, Reset)
}

func Magenta(message ...Message) {
	log.Println(FgMagenta, message, Reset)
}

func Red(message ...Message) {
	log.Println(FgRed, message, Reset)
}

func Black(message ...Message) {
	log.Println(BgWhite, FgBlack, message, Reset)
}

func Cyan(message ...Message) {
	log.Println(FgCyan, message, Reset)
}

func White(message ...Message) {
	log.Println(BgBlack, FgWhite, message, Reset)
}

func Yellow(message ...Message) {
	log.Println(FgYellow, message, Reset)
}
