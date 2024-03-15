package impl

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/productmgtmodel"
	"github.com/prakash-p-3121/productmgtms/service/product_service"
	"github.com/prakash-p-3121/restlib"
)

type ProductControllerImpl struct {
	ProductService product_service.ProductService
}

func (controller *ProductControllerImpl) CreateProduct(restCtx restlib.RestContext) {
	ginRestCtx, ok := restCtx.(*restlib.GinRestContext)
	if !ok {
		internalServerErr := errorlib.NewInternalServerError("Expected GinRestContext")
		internalServerErr.SendRestResponse(ginRestCtx.CtxGet())
		return
	}

	ctx := ginRestCtx.CtxGet()
	var req productmgtmodel.ProductCreateReq
	err := ctx.BindJSON(&req)
	if err != nil {
		badReqErr := errorlib.NewBadReqError("payload-serialization")
		badReqErr.SendRestResponse(ctx)
		return
	}

	idGenResp, appErr := controller.ProductService.CreateProduct(&req)
	if appErr != nil {
		appErr.SendRestResponse(ctx)
		return
	}
	restlib.OkResponse(ctx, idGenResp)
}

func (controller *ProductControllerImpl) FindProduct(restCtx restlib.RestContext) {
	ginRestCtx, ok := restCtx.(*restlib.GinRestContext)
	if !ok {
		internalServerErr := errorlib.NewInternalServerError("Expected GinRestContext")
		internalServerErr.SendRestResponse(ginRestCtx.CtxGet())
		return
	}

	ctx := ginRestCtx.CtxGet()
	productID := ctx.Query("product-id")
	if restlib.TrimAndCheckForEmptyString(&productID) {
		badReqErr := errorlib.NewBadReqError("product-id-empty")
		badReqErr.SendRestResponse(ctx)
		return
	}

	productPtr, appErr := controller.ProductService.FindProduct(productID)
	if appErr != nil {
		appErr.SendRestResponse(ctx)
		return
	}
	restlib.OkResponse(ctx, productPtr)
}
