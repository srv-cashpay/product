package dto

type ProductRequest struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	MerchantID   string `json:"merchant_id"`
	ProductName  string `json:"product_name"`
	Description  string `json:"description"`
	Stock        int    `json:"stock"`
	MinimalStock int    `json:"minimal_stock"`
	Price        int    `json:"price"`
	Status       bool   `json:"status"`
	CreatedBy    string `json:"created_by"`
}

type ProductResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	MerchantID   string `json:"merchant_id"`
	ProductName  string `json:"product_name"`
	Description  string `json:"description"`
	Stock        int    `json:"stock"`
	MinimalStock int    `json:"minimal_stock"`
	Price        int    `json:"price"`
	Status       bool   `json:"status"`
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
	DeletedBy string `json:"deleted_by"`
	Count     int    `json:"count"`
}

type ProductUpdateRequest struct {
	ID           string `json:"id"`
	ProductName  string `json:"product_name"`
	Stock        int    `json:"stock"`
	MinimalStock int    `json:"minimal_stock"`
	Price        int    `json:"price"`
	Status       bool   `json:"status"`
	UpdatedBy    string `json:"updated_by"`
	UserID       string `json:"user_id"`
	MerchantID   string `json:"merchant_id"`
	Description  string `json:"description"`
}

type ProductUpdateResponse struct {
	ID           string `json:"id"`
	ProductName  string `json:"product_name"`
	Stock        int    `json:"stock"`
	MinimalStock int    `json:"minimal_stock"`
	Price        int    `json:"price"`
	Status       bool   `json:"status"`
	UpdatedBy    string `json:"updated_by"`
	UserID       string `json:"user_id"`
	MerchantID   string `json:"merchant_id"`
	Description  string `json:"description"`
}
