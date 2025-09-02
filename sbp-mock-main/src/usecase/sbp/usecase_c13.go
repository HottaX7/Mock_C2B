package sbp

import (
	"context"
	"espp-mock/entity"
	"time"

	"github.com/rs/zerolog/log"
)

func (u *SBPUsecase) C13(ctx context.Context, c13 *entity.C13) error {
	transactionNumber := ctx.Value("transactionNumber").(string)

	go func() {
		time.Sleep(u.conf.CallbackDelay)
		// detach context
		newCtx := context.WithValue(context.Background(), "correlationID", ctx.Value("correlationID"))
		newCtx = context.WithValue(newCtx, "transactionNumber", ctx.Value("transactionNumber"))

		c14 := entity.NewC14(transactionNumber, c13)

		err := u.ESPP.C14(newCtx, c14)
		if err != nil {
			log.Error().Err(err).Msgf("failed to send C14")
		}
	}()

	return nil
}
