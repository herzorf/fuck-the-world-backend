package controller

import (
	"bookkeeping-server/database"
	"bookkeeping-server/internal/model"
	"bookkeeping-server/internal/pkg/email"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "数据读取失败",
		})
		return
	}
	err = database.DB.Create(&model.ValidationEmailCode{
		Email: aimEmail.Email,
		Code:  "123456",
	}).Error
	if err != nil {
		unit.HandleError("sendEmail接口数据库写入失败", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "数据库写入失败",
		})
		return
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
