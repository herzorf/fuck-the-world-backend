package controller

import (
	"fuck-the-world/database"
	"fuck-the-world/internal/model"
	"fuck-the-world/unit"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOperator(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		unit.RespondJSON(c, http.StatusBadRequest, "参数绑定失败", nil)
		return
	}
	if len(user.Username) == 0 || len(user.Password) == 0 {
		// 用户名密码不能为空
		unit.RespondJSON(c, http.StatusBadRequest, "用户名或密码不能为空", nil)
		return
	}
	//1.先检查用户名是否已经存在
	//2.如果已经存在则返回“用户名已存在”
	//3.如果用户名不存在则创建该用户
	var existingUser model.User
	if err := database.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		unit.RespondJSON(c, http.StatusBadRequest, "用户已存在", nil)
		return
	}
	if err := database.DB.Where("username = ?", user.Username).First(&user).Error; err == nil {

	}
}
