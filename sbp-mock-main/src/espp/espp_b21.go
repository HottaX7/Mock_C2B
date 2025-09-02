package espp

import (
	"context"
)

func (a *IPSAdapder) B21(ctx context.Context, request string) error {
	return a.SendBOperation(ctx, "B21", request)
}
