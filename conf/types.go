package conf

type InfoConfig struct {
	AppName string `json:"name"`
	Version string `json:"version"`
}

type ProcConfig struct {
	BackUp        bool     `json:"backup"`
	TickPerSec    float64  `json:"tick_per_sec"`
	BackUpExcepts []string `json:"backup_excepts"`
}

type LLMConfig struct {
	ContextLength          int64 `json:"context_length"`
	UseWebSearch           bool  `json:"use_web_search"`
	MaxResponseTokenLength int   `json:"max_response_token_length"`
	ShowToolCallArgs       bool  `json:"show_toolcall_args"`
	SysPrompt              string
	ApiKey                 string   `json:"api_key"`
	BaseURL                string   `json:"base_url"`
	ModelName              string   `json:"model_name"`
	Temperature            float64  `json:"temperature"`
	DirExcepts             []string `json:"dir_excepts"`
}

type Environment struct {
	System string
}

type EnvConfig struct {
	LLMConfig  `json:"llm"`
	ProcConfig `json:"proc"`
	InfoConfig `json:"info"`
	Environment
}
