package models

type TgConstructAttachment struct {
	Id           int    `json:"id" gorm:"AUTO_INCREMENT;primary_key;column:id"`
	OriginalName string `json:"original_name" gorm:"column:original_name"`
	Mime         string `json:"mime" gorm:"column:mime"`
	Extension    string `json:"extension" gorm:"column:extension"`
	Size         int    `json:"size" gorm:"column:size"`
	Sort         int    `json:"sort" gorm:"column:sort;default:0"`
	Path         string `json:"path" gorm:"column:path"`
	Description  string `json:"description" gorm:"column:description"`
	Alt          string `json:"alt" gorm:"column:alt"`
	Hash         string `json:"hash" gorm:"column:hash"`
	Disk         string `json:"disk" gorm:"column:disk"`
	Group        string `json:"group" gorm:"column:group"`
	UserId       uint   `json:"user_id" gorm:"column:user_id"`
	Type         string `json:"type" gorm:"column:type"`
	TgFileId     string `json:"tg_file_id" gorm:"column:tg_file_id"`
	CreatedAt    any    `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    any    `json:"updated_at" gorm:"column:updated_at"`
}

type TgConstructAttachmentable struct {
	Id                            int                    `json:"id" gorm:"AUTO_INCREMENT;primary_key;column:id"`
	TgConstructAttachmentableType string                 `json:"tg_construct_attachmentable_type" gorm:"column:tg_construct_attachmentable_type"`
	TgConstructAttachmentableId   uint                   `json:"tg_construct_attachmentable_id" gorm:"column:tg_construct_attachmentable_id"`
	TgConstructAttachmentId       uint                   `json:"tg_construct_attachment_id" gorm:"column:tg_construct_attachment_id"`
	TgConstructAttachment         *TgConstructAttachment `json:"TgConstructAttachmentId" gorm:"foreignKey:TgConstructAttachmentId;default:null"`
}
