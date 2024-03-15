package product_repository

import (
	"github.com/prakash-p-3121/productmgtms/database"
	"github.com/prakash-p-3121/productmgtms/repository/product_repository/impl"
)

func NewProductRepository() ProductRepository {
	repository := impl.ProductRepositoryImpl{ShardConnectionsMap: database.GetShardConnectionsMap()}
	return &repository
}
