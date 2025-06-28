package product

import (
	dto "github.com/srv-cashpay/product/dto"
)

func (u *productService) GetUrl(req dto.ProductRequest) (dto.UrlResponse, error) {
	// Validasi refresh token dan dapatkan user ID

	comments, err := u.Repo.GetUrl(req)
	if err != nil {
		return dto.UrlResponse{}, err
	}

	return comments, nil
}
