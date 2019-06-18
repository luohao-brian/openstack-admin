package main

import (
	"fmt"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/luohao-brian/openstack-admin/models"
	"github.com/luohao-brian/openstack-admin/pkg/logging"
	"github.com/luohao-brian/openstack-admin/pkg/setting"
	"github.com/luohao-brian/openstack-admin/routers"
)

func main() {
	setting.Setup()
	logging.Setup()
	models.Setup()

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	endless.DefaultReadTimeOut = readTimeout
	endless.DefaultWriteTimeOut = writeTimeout
	endless.DefaultMaxHeaderBytes = maxHeaderBytes
	server := endless.NewServer(endPoint, routersInit)
	server.BeforeBegin = func(add string) {
		logging.Info("Actual pid is", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		logging.Error("Server err: ", err)
	}
}
