package core

type ProductDto struct {
	Code      string `json:"code" msgpack:"code" xml:"code"`
	Name        string   `json:"name" msgpack:"name" xml:"name"`
	Price    int32 `json:"price" msgpack:"price" xml:"price"`
}


func(product ProductDto) GetProduct()Product{
	return Product{
		Code:  product.Code,
		Name:  product.Name,
		Price: product.Price,
	}
}