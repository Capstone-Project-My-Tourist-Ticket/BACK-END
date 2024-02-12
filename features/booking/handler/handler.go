package handler

import (
	"my-tourist-ticket/app/middlewares"
	"my-tourist-ticket/features/booking"
	"my-tourist-ticket/utils/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookingHandler struct {
	bookingService booking.BookingServiceInterface
}

func New(bs booking.BookingServiceInterface) *BookingHandler {
	return &BookingHandler{
		bookingService: bs,
	}
}

func (handler *BookingHandler) CreateBooking(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	newBooking := BookingRequest{}
	errBind := c.Bind(&newBooking)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data booking not valid", nil))
	}

	bookingCore := RequestToCoreBooking(newBooking, uint(userIdLogin))
	if newBooking.VoucherID == 0 {
		bookingCore.VoucherID = nil
	}
	payment, errInsert := handler.bookingService.CreateBooking(userIdLogin, bookingCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert booking", nil))
	}

	result := CoreToResponseBooking(payment)

	return c.JSON(http.StatusOK, responses.WebResponse("success insert booking", result))
}

func (handler *BookingHandler) WebhoocksNotification(c echo.Context) error {
	var reqNotif = WebhoocksRequest{}
	errBind := c.Bind(&reqNotif)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	bookingCore := WebhoocksRequestToCore(reqNotif)
	err := handler.bookingService.WebhoocksService(bookingCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error Notif "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("transaction success", nil))
}
