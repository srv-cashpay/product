package dto

type ProductRequest struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	ProductName  string `json:"product_name"`
	Stock        string `json:"stock"`
	MinimalStock string `json:"minimal_stock"`
	Price        string `json:"price"`
	CreatedBy    string `json:"created_by"`
}

type ProductResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	ProductName  string `json:"product_name"`
	Stock        string `json:"stock"`
	MinimalStock string `json:"minimal_stock"`
	Price        string `json:"price"`
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
	Stock        string `json:"stock"`
	MinimalStock string `json:"minimal_stock"`
	Price        string `json:"price"`
	UpdatedBy    string `json:"updated_by"`
}

type ProductUpdateResponse struct {
	ID           string `json:"id"`
	ProductName  string `json:"product_name"`
	Stock        string `json:"stock"`
	MinimalStock string `json:"minimal_stock"`
	Price        string `json:"price"`
	UpdatedBy    string `json:"updated_by"`
}
