package handler

import (
	"my-tourist-ticket/app/middlewares"
	packages "my-tourist-ticket/features/package"
	"my-tourist-ticket/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PackageHandler struct {
	packageService packages.PackageServiceInterface
}

func New(ps packages.PackageServiceInterface) *PackageHandler {
	return &PackageHandler{
		packageService: ps,
	}
}

func (handler *PackageHandler) CreatePackage(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	tourID, err := strconv.Atoi(c.Param("tour_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error parsing tour id", nil))
	}

	newPackage := PackageRequest{}
	errBind := c.Bind(&newPackage)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data, data not valid", nil))
	}

	packageCore := RequestToCore(newPackage, uint(tourID))
	errCreate := handler.packageService.Create(newPackage.Benefits, packageCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}
