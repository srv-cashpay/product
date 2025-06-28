package product

import (
	"github.com/labstack/echo/v4"
	m "github.com/srv-cashpay/middlewares/middlewares"
	dto "github.com/srv-cashpay/product/dto"

	r "github.com/srv-cashpay/product/repositories/product"
)

type ProductService interface {
	Get(context echo.Context, req *dto.Pagination) dto.Response
	Menu(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetByIdRequest) (*dto.ProductResponse, error)
}

type productService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewProductService(Repo r.DomainRepository, jwtS m.JWTService) ProductService {
	return &productService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
