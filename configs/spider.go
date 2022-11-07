package configs

type Spider struct {
	ProxyEnable bool     `mapstructure:"proxy-enable" json:"proxyEnable" yaml:"proxy-enable"`
	ProxyHost   string   `mapstructure:"proxy-host" json:"proxyHost" yaml:"proxy-host"`
	ProxyPort   string   `mapstructure:"proxy-port" json:"proxyPort" yaml:"proxy-port"`
	RoutineNum  int      `mapstructure:"routine-num" json:"routineNum" yaml:"routine-num"`
	UserAgent   []string `mapstructure:"user-agent" json:"userAgent" yaml:"user-agent"`
	FilePath    string   `mapstructure:"file-path" json:"filePath" yaml:"file-path"`
}
