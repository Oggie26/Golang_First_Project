package dto

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=2"`
	Code        string  `json:"code" validate:"required,min=3"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Stock       int     `json:"stock" validate:"required,gte=0"`
	CategoryID  uint    `json:"category_id" validate:"required"`
	Description string  `json:"description"`
}

type ProductResponse struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Code         string  `json:"code"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	CategoryName string  `json:"category_name,omitempty"`
}
