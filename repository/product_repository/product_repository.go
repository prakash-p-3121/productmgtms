package product_repository

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/productmgtmodel"
)

type ProductRepository interface {
	CreateProduct(shardID int64, idGenResp *idgenmodel.IDGenResp, req *productmgtmodel.ProductCreateReq) errorlib.AppError
	FindProduct(shardID int64, productID string) (*productmgtmodel.Product, errorlib.AppError)
}
