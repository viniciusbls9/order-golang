package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/viniciusbls9/order-golang/internal/infra/database"
	"github.com/viniciusbls9/order-golang/internal/usecase"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	orderRepository := database.NewOrderRepository(db)

	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "123",
		Price: 10.0,
		Tax:   1.0,
	}

	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	println(output)
}
