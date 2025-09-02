package sbp

import (
	"context"
	"espp-mock/entity"
	"time"

	"github.com/rs/zerolog/log"
)

func (u *SBPUsecase) C04(ctx context.Context, c04 *entity.C04) error {
	transactionNumber := ctx.Value("transactionNumber").(string)

	go func() {
		time.Sleep(u.conf.CallbackDelay)
		// detach context
		newCtx := context.WithValue(context.Background(), "correlationID", ctx.Value("correlationID"))
		newCtx = context.WithValue(newCtx, "transactionNumber", ctx.Value("transactionNumber"))

		c23 := entity.NewC23(transactionNumber, c04)

		err := u.ESPP.C23(newCtx, c23)
		if err != nil {
			log.Error().Err(err).Msgf("failed to send C23")
		}
	}()

	return nil
}
