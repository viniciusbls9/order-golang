package main

import "github.com/viniciusbls9/order-golang/internal/entity"

func main() {
	order, err := entity.NewOrder("1", 10, 1)
	if err != nil {
		println(err)
	} else {
		println(order.ID)
	}
}
