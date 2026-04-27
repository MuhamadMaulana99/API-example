package repository

import (
	"golang-api/config"
	"golang-api/internal/domain"
)

func CreateActivityLog(
	logData domain.ActivityLog,
) error {

	return config.DB.
		Create(
			&logData,
		).Error
}
