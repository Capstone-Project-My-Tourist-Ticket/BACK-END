package handler

import (
	"my-tourist-ticket/features/voucher"
	"my-tourist-ticket/utils/responses"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type VoucherHandler struct {
	voucherService voucher.VoucherServiceInterface
}

func New(vs voucher.VoucherServiceInterface) *VoucherHandler {
	return &VoucherHandler{
		voucherService: vs,
	}
}

func (handler *VoucherHandler) CreateVoucher(c echo.Context) error {
	newVoucher := VoucherRequest{}
	errBind := c.Bind(&newVoucher)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data, data not valid", nil))
	}

	voucherCore := RequestToCore(newVoucher)
	errInsert := handler.voucherService.Create(voucherCore)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "Error 1062 (23000): Duplicate entry") {
			return c.JSON(http.StatusBadRequest, responses.WebResponse("error insert data. "+errInsert.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data. "+errInsert.Error(), nil))
		}
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}

func (handler *VoucherHandler) GetAllVoucher(c echo.Context) error {
	vouchers, err := handler.voucherService.SelectAllVoucher()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error get data", nil))
	}

	vouchersResponses := CoreToResponseListGetAllVoucher(vouchers)

	return c.JSON(http.StatusOK, responses.WebResponse("success get data", vouchersResponses))
}
