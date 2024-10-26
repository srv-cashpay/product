package entity

type Satuan struct {
	ID     string `gorm:"primary_key,omitempty" json:"id"`
	UserID string `gorm:"type:varchar(36);index" json:"user_id"`
	Satuan string `gorm:"satuan,omitempty" json:"satuan"`
}
