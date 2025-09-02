package sbp

import (
	"context"
	"espp-mock/entity"
	"time"

	"github.com/rs/zerolog/log"
)

func (u *SBPUsecase) B05(ctx context.Context, b05 *entity.B05) error {
	transactionNumber := ctx.Value("transactionNumber").(string)

	go func() {
		time.Sleep(u.conf.CallbackDelay)
		// detach context
		newCtx := context.WithValue(context.Background(), "correlationID", ctx.Value("correlationID"))
		newCtx = context.WithValue(newCtx, "transactionNumber", ctx.Value("transactionNumber"))

		b06 := entity.NewB06(transactionNumber, b05)

		log.Info().
			Str("transactionNumber", transactionNumber).
			Msg("B05 usecase: формируется и отправляется B06")

		err := u.ESPP.B06(newCtx, b06)
		if err != nil {
			log.Error().Err(err).Msgf("B05 usecase: ошибка отправки B06")
		} else {
			log.Info().
				Str("transactionNumber", transactionNumber).
				Msg("B05 usecase: B06 успешно отправлен")
		}
	}()

	go func() {
		time.Sleep(2 * u.conf.CallbackDelay)
		// detach context
		newCtx := context.WithValue(context.Background(), "correlationID", ctx.Value("correlationID"))
		newCtx = context.WithValue(newCtx, "transactionNumber", ctx.Value("transactionNumber"))

		b21 := entity.NewB21(transactionNumber, b05)

		log.Info().
			Str("transactionNumber", transactionNumber).
			Msg("B05 usecase: формируется и отправляется B21")

		err := u.ESPP.B21(newCtx, b21)
		if err != nil {
			log.Error().Err(err).Msgf("B05 usecase: ошибка отправки B21")
		} else {
			log.Info().
				Str("transactionNumber", transactionNumber).
				Msg("B05 usecase: B21 успешно отправлен")
		}
	}()

	return nil
}
