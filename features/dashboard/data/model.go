package data

import (
	bd "my-tourist-ticket/features/booking/data"
	"my-tourist-ticket/features/dashboard"
	td "my-tourist-ticket/features/tour/data"
)

func BookingModelToDashboard(mb bd.Booking) dashboard.Booking {
	return dashboard.Booking{
		ID:          mb.ID,
		UserID:      mb.UserID,
		TourID:      mb.TourID,
		PackageID:   mb.PackageID,
		VoucherID:   mb.VoucherID,
		PaymentType: mb.PaymentType,
		GrossAmount: mb.GrossAmount,
		Status:      mb.Status,
		VaNumber:    mb.VaNumber,
		Bank:        mb.Bank,
		PhoneNumber: mb.PhoneNumber,
		Greeting:    mb.Greeting,
		FullName:    mb.FullName,
		Email:       mb.Email,
		Quantity:    mb.Quantity,
		ExpiredAt:   mb.ExpiredAt,
		CreatedAt:   mb.CreatedAt,
	}
}

func TourModelToDashboard(tour td.Tour) dashboard.Tour {
	return dashboard.Tour{
		ID:          tour.ID,
		CityId:      tour.CityId,
		UserId:      tour.UserId,
		TourName:    tour.TourName,
		Description: tour.Description,
		Image:       tour.Image,
		Thumbnail:   tour.Thumbnail,
		Addres:      tour.Addres,
		Latitude:    tour.Latitude,
		Longitude:   tour.Longitude,
		CreatedAt:   tour.CreatedAt,
		UpdatedAt:   tour.UpdatedAt,
	}
}
