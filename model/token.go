package model

import "time"

type TokenSchema struct {
	Token     string    `gorm:"not null" form:"token" json:"token"`
	CreatedAT time.Time `gorm:"not null" form:"createdAT" json:"createdAT"`
	Active    bool      `gorm:"not null" form:"active" json:"active"`
	Used      bool      `gorm:"not null" form:"used" json:"used"`
	CreatedBy string    `gorm:"not null" form:"createdBy" json:"createdBy"`
}
