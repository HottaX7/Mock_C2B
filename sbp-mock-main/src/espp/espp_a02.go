package espp

import (
	"context"
)

func (a *IPSAdapder) A02(ctx context.Context, request string) error {
	return a.sendRequest(ctx, "A02", request)
}
