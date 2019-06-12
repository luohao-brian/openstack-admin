package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luohao-brian/openstack-admin/models"
)

func InstanceHandler(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize := 10

	instances, totalCount, _ := models.GetInstances(pageNum, pageSize, getInstanceMaps())

	paginator := Paginator(c.Request.URL.Path, pageNum, pageSize, totalCount)
	c.HTML(http.StatusOK, "instances.html", gin.H{
		"active":    2,
		"Instances": instances,
		"paginator": paginator,
	})
	return
}

func getInstanceMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted"] = 0
	return maps
}
