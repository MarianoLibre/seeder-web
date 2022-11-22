package main

/*
Seeding order:
	- localities
	- carries
	- warehouses
	- buyers
	- employees
	- sellers
	- sections
	- products
	- product_batches
	- inbound_orders
	- product_records
	- purchase_orders
*/

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"syreclabs.com/go/faker"
)

type Repository interface {
	DropTables() error
	CreateTables() error
	SeedWarehouses(n int) error
	SeedBuyers(n int) error
	SeedSellers(n int) error
	SeedEmployees(n int) error
	SeedProducts(n int) error
	SeedSections(n int) error
	SeedLocalities(n int) error
	SeedCarries(n int) error
	SeedProductBatches(n int) error
	SeedProductRecords(n int) error
	SeedInboundOrders(n int) error
	SeedPurchaseOrders(n int) error
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
		purchase_orders_count) VALUES (?, ?, ?, ?)`
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
		expiration_rate,
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
		warehouse_id) VALUES (?, ?, ?, ?, ?)`
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
func (r *repository) SeedWarehouses(n int) error {
	codeGen := newCodeGen(3)

	for i := 0; i < n; i++ {
		stmt, err := r.db.Prepare(INSERT_WAREHOUSES)
		if err != nil {
			return err
		}

		res, err := stmt.Exec(
			faker.Address().StreetAddress(),
			faker.PhoneNumber().CellPhone(),
			codeGen(),
			faker.Number().NumberInt(2),
			faker.Number().Between(-100, 32),
		)

		if err != nil {
			return err
		}

		if ra, err := res.RowsAffected(); err != nil {
			return err
		} else if ra < 1 {
			return errors.New("[repository-seeder-warehouses] Oops! I fucked up!!!")
		}

	}
	return nil
}

func (r *repository) SeedBuyers(n int) error {
	for i := 0; i < n; i++ {
		stmt, err := r.db.Prepare(INSERT_BUYERS)
		if err != nil {
			return err
		}

		res, err := stmt.Exec(
			faker.Number().Between(20_000_000, 50_000_000),
			faker.Name().FirstName(),
			faker.Name().LastName(),
			0,
		)

		if err != nil {
			return err
		}

		if ra, err := res.RowsAffected(); err != nil {
			return err
		} else if ra < 1 {
			return errors.New("[repository-seeder-buyers] Oops! I fucked up!!!")
		}
	}

	return nil
}

func (r *repository) SeedSellers(n int) error {
	for i := 0; i < n; i++ {
		stmt, err := r.db.Prepare(INSERT_SELLERS)
		if err != nil {
			return err
		}

		res, err := stmt.Exec(
			i+1,
			faker.Company().Name(),
			faker.Address().StreetAddress(),
			faker.PhoneNumber().CellPhone(),
			rand.Intn(n)+1,
		)

		if err != nil {
			return err
		}

		if ra, err := res.RowsAffected(); err != nil {
			return err
		} else if ra < 1 {
			return errors.New("[repository-seeder-sellers] WTF!")
		}
	}
	return nil
}

func (r *repository) SeedEmployees(n int) error {
	for i := 0; i < n; i++ {
		stmt, err := r.db.Prepare(INSERT_EMPLOYEES)
		if err != nil {
			return err
		}

		res, err := stmt.Exec(
			faker.Number().Between(20_000_000, 50_000_000),
			faker.Name().FirstName(),
			faker.Name().LastName(),
			rand.Intn(n)+1,
		)

		if err != nil {
			return err
		}

		if ra, err := res.RowsAffected(); err != nil {
			return err
		} else if ra < 1 {
			return errors.New("[repository-seeder-employees] WTF!")
		}
	}
	return nil
}

