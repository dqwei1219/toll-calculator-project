package client

import (
	"context"

	"github.com/dqwei1219/toll-calculator-project/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	Endpoint string
	client   types.DistAggregatorClient
}

func NewGRPCClient(endpoint string) (*GRPCClient, error) {
	// dial endpoint with options
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := types.NewDistAggregatorClient(conn)
	return &GRPCClient{
		Endpoint: endpoint,
		client:   c,
	}, nil
}

func (c *GRPCClient) AggregateDist(ctx context.Context, req *types.AggregateDistReq) error {
	_, err := c.client.AggregateDist(ctx, req)
	return err
}
