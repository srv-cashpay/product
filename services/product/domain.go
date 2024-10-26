package product

import (
	m "github.com/srv-cashpay/middlewares/middlewares"
	dto "github.com/srv-cashpay/product/dto"
	entity "github.com/srv-cashpay/product/entity"

	r "github.com/srv-cashpay/product/repositories/product"
)

type ProductService interface {
	Create(req dto.ProductRequest) (dto.ProductResponse, error)
	Get() ([]entity.Product, error)
	GetById(req dto.GetByIdRequest) (*dto.ProductResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Update(req dto.ProductUpdateRequest) (dto.ProductUpdateResponse, error)
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
