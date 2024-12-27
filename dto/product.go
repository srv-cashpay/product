package dto

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
	Status       string `json:"status"`
	CreatedBy    string `json:"created_by"`
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
