package client

import (
	"context"

	"github.com/dr4g0n7ly/AutoTariff-Service/types"
)

type Client interface {
	Aggregate(context.Context, *types.AggregateRequest) error
}
