package db

import (
	"database/sql"
)

// SQLStore provides all functions to excute db queries and transactioncd
type Store interface {
	Querier
}

// SQLStore provides all functions to excute SQL queries and transaction
type SQLStore struct {
	*Queries // excute SQL queries
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		Queries: New(db), //New(db)返回*Queries
	}
}
