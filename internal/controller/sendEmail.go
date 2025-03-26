package controller

import (
	"crypto/rand"
	"fuck-the-world/database"
	"fuck-the-world/internal/model"
	"fuck-the-world/internal/pkg/email"
	"fuck-the-world/unit"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
)

// SendEmail godoc
// @Summary sendEmail
// @Description 发送验证码
// @Tags sendEmail
// @Accept json
// @Produce json
// @Success 200
// @error 500
// @Router /sendEmail [post]
func SendEmail(c *gin.Context) {
	type Email struct {
		Email string `json:"email"`
	}
	var aimEmail = &Email{}
	err := c.ShouldBindJSON(aimEmail)
	if err != nil {
		unit.HandleError("sendEmail接口数据读取失败", err)
		unit.RespondJSON(c, http.StatusBadRequest, "数据读取失败", nil)
		return
	}
	code, _ := generateCode(6)
	err = database.DB.Create(&model.ValidationEmailCode{
		Email: aimEmail.Email,
		Code:  code,
	}).Error
	if err != nil {
		unit.HandleError("sendEmail接口数据库写入失败", err)
		unit.RespondJSON(c, http.StatusInternalServerError, "数据库写入失败", nil)
		return
	}
	err = email.SendCode(aimEmail.Email, code)
	if err != nil {
		unit.HandleError("sendEmail接口发送邮件失败", err)
		unit.RespondJSON(c, http.StatusInternalServerError, "发送邮件失败", nil)
	} else {
		unit.RespondJSON(c, http.StatusOK, "发送邮件成功", nil)
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
