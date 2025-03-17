package controller

import (
	"bookkeeping-server/database"
	"bookkeeping-server/internal/model"
	"bookkeeping-server/internal/pkg/email"
	"bookkeeping-server/unit"
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
)

func SendEmail(c *gin.Context) {
	type Email struct {
		Email string `json:"email"`
	}
	var aimEmail = &Email{}
	err := c.ShouldBindJSON(aimEmail)
	if err != nil {
		unit.HandleError("sendEmail接口数据读取失败", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "数据读取失败",
		})
		return
	}
	code, _ := generateCode(6)
	err = database.DB.Create(&model.ValidationEmailCode{
		Email: aimEmail.Email,
		Code:  code,
	}).Error
	if err != nil {
		unit.HandleError("sendEmail接口数据库写入失败", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "数据库写入失败",
		})
		return
	}
	err = email.SendCode(aimEmail.Email, code)
	if err != nil {
		unit.HandleError("sendEmail接口发送邮件失败", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "发送失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "发送成功",
		})
	}
}

func generateCode(len int) (string, error) {
	key := make([]byte, len)
	var err1 error
	for i := range key {
		n, err := rand.Int(rand.Reader, big.NewInt(10))
		err1 = err
		key[i] = byte(n.Int64() + 48)
	}
	return string(key), err1
}
