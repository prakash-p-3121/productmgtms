package impl

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/productmgtmodel"
	"github.com/prakash-p-3121/productmgtms/service/category_service"
	"github.com/prakash-p-3121/restlib"
)

type CategoryControllerImpl struct {
	CategoryService category_service.CategoryService
}

func (controller *CategoryControllerImpl) CreateCategory(restCtx restlib.RestContext) {
	ginRestCtx, ok := restCtx.(*restlib.GinRestContext)
	if !ok {
		internalServerErr := errorlib.NewInternalServerError("Expected GinRestContext")
		internalServerErr.SendRestResponse(ginRestCtx.CtxGet())
		return
	}

	ctx := ginRestCtx.CtxGet()
	var req productmgtmodel.CategoryCreateReq
	err := ctx.BindJSON(&req)
	if err != nil {
		badReqErr := errorlib.NewBadReqError("payload-serialization")
		badReqErr.SendRestResponse(ctx)
		return
	}

	idResp, appErr := controller.CategoryService.CreateCategory(&req)
	if appErr != nil {
		appErr.SendRestResponse(ctx)
		return
	}
	restlib.OkResponse(ctx, idResp)
}

func (controller *CategoryControllerImpl) FindCategory(restCtx restlib.RestContext) {
	ginRestCtx, ok := restCtx.(*restlib.GinRestContext)
	if !ok {
		internalServerErr := errorlib.NewInternalServerError("Expected GinRestContext")
		internalServerErr.SendRestResponse(ginRestCtx.CtxGet())
		return
	}

	ctx := ginRestCtx.CtxGet()
	categoryID := ctx.Query("category-id")
	if restlib.TrimAndCheckForEmptyString(&categoryID) {
		badReqErr := errorlib.NewBadReqError("category-id-empty")
		badReqErr.SendRestResponse(ctx)
		return
	}

	categoryPtr, appErr := controller.CategoryService.FindCategory(categoryID)
	if appErr != nil {
		appErr.SendRestResponse(ctx)
		return
	}
	restlib.OkResponse(ctx, categoryPtr)
}