func (r *repository) SeedProducts(n int) error {
	codeGen := newCodeGen(3)
	for i := 0; i < n; i++ {
		stmt, err := r.db.Prepare(INSERT_PRODUCTS)
		if err != nil {
			return err
		}

		res, err := stmt.Exec(
			faker.Commerce().ProductName(), // description
			faker.Number().Decimal(3, 2),   // expiration_rate
			faker.Number().Decimal(3, 2),   // freezing_rate
			faker.Number().Decimal(3, 2),   // height
			faker.Number().Decimal(3, 2),   // lenght
			faker.Number().Decimal(3, 2),   // netweight
			codeGen(),                      // product_code
			faker.Number().Decimal(3, 2),   // recommended_freezing_temperature
			faker.Number().Decimal(3, 2),   // width
			rand.Intn(n)+1,                 // id_product_type
			rand.Intn(n)+1,                 // id_seller
		)

		if err != nil {
			return err
		}

		if ra, err := res.RowsAffected(); err != nil {
			return err
		} else if ra < 1 {
			return errors.New("[repository-seeder-products] WTF!")
		}
	}
	return nil
}

func (r *repository) SeedSections(n int) error {
	for i := 0; i < n; i++ {
		stmt, err := r.db.Prepare(INSERT_SECTIONS)
		if err != nil {
			return err
		}

		sectionNumber := i + 1
		currentTemp := faker.Number().Between(-10, 45)
		minTemp := faker.Number().Between(-10, 45)
		minCap := faker.Number().Between(1, 50)
		maxCap := faker.Number().Between(51, 500)
		min, _ := strconv.Atoi(minCap)
		max, _ := strconv.Atoi(maxCap)
		currentCap := faker.Number().Between(min, max)
		warehoseId := rand.Intn(n) + 1
		idProdType := rand.Intn(n) + 1

		res, err := stmt.Exec(
			sectionNumber,
			currentTemp,
			minTemp,
			currentCap,
			minCap,
			maxCap,
			warehoseId,
			idProdType,
		)

		if err != nil {
			return err
		}

		if ra, err := res.RowsAffected(); err != nil {
			return err
		} else if ra < 1 {
			return errors.New("[repository-seeder-employees] WTF!")
		}
	}
	return nil
}

func (r *repository) SeedLocalities(n int) error {
	var zips []string
	for i := 0; i < n; i++ {
		stmt, err := r.db.Prepare(INSERT_LOCALITIES)
		if err != nil {
			return err
		}

		var zipCode string
		var address faker.FakeAddress
		for {
			// make sure we've got a new zip code (UNIQUE)
			address = faker.Address()
			zipCode = address.ZipCode()
			if isIn(zipCode, zips) {
				continue
			}
			break
		}
		zips = append(zips, zipCode)

		res, err := stmt.Exec(
			address.City(),
			address.State(),
			address.Country(),
			zipCode,
		)

		if err != nil {
			return err
		}

		if ra, err := res.RowsAffected(); err != nil {
			return err
		} else if ra < 1 {
			return errors.New("[localities] seeder failed...")
		}
	}
	return nil
}

func (r *repository) SeedCarries(n int) error {
	for i := 0; i < n; i++ {
		stmt, err := r.db.Prepare(INSERT_CARRIES)
		if err != nil {
			return err
		}

		res, err := stmt.Exec(
			i+1,
			faker.Company().Name(),
			faker.Address().StreetAddress(),
			faker.PhoneNumber().CellPhone(),
			rand.Intn(n)+1,
		)

		if err != nil {
			return err
		}

		if ra, err := res.RowsAffected(); err != nil {
			return err
		} else if ra < 1 {
			return errors.New("[carries] seeder failed...")
		}
	}
	return nil
}

