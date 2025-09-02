package sbp

import (
	"context"
	"espp-mock/entity"
	"time"

	"github.com/rs/zerolog/log"
)

func (u *SBPUsecase) C11(ctx context.Context, c11 *entity.C11) error {
	transactionNumber := ctx.Value("transactionNumber").(string)

	go func() {
		time.Sleep(u.conf.CallbackDelay)
		// detach context
		newCtx := context.WithValue(context.Background(), "correlationID", ctx.Value("correlationID"))
		newCtx = context.WithValue(newCtx, "transactionNumber", ctx.Value("transactionNumber"))

		c12 := entity.NewC12(transactionNumber, c11)

		err := u.ESPP.C12(newCtx, c12)
		if err != nil {
			log.Error().Err(err).Msgf("failed to send C12")
		}
	}()

	return nil
}
