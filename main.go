package main

import (
	"fmt"
	"net/http"
	"sky_take_out/controller"
	"sky_take_out/database"
	"sky_take_out/utils"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	utils.InitLogger()

	utils.InitViper()

	database.InitDB()
	defer database.CloseDB()

	controller.EmployeeMakeHandler(database.GetDB())
	controller.CategoryMakeHandler(database.GetDB())

	host := viper.GetString("server.host")
	port := viper.GetString("server.port")
	addr := fmt.Sprintf("%s:%s", host, port)
	utils.Logger.Info("服务开启...", zap.String("addr", addr))
	if err := http.ListenAndServe(addr, nil); err != nil { // 监听成功就会阻塞，不会执行后面的代码
		utils.Logger.Error("服务监听失败", zap.Error(err))
	}
}
