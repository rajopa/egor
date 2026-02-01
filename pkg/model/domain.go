package domain

import "time"

type Target struct {
	Id        int       `json:"id" db:"id"`
	UserId    int       `json:"user_id" db:"user_id"`
	Title     string    `json:"title" db:"title" binding:"required"`
	URL       string    `json:"URL" db:"url" binding:"required"`
	Interval  int       `json:"interval" db:"interval"`
	Status    bool      `json:"status" db:"status" `
	LastCheck time.Time `json:"checklast" db:"last_check"`
}

type UserTarget struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

type UpdateTargetInput struct {
	Title    *string `json:"title"`
	URL      *string `json:"URL" `
	Interval *int    `json:"interval"`
}
