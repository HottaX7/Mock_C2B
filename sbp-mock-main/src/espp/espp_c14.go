package espp

import (
	"context"
)

func (a *IPSAdapder) C14(ctx context.Context, request string) error {
	return a.sendRequest(ctx, "C14", request)
}
