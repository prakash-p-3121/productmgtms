package marketplace_listing_service

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/productmgtmodel"
)

type MarketplaceListingService interface {
	CreateMarketplaceListing(req *productmgtmodel.MarketplaceListingCreateReq) (*idgenmodel.IDGenResp, errorlib.AppError)
	FindMarketplaceListingByID(listingID string) (*productmgtmodel.MarketplaceListing, errorlib.AppError)
}
