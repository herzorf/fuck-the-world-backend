package controller

import (
	"fuck-the-world/database"
	"fuck-the-world/internal/model"
	"fuck-the-world/internal/pkg/jwt"
	"fuck-the-world/unit"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	type LoginInfo struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var loginInfo = &LoginInfo{}
	err := c.ShouldBindJSON(loginInfo)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "数据读取失败",
		})
		return
	}
	var user model.User
	result := database.DB.Where("username = ? AND is_deleted = ?", loginInfo.Username, false).First(&user)
	if result.Error != nil {
		unit.RespondJSON(c, http.StatusBadRequest, "用户不存在", nil)
		return
	} else {
		if !user.CheckPassword(loginInfo.Password) {
			unit.RespondJSON(c, http.StatusBadRequest, "密码错误", nil)
			return
		}
		if !user.IsActive {
			unit.RespondJSON(c, http.StatusForbidden, "该用户不可用，请联系管理员", nil)
			return
		}
		jwtString, err := FTWJwt.GenerateJWT(user)
		if err != nil {
			unit.HandleError("生成JWT失败", err)
		}
		unit.RespondJSON(c, http.StatusOK, "登录成功", gin.H{
			"jwt":      jwtString,
			"username": user.Username,
		})
		return
	}
}
