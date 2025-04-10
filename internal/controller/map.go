package controller

import (
	"fuck-the-world/database"
	"fuck-the-world/internal/model"
	"fuck-the-world/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateMap(c *gin.Context) {
	var mapBody model.Map
	err := c.ShouldBindJSON(&mapBody)
	utils.HandleError("参数绑定失败", err)
	if err := database.DB.Create(&mapBody).Error; err != nil {
		utils.HandleError("地图创建失败", err)
		utils.RespondJSON(c, http.StatusInternalServerError, "地图创建失败", nil)
		return
	}
	utils.RespondJSON(c, http.StatusOK, "地图创建成功", nil)
}

func QueryMapList(c *gin.Context) {
	type queryBodyParams struct {
		PageSize int    `json:"pageSize" binding:"gte=1,lte=100"`
		PageNo   int    `json:"pageNo" binding:"gte=1"`
		Author   string `json:"author"`
	}
	var body queryBodyParams
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.HandleError("参数绑定失败", err)
		return
	}

	var (
		mapList []model.Map
		total   int64
		query   = database.DB.Model(&model.Map{})
	)

	if body.Author != "" {
		query = query.Where("author LIKE ?", "%"+body.Author+"%")
	}

	query.Count(&total)

	err := query.
		Order("created_at DESC").
		Offset((body.PageNo - 1) * body.PageSize).
		Limit(body.PageSize).
		Find(&mapList).Error

	if err != nil {
		utils.HandleError("查询地图失败", err)
		return
	}

	utils.RespondJSON(c, http.StatusOK, "", gin.H{
		"total": total,
		"list":  mapList,
	})

}

func DeleteMap(c *gin.Context) {
	var mapBody model.Map
	err := c.ShouldBindJSON(&mapBody)
	if err != nil {
		utils.HandleError("参数绑定失败", err)
		utils.RespondJSON(c, http.StatusBadRequest, "参数绑定失败", nil)
		return
	}
	if err := database.DB.Where("id = ?", mapBody.ID).Delete(&model.Map{}).Error; err != nil {
		utils.HandleError("地图删除失败", err)
		utils.RespondJSON(c, http.StatusInternalServerError, "地图删除失败", nil)
		return
	}
	utils.RespondJSON(c, http.StatusOK, "地图删除成功", nil)
}

func UpdateMap(c *gin.Context) {
	var mapBody model.Map
	err := c.ShouldBindJSON(&mapBody)
	if err != nil {
		utils.HandleError("参数绑定失败", err)
		utils.RespondJSON(c, http.StatusBadRequest, "参数绑定失败", nil)
		return
	}
	log.Println(mapBody)
	if err := database.DB.Model(&mapBody).Where("id = ?", mapBody.ID).Omit("created_at").Select("*").Updates(&mapBody).Error; err != nil {
		utils.HandleError("地图更新失败", err)
		utils.RespondJSON(c, http.StatusInternalServerError, "地图更新失败", nil)
		return
	}
	utils.RespondJSON(c, http.StatusOK, "地图更新成功", nil)
}
