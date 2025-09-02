package sbp

import (
	"context"
	"espp-mock/entity"
	"time"

	"github.com/rs/zerolog/log"
)

func (u *SBPUsecase) C01(ctx context.Context, c01 *entity.C01) error {
	transactionNumber := ctx.Value("transactionNumber").(string)

	go func() {
		time.Sleep(u.conf.CallbackDelay)
		// detach context
		newCtx := context.WithValue(context.Background(), "correlationID", ctx.Value("correlationID"))
		newCtx = context.WithValue(newCtx, "transactionNumber", ctx.Value("transactionNumber"))

		c02 := entity.NewC02(transactionNumber, c01)

		err := u.ESPP.C02(newCtx, c02)
		if err != nil {
			log.Error().Err(err).Msgf("failed to send C02")
		}
	}()

	return nil
}
