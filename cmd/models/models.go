package models

// Define a Product type to hold the information about an individual product.
type Product struct {
	ProductID int
	Qty       int
}

type ProductJson struct {
	ProductID int `json:"product_id"`
	Qty       int `json:"qty"`
}

// For convenience we also define a Products type, which is a slice for holding
// multiple Product objects.
type Products []*Product
