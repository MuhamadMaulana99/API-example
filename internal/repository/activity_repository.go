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

func GetActivityLogs(
	page int,
	limit int,
	search string,
) (
	[]domain.ActivityLog,
	int64,
	error,
) {

	var logs []domain.ActivityLog
	var total int64

	offset := (page - 1) * limit

	query := config.DB.Model(
		&domain.ActivityLog{},
	)

	if search != "" {

		query = query.Where(
			"user_name ILIKE ?",
			"%"+search+"%",
		)
	}

	query.Count(&total)

	err := query.
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Find(&logs).Error

	return logs, total, err
}
