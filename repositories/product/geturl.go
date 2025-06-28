package product

import (
	dto "github.com/srv-cashpay/product/dto"
	"github.com/srv-cashpay/product/entity"
)

func (u *productRepository) GetUrl(req dto.ProductRequest) (dto.ProductResponse, error) {
	var existingUser entity.Product

	if err := u.DB.Where("merchant_id = ?", req.MerchantID).Find(&existingUser).Error; err != nil {
		return dto.ProductResponse{}, err
	}

	resp := dto.ProductResponse{
		MerchantID: "https://cashpay.my.id/menu?merchant_id=" + existingUser.MerchantID,
	}

	return resp, nil
}
