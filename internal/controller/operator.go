package controller

import (
	"fuck-the-world/database"
	"fuck-the-world/internal/model"
	"fuck-the-world/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateOperator(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "参数绑定失败", nil)
		return
	}
	if len(user.Username) == 0 || len(user.Password) == 0 {
		utils.RespondJSON(c, http.StatusBadRequest, "用户名或密码不能为空", nil)
		return
	}
	var existingUser model.User
	if err := database.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		utils.RespondJSON(c, http.StatusBadRequest, "用户已存在", nil)
		return
	}
	if err := user.HashPassword(); err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "密码加密失败", nil)
		return
	}
	newUser := model.User{
		Username: user.Username,
		Password: user.Password,
		IsActive: user.IsActive,
		Role:     model.RoleOperator,
	}
	if err := database.DB.Create(&newUser).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "创建用户失败", nil)
		return
	}
	utils.RespondJSON(c, http.StatusOK, "创建用户成功", nil)
}

func DeleteOperator(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "参数绑定失败", nil)
		return
	}
	if user.ID == 0 {
		utils.RespondJSON(c, http.StatusBadRequest, "用户ID不能为空", nil)
		return
	}
	var existingUser model.User
	if err := database.DB.Where("id = ?", user.ID).First(&existingUser).Error; err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "用户不存在", nil)
		return
	}
	if *existingUser.IsDeleted == true {
		utils.RespondJSON(c, http.StatusBadRequest, "用户已删除", nil)
		return
	}
	*existingUser.IsDeleted = true
	*existingUser.IsActive = false
	if err := database.DB.Save(&existingUser).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "删除用户失败", nil)
		return
	}
	utils.RespondJSON(c, http.StatusOK, "删除用户成功", nil)
}
func UpdateOperator(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "参数绑定失败", nil)
		return
	}
	if user.ID == 0 {
		utils.RespondJSON(c, http.StatusBadRequest, "用户id不能为空", nil)
		return
	}
	var existingUser model.User
	if err := database.DB.Where("id = ?", user.ID).First(&existingUser).Error; err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "用户不存在", nil)
		return
	}
	if *existingUser.IsDeleted == true {
		utils.RespondJSON(c, http.StatusBadRequest, "用户已删除", nil)
		return
	}

	if err := database.DB.Model(&user).Where("id = ?", user.ID).Update("is_active", user.IsActive).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "修改用户失败", nil)
		return
	}
	utils.RespondJSON(c, http.StatusOK, "修改用户成功", nil)
}
func QueryOperatorList(c *gin.Context) {
	type queryBodyParams struct {
		PageNo   int    `json:"pageNo" binding:"gte=1"`
		PageSize int    `json:"pageSize" binding:"gte=1"`
		Username string `json:"username"`
	}
	var body queryBodyParams
	err := c.ShouldBindJSON(&body)
	if err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "参数绑定失败", nil)
		return
	}
	type User struct {
		ID        uint   `json:"id"`
		Username  string `json:"username"`
		UpdatedAt string `json:"updatedAt"`
		IsActive  *bool  `json:"isActive"`
	}
	var users []model.User
	var total int64
	var query = database.DB.Model(&model.User{}).Where("is_deleted = ? AND role= ?", false, model.RoleOperator)
	if len(body.Username) > 0 {
		query = query.Where("username LIKE ?", "%"+body.Username+"%")
	}
	query.Count(&total)
	query = query.Order("created_at DESC")
	query = query.Offset((body.PageNo - 1) * body.PageSize).Limit(body.PageSize)
	if err := query.Select("id, username, role, updated_at, is_active").Find(&users).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "查询用户失败", nil)
		return
	}
	var responseUsers = make([]User, 0)
	for i := range users {
		responseUsers = append(responseUsers, User{
			ID:        users[i].ID,
			Username:  users[i].Username,
			UpdatedAt: users[i].UpdatedAt.Format(time.DateTime),
			IsActive:  users[i].IsActive,
		})

	}
	utils.RespondJSON(c, http.StatusOK, "", gin.H{
		"total": total,
		"list":  responseUsers,
	})
}
