package handler

import (
	"my-tourist-ticket/app/middlewares"
	"my-tourist-ticket/features/tour"
	"my-tourist-ticket/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TourHandler struct {
	tourService tour.TourServiceInterface
}

func NewTour(service tour.TourServiceInterface) *TourHandler {
	return &TourHandler{
		tourService: service,
	}
}

func (handler *TourHandler) CreateTour(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	userRole, err := handler.tourService.GetUserRoleById(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Internal Server Error", nil))
	}
	if userRole != "pengelola" {
		return c.JSON(http.StatusForbidden, responses.WebResponse("Forbidden - User is not a pengelola", nil))
	}

	var tourReq TourRequest

	if err := c.Bind(&tourReq); err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	// Mendapatkan file gambar dan thumbnail dari formulir
	_, imageHeader, err := c.Request().FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error retrieving the image file", nil))
	}

	_, thumbnailHeader, err := c.Request().FormFile("thumbnail")
	if err != nil && err != http.ErrMissingFile {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error retrieving the thumbnail file", nil))
	}

	tourCore := RequestToCore(tourReq)
	tourCore.UserId = uint(userId)
	// Memanggil tourService.Insert dengan argumen yang sesuai, termasuk ID pengguna
	err = handler.tourService.Insert(uint(userId), tourCore, imageHeader, thumbnailHeader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error creating tour/ city does not exist", nil))
	}

	return c.JSON(http.StatusCreated, responses.WebResponse("tour created successfully", nil))
}
