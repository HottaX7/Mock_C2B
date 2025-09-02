package espp

import (
	"context"
)

func (a *IPSAdapder) C12(ctx context.Context, request string) error {
	return a.sendRequest(ctx, "C12", request)
}
