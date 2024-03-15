package product_service

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/productmgtmodel"
)

type ProductService interface {
	CreateProduct(req *productmgtmodel.ProductCreateReq) (*idgenmodel.IDGenResp, errorlib.AppError)
	FindProduct(productID string) (*productmgtmodel.Product, errorlib.AppError)
}