func (r *repository) SeedProductBatches(n int) error {
	for i := 0; i < n; i++ {
		stmt, err := r.db.Prepare(INSERT_PRODUCT_BATCHES)
		if err != nil {
			return err
		}

		res, err := stmt.Exec(
			i+1,                             // batch_number
			faker.Number().Between(1, 100),  // current_quantity
			faker.Number().Between(-10, 45), // current_temperature
			faker.Date().Forward(1_000_000), // due_date
			faker.Number().Between(1, 100),  // initial_quantity
			faker.Date().Birthday(1, 18),    // manufacturing_date
			faker.Number().Between(0, 23),   // manufacturing_hour
			faker.Number().Between(-10, 22), // minimum_temperature
			rand.Intn(n)+1,                  // product_id
			rand.Intn(n)+1,                  // section_id
		)

		if err != nil {
			return err
		}

		if ra, err := res.RowsAffected(); err != nil {
			return err
		} else if ra < 1 {
			return errors.New("[product_batches] seeder failed...")
		}
	}
	return nil
}

func (r *repository) SeedProductRecords(n int) error {
	for i := 0; i < n; i++ {
		stmt, err := r.db.Prepare(INSERT_PRODUCT_RECORDS)
		if err != nil {
			return err
		}

		purchasePrice := faker.Number().Decimal(4, 2)
		base, err := strconv.ParseFloat(purchasePrice, 32)
		if err != nil {
			return err
		}
		res, err := stmt.Exec(
			faker.Date().Backward(1_000_000), // last_updated_date
			purchasePrice,                    // purchase_price
			base*2,                           // sale_price
			rand.Intn(n)+1,
		)

		if err != nil {
			return err
		}

		if ra, err := res.RowsAffected(); err != nil {
			return err
		} else if ra < 1 {
			return errors.New("[product_records] seeder failed...")
		}
	}
	return nil
}

func (r *repository) SeedInboundOrders(n int) error {
	for i := 0; i < n; i++ {
		stmt, err := r.db.Prepare(INSERT_INBOUND_ORDERS)
		if err != nil {
			return err
		}

		res, err := stmt.Exec(
			faker.Date().Backward(1_000_000), // order_date
			fmt.Sprintf("order#%d", i+1),     // order_number
			rand.Intn(n)+1,                   // employee_id
			rand.Intn(n)+1,                   // product_batch_id
			rand.Intn(n)+1,                   // warehouse_id
		)

		if err != nil {
			return err
		}

		if ra, err := res.RowsAffected(); err != nil {
			return err
		} else if ra < 1 {
			return errors.New("[inbound_orders] seeder failed...")
		}
	}
	return nil
}

func (r *repository) SeedPurchaseOrders(n int) error {
	codeGen := newCodeGen(6)

	for i := 0; i < n; i++ {
		stmt, err := r.db.Prepare(INSERT_PURCHASE_ORDERS)
		if err != nil {
			return err
		}

		res, err := stmt.Exec(
			i+1, // order_number
			faker.Date().Backward(
				time.Duration(rand.Intn(1_000_000_000_000)),
			), // order_date
			codeGen(),      // tracking_code
			rand.Intn(n)+1, // buyer_id
			rand.Intn(n)+1, // product_record_id
			rand.Intn(n)+1, // order_status_id
		)

		if err != nil {
			return err
		}

		if ra, err := res.RowsAffected(); err != nil {
			return err
		} else if ra < 1 {
			return errors.New("[purchase_orders] seeder failed...")
		}
	}
	return nil
}

// Create a code generator func that return a unique alphabetic code (allcaps) of the specified length
func newCodeGen(length int) func() string {
	var newGenerator []int
	for i := 0; i < length; i++ {
		newGenerator = append(newGenerator, 65)
	}
	return func() string {
		code := ""
		for i := 0; i < len(newGenerator); i++ {
			code = fmt.Sprintf("%c", newGenerator[i]) + code
		}

		for i, v := range newGenerator {
			if v < 90 {
				newGenerator[i] = v + 1
				break
			} else {
				newGenerator[i] = 65
			}
		}

		return code
	}
}

func isIn(word string, wordList []string) bool {
	for _, w := range wordList {
		if w == word {
			return true
		}
	}
	return false
}
