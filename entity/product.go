package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID           string         `gorm:"primary_key,omitempty" json:"id"`
	UserID       string         `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID   string         `gorm:"type:varchar(36);index" json:"merchant_id"`
	ProductName  string         `gorm:"product_name,omitempty" json:"product_name"`
	Stock        int            `gorm:"stock,omitempty" json:"stock"`
	MinimalStock int            `gorm:"minimal_stock,omitempty" json:"minimal_stock"`
	Price        string         `gorm:"price,omitempty" json:"price"`
	Description  string         `gorm:"description,omitempty" json:"description"`
	CreatedBy    string         `gorm:"created_by" json:"created_by"`
	UpdatedBy    string         `gorm:"updated_by" json:"updated_by"`
	DeletedBy    string         `gorm:"deleted_by" json:"deleted_by"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
