package marketplace_listing_controller

import "github.com/prakash-p-3121/restlib"

type MarketPlaceListingController interface {
	CreateMarketplaceListing(restCtx restlib.RestContext)
	FindMarketplaceListingByID(restCtx restlib.RestContext)
}
