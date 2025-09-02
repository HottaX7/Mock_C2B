package sbp

import (
	"context"
	"encoding/xml"
	"espp-mock/entity"
	"espp-mock/metrics"
	"io"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
)

/*
func (h *Handler) HandleB05(w http.ResponseWriter, r *http.Request) {
	const op = "sbp.HandleB05"
	metrics.AddRequest("B05")

	correlationID := r.Header.Get("X-Correlation-ID")
	transactionNumber := "X32123420144" + strconv.Itoa(rand.Intn(9999999999-1000000000)+1000000000) + "IA00000001"

	ctx := context.WithValue(r.Context(), "correlationID", correlationID)
	ctx = context.WithValue(ctx, "transactionNumber", transactionNumber)

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error().Err(err).Str("op", op).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b05Request := &entity.B05Request{}
	if err := xml.Unmarshal(bodyBytes, b05Request); err != nil {
		log.Error().Err(err).Str("op", op).Str("body", string(bodyBytes)).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Вызов бизнес-логики, если нужно (например, отправка B21)
	if err := h.SBPUsecase.B05(ctx, b05Request); err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Формируем B06 из шаблона
	responseXML := entity.NewB06(transactionNumber, b05Request)

	// Отправляем клиенту
	w.Header().Set("Content-Type", "application/xml")
	w.Header().Set("X-Sbp-Trn-Num", transactionNumber)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseXML))
}
*/

func (h *Handler) HandleB05(w http.ResponseWriter, r *http.Request) {
	const op = "sbp.HandleB05"
	log := log.With().Str("op", op).Logger()

	metrics.AddRequest("B05")

	// Получаем заголовки
	correlationID := r.Header.Get("X-Correlation-ID")
	transactionNumber := "X32123420144" + strconv.Itoa(rand.Intn(9999999999-1000000000)+1000000000) + "IA00000001"

	// Устанавливаем контекст
	ctx := context.WithValue(r.Context(), "correlationID", correlationID)
	ctx = context.WithValue(ctx, "transactionNumber", transactionNumber)

	log.Info().
		Str("correlationID", correlationID).
		Str("transactionNumber", transactionNumber).
		Msg("B05: сообщение получено")

	// Читаем тело
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read request body")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Парсим в существующую структуру B05
	var b05 entity.B05
	if err := xml.Unmarshal(bodyBytes, &b05); err != nil {
		log.Error().Err(err).Str("body", string(bodyBytes)).Msg("Failed to parse B05")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	log.Info().
		Str("transactionNumber", transactionNumber).
		Msg("B05: сообщение успешно разобрано и принято к обработке")

	// Вызываем бизнес-логику (например, асинхронную отправку B21/B22)
	if err := h.SBPUsecase.B05(ctx, &b05); err != nil {
		log.Error().Err(err).Msg("Business logic failed")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Info().
		Str("transactionNumber", transactionNumber).
		Msg("B05: сообщение успешно обработано")

	// Формируем B06 на основе данных из B05
	responseXML := entity.NewB06(transactionNumber, &b05)

	// Логируем ответ
	log.Info().
		Str("transactionNumber", transactionNumber).
		Str("response_B06", responseXML).
		Msg("B05: ответ сформирован и отправляется клиенту")

	// Отправляем клиенту
	w.Header().Set("Content-Type", "application/xml")
	w.Header().Set("X-Sbp-Trn-Num", transactionNumber)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseXML))

	log.Info().
		Str("transactionNumber", transactionNumber).
		Msg("B06: ответ успешно отправлен клиенту")
}
