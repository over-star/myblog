package models

import "time"

type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
}

// User 用户模型
type User struct {
	BaseModel
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	LastLoginIp string `gorm:"column:last_login_ip;"`
	CommonTimestampsField
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"created_at,omitempty"`
}

func (User) TableName() string {
	return "user"
}
