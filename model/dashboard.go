package model

type DashBoard struct {
	UsedToken    []TokenSchema `gorm:"not null" form:"used_token" json:"used_token"`
	PendingToken []TokenSchema `gorm:"not null" form:"pending_token" json:"pending_token"`
	ExpiredToken []TokenSchema `gorm:"not null" form:"expired_token" json:"expired_token"`
	RevokedToken []TokenSchema `gorm:"not null" form:"revoked_token" json:"revoked_token"`
}
