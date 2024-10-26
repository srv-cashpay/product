package entity

type Type struct {
	ID     string `gorm:"primary_key,omitempty" json:"id"`
	UserID string `gorm:"type:varchar(36);index" json:"user_id"`
	Type   string `gorm:"type,omitempty" json:"type"`
}
