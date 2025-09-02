package espp

import (
	"context"
)

func (a *IPSAdapder) C02(ctx context.Context, request string) error {
	return a.sendRequest(ctx, "C02", request)
}
