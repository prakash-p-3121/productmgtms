package category_service

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/productmgtmodel"
	restlib_model "github.com/prakash-p-3121/restlib/model"
)

type CategoryService interface {
	CreateCategory(req *productmgtmodel.CategoryCreateReq) (*restlib_model.IDResponse, errorlib.AppError)
	FindCategory(categoryID string) (*productmgtmodel.Category, errorlib.AppError)
}
