package dependency

import (
	"hexagonal-architecture-example/config"
	"hexagonal-architecture-example/core/logic"
	"hexagonal-architecture-example/core/serializer"
	"hexagonal-architecture-example/core/service"
	"hexagonal-architecture-example/repository/in_memory"
	"hexagonal-architecture-example/repository/mongo"
	"hexagonal-architecture-example/repository/mysql"
	"hexagonal-architecture-example/repository/redis"
	"hexagonal-architecture-example/serializer/json"
	"hexagonal-architecture-example/serializer/msgpack"
	"hexagonal-architecture-example/serializer/xml"
)

func GetProductService() service.ProductService {
	if config.ServiceVersion==string(config.V1){
		if config.Database == string(config.MONGO) {
			return logic.NewV1ProductService(mongo.NewProductRepository(5000))
		}else if config.Database == string(config.IN_MEMORY) {
			return logic.NewV1ProductService(in_memory.NewProductRepository())
		}else if config.Database == string(config.MYSQL) {
			return logic.NewV1ProductService(mysql.NewProductRepository())
		}else if config.Database == string(config.REDIS) {
			return logic.NewV1ProductService(redis.NewProductRepository())
		}  else {
			return nil
		}
	}else {
		return nil
	}
}
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