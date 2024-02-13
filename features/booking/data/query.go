package data

import (
	"errors"
	"my-tourist-ticket/features/booking"
	pd "my-tourist-ticket/features/package/data"
	"my-tourist-ticket/features/user"
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

// GetUserRoleById implements booking.BookingDataInterface.
func (repo *bookingQuery) GetUserRoleById(userId int) (string, error) {
	var user user.Core
	if err := repo.db.Table("users").Where("id = ?", userId).First(&user).Error; err != nil {
		return "", err
	}

	return user.Role, nil
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

func (repo *bookingQuery) CancleBooking(userIdLogin int, bookingId string, bookingCore booking.Core) error {
	if bookingCore.Status == "cancelled" {
		repo.paymentMidtrans.CancelBookingPayment(bookingId)
	}

	dataGorm := CoreToModelBookingCancle(bookingCore)
	tx := repo.db.Model(&Booking{}).Where("id = ? AND user_id = ?", bookingId, userIdLogin).Updates(dataGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
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

// Update implements booking.BookingDataInterface.
func (repo *bookingQuery) WebhoocksData(reqNotif booking.Core) error {
	dataGorm := CoreToModel(reqNotif)
	tx := repo.db.Model(&Booking{}).Where("id = ?", reqNotif.ID).Updates(dataGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}

// SelectAllBooking implements booking.BookingDataInterface.
func (repo *bookingQuery) SelectAllBooking(page int, limit int) ([]booking.Core, int, error) {
	var bookingGorm []Booking
	query := repo.db.Order("created_at desc")

	var totalData int64
	err := query.Model(&Booking{}).Count(&totalData).Error
	if err != nil {
		return nil, 0, err
	}

	totalPage := int((totalData + int64(limit) - 1) / int64(limit))

	// Retrieve booking data with associated user, tour, and package
	err = query.Limit(limit).Offset((page - 1) * limit).Preload("Package").Preload("Tour").Find(&bookingGorm).Error
	if err != nil {
		return nil, 0, err
	}

	// Convert booking data to booking.Core
	bookingCore, err := ModelToCoreList(bookingGorm)
	if err != nil {
		return nil, 0, err
	}

	// for i := range bookingCore {
	// 	bookingCore[i].Package.ID = bookingGorm[i].PackageID
	// }

	return bookingCore, totalPage, nil
}
