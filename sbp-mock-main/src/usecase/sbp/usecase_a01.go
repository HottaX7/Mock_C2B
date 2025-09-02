package sbp

import (
	"context"
	"espp-mock/entity"
	"time"

	"github.com/rs/zerolog/log"
)

func (u *SBPUsecase) A01(ctx context.Context, a01 *entity.A01) error {
	transactionNumber := ctx.Value("transactionNumber").(string)

	go func() {
		time.Sleep(u.conf.CallbackDelay)
		// detach context
		newCtx := context.WithValue(context.Background(), "correlationID", ctx.Value("correlationID"))
		newCtx = context.WithValue(newCtx, "transactionNumber", ctx.Value("transactionNumber"))

		a02 := entity.NewA02(transactionNumber, a01)

		err := u.ESPP.A02(newCtx, a02)
		if err != nil {
			log.Error().Err(err).Msgf("failed to send A02")
		}
	}()

	return nil
}
