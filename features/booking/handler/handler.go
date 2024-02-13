package handler

import (
	"my-tourist-ticket/app/middlewares"
	"my-tourist-ticket/features/booking"
	"my-tourist-ticket/utils/responses"
	"net/http"
	"strconv"
	"strings"

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

func (handler *BookingHandler) CancleBookingById(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	bookingId := c.Param("id")

	updateBookingStatus := CancleBookingRequest{}
	errBind := c.Bind(&updateBookingStatus)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	bookingCore := CancleRequestToCoreBooking(updateBookingStatus)
	errCancle := handler.bookingService.CancleBooking(userIdLogin, bookingId, bookingCore)
	if errCancle != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error cancle order", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success cancle order", nil))
}

func (handler *BookingHandler) CreateBookingReview(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	bookingID := c.Param("booking_id")
	if bookingID == "" {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("Invalid booking ID", nil))
	}

	var reviewReq ReviewRequest
	if err := c.Bind(&reviewReq); err != nil {
		return err
	}
	reviewCore := RequestToCoreBookingReview(reviewReq)
	reviewCore.UserID = uint(userIdLogin)
	reviewCore.BookingID = bookingID
	// reviewCore.TextReview = reviewReq.TextReview

	err := handler.bookingService.CreateBookingReview(reviewCore)
	if err != nil {
		if strings.Contains(err.Error(), "is required") {
			return c.JSON(http.StatusBadRequest, responses.WebResponse(err.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error creating review", nil))
	}

	return c.JSON(http.StatusCreated, responses.WebResponse("review created successfully", nil))
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

func (handler *BookingHandler) GetBookingUser(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	results, errSelect := handler.bookingService.GetBookingUser(userIdLogin)
	if errSelect != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errSelect.Error(), nil))
	}

	var bookingResult = CoreToResponseListUser(results)
	return c.JSON(http.StatusOK, responses.WebResponse("success read data.", bookingResult))
}

func (handler *BookingHandler) GetAllBooking(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)

	userRole, err := handler.bookingService.GetUserRoleById(userIdLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Internal Server Error", nil))
	}
	if userRole != "admin" {
		return c.JSON(http.StatusForbidden, responses.WebResponse("Forbidden - User is not an admin", nil))
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	bookings, totalPage, err := handler.bookingService.SelectAllBooking(page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error get data", nil))
	}

	bookingResponses := CoreToResponseList(bookings)

	return c.JSON(http.StatusOK, responses.WebResponsePagination("success get data", bookingResponses, totalPage))
}
