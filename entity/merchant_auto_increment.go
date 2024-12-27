package entity

type MerchantAutoIncrement struct {
	MerchantID    string `gorm:"primary_key"`
	NextIncrement int    `gorm:"not null;default:1"`
}

type MerkAutoIncrement struct {
	MerkID        string `gorm:"primary_key"`
	NextIncrement int    `gorm:"not null;default:1"`
}

type CategoryAutoIncrement struct {
	CategoryID    string `gorm:"primary_key"`
	NextIncrement int    `gorm:"not null;default:1"`
}
