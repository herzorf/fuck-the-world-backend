package controller

import (
	"bookkeeping-server/database"
	"bookkeeping-server/internal/model"
	"bookkeeping-server/internal/pkg/jwt"
	"bookkeeping-server/unit"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
		jwt, err := jwt.GenerateJWT(loginInfo.Email)
		var user model.User
		result := database.DB.Where("email = ?", loginInfo.Email).First(&user)
		if result.RowsAffected == 0 {
			user = model.User{Email: loginInfo.Email}
			database.DB.Create(&user)
			log.Println("没找到", user)
		}
		log.Println(user)
		if err != nil {
			unit.HandleError("生成JWT失败", err)
		}
		unit.RespondJSON(c, http.StatusOK, "登录成功", gin.H{
			"jwt":    jwt,
			"userId": user.ID,
		})
		return
	}
}
