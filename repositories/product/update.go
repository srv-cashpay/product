package product

import (
	dto "github.com/srv-cashpay/product/dto"
	"github.com/srv-cashpay/product/entity"
)

func (b *productRepository) Update(req dto.ProductUpdateRequest) (dto.ProductUpdateResponse, error) {
	tr := dto.GetByIdRequest{
		ID: req.ID,
	}

	request := entity.Product{
		ProductName:  req.ProductName,
		Stock:        req.Stock,
		MinimalStock: req.MinimalStock,
		Price:        req.Price,
		UpdatedBy:    req.UpdatedBy,
	}

	product, err := b.GetById(tr)
	if err != nil {
		return dto.ProductUpdateResponse{}, err
	}

	err = b.DB.Where("ID = ?", req.ID).Updates(entity.Product{
		ProductName:  request.ProductName,
		Stock:        request.Stock,
		MinimalStock: request.MinimalStock,
		Price:        request.Price,
		UpdatedBy:    request.UpdatedBy,
	}).Error
	if err != nil {
		return dto.ProductUpdateResponse{}, err
	}

	response := dto.ProductUpdateResponse{
		ProductName:  request.ProductName,
		Stock:        request.Stock,
		MinimalStock: request.MinimalStock,
		Price:        request.Price,
		ID:           product.ID,
		UpdatedBy:    request.UpdatedBy,
	}

	return response, nil
}
