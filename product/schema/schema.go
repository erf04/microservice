package schema

type GetProductByIDSchema struct {
	ID string `json:"id"`
}

type GetProductsSchema struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}