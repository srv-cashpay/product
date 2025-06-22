package product

import (
	"fmt"
	"math"
	"strings"

	dto "github.com/srv-cashpay/product/dto"
	"github.com/srv-cashpay/product/entity"
	"github.com/srv-cashpay/product/helpers"
)

func (r *productRepository) Get(req *dto.Pagination) (RepositoryResult, int) {
	var products []entity.Product
	var totalRows int64
	totalPages, fromRow, toRow := 0, 0, 0
	offset := (req.Page - 1) * req.Limit

	find := r.DB.Preload("Category").Preload("Image").
		Where("merchant_id = ? AND status = ?", req.MerchantID, 1).
		Limit(req.Limit).
		Offset(offset).
		Order(req.Sort)

	// Filtering pencarian
	if req.Searchs != nil {
		for _, value := range req.Searchs {
			column := value.Column
			action := value.Action
			query := value.Query

			// Default field filtering
			switch action {
			case "equals":
				find = find.Where(fmt.Sprintf("%s = ?", column), query)
			case "contains":
				find = find.Where(fmt.Sprintf("%s LIKE ?", column), "%"+query+"%")
			case "in":
				find = find.Where(fmt.Sprintf("%s IN (?)", column), strings.Split(query, ","))
			}
		}
	}

	find = find.Find(&products)

	// Periksa jika ada error saat pengambilan data
	if errFind := find.Error; errFind != nil {
		return RepositoryResult{Error: errFind}, totalPages
	}

	// Hitung total data
	if errCount := r.DB.Model(&entity.Product{}).Where("merchant_id = ?", req.MerchantID).Count(&totalRows).Error; errCount != nil {
		return RepositoryResult{Error: errCount}, totalPages
	}
	for i := range products {
		products[i].ProductName = helpers.TruncateString(products[i].ProductName, 47)
	}

	// Pagination info
	req.TotalRows = int(totalRows)
	totalPages = int(math.Ceil(float64(totalRows) / float64(req.Limit)))
	req.TotalPages = totalPages

	fromRow = offset + 1
	toRow = offset + req.Limit
	if toRow > int(totalRows) {
		toRow = int(totalRows)
	}

	req.FromRow = fromRow
	req.ToRow = toRow

	return RepositoryResult{Result: req}, totalPages
}
