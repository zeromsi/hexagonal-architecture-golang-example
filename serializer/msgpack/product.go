package msgpack

import (
	"hexagonal-architecture-example/core"
	"hexagonal-architecture-example/core/serializer"
)

type productSerializer struct {
	productSerializer serializer.ProductSerializer
}

func (p productSerializer) Decode(input []byte) (*core.ProductDto, error) {
	panic("implement me")
}

func (p productSerializer) Encode(input *core.ProductDto) ([]byte, error) {
	panic("implement me")
}

func NewProductSerializer()  serializer.ProductSerializer {
	return &productSerializer{}
}

