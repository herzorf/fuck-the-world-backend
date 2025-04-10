package middleware

import (
	"fuck-the-world/internal/model"
	"fuck-the-world/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthAdminMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != model.RoleAdmin {
			utils.RespondJSON(c, http.StatusForbidden, "无权限，必须是管理员", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
