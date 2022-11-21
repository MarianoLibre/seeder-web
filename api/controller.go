package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Seeder struct {
	repository Repository
}

func NewSeeder(r Repository) *Seeder {
	return &Seeder{
		repository: r,
	}
}

func (s *Seeder) SayHi() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	}
}

func (s *Seeder) Drop() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := s.repository.DropTables()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, "All tables dropped")
	}
}

func (s *Seeder) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := s.repository.CreateTables()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, "All tables created")
		return
	}
}

// Seeders
func (s *Seeder) SeedWarehouses() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		qty := ctx.Query("qty")
		if qty == "" {
			qty = "10"
		}

		n, err := strconv.ParseInt(qty, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Invalid quantity")
			return
		}

		err = s.repository.SeedWarehouses(int(n))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "Oops! It seems I fucked up!")
			return
		}

		ctx.JSON(http.StatusCreated, "table 'warehouses': SEEDED")
	}
}

func (s *Seeder) SeedBuyers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		qty := ctx.Query("qty")
		if qty == "" {
			qty = "10"
		}

		n, err := strconv.ParseInt(qty, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Invalid quantity")
			return
		}

		err = s.repository.SeedBuyers(int(n))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "Oops! It seems I fucked up!")
			return
		}

		ctx.JSON(http.StatusCreated, "table 'buyers': SEEDED")

	}
}

func (s *Seeder) SeedSellers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		qty := ctx.Query("qty")
		if qty == "" {
			qty = "10"
		}

		n, err := strconv.ParseInt(qty, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Invalid quantity")
			return
		}

		err = s.repository.SeedSellers(int(n))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "Oops! It seems I fucked up!")
			return
		}

		ctx.JSON(http.StatusCreated, "table 'sellers': SEEDED")

	}
}

func (s *Seeder) SeedEmployees() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		qty := ctx.Query("qty")
		if qty == "" {
			qty = "10"
		}

		n, err := strconv.ParseInt(qty, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Invalid quantity")
			return
		}

		err = s.repository.SeedEmployees(int(n))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "Oops! It seems I fucked up!")
			return
		}

		ctx.JSON(http.StatusCreated, "table 'employees': SEEDED")

	}
}

func (s *Seeder) SeedProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		qty := ctx.Query("qty")
		if qty == "" {
			qty = "10"
		}

		n, err := strconv.ParseInt(qty, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Invalid quantity")
			return
		}

		err = s.repository.SeedProducts(int(n))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "Oops! It seems I fucked up!")
			return
		}

		ctx.JSON(http.StatusCreated, "table 'products': SEEDED")

	}
}

func (s *Seeder) SeedSections() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		qty := ctx.Query("qty")
		if qty == "" {
			qty = "10"
		}

		n, err := strconv.ParseInt(qty, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Invalid quantity")
			return
		}

		err = s.repository.SeedSections(int(n))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "Oops! It seems I fucked up!")
			return
		}

		ctx.JSON(http.StatusCreated, "table 'sections': SEEDED")

	}
}

func (s *Seeder) SeedLocalities() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		qty := ctx.Query("qty")
		if qty == "" {
			qty = "10"
		}

		n, err := strconv.ParseInt(qty, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Invalid quantity")
			return
		}

		err = s.repository.SeedLocalities(int(n))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "Oops! It seems I fucked up!")
			return
		}

		ctx.JSON(http.StatusCreated, "table 'localities': SEEDED")

	}
}

func (s *Seeder) SeedCarries() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		qty := ctx.Query("qty")
		if qty == "" {
			qty = "10"
		}

		n, err := strconv.ParseInt(qty, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Invalid quantity")
			return
		}

		err = s.repository.SeedCarries(int(n))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "Oops! It seems I fucked up!")
			return
		}

		ctx.JSON(http.StatusCreated, "table 'carries': SEEDED")

	}
}

func (s *Seeder) SeedProductBatches() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		qty := ctx.Query("qty")
		if qty == "" {
			qty = "10"
		}

		n, err := strconv.ParseInt(qty, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Invalid quantity")
			return
		}

		err = s.repository.SeedProductBatches(int(n))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "Oops! It seems I fucked up!")
			return
		}

		ctx.JSON(http.StatusCreated, "table 'product_batches': SEEDED")

	}
}

func (s *Seeder) SeedProductRecords() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		qty := ctx.Query("qty")
		if qty == "" {
			qty = "10"
		}

		n, err := strconv.ParseInt(qty, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Invalid quantity")
			return
		}

		err = s.repository.SeedProductRecords(int(n))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "Oops! It seems I fucked up!")
			return
		}

		ctx.JSON(http.StatusCreated, "table 'product_records': SEEDED")

	}
}

func (s *Seeder) SeedInboundOrders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		qty := ctx.Query("qty")
		if qty == "" {
			qty = "10"
		}

		n, err := strconv.ParseInt(qty, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Invalid quantity")
			return
		}

		err = s.repository.SeedInboundOrders(int(n))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "Oops! It seems I fucked up!")
			return
		}

		ctx.JSON(http.StatusCreated, "table 'inbound_orders': SEEDED")

	}
}

func (s *Seeder) SeedPurchaseOrders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		qty := ctx.Query("qty")
		if qty == "" {
			qty = "10"
		}

		n, err := strconv.ParseInt(qty, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "Invalid quantity")
			return
		}

		err = s.repository.SeedPurchaseOrders(int(n))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "Oops! It seems I fucked up!")
			return
		}

		ctx.JSON(http.StatusCreated, "table 'purchase_orders': SEEDED")

	}
}
