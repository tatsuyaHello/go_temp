package model

import "time"

// RequestUser はログイン/サインアップユーザ用のモデル
type RequestUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// User はユーザ用のモデル
type User struct {
	ID        uint64    `json:"id"`
	Email     *string   `json:"email"`
	Password  *string   `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
