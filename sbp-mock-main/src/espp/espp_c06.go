package espp

import (
	"context"
)

func (a *IPSAdapder) C06(ctx context.Context, request string) error {
	return a.sendRequest(ctx, "C06", request)
}
