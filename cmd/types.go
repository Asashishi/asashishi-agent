package cmd

type CmdMap[P any, R any] map[string]func(...P) R
