package controller

import (
	"fuck-the-world/database"
	"fuck-the-world/internal/model"
	"fuck-the-world/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateMap(c *gin.Context) {
	var mapBody model.Map
	err := c.ShouldBindJSON(&mapBody)
	utils.HandleError("参数绑定失败", err)
	if err := database.DB.Create(&mapBody).Error; err != nil {
		utils.HandleError("地图创建失败", err)
		utils.RespondJSON(c, http.StatusInternalServerError, "地图创建失败", nil)
		return
	}
	utils.RespondJSON(c, http.StatusOK, "地图创建成功", nil)
}
