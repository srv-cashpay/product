package product

import (
	dto "github.com/srv-cashpay/product/dto"
	util "github.com/srv-cashpay/util/s"
)

func (s *productService) Create(req dto.ProductRequest) (dto.ProductResponse, error) {
	create := dto.ProductRequest{
		ID:           util.GenerateRandomString(),
		ProductName:  req.ProductName,
		Description:  req.Description,
		Stock:        req.Stock,
		MinimalStock: req.MinimalStock,
		Price:        req.Price,
		Status:       req.Status,
		UserID:       req.UserID,
		MerchantID:   req.MerchantID,
		CreatedBy:    req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	response := dto.ProductResponse{
		ID:           created.ID,
		UserID:       created.UserID,
		ProductName:  created.ProductName,
		Description:  created.Description,
		Stock:        created.Stock,
		MinimalStock: created.MinimalStock,
		Price:        created.Price,
		Status:       created.Status,
		MerchantID:   created.MerchantID,
		CreatedBy:    created.CreatedBy,
	}

	return response, nil
}
