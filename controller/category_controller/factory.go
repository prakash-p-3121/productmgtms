package category_controller

import (
	"github.com/prakash-p-3121/productmgtms/controller/category_controller/impl"
	"github.com/prakash-p-3121/productmgtms/service/category_service"
)

func NewCategoryController() CategoryController {
	service := category_service.NewCategoryService()
	return &impl.CategoryControllerImpl{CategoryService: service}
}
