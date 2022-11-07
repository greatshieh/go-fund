package configs

type System struct {
	Debug  bool   `mapstructure:"debug" json:"debug" yaml:"debug"`
	DbType string `mapstructure:"db-type" json:"dbType" yaml:"db-type"` // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
}
