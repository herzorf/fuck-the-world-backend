package controller

import (
	"fuck-the-world/database"
	"fuck-the-world/internal/model"
	"fuck-the-world/unit"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOperator(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		unit.RespondJSON(c, http.StatusBadRequest, "参数绑定失败", nil)
		return
	}
	if len(user.Username) == 0 || len(user.Password) == 0 {
		unit.RespondJSON(c, http.StatusBadRequest, "用户名或密码不能为空", nil)
		return
	}
	var existingUser model.User
	if err := database.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		unit.RespondJSON(c, http.StatusBadRequest, "用户已存在", nil)
		return
	}
	if err := user.HashPassword(); err != nil {
		unit.RespondJSON(c, http.StatusInternalServerError, "密码加密失败", nil)
		return
	}
	newUser := model.User{
		Username: user.Username,
		Password: user.Password,
		Role:     model.RoleOperator,
	}
	if err := database.DB.Create(&newUser).Error; err != nil {
		unit.RespondJSON(c, http.StatusInternalServerError, "创建用户失败", nil)
		return
	}
	unit.RespondJSON(c, http.StatusOK, "创建用户成功", nil)
}

func DeleteOperator(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		unit.RespondJSON(c, http.StatusBadRequest, "参数绑定失败", nil)
		return
	}
	if user.ID == 0 {
		unit.RespondJSON(c, http.StatusBadRequest, "用户ID不能为空", nil)
		return
	}
	var existingUser model.User
	if err := database.DB.Where("id = ?", user.ID).First(&existingUser).Error; err != nil {
		unit.RespondJSON(c, http.StatusBadRequest, "用户不存在", nil)
		return
	}
	if existingUser.IsDeleted == true {
		unit.RespondJSON(c, http.StatusBadRequest, "用户已删除", nil)
		return
	}
	existingUser.IsDeleted = true
	existingUser.IsActive = false
	if err := database.DB.Save(&existingUser).Error; err != nil {
		unit.RespondJSON(c, http.StatusInternalServerError, "删除用户失败", nil)
		return
	}
	unit.RespondJSON(c, http.StatusOK, "删除用户成功", nil)
}
func UpdateOperator(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		unit.RespondJSON(c, http.StatusBadRequest, "参数绑定失败", nil)
		return
	}
	if len(user.Username) == 0 {
		unit.RespondJSON(c, http.StatusBadRequest, "用户名不能为空", nil)
		return
	}
	var existingUser model.User
	if err := database.DB.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
		unit.RespondJSON(c, http.StatusBadRequest, "用户不存在", nil)
		return
	}
	if len(user.Password) != 0 {
		existingUser.Password = user.Password
	}
	if len(user.Role) != 0 {
		existingUser.Role = user.Role
	}
	if err := database.DB.Save(&existingUser).Error; err != nil {
		unit.RespondJSON(c, http.StatusInternalServerError, "修改用户失败", nil)
		return
	}
	unit.RespondJSON(c, http.StatusOK, "修改用户成功", nil)
}
func QueryOperatorList(c *gin.Context) {
	var users []model.User
	if err := database.DB.Where("role = ? AND is_deleted = ?", model.RoleOperator, false).Find(&users).Error; err != nil {
		unit.RespondJSON(c, http.StatusInternalServerError, "查询用户列表失败", nil)
	}
	// 删除密码字段
	for i := range users {
		users[i].Password = ""
	}
	unit.RespondJSON(c, http.StatusOK, "", users)
}
