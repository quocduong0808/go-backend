package setting

type Config struct {
	Mysql MysqlSetting `mapstructure:"mysql"`
}

type MysqlSetting struct {
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleCons     string `mapstructure:"maxIdleCons"`
	MaxOpenCons     string `mapstructure:"maxOpenCons"`
	ConnMaxLifeTime string `mapstructure:"connMaxLifeTime"`
}
