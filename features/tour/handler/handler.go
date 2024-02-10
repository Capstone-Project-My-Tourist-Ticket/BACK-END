package handler

import (
	"my-tourist-ticket/app/middlewares"
	"my-tourist-ticket/features/tour"
	"my-tourist-ticket/utils/responses"
	"net/http"
	"strconv"

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

func (handler *TourHandler) UpdateTour(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	userRole, err := handler.tourService.GetUserRoleById(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Internal Server Error", nil))
	}
	if userRole != "pengelola" {
		return c.JSON(http.StatusForbidden, responses.WebResponse("Forbidden - User is not an pengelola", nil))
	}

	tourID, err := strconv.Atoi(c.Param("tour_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("Invalid city ID", nil))
	}

	var tourReq TourRequest
	if err := c.Bind(&tourReq); err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("Error binding data. Data not valid", nil))
	}

	// Ubah request menjadi core model
	tourCore := RequestToCore(tourReq)

	// Dapatkan file gambar dan thumbnail dari form
	_, imageHeader, err := c.Request().FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("Error retrieving the image file", nil))
	}

	_, thumbnailHeader, err := c.Request().FormFile("thumbnail")
	if err != nil && err != http.ErrMissingFile {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("Error retrieving the thumbnail file", nil))
	}

	err = handler.tourService.Update(tourID, tourCore, imageHeader, thumbnailHeader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Error updating city", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Tour updated successfully", nil))
}

func (handler *TourHandler) GetTourById(c echo.Context) error {
	tourID, err := strconv.Atoi(c.Param("tour_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("Invalid Tour Id", nil))
	}

	tourData, err := handler.tourService.SelectTourById(tourID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Error retrieving tour data", nil))
	}

	tourResponse := ModelToResponse(tourData)

	return c.JSON(http.StatusOK, responses.WebResponse("Tour data retrieved successfully", tourResponse))
}

func (handler *TourHandler) DeleteTour(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	userRole, err := handler.tourService.GetUserRoleById(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Internal Server Error", nil))
	}
	if userRole != "pengelola" {
		return c.JSON(http.StatusForbidden, responses.WebResponse("Forbidden - User is not an pengelola", nil))
	}

	id := c.Param("tour_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error. id should be number", nil))
	}

	err = handler.tourService.Delete(idParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error delete tour. delete failed "+err.Error(), nil))
	}
	return c.JSON(http.StatusOK, responses.WebResponse("success delete tour", nil))
}
