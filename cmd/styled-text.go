package cmd

import (
	"asashishi-agent/global"
)

var FileContentMark string = global.LineBreakString + global.GetStyledSystemComent("-- File Content:")

var EOFMark string = global.LineBreakString + global.GetStyledSystemComent("-- EOF")
