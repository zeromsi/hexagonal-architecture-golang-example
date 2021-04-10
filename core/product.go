package core

type Product struct {
	Code      string `bson:"code"`
	Name        string   `bson:"name"`
	Price    int32 `bson:"price"`
}

func(product Product) GetProductDto() ProductDto{
	return ProductDto{
		Code:  product.Code,
		Name:  product.Name,
		Price: product.Price,
	}
}
