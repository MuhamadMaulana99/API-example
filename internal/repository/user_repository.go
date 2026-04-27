package repository

import (
	"golang-api/config"
	"golang-api/internal/domain"
)

func CreateUser(
	user domain.User,
) (domain.User, error) {

	err :=
		config.DB.
			Create(&user).
			Error

	return user, err
}

func GetAllUsers() (
	[]domain.User,
	error,
) {

	var users []domain.User

	err :=
		config.DB.
			Find(
				&users,
			).Error

	return users, err
}

func UpdateUser(
	id uint,
	name string,
	email string,
) (domain.User, error) {

	var user domain.User

	err :=
		config.DB.
			First(
				&user,
				id,
			).Error

	if err != nil {
		return user, err
	}

	user.Name = name
	user.Email = email

	err =
		config.DB.
			Save(
				&user,
			).Error

	return user, err
}

func FindByEmail(
	email string,
) (domain.User, error) {

	var user domain.User

	err :=
		config.DB.
			Where(
				"email=?",
				email,
			).
			First(
				&user,
			).Error

	return user, err
}

func DeleteUser(
	id uint,
) error {

	return config.DB.
		Delete(
			&domain.User{},
			id,
		).Error
}
