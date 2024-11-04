package product

import (
	"sync"

	dto "github.com/srv-cashpay/product/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.ProductRequest) (dto.ProductRequest, error)
	Get(req *dto.Pagination) (RepositoryResult, int)
	GetById(req dto.GetByIdRequest) (*dto.ProductResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Update(req dto.ProductUpdateRequest) (dto.ProductUpdateResponse, error)
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
