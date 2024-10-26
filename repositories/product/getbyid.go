package product

import (
	dto "github.com/srv-cashpay/product/dto"
	"github.com/srv-cashpay/product/entity"
)

func (b *productRepository) GetById(req dto.GetByIdRequest) (*dto.ProductResponse, error) {
	tr := entity.Product{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.ProductResponse{
		ProductName: tr.ProductName,
	}

	return response, nil
}
