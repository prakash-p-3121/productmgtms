package impl

import (
	"database/sql"
	"errors"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/mysqllib"
	"github.com/prakash-p-3121/productmgtmodel"
)

type MarketplaceListingRepositoryImpl struct {
	SingleStoreConnection *sql.DB
}

func (repository *MarketplaceListingRepositoryImpl) CreateMarketplaceListing(idGenResp *idgenmodel.IDGenResp,
	req *productmgtmodel.MarketplaceListingCreateReq) errorlib.AppError {
	db := repository.SingleStoreConnection
	qry := `INSERT INTO marketplace_listings (id, 
            id_bit_count, 
            product_id, 
            seller_id, 
            selling_price, 
            currency, 
            stock_count,
            return_policy
            ) VALUES (?,?,?,?,?,?,?,?); `
	_, err := db.Exec(qry,
		idGenResp.ID,
		idGenResp.BitCount,
		*req.ProductID,
		*req.SellerID,
		*req.SellingPrice,
		*req.Currency,
		*req.StockCount,
		*req.ReturnPolicy,
	)
	if err != nil && mysqllib.IsConflictError(err) {
		return errorlib.NewConflictError("marketplace-listing-unq-key-conflict")
	}
	if err != nil {
		return errorlib.NewInternalServerError(err.Error())
	}
	return nil
}

func (repository *MarketplaceListingRepositoryImpl) FindMarketplaceListingByID(listingID string) (*productmgtmodel.MarketplaceListing,
	errorlib.AppError) {
	db := repository.SingleStoreConnection
	qry := `SELECT id, 
            id_bit_count, 
            product_id, 
            seller_id, 
            currency,
            selling_price,  
            stock_count, 
            return_policy,
            created_at,
            updated_at FROM marketplace_listings WHERE id=?;`
	row := db.QueryRow(qry, listingID)
	var resp productmgtmodel.MarketplaceListing
	err := row.Scan(&resp.ID,
		&resp.IDBitCount,
		&resp.ProductID,
		&resp.SellerID,
		&resp.Currency,
		&resp.SellingPrice,
		&resp.StockCount,
		&resp.ReturnPolicy,
		&resp.CreatedAt,
		&resp.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errorlib.NewNotFoundError("listing-id=" + listingID)
	}
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	return &resp, nil
}
