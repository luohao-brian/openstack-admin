package routers

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/luohao-brian/openstack-admin/handler"
	"github.com/luohao-brian/openstack-admin/middleware"
	"github.com/luohao-brian/openstack-admin/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	r.StaticFS("/static", http.Dir("static"))
	r.StaticFS("/img", http.Dir("static/img"))
	r.LoadHTMLGlob("templates/admin/*")

	store, _ := sessions.NewRedisStore(
		setting.RedisSetting.MaxIdle,
		"tcp",
		setting.RedisSetting.Host,
		setting.RedisSetting.Password,
		[]byte("secret"))

	r.Use(sessions.Sessions("mysession", store))

	g1 := r.Group("/")
	g1.Use(middleware.AuthRequired())
	g1.GET("/", handler.HostHandler)
	g1.GET("/instances", handler.InstanceHandler)

	r.GET("/login", handler.LoginHandler)
	r.POST("/login", handler.LoginPostHandler)
	return r
}
