package configs

type Server struct {
	Zap   Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`
	Email Email `mapstructure:"email" json:"email" yaml:"email"`

	System System `mapstructure:"system" json:"system" yaml:"system"`

	// gorm
	PgSQL  PgSQL `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	MySQL  MySQL `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	DBList []DB  `mapstructure:"db-list" json:"db-list" yaml:"db-list"`

	// spider
	Spider Spider `mapstructure:"spider" json:"spider" yaml:"spider"`
}
