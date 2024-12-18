package initialize

import (
	"fmt"
	"go/go-backend-api/global"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRegis()
	r := InitRouter()
	r.Run(fmt.Sprintf(":%v", global.Config.Server.Port))
}
