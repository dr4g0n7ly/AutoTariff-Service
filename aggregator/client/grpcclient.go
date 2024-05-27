package client

import (
	"github.com/dr4g0n7ly/AutoTariff-Service/types"
	"google.golang.org/grpc"
)

type GRPCClient struct {
	endpoint string
	types.AggregatorClient
}

func NewGRPCClient(endpoint string) (*GRPCClient, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := types.NewAggregatorClient(conn)
	return &GRPCClient{
		endpoint:         endpoint,
		AggregatorClient: c,
	}, nil
}
