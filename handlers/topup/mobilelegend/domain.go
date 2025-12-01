package mobilelegend

import (
	s "github.com/srv-cashpay/product/services/topup/mobilelegend"

	"github.com/labstack/echo/v4"
)

type MobileLegendHandler interface {
	TopUp(c echo.Context) error
}

type mobilelegendHandler struct {
	serviceMobileLegend s.MobileLegendService
}

func NewMobileLegendHandler(service s.MobileLegendService) MobileLegendHandler {
	return &mobilelegendHandler{
		serviceMobileLegend: service,
	}
}
