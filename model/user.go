package model

import "time"

type Users struct {
	EmailID   string `gorm:"unique"  form:"email_id" json:"email_id"`
	Password  string `gorm:"not null" form:"password" json:"password"`
	CreatedAT time.Time `gorm:"not null" form:"created_at" json:"date"`
}
