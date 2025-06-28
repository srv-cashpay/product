package product

import (
	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/product/helpers"
	res "github.com/srv-cashpay/util/s/response"
)

func (b *domainHandler) Menu(c echo.Context) error {
	paginationDTO := helpers.GeneratePaginationRequest(c)

	merchantID := c.QueryParam("merchant_id")
	if merchantID == "" {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	paginationDTO.MerchantID = merchantID

	if err := c.Bind(&paginationDTO); err != nil {
		return c.JSON(400, "Invalid request")
	}

	users := b.serviceProduct.Menu(c, paginationDTO)

	return c.JSON(200, users)
}
