package dto

import "time"

type ProductRequest struct {
	ID           string `json:"id"`
	Barcode      string `json:"barcode"`
	UserID       string `json:"user_id"`
	MerchantID   string `json:"merchant_id"`
	MerkID       string `json:"merk_id"`
	CategoryID   string `json:"category_id"`
	ProductName  string `json:"product_name"`
	Description  string `json:"description"`
	Stock        int    `json:"stock"`
	MinimalStock int    `json:"minimal_stock"`
	Price        int    `json:"price"`
	Status       int    `json:"status"`
	CreatedBy    string `json:"created_by"`
}

type ProductResponse struct {
	ID           string    `json:"id"`
	Barcode      string    `json:"barcode"`
	UserID       string    `json:"user_id"`
	MerchantID   string    `json:"merchant_id"`
	MerkID       string    `json:"merk_id"`
	CategoryID   string    `json:"category_id"`
	CategoryName string    `json:"category_name"`
	ProductName  string    `json:"product_name"`
	Description  string    `json:"description"`
	Stock        int       `json:"stock"`
	MinimalStock int       `json:"minimal_stock"`
	Price        int       `json:"price"`
	Status       int       `json:"status"`
	CreatedBy    string    `json:"created_by"`
	CreatedAt    Timestamp `json:"created_at"`
}

type UrlResponse struct {
	MerchantID string `json:"merchant_id"`
}

type ProductPaginationResponse struct {
	Limit        int               `json:"limit"`
	Page         int               `json:"page"`
	Sort         string            `json:"sort"`
	TotalRows    int               `json:"total_rows"`
	TotalPages   int               `json:"total_page"`
	FirstPage    string            `json:"first_page"`
	PreviousPage string            `json:"previous_page"`
	NextPage     string            `json:"next_page"`
	LastPage     string            `json:"last_page"`
	FromRow      int               `json:"from_row"`
	ToRow        int               `json:"to_row"`
	Data         []ProductResponse `json:"data"`
	Searchs      []Search          `json:"searchs"`
}

type GetByIdRequest struct {
	ID string `param:"id" validate:"required"`
}

type DeleteRequest struct {
	ID        string `param:"id" validate:"required"`
	DeletedBy string `json:"deleted_by"`
}

type DeleteResponse struct {
	ID        string `json:"id"`
	DeletedBy string `json:"deleted_by"`
}

type BulkDeleteRequest struct {
	ID        []string `json:"id"`
	DeletedBy string   `json:"deleted_by"`
}

type BulkDeleteResponse struct {
	ID        []string `json:"id"`
	DeletedBy string   `json:"deleted_by"`
	Count     int      `json:"count"`
}

type ProductUpdateRequest struct {
	ID           string `json:"id"`
	ProductName  string `json:"product_name"`
	Barcode      string `json:"barcode"`
	Stock        int    `json:"stock"`
	MinimalStock int    `json:"minimal_stock"`
	Price        int    `json:"price"`
	Status       int    `json:"status"`
	UpdatedBy    string `json:"updated_by"`
	UserID       string `json:"user_id"`
	MerchantID   string `json:"merchant_id"`
	Description  string `json:"description"`
}

type ProductUpdateResponse struct {
	ID           string `json:"id"`
	ProductName  string `json:"product_name"`
	Barcode      string `json:"barcode"`
	Stock        int    `json:"stock"`
	MinimalStock int    `json:"minimal_stock"`
	Price        int    `json:"price"`
	Status       int    `json:"status"`
	UpdatedBy    string `json:"updated_by"`
	UserID       string `json:"user_id"`
	MerchantID   string `json:"merchant_id"`
	Description  string `json:"description"`
}

// Formatter untuk timestamp dengan nama bulan
type Timestamp time.Time

// Format waktu: 19 December 2024, 09:28:17
const timeFormat = "02 January 2006, 15:04:05"

// MarshalJSON untuk memformat waktu
func (t Timestamp) MarshalJSON() ([]byte, error) {
	// Konversi waktu ke zona waktu lokal
	localTime := time.Time(t).Local()
	formattedTime := "\"" + localTime.Format(timeFormat) + "\""
	return []byte(formattedTime), nil
}

// UnmarshalJSON untuk parsing waktu dari JSON (opsional)
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	parsedTime, err := time.Parse("\""+timeFormat+"\"", string(data))
	if err != nil {
		return err
	}
	*t = Timestamp(parsedTime)
	return nil
}

// ToTime untuk mengonversi kembali ke time.Time
func (t Timestamp) ToTime() time.Time {
	return time.Time(t)
}
