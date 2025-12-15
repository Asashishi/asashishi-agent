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
	var styled string = ""
	if style.Bold {
		styled += Bold
	}
	if style.Dim {
		styled += Dim
	}
	if style.Italic {
		styled += Italic
	}
	if style.Underline {
		styled += Underline
	}
	if style.Blink {
		styled += Blink
	}
	if style.FastBlink {
		styled += FastBlink
	}
	if style.Reverse {
		styled += Reverse
	}
	if style.Hidden {
		styled += Hidden
	}
	if style.Strikethrough {
		styled += Strikethrough
	}
	if style.Fg != "" {
		styled += FgColorBase + hexToANSI(style.Fg)
	}
	if style.Bg != "" {
		styled += BgColorBase + hexToANSI(style.Bg)
	}
	return styled + content + Reset
}
