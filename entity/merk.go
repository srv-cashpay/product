package entity

type Merk struct {
	ID     string `gorm:"primary_key,omitempty" json:"id"`
	UserID string `gorm:"type:varchar(36);index" json:"user_id"`
	Merk   string `gorm:"merk,omitempty" json:"merk"`
}
