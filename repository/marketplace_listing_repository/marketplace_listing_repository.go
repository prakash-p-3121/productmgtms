package marketplace_listing_repository

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/productmgtmodel"
)

type MarketplaceListingRepository interface {
	CreateMarketplaceListing(idGenResp *idgenmodel.IDGenResp, req *productmgtmodel.MarketplaceListingCreateReq) errorlib.AppError
	FindMarketplaceListingByID(listingID string) (*productmgtmodel.MarketplaceListing, errorlib.AppError)
}
