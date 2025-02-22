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
		ID:           tr.ID,
		Barcode:      tr.Barcode,
		UserID:       tr.UserID,
		MerchantID:   tr.MerchantID,
		MerkID:       tr.MerkID,
		CategoryID:   tr.CategoryID,
		ProductName:  tr.ProductName,
		Description:  tr.Description,
		Stock:        tr.Stock,
		MinimalStock: tr.MinimalStock,
		Price:        tr.Price,
		Status:       tr.Status,
		CreatedBy:    tr.CreatedBy,
		CreatedAt:    dto.Timestamp(tr.CreatedAt),
	}

	return response, nil
}
