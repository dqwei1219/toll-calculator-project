package main

import (
	"context"

	"github.com/dqwei1219/toll-calculator-project/types"
)

type GRPCServer struct {
	types.UnimplementedDistAggregatorServer
	svc Aggregator
}

func NewGRPCServer(svc Aggregator) *GRPCServer {
	return &GRPCServer{svc: svc}
}

func (s *GRPCServer) AggregateDist(ctx context.Context, req *types.AggregateDistReq) (*types.None, error) {
	distance := types.Distance{
		UnitId: int(req.UnitId),
		Value:  float64(req.Value),
		Unix:   req.Unix,
	}
	return &types.None{}, s.svc.AggregateDistance(distance)
}
