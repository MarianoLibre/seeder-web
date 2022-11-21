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
	INSERT_WAREHOUSES = `INSERT INTO warehouses (
		address,
		telephone,
		warehouse_code,
		minimum_capacity,
		minimum_temperature) VALUES (?, ?, ?, ?, ?)`
	INSERT_BUYERS = `INSERT INTO buyers (
		card_number_id,
		first_name,
		last_name,
		purchase_order_count) VALUES (?, ?, ?, ?)`
	INSERT_SELLERS = `INSERT INTO sellers (
		cid,
		company_name,
		address,
		telephone,
		locality_id) VALUES (?, ?, ?, ?, ?)`
	INSERT_SECTIONS = `INSERT INTO sections (
		section_number,
		current_temperature,
		minimum_temperature,
		current_capacity,
		minimum_capacity,
		maximum_capacity,
		warehouse_id,
		id_product_type) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	INSERT_EMPLOYEES = `INSERT INTO employees (
		card_number_id,
		first_name,
		last_name,
		warehouse_id) VALUES (?, ?, ?, ?)`
	INSERT_PRODUCTS = `INSERT INTO products (
		description,
		expiration_date,
		freezing_rate,
		height,
		lenght,
		netweight,
		product_code,
		recommended_freezing_temperature,
		width,
		id_product_type,
		id_seller) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	INSERT_LOCALITIES = `INSERT INTO localities (
		locality_name,
		province_name,
		country_name,
		zip_code) VALUES (?, ?, ?, ?)`
	INSERT_CARRIES = `INSERT INTO carries (
		cid,
		company_name,
		address,
		telephone,
		locality_id) VALUES (?, ?, ?, ?, ?)`
	INSERT_INBOUND_ORDERS = `INSERT INTO inbound_orders (
		order_date,
		order_number,
		employee_id,
		product_batch_id,
		warehouse_id) VALUES (?, ?, ?, ?, ?`
	INSERT_PRODUCT_BATCHES = `INSERT INTO product_batches (
		batch_number,
		current_quantity,
		current_temperature,
		due_date,
		initial_quantity,
		manufacturing_date,
		manufacturing_hour,
		minimum_temperature,
		product_id,
		section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	INSERT_PRODUCT_RECORDS = `INSERT INTO product_records (
		last_updated_date,
		purchase_price,
		sale_price,
		product_id) VALUES (?, ?, ?, ?)`
	INSERT_PURCHASE_ORDERS = `INSERT INTO purchase_orders (
		order_number,
		order_date,
		tracking_code,
		buyer_id,
		product_record_id,
		order_status_id) VALUES (?, ?, ?, ?, ?, ?)`
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

// Seeders
func SeedWarehouses(n int) error {

	return nil
}

func SeedBuyers(n int) error {

	return nil
}

func SeedSellers(n int) error {

	return nil
}

func SeedEmployees(n int) error {

	return nil
}

func SeedProducts(n int) error {

	return nil
}

func SeedSections(n int) error {

	return nil
}

func SeedLocalities(n int) error {

	return nil
}

func SeedCarries(n int) error {

	return nil
}

func SeedProductBatches(n int) error {

	return nil
}

func SeedProductOrders(n int) error {

	return nil
}

func SeedInboundOrders(n int) error {

	return nil
}

func SeedPurchaseOrders(n int) error {

	return nil
}

