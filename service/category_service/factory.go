package category_service

import (
	"github.com/prakash-p-3121/productmgtms/database"
	"github.com/prakash-p-3121/productmgtms/repository/category_repository"
	"github.com/prakash-p-3121/productmgtms/service/category_service/impl"
)

func NewCategoryService() CategoryService {
	repository := category_repository.NewCategoryRepository(database.GetShardConnectionsMap())
	return &impl.CategoryServiceImpl{CategoryRepository: repository}
}
