package impl

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/productmgtmodel"
	"github.com/prakash-p-3121/productmgtms/service/marketplace_listing_service"
	"github.com/prakash-p-3121/restlib"
)

type MarketplaceListingControllerImpl struct {
	MarketplaceListingService marketplace_listing_service.MarketplaceListingService
}

func (controller *MarketplaceListingControllerImpl) CreateMarketplaceListing(restCtx restlib.RestContext) {
	ginRestCtx, ok := restCtx.(*restlib.GinRestContext)
	if !ok {
		internalServerErr := errorlib.NewInternalServerError("Expected GinRestContext")
		internalServerErr.SendRestResponse(ginRestCtx.CtxGet())
		return
	}

	ctx := ginRestCtx.CtxGet()
	var req productmgtmodel.MarketplaceListingCreateReq
	err := ctx.BindJSON(&req)
	if err != nil {
		badReqErr := errorlib.NewBadReqError("payload-serialization")
		badReqErr.SendRestResponse(ctx)
		return
	}
	idGenResp, appErr := controller.MarketplaceListingService.CreateMarketplaceListing(&req)
	if appErr != nil {
		appErr.SendRestResponse(ctx)
		return
	}
	restlib.OkResponse(ctx, idGenResp)
}

func (controller *MarketplaceListingControllerImpl) FindMarketplaceListingByID(restCtx restlib.RestContext) {
	ginRestCtx, ok := restCtx.(*restlib.GinRestContext)
	if !ok {
		internalServerErr := errorlib.NewInternalServerError("Expected GinRestContext")
		internalServerErr.SendRestResponse(ginRestCtx.CtxGet())
		return
	}

	ctx := ginRestCtx.CtxGet()
	listingID := ctx.Query("listing-id")
	if restlib.TrimAndCheckForEmptyString(&listingID) {
		badReqErr := errorlib.NewBadReqError("listing-id-empty")
		badReqErr.SendRestResponse(ctx)
		return
	}
	marketplaceListing, appErr := controller.MarketplaceListingService.FindMarketplaceListingByID(listingID)
	if appErr != nil {
		appErr.SendRestResponse(ctx)
		return
	}
	restlib.OkResponse(ctx, marketplaceListing)
}
