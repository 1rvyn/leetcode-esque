package models

import "time"

type Session struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Browser   string `json:"browser"`
	UserAgent string `json:"user_agent"`
	Cookie    string `json:"cookie" gorm:"unique"`
	Email     string `json:"email"`
	IP        string `json:"ip"`
}
