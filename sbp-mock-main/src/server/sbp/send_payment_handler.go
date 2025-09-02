// send_payment_handler.go

package sbp

import (
	"encoding/xml"
	"io"
	"net/http"

	"espp-mock/entity"

	"github.com/rs/zerolog/log"
)

func (h *Handler) HandleSendPayment(w http.ResponseWriter, r *http.Request) {
	const op = "handler.HandleSendPayment"
	logger := log.With().Str("op", op).Logger()

	logger.Info().Msg("SendPayment: запрос получен")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error().Err(err).Msg("SendPayment: ошибка чтения тела запроса")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var req entity.SendPaymentRequest
	if err := xml.Unmarshal(body, &req); err != nil {
		logger.Error().Err(err).Str("body", string(body)).Msg("SendPayment: ошибка парсинга XML")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	logger.Info().Interface("request", req).Msg("SendPayment: запрос успешно разобран и принят к обработке")

	// Здесь могла бы быть бизнес-логика
	logger.Info().Msg("SendPayment: запрос успешно обработан")

	// Формируем ответ
	response := entity.SendPaymentResponse{
		PaymentId:           "1259891",
		PaymentTime:         "2024-11-21T07:52:18.2+03:00",
		State:               "5",
		StateComment:        "Завершена",
		ProcessingErrorCode: "0",
		Currency:            "RUR",
	}

	respBody, _ := xml.MarshalIndent(response, "", "  ")

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)

	logger.Info().Msg("SendPayment: ответ успешно отправлен клиенту")
}
