package service

import (
	"errors"

	"golang-api/config"
	"golang-api/internal/domain"
	"golang-api/internal/dto"
	"golang-api/internal/repository"
	"golang-api/pkg/utils"
)

func Register(
	req dto.RegisterDTO,
) (domain.User, error) {

	hash, _ :=
		utils.HashPassword(
			req.Password,
		)

	user := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hash,
		Role:     req.Role,
	}

	createdUser, err :=
		repository.CreateUser(
			user,
		)

	if err != nil {
		return createdUser, err
	}

	// ==================
	// AUDIT LOG REGISTER
	// ==================
	SaveActivity(
		createdUser.ID,
		"CREATE_USER",
		"/api/register",
		"POST",
		"",
		"User registered",
	)

	return createdUser, nil
}

func GetUsersPaginated(
	page int,
	limit int,
	search string,
) (
	[]domain.User,
	int64,
	error,
) {

	return repository.
		GetUsersPaginated(
			page,
			limit,
			search,
		)
}

func DeleteUser(
	targetID uint,
	actorID uint,
) error {

	err :=
		repository.DeleteUser(
			targetID,
		)

	if err != nil {
		return err
	}

	// =================
	// AUDIT DELETE LOG
	// =================
	SaveActivity(
		actorID,
		"DELETE_USER",
		"/api/users",
		"DELETE",
		"",
		"Deleted user",
	)

	return nil
}

func Login(
	req dto.LoginDTO,
) (string, error) {

	user, err :=
		repository.FindByEmail(
			req.Email,
		)

	if err != nil {
		return "", errors.New(
			"user not found",
		)
	}

	if !utils.CheckPassword(
		user.Password,
		req.Password,
	) {
		return "", errors.New(
			"wrong password",
		)
	}

	SaveActivity(
		user.ID,
		"LOGIN",
		"/api/login",
		"POST",
		"",
		"User Login",
	)

	token, err :=
		config.GenerateToken(
			user,
		)

	return token, nil
}
