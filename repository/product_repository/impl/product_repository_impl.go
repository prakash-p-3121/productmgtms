package impl

import (
	"database/sql"
	"errors"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/mysqllib"
	"github.com/prakash-p-3121/productmgtmodel"
	"sync"
)

type ProductRepositoryImpl struct {
	ShardConnectionsMap *sync.Map
}

func (repository *ProductRepositoryImpl) CreateProduct(shardID int64,
	idResp *idgenmodel.IDGenResp,
	req *productmgtmodel.ProductCreateReq) errorlib.AppError {
	db, err := mysqllib.RetrieveShardConnectionByShardID(repository.ShardConnectionsMap, shardID)
	if err != nil {
		return errorlib.NewInternalServerError(err.Error())
	}

	qry := `INSERT INTO products (id, 
                      id_bit_count, 
                      name, 
                      description, 
                      currency, 
                      cost_price, 
                      category_id, 
                      media_type,
                      media_path
            );`
	_, err = db.Exec(qry,
		idResp.ID,
		idResp.BitCount,
		*req.Name,
		*req.Description,
		*req.Currency,
		*req.CostPrice,
		*req.CategoryID,
		*req.MediaType,
		*req.MediaPath,
	)
	if err != nil && mysqllib.IsConflictError(err) {
		return errorlib.NewConflictError("name-unique-index-violation")
	}
	if err != nil {
		return errorlib.NewInternalServerError(err.Error())
	}
	return nil
}

func (repository *ProductRepositoryImpl) FindProduct(shardID int64, productID string) (*productmgtmodel.Product, errorlib.AppError) {
	db, err := mysqllib.RetrieveShardConnectionByShardID(repository.ShardConnectionsMap, shardID)
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	qry := `SELECT id, 
			id_bit_count, 
			name, 
			description, 
			currency, 
			cost_price, 
			category_id,
			media_type,
			media_path, 
			created_at, 
			updated_at WHERE id=?;`
	row := db.QueryRow(qry, productID)
	var resp productmgtmodel.Product
	err = row.Scan(&resp.ID,
		&resp.IDBitCount,
		&resp.Name,
		&resp.Description,
		&resp.Currency,
		&resp.CostPrice,
		&resp.CategoryID,
		&resp.MediaType,
		&resp.MediaPath,
		&resp.CreatedAt,
		&resp.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errorlib.NewNotFoundError("product-id=" + productID)
	}
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	return &resp, nil
}
