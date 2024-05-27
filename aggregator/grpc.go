package main

import (
	"context"

	"github.com/dr4g0n7ly/AutoTariff-Service/types"
)

type AggregatorGRPCServer struct {
	types.UnimplementedAggregatorServer
	svc Aggregator
}

func NewAggregatorGRPCServer(svc Aggregator) *AggregatorGRPCServer {
	return &AggregatorGRPCServer{
		svc: svc,
	}
}
func (s *AggregatorGRPCServer) Aggregate(ctx context.Context, req *types.AggregateRequest) (*types.None, error) {
	distance := types.Distance{
		OBUID:    int(req.ObuID),
		Distance: req.Distance,
		Unix:     req.Unix,
	}
	return &types.None{}, s.svc.AggregateDistance(distance)
}
