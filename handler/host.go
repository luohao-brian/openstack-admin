package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luohao-brian/openstack-admin/models"
)

func HostHandler(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize := 10

	services, totalCount, _ := models.GetServices(pageNum, pageSize, getHostMaps())

	paginator := Paginator(c.Request.URL.Path, pageNum, pageSize, totalCount)
	c.HTML(http.StatusOK, "hosts.html", gin.H{
		"active":    1,
		"Services":  services,
		"paginator": paginator,
	})
	return
}

func getHostMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted"] = 0
	return maps
}
