package espp

import (
	"context"
)

func (a *IPSAdapder) C21(ctx context.Context, request string) error {
	return a.sendRequest(ctx, "C21", request)
}
