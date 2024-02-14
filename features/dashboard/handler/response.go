package handler

import "my-tourist-ticket/features/dashboard"

type DashboardResponse struct {
	TotalCustomer    int `json:"total_user"`
	TotalPengelola   int `json:"total_pengelola"`
	TotalTransaction int `json:"total_transaction"`
	TotalTour        int `json:"total_tour"`

	RecentBooking []BookingResponse `json:"recent_booking"`

	TopTours []TourResponse `json:"top_tours"`
}
type BookingResponse struct {
	ID          string `json:"booking_id"`
	UserID      uint   `json:"user_id"`
	TourID      uint   `json:"tour_id"`
	PackageID   uint   `json:"package_id"`
	VoucherID   *uint  `json:"voucher_id"`
	PaymentType string `json:"payment_type"`
	GrossAmount int    `json:"gross_amount"`
	Status      string `json:"status"`
	VaNumber    string `json:"va_number"`
	Bank        string `json:"bank"`
	PhoneNumber string `json:"phone_number"`
	Greeting    string `json:"greeting"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	Quantity    int    `json:"quantity"`
	ExpiredAt   string `json:"payment_expired"`
	CreatedAt   string `json:"created_at"`
}

type TourResponse struct {
	ID          uint    `json:"id"`
	CityId      uint    `json:"city_id"`
	UserId      uint    `json:"user_id"`
	TourName    string  `json:"tour_name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Thumbnail   string  `json:"thumbnail"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func CoreToBookingResponseList(bookings []dashboard.Booking) []BookingResponse {
	var responseList []BookingResponse
	for _, booking := range bookings {
		responseList = append(responseList, BookingResponse{
			ID:          booking.ID,
			UserID:      booking.UserID,
			TourID:      booking.TourID,
			PackageID:   booking.PackageID,
			VoucherID:   booking.VoucherID,
			PaymentType: booking.PaymentType,
			GrossAmount: booking.GrossAmount,
			Status:      booking.Status,
			VaNumber:    booking.VaNumber,
			Bank:        booking.Bank,
			PhoneNumber: booking.PhoneNumber,
			Greeting:    booking.Greeting,
			FullName:    booking.FullName,
			Email:       booking.Email,
			Quantity:    booking.Quantity,
			ExpiredAt:   booking.ExpiredAt.Format("2006-01-02 15:04:05"),
			CreatedAt:   booking.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return responseList
}

func CoreToTourResponseList(tours []dashboard.Tour) []TourResponse {
	var responseList []TourResponse
	for _, tour := range tours {
		responseList = append(responseList, TourResponse{
			ID:          tour.ID,
			CityId:      tour.CityId,
			UserId:      tour.UserId,
			TourName:    tour.TourName,
			Description: tour.Description,
			Image:       tour.Image,
			Thumbnail:   tour.Thumbnail,
			Address:     tour.Addres,
			Latitude:    tour.Latitude,
			Longitude:   tour.Longitude,
			CreatedAt:   tour.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   tour.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return responseList
}

func CoreToDashboardResponse(core *dashboard.Dashboard) DashboardResponse {
	return DashboardResponse{
		TotalCustomer:    core.TotalCustomer,
		TotalPengelola:   core.TotalPengelola,
		TotalTransaction: core.TotalTransaction,
		TotalTour:        core.TotalTour,
		RecentBooking:    CoreToBookingResponseList(core.RecentBooking),
		TopTours:         CoreToTourResponseList(core.TopTours),
	}
}
