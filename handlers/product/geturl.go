package product

import (
	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/product/dto"
	res "github.com/srv-cashpay/util/s/response"
)

func (h *domainHandler) GetUrl(c echo.Context) error {
	var req dto.ProductRequest
	var resp dto.ProductResponse

	userid, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	merchantId, ok := c.Get("MerchantId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	req.MerchantID = merchantId
	req.UserID = userid

	resp, err := h.serviceProduct.GetUrl(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)
}
