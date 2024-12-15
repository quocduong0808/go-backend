package initialize

func Run() {
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRegis()
	r := InitRouter()
	r.Run(":8002")
}
