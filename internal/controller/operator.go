package controller

import (
	"fuck-the-world/unit"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOperator(c *gin.Context) {
	var operatorInfo struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	err := c.ShouldBindJSON(&operatorInfo)
	if err != nil {
		unit.RespondJSON(c, http.StatusBadRequest, "参数绑定失败", nil)
		return
	}
	if len(operatorInfo.Username) == 0 || len(operatorInfo.Password) == 0 {
		// 用户名密码不能为空
		unit.RespondJSON(c, http.StatusBadRequest, "用户名或密码不能为空", nil)
		return
	}
	//将密码进行加密后存到数据库中
	//var md5Password = md5Hash(operatorInfo.Password)
	//使用gorm查询users表看看用户名是否重复
}
