package category_repository

import (
	"github.com/prakash-p-3121/productmgtms/repository/category_repository/impl"
	"sync"
)

func NewCategoryRepository(shardConnectionsMap *sync.Map) CategoryRepository {
	repository := impl.CategoryRepositoryImpl{ShardConnectionsMap: shardConnectionsMap}
	return &repository
}
