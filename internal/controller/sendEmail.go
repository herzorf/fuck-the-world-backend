package controller

import (
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
	c.ShouldBindJSON(aimEmail)
	log.Println("发送邮件给", aimEmail)
	c.String(http.StatusBadRequest, fmt.Sprintf("发送邮件给:%s", aimEmail.Email))
}
