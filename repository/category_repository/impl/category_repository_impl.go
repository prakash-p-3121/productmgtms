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

type CategoryRepositoryImpl struct {
	ShardConnectionsMap *sync.Map
}

func (repository *CategoryRepositoryImpl) CreateCategory(shardID int64,
	idResp *idgenmodel.IDGenResp,
	req *productmgtmodel.CategoryCreateReq) errorlib.AppError {
	db, err := mysqllib.RetrieveShardConnectionByShardID(repository.ShardConnectionsMap, shardID)
	if err != nil {
		return errorlib.NewInternalServerError(err.Error())
	}

	qry := `INSERT INTO categories (id, id_bit_count, name) VALUES (?, ?, ?);`
	_, err = db.Exec(qry, idResp.ID, idResp.BitCount, *req.Name)
	if err != nil && mysqllib.IsConflictError(err) {
		return errorlib.NewConflictError("name-unique-constraint-violation")
	}
	if err != nil {
		return errorlib.NewInternalServerError(err.Error())
	}
	return nil
}

func (repository *CategoryRepositoryImpl) FindCategory(shardID int64,
	categoryID string) (*productmgtmodel.Category, errorlib.AppError) {
	db, err := mysqllib.RetrieveShardConnectionByShardID(repository.ShardConnectionsMap, shardID)
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}

	qry := `SELECT id, id_bit_count, name, created_at, updated_at FROM categories WHERE id=?;`
	row := db.QueryRow(qry, categoryID)
	var resp productmgtmodel.Category
	err = row.Scan(&resp.ID, &resp.IDBitCount, &resp.Name, &resp.CreatedAt, &resp.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errorlib.NewNotFoundError("categoryID=" + categoryID)
	}
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}
	return &resp, nil
}
