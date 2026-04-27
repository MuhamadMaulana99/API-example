package service

import (
	"golang-api/internal/domain"
	"golang-api/internal/repository"
)

func SaveActivity(
	userID uint,
	action string,
	endpoint string,
	method string,
	ip string,
	description string,
) error {

	activity := domain.ActivityLog{
		UserID: userID,

		Action: action,

		Endpoint: endpoint,

		Method: method,

		IPAddress: ip,

		Description: description,
	}

	return repository.
		CreateActivityLog(
			activity,
		)
}
