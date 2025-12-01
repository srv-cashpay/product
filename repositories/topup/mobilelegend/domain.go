package mobilelegend

import (
	"sync"

	dto "github.com/srv-cashpay/product/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	MobileLegend(req dto.MobileLegendRequest) (dto.MobileLegendResponse, error)
}

type topupRepository struct {
	DB *gorm.DB
	mu sync.Mutex
}

func NewTopUpRepository(DB *gorm.DB) DomainRepository {
	return &topupRepository{
		DB: DB,
	}
}
