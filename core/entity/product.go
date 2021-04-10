package entity

import "hexagonal-architecture-example/core/dto"

type Product struct {
	Code      string `bson:"code"`
	Name        string   `bson:"name"`
	Price    int32 `bson:"price"`
}

func(product Product) GetProductDto() dto.Product{
	return dto.Product{
		Code:  product.Code,
		Name:  product.Name,
		Price: product.Price,
	}
}
