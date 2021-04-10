package json

import (
	"encoding/json"
	"hexagonal-architecture-example/core"
	"hexagonal-architecture-example/core/serializer"
	"log"
)

type productSerializer struct {
	productSerializer serializer.ProductSerializer
}

func (p productSerializer) Decode(input []byte) (*core.ProductDto, error) {
	var obj core.ProductDto
	if err := json.Unmarshal(input, &obj); err != nil {
		log.Println(err.Error())
		return nil,err
	}
return &obj,nil
}

func (p productSerializer) Encode(input *core.ProductDto) ([]byte, error) {
	b, err:= json.Marshal(input)
	if err != nil {
		log.Println(err.Error())
		return nil,err
	}
	return b,nil
}

func NewProductSerializer()  serializer.ProductSerializer {
	return &productSerializer{}
}
