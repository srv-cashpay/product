package product

import (
	dto "github.com/srv-cashpay/product/dto"
)

func (b *productService) Delete(req dto.DeleteRequest) (dto.DeleteResponse, error) {
	transactionBody := dto.DeleteRequest{
		ID:        req.ID,
		DeletedBy: req.DeletedBy,
	}

	_, err := b.Repo.Delete(req)
	if err != nil {
		return dto.DeleteResponse{}, err
	}

	response := dto.DeleteResponse{
		ID:        transactionBody.ID,
		DeletedBy: transactionBody.DeletedBy,
	}

	return response, nil
}
