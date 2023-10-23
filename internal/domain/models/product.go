package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	SKU   string  `form:"sku" gorm:"unique; not null"`
	Name  string  `gorm:"type:varchar(50);not null"`
	Brand string  `gorm:"type:varchar(50);not null"`
	Size  string  `gorm:"type:varchar(50);not null"`
	Price float32 `gorm:"type:decimal(8,2)"`
}
