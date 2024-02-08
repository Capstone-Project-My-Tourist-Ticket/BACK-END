package handler

import "my-tourist-ticket/features/user"

type UserResponse struct {
	FullName    string `json:"full_name" form:"full_name"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Email       string `json:"email" form:"email"`
	Image       string `json:"image" form:"image"`
	Role        string `json:"role" form:"role"`
	Status      string `json:"status" form:"status"`
}

type UserResponsePengelola struct {
	FullName    string `json:"full_name" form:"full_name"`
	NoKtp       string `json:"no_ktp" form:"no_ktp"`
	Address     string `json:"address" form:"address"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Email       string `json:"email" form:"email"`
	Image       string `json:"image" form:"image"`
	Role        string `json:"role" form:"role"`
	Status      string `json:"status" form:"status"`
}

type UserResponseLogin struct {
	FullName string `json:"full_name" form:"full_name"`
	Role     string `json:"role" form:"role"`
	Token    string `json:"token" form:"token"`
}

func CoreToResponseUser(data *user.Core) UserResponse {
	var result = UserResponse{
		FullName:    data.FullName,
		PhoneNumber: data.PhoneNumber,
		Email:       data.Email,
		Image:       data.Image,
		Role:        data.Role,
		Status:      data.Status,
	}
	return result
}

func CoreToResponsePengelola(data *user.Core) UserResponsePengelola {
	var result = UserResponsePengelola{
		FullName:    data.FullName,
		NoKtp:       data.NoKtp,
		Address:     data.Address,
		PhoneNumber: data.PhoneNumber,
		Email:       data.Email,
		Image:       data.Image,
		Role:        data.Role,
		Status:      data.Status,
	}
	return result
}

func CoreToResponseLogin(data *user.Core, token string) UserResponseLogin {
	var result = UserResponseLogin{
		FullName: data.FullName,
		Role:     data.Role,
		Token:    token,
	}
	return result
}
