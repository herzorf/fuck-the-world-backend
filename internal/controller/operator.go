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
		unit.RespondJSON(c, http.StatusBadRequest, "用户名或密码不能为空", nil)
		return
	}
	var existingUser model.User
	if err := database.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		unit.RespondJSON(c, http.StatusBadRequest, "用户已存在", nil)
		return
	}
	if err := user.HashPassword(); err != nil {
		unit.RespondJSON(c, http.StatusInternalServerError, "密码加密失败", nil)
		return
	}
	newUser := model.User{
		Username: user.Username,
		Password: user.Password,
		Role:     model.RoleOperator,
	}
	if err := database.DB.Create(&newUser).Error; err != nil {
		unit.RespondJSON(c, http.StatusInternalServerError, "创建用户失败", nil)
		return
	}
	unit.RespondJSON(c, http.StatusOK, "创建用户成功", nil)
}
