package domain

import "gorm.io/gorm"

type ActivityLog struct {
	gorm.Model

	UserID uint

	Action      string
	Endpoint    string
	Method      string
	IPAddress   string
	Description string
}
