package main

import "github.com/dr4g0n7ly/AutoTariff-Service/types"

type AggregatorGRPCServer struct {
	types.UnimplementedAggregatorServer
	svc Aggregator
}

func NewAggregatorGRPCServer(svc Aggregator) *AggregatorGRPCServer {
	return &AggregatorGRPCServer{
		svc: svc,
	}
}
func (s *AggregatorGRPCServer) AggregateDistance(req *types.AggregateRequest) error {
	distance := types.Distance{
		OBUID:    int(req.ObuID),
		Distance: req.Distance,
		Unix:     req.Unix,
	}
	return s.svc.AggregateDistance(distance)
}
