package main

import (
	"database/sql"
	"io/ioutil"
	"strings"
)


type Repository interface {
	DropTables() error
	CreateTables() error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

const (
	DROP_TABLES = `DROP TABLE IF EXISTS
		carries,
		inbound_orders,
		product_batches,
		purchase_orders,
		product_records,
		products,
		buyers,
		employees,
		sellers,
		localities,
		warehouses,
		sections
	`
)


func (r *repository) DropTables() error {
	_, err := r.db.Exec(DROP_TABLES)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) CreateTables() error {
	data, err := ioutil.ReadFile("./db.sql")
	if err != nil {
		return err
	}

	queryList := strings.Split(strings.Trim(string(data), ";"), ";")
	for _, q := range queryList {
		query := strings.TrimLeft(q, "\n")
		if query != "" {
			query = strings.ReplaceAll(query, "\n", " ")
			_, err := r.db.Exec(query)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
