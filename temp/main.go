package main

import (
	"context"
	"log"
	"time"

	"github.com/dqwei1219/toll-calculator-project/aggregator/client"
	"github.com/dqwei1219/toll-calculator-project/types"
)

func main() {
	c, err := client.NewGRPCClient(":3001")
	if err != nil {
		log.Fatal(err)
	}
	if err := c.AggregateDist(context.Background(), &types.AggregateDistReq{
		UnitId: 1,
		Value:  1.0,
		Unix:   time.Now().Unix(),
	}); err != nil {
		log.Fatal(err)
	}
}
