package product

import (
	dto "github.com/srv-cashpay/product/dto"
)

func (b *productService) GetById(req dto.GetByIdRequest) (*dto.ProductResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
