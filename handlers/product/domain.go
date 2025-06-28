package product

import (
	s "github.com/srv-cashpay/product/services/product"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Get(c echo.Context) error
	GetById(c echo.Context) error
	Menu(c echo.Context) error
}

type domainHandler struct {
	serviceProduct s.ProductService
}

func NewProductHandler(service s.ProductService) DomainHandler {
	return &domainHandler{
		serviceProduct: service,
	}
}
