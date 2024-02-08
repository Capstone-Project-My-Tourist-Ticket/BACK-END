package service

import (
	"errors"
	"my-tourist-ticket/app/middlewares"
	"my-tourist-ticket/features/user"
	"my-tourist-ticket/utils/encrypts"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData    user.UserDataInterface
	hashService encrypts.HashInterface
	validate    *validator.Validate
}

// dependency injection
func New(repo user.UserDataInterface, hash encrypts.HashInterface) user.UserServiceInterface {
	return &userService{
		userData:    repo,
		hashService: hash,
		validate:    validator.New(),
	}
}

// Create implements user.UserServiceInterface.
func (service *userService) Create(input user.Core) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	if input.Password != "" {
		hashedPass, errHash := service.hashService.HashPassword(input.Password)
		if errHash != nil {
			return errors.New("Error hash password.")
		}
		input.Password = hashedPass
	}

	if input.Role == "pengelola" {
		input.Status = "pending"
	} else {
		input.Status = "approved"
	}

	err := service.userData.Insert(input)
	return err
}

// Login implements user.UserServiceInterface.
func (service *userService) Login(email string, password string) (data *user.Core, token string, err error) {
	if email == "" && password == "" {
		return nil, "", errors.New("email dan password wajib diisi.")
	}
	if email == "" {
		return nil, "", errors.New("email wajib diisi.")
	}
	if password == "" {
		return nil, "", errors.New("password wajib diisi.")
	}

	data, err = service.userData.Login(email, password)
	if err != nil {
		return nil, "", err
	}
	isValid := service.hashService.CheckPasswordHash(data.Password, password)
	if !isValid {
		return nil, "", errors.New("password tidak sesuai.")
	}

	token, errJwt := middlewares.CreateToken(int(data.ID))
	if errJwt != nil {
		return nil, "", errJwt
	}

	return data, token, err
}

// GetById implements user.UserServiceInterface.
func (service *userService) GetById(userIdLogin int) (*user.Core, error) {
	result, err := service.userData.SelectById(userIdLogin)
	return result, err
}

// Update implements user.UserServiceInterface.
func (service *userService) Update(userIdLogin int, input user.Core) error {
	if userIdLogin <= 0 {
		return errors.New("invalid id.")
	}

	if input.Password != "" {
		hashedPass, errHash := service.hashService.HashPassword(input.Password)
		if errHash != nil {
			return errors.New("Error hash password.")
		}
		input.Password = hashedPass
	}

	err := service.userData.Update(userIdLogin, input)
	return err
}
