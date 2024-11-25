package repository

import (
	"database/sql"

	"go.uber.org/zap"
)

type AllRepository struct {
	CustomerRep CustomerRepoInterface
}

func NewAllRepository(db *sql.DB, log *zap.Logger) AllRepository {
	return AllRepository{
		CustomerRep: NewCustomerRepository(db, log),
	}
}
