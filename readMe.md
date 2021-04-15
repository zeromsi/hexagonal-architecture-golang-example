## Project Hierarchy 
```
.
├── api
│   ├── base_router.go
│   ├── product_handler.go
│   ├── server.go
│   └── utility
│       ├── error_responses.go
│       ├── middleware
│       │   ├── middleware.go
│       │   └── token_validator_middleware.go
│       └── success_responses.go
├── application.go
├── config
│   ├── config.go
│   ├── load_env.go
│   └── run_mode.go
├── core
│   ├── dto
│   ├── entity
│   ├── logic
│   │   └── v1_product_service_impl.go
│   ├── product_dto.go
│   ├── product.go
│   ├── repository
│   │   └── productRepository.go
│   ├── serializer
│   │   └── product.go
│   └── service
│       └── productService.go
├── dependency
│   └── service_loader.go
├── Dockerfile
├── go.mod
├── go.sum
├── readMe.md
├── repository
│   ├── in_memory
│   │   ├── data.go
│   │   └── product_repository_impl.go
│   ├── mongo
│   │   ├── odm.go
│   │   └── product_repository_impl.go
│   ├── mysql
│   │   ├── connector.go
│   │   └── product_repository_impl.go
│   └── redis
│       ├── connector.go
│       └── product_repository_impl.go
└── serializer
    ├── json
    │   └── product.go
    ├── msgpack
    │   └── product.go
    └── xml
        └── product.go


```