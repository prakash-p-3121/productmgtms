package category_controller

import (
	"github.com/prakash-p-3121/restlib"
)

type CategoryController interface {
	CreateCategory(restCtx restlib.RestContext)
	FindCategory(restCtx restlib.RestContext)
}
