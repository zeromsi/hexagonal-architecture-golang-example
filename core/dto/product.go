package dto
type Product struct {
	Code      string `json:"code" msgpack:"code" xml:"code"`
	Name        string   `json:"name" msgpack:"name" xml:"name"`
	Price    int32 `json:"price" msgpack:"price" xml:"price"`
}

