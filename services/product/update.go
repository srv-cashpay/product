package product

import "github.com/srv-cashpay/product/dto"

func (b *productService) Update(req dto.ProductUpdateRequest) (dto.ProductUpdateResponse, error) {
	request := dto.ProductUpdateRequest{
		ProductName:  req.ProductName,
		Stock:        req.Stock,
		MinimalStock: req.MinimalStock,
		Price:        req.Price,
		UpdatedBy:    req.UpdatedBy,
	}

	product, err := b.Repo.Update(req)
	if err != nil {
		return product, err
	}

	response := dto.ProductUpdateResponse{
		ProductName:  request.ProductName,
		Stock:        request.Stock,
		MinimalStock: request.MinimalStock,
		Price:        request.Price,
		UpdatedBy:    request.UpdatedBy,
	}

	return response, nil
}
