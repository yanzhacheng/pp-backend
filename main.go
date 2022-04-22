package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pp-backend/models"
	"pp-backend/pkg/logging"
	"pp-backend/pkg/setting"
	"pp-backend/pkg/util"
	"pp-backend/routers"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	util.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description Ping Pong Management System Backend
// @termsOfService https://github.com/yanzhacheng/pp-backend
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf("127.0.0.1:%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
