package product

import (
	"github.com/srv-cashpay/product/entity"
)

func (r *productRepository) Get() ([]entity.Product, error) {
	var data []entity.Product

	if err := r.DB.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
