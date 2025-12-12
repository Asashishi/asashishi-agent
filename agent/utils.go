package agent

func IsUseNewLine(cli *AgentClient) bool {
	var isHasContentAtOutput bool = false
	if cli.StdOutHasContent {
		isHasContentAtOutput = true
		cli.StdOutHasContent = false
	}
	return cli.MsgContext[len(cli.MsgContext)-1].OfTool == nil || isHasContentAtOutput
}
