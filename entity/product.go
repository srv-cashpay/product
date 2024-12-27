package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID           string         `gorm:"primary_key,omitempty" json:"id"`
	Barcode      string         `gorm:"barcode" json:"barcode"`
	UserID       string         `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID   string         `gorm:"type:varchar(36);index" json:"merchant_id"`
	MerkID       string         `gorm:"merk_id" json:"merk_id"`
	PictureID    string         `gorm:"picture_id" json:"picture_id"`
	CategoryID   string         `gorm:"category_id" json:"category_id"`
	ProductName  string         `gorm:"product_name,omitempty;type:varchar(70)" json:"product_name"`
	Stock        int            `gorm:"stock,omitempty" json:"stock"`
	MinimalStock int            `gorm:"minimal_stock,omitempty" json:"minimal_stock"`
	Price        int            `gorm:"price,omitempty" json:"price"`
	Status       int            `gorm:"status" json:"status"`
	Description  string         `gorm:"description" json:"description"`
	CreatedBy    string         `gorm:"created_by" json:"created_by"`
	UpdatedBy    string         `gorm:"updated_by" json:"updated_by"`
	DeletedBy    string         `gorm:"deleted_by" json:"deleted_by"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
