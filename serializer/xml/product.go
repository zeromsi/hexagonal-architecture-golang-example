package xml

import (
	"encoding/xml"
	"hexagonal-architecture-example/core"
	"hexagonal-architecture-example/core/serializer"
	"log"
)

type productSerializer struct {
	productSerializer serializer.ProductSerializer
}

func (p productSerializer) Decode(input []byte) (*core.ProductDto, error) {
	var obj core.ProductDto
	if err := xml.Unmarshal(input, &obj); err != nil {
		log.Println(err.Error())
		return nil,err
	}
	return &obj,nil
}

func (p productSerializer) Encode(input *core.ProductDto) ([]byte, error) {
	b, err:= xml.Marshal(input)
	if err != nil {
		log.Println(err.Error())
		return nil,err
	}
	return b,nil
}

func NewProductSerializer()  serializer.ProductSerializer {
	return &productSerializer{}
}

