package controller

import (
	"bookkeeping-server/unit"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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
	log.Println("发送邮件给", aimEmail)
	c.String(http.StatusBadRequest, fmt.Sprintf("发送邮件给:%s", aimEmail.Email))
}
