package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"unique;not null" json:"username"`
	Nickname     string    `gorm:"size:50" json:"nickname"` // 昵称（展示用）
	Email        string    `gorm:"unique;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Role         int       `gorm:"default:0" json:"role"` // 0=普通用户,1=管理员
	AvatarURL    string    `json:"avatar_url"`
	Bio          string    `json:"bio"`
	Status       int       `gorm:"default:1" json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	LastLogin    time.Time `json:"last_login"`
}
