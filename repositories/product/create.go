package product

import (
	dto "github.com/srv-cashpay/product/dto"
	"github.com/srv-cashpay/product/entity"
)

func (r *productRepository) Create(req dto.ProductRequest) (dto.ProductRequest, error) {

	create := entity.Product{
		ID:           req.ID,
		ProductName:  req.ProductName,
		Stock:        req.Stock,
		MinimalStock: req.MinimalStock,
		Price:        req.Price,
		UserID:       req.UserID,
		CreatedBy:    req.CreatedBy,
	}

	if err := r.DB.Save(&create).Error; err != nil {
		return dto.ProductRequest{}, err
	}

	response := dto.ProductRequest{
		ID:           req.ID,
		UserID:       req.UserID,
		ProductName:  create.ProductName,
		Stock:        create.Stock,
		MinimalStock: create.MinimalStock,
		Price:        create.Price,
		CreatedBy:    req.CreatedBy,
	}

	return response, nil

}
