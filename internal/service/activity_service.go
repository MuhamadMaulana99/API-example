package service

import (
	"golang-api/internal/domain"
	"golang-api/internal/repository"
)

func GetAllUsers() (
	[]domain.User,
	error,
) {

	return repository.
		GetAllUsers()
}

func UpdateUser(
	id uint,
	name string,
	email string,
	actorID uint,
) (domain.User, error) {

	user, err :=
		repository.UpdateUser(
			id,
			name,
			email,
		)

	if err != nil {
		return user, err
	}

	// audit log
	SaveActivity(
		actorID,
		"UPDATE_USER",
		"/api/users",
		"PUT",
		"",
		"Updated user",
	)

	return user, nil
}
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
