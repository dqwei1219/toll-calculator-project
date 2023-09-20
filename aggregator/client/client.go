package client

import (
	"context"

	"github.com/dqwei1219/toll-calculator-project/types"
)

type Client interface {
	AggregateDist(context.Context, *types.AggregateDistReq) error
}
