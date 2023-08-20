package graph

import (
	"github.com/jmoiron/sqlx"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func NewResolver(db *sqlx.DB) *Resolver {
	return &Resolver{db: db}
}

type Resolver struct {
	db *sqlx.DB
}
