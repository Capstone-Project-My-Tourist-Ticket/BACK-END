package data

import (
	"errors"
	"my-tourist-ticket/features/booking"
	pd "my-tourist-ticket/features/package/data"
	vd "my-tourist-ticket/features/voucher/data"
	"my-tourist-ticket/utils/externalapi"

	"gorm.io/gorm"
)

type bookingQuery struct {
	db              *gorm.DB
	paymentMidtrans externalapi.MidtransInterface
}

func New(db *gorm.DB, mi externalapi.MidtransInterface) booking.BookingDataInterface {
	return &bookingQuery{
		db:              db,
		paymentMidtrans: mi,
	}
}

// InsertBooking implements booking.BookingDataInterface.
func (repo *bookingQuery) InsertBooking(userIdLogin int, inputBooking booking.Core) (*booking.Core, error) {

	var totalHargaKeseluruhan int
	var packageGorm pd.Package
	ts := repo.db.Where("tour_id = ? AND id = ?", inputBooking.TourID, inputBooking.PackageID).First(&packageGorm)
	if ts.Error != nil {
		return nil, ts.Error
	}

	if inputBooking.VoucherID != nil {
		var voucherGorm vd.Voucher
		ts := repo.db.Where("id = ?", inputBooking.VoucherID).First(&voucherGorm)
		if ts.Error != nil {
			return nil, ts.Error
		}
		totalHargaKeseluruhan = ((packageGorm.JumlahTiket * packageGorm.Price) * inputBooking.Quantity) - voucherGorm.DiscountValue
	} else {
		totalHargaKeseluruhan = (packageGorm.JumlahTiket * packageGorm.Price) * inputBooking.Quantity
	}

	inputBooking.GrossAmount = totalHargaKeseluruhan

	payment, errPay := repo.paymentMidtrans.NewBookingPayment(inputBooking)
	if errPay != nil {
		return nil, errPay
	}

	bookingPaymentModel := CoreToModelBooking(inputBooking)
	bookingPaymentModel.PaymentType = payment.PaymentType
	bookingPaymentModel.Status = payment.Status
	bookingPaymentModel.VaNumber = payment.VaNumber
	bookingPaymentModel.ExpiredAt = payment.ExpiredAt

	tx := repo.db.Create(&bookingPaymentModel)
	if tx.Error != nil {
		return nil, tx.Error
	}

	bookingCore := ModelToCoreBooking(bookingPaymentModel)

	return &bookingCore, nil
}

// InsertBookingReview implements booking.BookingDataInterface.
func (repo *bookingQuery) InsertBookingReview(inputReview booking.ReviewCore) error {
	dataGorm := CoreReviewToModelReview(inputReview)

	tx := repo.db.Create(&dataGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}
