package products

var products []Product

func init() {
	products = []Product{
		Product{
			ID:          0,
			Name:        "Bag",
			Description: "Size: 38cm x 42cm",
			Price:       100,
		},
		Product{
			ID:          1,
			Name:        "T-shirt",
			Description: "Color: Black; Sizes: S, M, L, XL",
			Price:       200,
		},
	}
}

// GetProduct return product by id
func GetProduct(id int64) Product {
	return products[id]
}

// GetProducts return products
func GetProducts() []Product {
	return products
}
