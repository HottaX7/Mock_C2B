package sbp

import (
	"encoding/json"
	"espp-mock/metrics"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func (h *Handler) HandlePaymentData(w http.ResponseWriter, r *http.Request) {
	uniQrID := chi.URLParam(r, "id")         // Получаем {id} из маршрута
	amountStr := r.URL.Query().Get("amount") // Получаем amount из query-параметров

	metrics.AddRequest("C2BQRD")

	log.Info().
		Str("uniQrId", uniQrID).
		Str("amountStr", amountStr).
		Msg("C2B: запрос получен")

	if amountStr == "" {
		amountStr = "1100"
	}

	// Заменяем запятую на точку, если это необходимо
	amountStr = strings.ReplaceAll(amountStr, ",", ".")
	log.Info().Str("amountStr", amountStr).Msg("C2B: amount после обработки")

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		log.Error().Err(err).Str("amountStr", amountStr).Msg("C2B: некорректный формат amount")
		http.Error(w, `{"error": "invalid amount format"}`, http.StatusBadRequest)
		return
	}
	log.Info().Str("uniQrId", uniQrID).Float64("amount", amount).Msg("C2B: amount успешно преобразован")

	// Формируем JSON-ответ
	response := map[string]interface{}{
		"code":    "RQ00000",
		"message": "Запрос обработан успешно",
		"data": map[string]interface{}{
			"uniQrId": uniQrID,
			/*"uniQrId":        "BS10006LQEGL0BL78D9R7SO6O93SGJVN", //жестко зашиваем QR*/
			"uniQrType": "02",
			"scenario":  "C2B",
			"legalName": "ООО \"ТД АМТ\"",
			"memberId":  "120000000020",
			"brandName": "тест_2",
			/*"amount":         amount, // Отправляем amount (либо переданный, либо дефолтный)*/
			"amount":         strconv.FormatFloat(amount, 'f', -1, 64),
			"paymentPurpose": "Rath, Oberbrunner and Howell",
			"address":        "Уфа",
			"mcc":            "5047",
			"ogrn":           "1200200041360",
			"inn":            "0277950370",
			"redirectUrl":    nil,
			"agentId":        "A11000000072",
			"merchantId":     "MA0000018139",
			"countryCode":    "RU",
			"receiverBic":    "044525111",
			"responseId":     "79f9c5fe3af74300b6dd13fea2c09d27",
			"paramsId":       nil,
			"paymentMethodData": []map[string]interface{}{
				{
					"paymentServiceId": "PS0000000001",
					"version":          "v1",
					"additionalData": map[string]string{
						"fraudScore":      "FFFFFDDDDDDDDDDD",
						"receiverAccount": "40702810041000001132",
					},
				},
			},
		},
	}

	log.Info().
		Str("uniQrId", uniQrID).
		Float64("amount", amount).
		Msg("C2B: ответ сформирован")

	responseBytes, err := json.Marshal(response)
	if err != nil {
		log.Error().Err(err).Msg("C2B: ошибка сериализации JSON-ответа")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBytes)

	log.Info().
		Str("uniQrId", uniQrID).
		RawJSON("response", responseBytes).
		Msg("C2B: ответ успешно отправлен клиенту")
}
