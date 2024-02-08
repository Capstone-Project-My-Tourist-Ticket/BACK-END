package handler

import (
	"my-tourist-ticket/app/middlewares"
	"my-tourist-ticket/features/user"
	"my-tourist-ticket/utils/cloudinary"
	"my-tourist-ticket/utils/responses"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
	cld         cloudinary.CloudinaryUploaderInterface
}

func New(service user.UserServiceInterface, cloudinaryUploader cloudinary.CloudinaryUploaderInterface) *UserHandler {
	return &UserHandler{
		userService: service,
		cld:         cloudinaryUploader,
	}
}

func (handler *UserHandler) RegisterUser(c echo.Context) error {
	newUser := UserRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data, data not valid", nil))
	}

	userCore := RequestToCore(newUser)
	errInsert := handler.userService.Create(userCore)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "Error 1062 (23000): Duplicate entry") {
			return c.JSON(http.StatusBadRequest, responses.WebResponse("error insert data. "+errInsert.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data. "+errInsert.Error(), nil))
		}
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}

func (handler *UserHandler) Login(c echo.Context) error {
	var reqData = LoginRequest{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data, data not valid", nil))
	}
	result, token, err := handler.userService.Login(reqData.Email, reqData.Password)
	if err != nil {
		if strings.Contains(err.Error(), "email wajib diisi.") {
			return c.JSON(http.StatusUnauthorized, responses.WebResponse("error login. "+err.Error(), nil))
		} else if strings.Contains(err.Error(), "password wajib diisi.") {
			return c.JSON(http.StatusUnauthorized, responses.WebResponse("error login. "+err.Error(), nil))
		} else if strings.Contains(err.Error(), "password tidak sesuai.") {
			return c.JSON(http.StatusUnauthorized, responses.WebResponse("error login. "+err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, responses.WebResponse("error login. "+err.Error(), nil))
		}
	}
	var responseData = CoreToResponseLogin(result, token)
	return c.JSON(http.StatusOK, responses.WebResponse("success login", responseData))
}

func (handler *UserHandler) GetUser(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)

	result, errSelect := handler.userService.GetById(userIdLogin)
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errSelect.Error(), nil))
	}

	if result.NoKtp == "" {
		var userResult = CoreToResponseUser(result)
		return c.JSON(http.StatusOK, responses.WebResponse("success read data", userResult))
	} else {
		var userResult = CoreToResponsePengelola(result)
		return c.JSON(http.StatusOK, responses.WebResponse("success read data", userResult))
	}
}
