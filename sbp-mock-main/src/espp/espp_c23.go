package espp

import (
	"context"
)

func (a *IPSAdapder) C23(ctx context.Context, request string) error {
	return a.sendRequest(ctx, "C23", request)
}
