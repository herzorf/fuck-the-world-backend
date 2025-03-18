package controller

import (
	"bookkeeping-server/database"
	"bookkeeping-server/internal/model"
	"github.com/gin-gonic/gin"
	"log"
)

func Login(c *gin.Context) {
	type LoginInfo struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}
	var loginInfo = &LoginInfo{}
	err := c.ShouldBindJSON(loginInfo)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "数据读取失败",
		})
		return
	}
	result := database.DB.Where("email = ? AND code = ?", loginInfo.Email, loginInfo.Code).First(&model.ValidationEmailCode{})
	if result.Error != nil {
		log.Println(result.RowsAffected)
		c.JSON(400, gin.H{
			"message": "邮箱和验证码不匹配",
		})
		return
	} else {
		JWT := "12121212121"
		c.JSON(200, gin.H{
			"message": "登录成功",
			"result":  JWT,
		})
		return
	}
}
