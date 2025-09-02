package espp

import (
	"context"
)

// B22 — метод для отправки B22 операции
func (a *IPSAdapder) B22(ctx context.Context, response string) error {
	return a.SendBOperation(ctx, "B22", response)
}
