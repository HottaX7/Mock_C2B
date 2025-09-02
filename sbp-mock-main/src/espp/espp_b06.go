package espp

import (
	"context"
)

func (a *IPSAdapder) B06(ctx context.Context, request string) error {
	return a.SendBOperation(ctx, "B06", request)
}
