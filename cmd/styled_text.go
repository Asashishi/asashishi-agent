package cmd

import (
	"asashishi-agent/global"
)

var EOFMark string = global.LineBreakString + global.GetStyledSystemComent("-- EOF")
var FileContentMark string = global.LineBreakString + global.GetStyledSystemComent("-- File Content:")
