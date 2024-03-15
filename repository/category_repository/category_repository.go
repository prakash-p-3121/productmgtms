package category_repository

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/productmgtmodel"
)

type CategoryRepository interface {
	CreateCategory(shardID int64, idResp *idgenmodel.IDGenResp, req *productmgtmodel.CategoryCreateReq) errorlib.AppError
	FindCategory(shardID int64, categoryID string) (*productmgtmodel.Category, errorlib.AppError)
}
