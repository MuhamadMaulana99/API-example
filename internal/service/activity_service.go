package service

import (
	"fmt"
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
	ip string,
) (domain.User, error) {

	user, err := repository.UpdateUser(id, name, email)

	if err != nil {
		return user, err
	}

	// 🔥 audit log
	_ = SaveActivity(
		actorID,
		"UPDATE_USER",
		"/api/users/"+fmt.Sprint(id),
		"PUT",
		ip,
		"Updated user",
		id,
	)

	return user, nil
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

	return repository.GetActivityLogs(
		page,
		limit,
		search,
	)
}
func SaveActivity(
	userID uint,
	action string,
	endpoint string,
	method string,
	ip string,
	description string,
	targetID uint,
) error {

	activity := domain.ActivityLog{
		UserID:      userID,
		Action:      action,
		Endpoint:    endpoint,
		Method:      method,
		IPAddress:   ip,
		Description: description,
		TargetID:    targetID,
	}

	return repository.CreateActivityLog(activity)
}
