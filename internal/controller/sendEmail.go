package controller

import (
	"bookkeeping-server/internal/email"
	"bookkeeping-server/unit"
	"github.com/gin-gonic/gin"
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
	}
	err = email.SendCode(aimEmail.Email, "123456")
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
