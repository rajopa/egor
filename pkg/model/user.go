package domain

import "time"

type User struct {
	Id        int        `json:"id" db:"id"`
	Username  string     `json:"username" db:"username" binding:"required"`
	Password  string     `json:"password" db:"password_hash" binding:"required"`
	LastLogin *time.Time `json:"last_login,omitempty" db:"last_login"`
	IsDeleted bool       `json:"-" db:"is_deleted"`
}
