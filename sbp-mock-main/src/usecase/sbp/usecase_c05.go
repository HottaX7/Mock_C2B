package sbp

import (
	"context"
	"espp-mock/entity"
	"time"

	"github.com/rs/zerolog/log"
)

func (u *SBPUsecase) C05(ctx context.Context, c05 *entity.C05) error {
	transactionNumber := ctx.Value("transactionNumber").(string)

	go func() {
		time.Sleep(u.conf.CallbackDelay)
		// detach context
		newCtx := context.WithValue(context.Background(), "correlationID", ctx.Value("correlationID"))
		newCtx = context.WithValue(newCtx, "transactionNumber", ctx.Value("transactionNumber"))

		c06 := entity.NewC06(transactionNumber, c05)

		err := u.ESPP.C06(newCtx, c06)
		if err != nil {
			log.Error().Err(err).Msgf("failed to send C06")
		}
	}()

	go func() {
		time.Sleep(2 * u.conf.CallbackDelay)
		// detach context
		newCtx := context.WithValue(context.Background(), "correlationID", ctx.Value("correlationID"))
		newCtx = context.WithValue(newCtx, "transactionNumber", ctx.Value("transactionNumber"))

		c21 := entity.NewC21(transactionNumber, c05)

		err := u.ESPP.C21(newCtx, c21)
		if err != nil {
			log.Error().Err(err).Msgf("failed to send C21")
		}
	}()

	return nil
}
