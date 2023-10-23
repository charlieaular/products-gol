package dtos

type CreateProductDTO struct {
	SKU   string  `json:"sku" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Brand string  `json:"brand" binding:"required"`
	Price float32 `json:"price" binding:"required"`
	Size  string  `json:"size" binding:"required"`
}
