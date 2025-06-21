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

	// Ubah offset agar sesuai dengan page yang dimulai dari 1
	offset := (req.Page - 1) * req.Limit

	// Ambil data sesuai limit, offset, dan urutan
	find := r.DB.Preload("Category").Preload("Image").
		Where("merchant_id = ? AND status = ?", req.MerchantID, 1).
		Limit(req.Limit).
		Offset(offset).
		Order(req.Sort)

	// Generate where query untuk search
	if req.Searchs != nil {
		for _, value := range req.Searchs {
			column := value.Column
			action := value.Action
			query := value.Query

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
		return dto.ProductPaginationResponse{}, totalPages
	}

	req.Rows = products

	if errCount := r.DB.Model(&entity.Product{}).Where("merchant_id = ?", req.MerchantID).Count(&totalRows).Error; errCount != nil {
		return dto.ProductPaginationResponse{}, totalPages
	}

	for i := range products {
		products[i].ProductName = helpers.TruncateString(products[i].ProductName, 47)
	}

	req.TotalRows = int(totalRows)

	// Hitung total halaman berdasarkan limit
	totalPages = int(math.Ceil(float64(totalRows) / float64(req.Limit)))
	req.TotalPages = totalPages
	// Hitung `fromRow` dan `toRow` untuk page saat ini
	if req.Page == 1 {
		// Untuk halaman pertama
		fromRow = 1
		toRow = req.Limit
	} else {
		if req.Page <= totalPages {
			fromRow = (req.Page-1)*req.Limit + 1
			toRow = req.Page * req.Limit
		}
	}

	// Pastikan `toRow` tidak melebihi `totalRows`
	if toRow > int(totalRows) {
		toRow = int(totalRows)
	}

	// Set hasil akhir
	req.FromRow = fromRow
	req.ToRow = toRow

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

	response := dto.ProductPaginationResponse{
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
	}
	return response, totalPages

	// return RepositoryResult{Result: req}, totalPages
}
