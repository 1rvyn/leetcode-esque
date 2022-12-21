package models

import "time"

type Session struct {
	ID            uint `json:"id"`
	LoginCreation time.Time
	Browser       string `json:"browser"`
	UserAgent     string `json:"user_agent"`
	Cookie        string `json:"cookie"`
	Email         string `json:"email"`
	IP            string `json:"ip"`
}
