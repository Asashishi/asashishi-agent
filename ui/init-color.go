package ui

import (
	"fmt"
	"strconv"
)

func hexToANSI(hex string) string {
	var (
		R   int64
		G   int64
		B   int64
		err error
	)
	if hex[0] == '#' {
		hex = hex[1:]
	}
	if R, err = strconv.ParseInt(hex[0:2], 16, 0); err != nil {
		panic(ExceptionAtColorHex)
	} else if G, err = strconv.ParseInt(hex[2:4], 16, 0); err != nil {
		panic(ExceptionAtColorHex)
	} else if B, err = strconv.ParseInt(hex[4:6], 16, 0); err != nil {
		panic(ExceptionAtColorHex)
	}
	return fmt.Sprintf("2;%d;%d;%dm", R, G, B)
}

func ResetStyle() {
	fmt.Print(Reset)
}

func WidthStyle(content string, style Style) string {
	var result string = ""
	if style.Bold {
		result += Bold
	}
	if style.Dim {
		result += Dim
	}
	if style.Italic {
		result += Italic
	}
	if style.Underline {
		result += Underline
	}
	if style.Blink {
		result += Blink
	}
	if style.FastBlink {
		result += FastBlink
	}
	if style.Reverse {
		result += Reverse
	}
	if style.Hidden {
		result += Hidden
	}
	if style.Strikethrough {
		result += Strikethrough
	}
	if style.Fg != "" {
		result += FgColorBase + hexToANSI(style.Fg)
	}
	if style.Bg != "" {
		result += BgColorBase + hexToANSI(style.Bg)
	}
	return result + content + Reset
}
