package sbp

import (
	"context"
	"encoding/xml"
	"time"

	"espp-mock/entity"

	"github.com/rs/zerolog/log"
)

func (u *SBPUsecase) B21(ctx context.Context, b21 *entity.B21) error {
	const op = "sbp.B21"

	log.Info().
		Str("op", op).
		Str("BizMsgIdr", b21.AppHdr.BizMsgIdr).
		Msg("B21 usecase: сообщение получено")

	// Сериализуем B21 в XML для отправки
	xmlData, err := xml.Marshal(b21)
	if err != nil {
		log.Error().Err(err).Str("op", op).Msg("B21 usecase: ошибка преобразования B21")
		return err
	}

	log.Info().
		Str("op", op).
		Msg("B21 usecase: сообщение сериализовано, отправляется во внешнюю систему")

	// Отправляем B21 во внешнюю систему
	if err := u.ESPP.B21(ctx, string(xmlData)); err != nil {
		log.Error().Err(err).Str("op", op).Msg("B21 usecase: ошибка отправки B21 во внешнюю систему")
		return err
	}

	log.Info().
		Str("op", op).
		Msg("B21 usecase: B21 успешно отправлен во внешнюю систему")

	// Асинхронно отправляем B22
	go func() {
		time.Sleep(u.conf.CallbackDelay)

		// Создаем новый контекст
		newCtx := context.WithValue(context.Background(), "correlationID", ctx.Value("correlationID"))
		newCtx = context.WithValue(newCtx, "transactionNumber", ctx.Value("transactionNumber"))

		// Проверяем наличие transactionNumber
		transactionNumber, ok := ctx.Value("transactionNumber").(string)
		if !ok {
			log.Error().Str("op", op).Msg("B21 usecase: не найден transactionNumber в контексте")
			return
		}

		// Извлекаем fromId и toId из b21.AppHdr
		// Предполагаем, что Fr и To имеют вложенную структуру, как в B22
		fromId := b21.AppHdr.Fr.FIId.FinInstnId.Othr.ID
		toId := b21.AppHdr.To.FIId.FinInstnId.Othr.ID
		bizMsgIdr := b21.AppHdr.BizMsgIdr

		log.Info().
			Str("op", op).
			Str("transactionNumber", transactionNumber).
			Msg("B21 usecase: формируется и отправляется B22")

		response := entity.NewB22Response(transactionNumber, fromId, toId, bizMsgIdr)

		// Отправляем B22 (предполагается, что метод B22 реализован)
		if err := u.ESPP.B22(newCtx, response); err != nil {
			log.Error().Err(err).Str("op", op).Msg("B21 usecase: ошибка отправки B22")
		} else {
			log.Info().
				Str("op", op).
				Str("transactionNumber", transactionNumber).
				Msg("B21 usecase: B22 успешно отправлен")
		}
	}()

	return nil
}
