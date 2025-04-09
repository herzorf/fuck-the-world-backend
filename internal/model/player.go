package model

type Player struct {
	ID       uint   ` json:"id" gorm:"primaryKey"`                                          // 主键
	Username string `json:"username"gorm:"type:varchar(100);unique;not null;comment:玩家用户名"` // 用户名
}
