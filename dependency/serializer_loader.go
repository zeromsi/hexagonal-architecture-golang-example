package dependency

import (
	"hexagonal-architecture-example/config"
	"hexagonal-architecture-example/core/serializer"
	"hexagonal-architecture-example/serializer/json"
	"hexagonal-architecture-example/serializer/msgpack"
	"hexagonal-architecture-example/serializer/xml"
)

func GetProductSerializer() serializer.ProductSerializer{
	if config.Serializer==string(config.JSON){
		return json.NewProductSerializer()
	}else if config.Serializer==string(config.XML){
		return xml.NewProductSerializer()
	}else if config.Serializer==string(config.MSGPACK){
		return msgpack.NewProductSerializer()
	}else{
		return nil
	}
}

