package product_controller

import (
	"github.com/prakash-p-3121/productmgtms/controller/product_controller/impl"
	"github.com/prakash-p-3121/productmgtms/service/product_service"
)

func NewProductController() ProductController {
	service := product_service.NewProductService()
	return &impl.ProductControllerImpl{ProductService: service}
}
