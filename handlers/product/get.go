package product

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/auth/dto/auth"
	res "github.com/srv-cashpay/util/s/response"
)

func (b *domainHandler) Get(c echo.Context) error {
	var req dto.SignupRequest

	userId, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	req.ID = userId

	users, err := b.serviceProduct.Get()
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)

	}
	return res.SuccessResponse(users).Send(c)

}
