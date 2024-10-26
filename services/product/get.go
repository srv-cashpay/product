package product

import (
	entity "github.com/srv-cashpay/product/entity"
)

func (s *productService) Get() ([]entity.Product, error) {
	// Fetch comments from the repository layer based on post_id
	comments, err := s.Repo.Get()
	if err != nil {
		return nil, err
	}

	return comments, nil
}
