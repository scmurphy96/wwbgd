package main

import (
	"context"
	"log"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func main() {
	c := polygon.New("qhMVh4IRHFYn7gus5BmmUrd8JU7jVVD7")

	iter := c.VX.ListStockFinancials(context.Background(), models.ListStockFinancialsParams{}.WithTicker("AAPL"))

	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
			iter.Err()
	}
}
