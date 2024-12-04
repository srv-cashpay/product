package entity

type MerchantAutoIncrement struct {
	MerchantID    string `gorm:"primary_key"`
	NextIncrement int    `gorm:"not null;default:1"`
}
