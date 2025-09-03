package sbp

import (
	"context"
	"espp-mock/entity"
	"time"

	"github.com/rs/zerolog/log"
)

// B05 строго по сценарию:
// 1. Формируется и отправляется B06
// 2. Дожидаемся успешного ответа (200 OK)
// 3. Формируется и отправляется B21
func (u *SBPUsecase) B05(ctx context.Context, b05 *entity.B05) error {
	transactionNumber := ctx.Value("transactionNumber").(string)

	// Формируем B06
	b06 := entity.NewB06(transactionNumber, b05)
	log.Info().
		Str("transactionNumber", transactionNumber).
		Msg("B05 usecase: формируется и отправляется B06")

	// Отправка B06 с задержкой CallbackDelay
	time.Sleep(u.conf.CallbackDelay)
	if err := u.ESPP.B06(ctx, b06); err != nil {
		log.Error().
			Err(err).
			Str("transactionNumber", transactionNumber).
			Msg("B05 usecase: ошибка отправки B06")
		return err
	}
	log.Info().
		Str("transactionNumber", transactionNumber).
		Msg("B05 usecase: B06 успешно отправлен")

	// Формируем B21
	b21 := entity.NewB21(transactionNumber, b05)
	log.Info().
		Str("transactionNumber", transactionNumber).
		Msg("B05 usecase: формируется и отправляется B21")

	// Отправка B21 с задержкой 2*CallbackDelay
	time.Sleep(2 * u.conf.CallbackDelay)
	if err := u.ESPP.B21(ctx, b21); err != nil {
		log.Error().
			Err(err).
			Str("transactionNumber", transactionNumber).
			Msg("B05 usecase: ошибка отправки B21")
		return err
	}
	log.Info().
		Str("transactionNumber", transactionNumber).
		Msg("B05 usecase: B21 успешно отправлен")

	return nil
}
