package models

import "time"

type Submission struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time
	Code       string `json:"code" gorm:"unique"`
	Cookie     string `json:"cookie"`
	Email      string `json:"email"`
	IP         string `json:"ip"`
	Successout string `json:"successout"`
	Errorout   string `json:"errorout"`
	//MetaData string `json:"meta_data"`
}
