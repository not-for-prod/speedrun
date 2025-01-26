package peach_repository

import (
	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	"github.com/jmoiron/sqlx"
)

type PeachRepository struct {
	db        *sqlx.DB
	ctxGetter *trmsqlx.CtxGetter
	txManager *manager.Manager
}

func New(
	db *sqlx.DB,
	ctxGetter *trmsqlx.CtxGetter,
	txManager *manager.Manager,
) *PeachRepository {
	return &PeachRepository{
		db:        db,
		ctxGetter: ctxGetter,
		txManager: txManager,
	}
}
