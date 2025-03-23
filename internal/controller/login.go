package controller

import (
	"bookkeeping-server/database"
	"bookkeeping-server/internal/model"
	"bookkeeping-server/internal/pkg/jwt"
	"bookkeeping-server/unit"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func md5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
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
	result := database.DB.Where("username = ?", loginInfo.Username).First(&user)
	if result.Error != nil {
		log.Println(result.RowsAffected)
		c.JSON(400, gin.H{
			"message": "用户不存在",
		})
		return
	} else {
		log.Println(md5Hash("123456"))
		if user.Password != md5Hash(loginInfo.Password) {
			unit.RespondJSON(c, http.StatusBadRequest, "密码错误", nil)
			return
		}
		jwt, err := jwt.GenerateJWT(user.ID)
		if err != nil {
			unit.HandleError("生成JWT失败", err)
		}
		unit.RespondJSON(c, http.StatusOK, "登录成功", gin.H{
			"jwt": jwt,
		})
		return
	}
}
