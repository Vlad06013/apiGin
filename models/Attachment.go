package models

type Attachment struct {
	Id           int    `json:"id" gorm:"AUTO_INCREMENT;primary_key;column:id"`
	OriginalName string `json:"original_name" gorm:"column:original_name"`
	Mime         string `json:"mime" gorm:"column:mime"`
	Extension    string `json:"extension" gorm:"column:extension"`
	Size         int    `json:"size" gorm:"column:size"`
	Sort         int    `json:"sort" gorm:"column:sort;default:0"`
	Path         int    `json:"path" gorm:"column:path"`
	Description  string `json:"description" gorm:"column:description"`
	Alt          string `json:"alt" gorm:"column:alt"`
	Hash         string `json:"hash" gorm:"column:hash"`
	Disk         string `json:"disk" gorm:"column:disk"`
	Group        string `json:"group" gorm:"column:group"`
	UserId       uint   `json:"user_id" gorm:"column:user_id"`
	CreatedAt    any    `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    any    `json:"updated_at" gorm:"column:updated_at"`
}
