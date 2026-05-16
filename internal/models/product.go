package models

type Product struct {
	ProductID       string  `json:"product_id"`
	ProductName     string  `json:"product_name"`
	ProductCode     string  `json:"product_code"`
	ProductCategory string  `json:"product_category"`
	ProductPrice    float64 `json:"product_price"`
}
