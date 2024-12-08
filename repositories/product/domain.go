package product

import (
	"sync"

	dto "github.com/srv-cashpay/product/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Get(req *dto.Pagination) (RepositoryResult, int)
	GetById(req dto.GetByIdRequest) (*dto.ProductResponse, error)
}

type productRepository struct {
	DB *gorm.DB
	mu sync.Mutex
}

func NewProductRepository(DB *gorm.DB) DomainRepository {
	return &productRepository{
		DB: DB,
	}
}
