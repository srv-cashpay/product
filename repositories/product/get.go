package product

import (
	"fmt"
	"math"
	"strings"

	dto "github.com/srv-cashpay/product/dto"
	"github.com/srv-cashpay/product/entity"
	"github.com/srv-cashpay/product/helpers"
)

func (r *productRepository) Get(req *dto.Pagination) (dto.ProductPaginationResponse, int) {
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
				list := strings.Split(query, ",")
				find = find.Where(fmt.Sprintf("%s IN ?", column), list)
			}
		}
	}

	if err := find.Find(&products).Error; err != nil {
		return dto.ProductPaginationResponse{}, totalPages
	}

	// Hitung total data
	if errCount := r.DB.Model(&entity.Product{}).Where("merchant_id = ?", req.MerchantID).Count(&totalRows).Error; errCount != nil {
		return dto.ProductPaginationResponse{}, totalPages
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

	// Mapping ke DTO response
	var productResponses []dto.ProductResponse
	for _, p := range products {
		productResponses = append(productResponses, dto.ProductResponse{
			ID:           p.ID,
			Barcode:      p.Barcode,
			UserID:       p.UserID,
			MerchantID:   p.MerchantID,
			MerkID:       p.MerkID,
			CategoryID:   p.CategoryID,
			CategoryName: p.Category.CategoryName,
			ProductName:  p.ProductName,
			Description:  p.Description,
			Stock:        p.Stock,
			MinimalStock: p.MinimalStock,
			Price:        p.Price,
			Status:       p.Status,
			CreatedBy:    p.CreatedBy,
		})
	}

	return dto.ProductPaginationResponse{
		Limit:        req.Limit,
		Page:         req.Page,
		Sort:         req.Sort,
		TotalRows:    req.TotalRows,
		TotalPages:   req.TotalPages,
		FirstPage:    req.FirstPage,
		PreviousPage: req.PreviousPage,
		NextPage:     req.NextPage,
		LastPage:     req.LastPage,
		FromRow:      req.FromRow,
		ToRow:        req.ToRow,
		Data:         productResponses,
		Searchs:      req.Searchs,
	}, totalPages
}
