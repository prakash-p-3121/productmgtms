package product_service

import (
	"github.com/prakash-p-3121/productmgtms/repository/product_repository"
	"github.com/prakash-p-3121/productmgtms/service/product_service/impl"
)

func NewProductService() ProductService {
	repository := product_repository.NewProductRepository()
	return &impl.ProductServiceImpl{ProductRepository: repository}
}
