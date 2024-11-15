package product

import (
	dto "github.com/srv-cashpay/product/dto"
	"github.com/srv-cashpay/product/entity"
)

func (r *productRepository) Create(req dto.ProductRequest) (dto.ProductResponse, error) {

	create := entity.Product{
		ID:           req.ID,
		ProductName:  req.ProductName,
		Stock:        req.Stock,
		MinimalStock: req.MinimalStock,
		Price:        req.Price,
		Status:       req.Status,
		UserID:       req.UserID,
		CreatedBy:    req.CreatedBy,
		Description:  req.Description,
	}

	if err := r.DB.Save(&create).Error; err != nil {
		return dto.ProductResponse{}, err
	}

	response := dto.ProductResponse{
		ID:           req.ID,
		UserID:       req.UserID,
		ProductName:  create.ProductName,
		Description:  create.Description,
		Stock:        create.Stock,
		MinimalStock: create.MinimalStock,
		Price:        create.Price,
		Status:       create.Status,
		CreatedBy:    req.CreatedBy,
	}

	return response, nil

}
