package model

import "time"

type Map struct {
	ID        uint      `json:"id";gorm:"primaryKey"`
	Author    string    `json:"author"gorm:"type:varchar(100);not null;comment:地图作者"`
	Info      string    `json:"info"gorm:"type:text;comment:地图描述"`
	Remark    string    `json:"remark"gorm:"type:varchar(255);comment:备注"`
	IsActive  bool      `json:"isActive"gorm:"default:true;comment:是否启用"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:'创建时间'"`                             // 确保格式
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'更新时间'"` // 自动更新

}
