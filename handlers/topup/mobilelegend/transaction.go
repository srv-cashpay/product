package mobilelegend

import (
	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/product/dto"
	res "github.com/srv-cashpay/util/s/response"
)

func (h *mobilelegendHandler) TopUp(c echo.Context) error {
	var req dto.MobileLegendRequest

	if err := c.Bind(&req); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	resp, err := h.serviceMobileLegend.MobileLegend(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	return c.JSON(200, resp)
}
