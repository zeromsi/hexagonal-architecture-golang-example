package serializer

import (
	"hexagonal-architecture-example/core"
)

type ProductSerializer interface {
	Decode(input []byte) (*core.ProductDto, error)
	Encode(input *core.ProductDto) ([]byte, error)
}

