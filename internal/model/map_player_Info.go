package model

import "time"

type MapPlayerSave struct {
	ID        uint      `gorm:"primaryKey" json:"id"`                                              // 主键
	MapID     uint      `gorm:"not null;index:idx_map_player,unique;comment:地图ID" json:"mapId"`    // 地图 ID
	PlayerID  uint      `gorm:"not null;index:idx_map_player,unique;comment:玩家ID" json:"playerId"` // 玩家 ID
	SaveData  string    `gorm:"type:json;comment:玩家在地图上的存档数据" json:"saveData"`                     // 存档内容
	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:更新时间" json:"updatedAt"`                      // 更新时间

	Map    Map    `gorm:"foreignKey:MapID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`    // 外键
	Player Player `gorm:"foreignKey:PlayerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"` // 外键
}
