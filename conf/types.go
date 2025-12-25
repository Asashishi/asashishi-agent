package conf

type InfoConfig struct {
	AppName string `json:"name"`
	Version string `json:"version"`
}

type WebCrosConfig struct {
	AllowHeaders string   `json:"allow_headers"`
	AllowMethods string   `json:"allow_methods"`
	AllowOrigins []string `json:"allow_origins"`
}

type WebConfig struct {
	WebCrosConfig  `json:"cros"`
	WebMode        bool   `json:"web_mode"`
	ServerListen   string `json:"server_listen"`
	StatusRoute    string `json:"status_route"`
	WebsocketRoute string `json:"websocket_route"`
	ServerRootPath string `json:"server_root_path"`
	ServerBaseURL  string `json:"server_base_url"`
}

type ProcConfig struct {
	WebConfig         `json:"web"`
	BackUp            bool     `json:"backup"`
	TickPerSec        float64  `json:"tick_per_sec"`
	TerminalCodeStyle string   `json:"terminal_code_style"`
	BackUpExcepts     []string `json:"backup_excepts"`
}

type LLMConfig struct {
	ContextLength          int64 `json:"context_length"`
	UseWebSearch           bool  `json:"use_web_search"`
	MaxResponseTokenLength int   `json:"max_response_token_length"`
	ShowToolCallArgs       bool  `json:"show_toolcall_args"`
	SysPrompt              string
	ApiKey                 string   `json:"api_key"`
	ModleBaseURL           string   `json:"model_base_url"`
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
