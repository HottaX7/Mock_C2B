package espp

import (
	"context"
	"espp-mock/configs"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

type IPSAdapder struct {
	conf   *configs.Server
	Client *http.Client
}

// Обработка "C" and "A" операций
func New(conf *configs.Server, client *http.Client) *IPSAdapder {
	log.Info().
		Str("component", "IPSAdapder").
		Str("callbackAddress", conf.CallbackAddress).
		Msg("Инициализация IPSAdapder")
	return &IPSAdapder{
		conf:   conf,
		Client: client,
	}
}

// sendRequest — универсальный метод для отправки C2CPush сообщений (C04, C24 и т.д.)
func (a *IPSAdapder) sendRequest(ctx context.Context, call string, requestBody string) error {
	const op = "espp.sendRequest"
	log.Info().Str("op", op).Str("call", call).Msg("Начало sendRequest")
	log.Debug().Str("requestBody", requestBody).Msg("sendRequest: тело запроса")

	transactionNumber, ok := ctx.Value("transactionNumber").(string)
	if !ok {
		log.Error().Str("op", op).Msg("Missing transactionNumber in context")
		return fmt.Errorf("%s: missing transactionNumber in context", op)
	}
	log.Debug().Str("transactionNumber", transactionNumber).Msg("Got transaction number from context")

	url := a.conf.CallbackAddress + "/v01/request/C2CPush/1500020/" + transactionNumber + "/" + call
	log.Info().Str("url", url).Msg("Constructed request URL")

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(requestBody))
	if err != nil {
		log.Error().Err(err).Str("op", op).Msg("Failed to create HTTP request")
		return fmt.Errorf("%s: %w", op, err)
	}

	// Установка заголовков
	correlationID, ok := ctx.Value("correlationID").(string)
	if !ok {
		log.Error().Str("op", op).Msg("Missing correlationID in context")
		return fmt.Errorf("%s: missing correlationID in context", op)
	}
	req.Header.Set("X-Correlation-ID", correlationID)
	req.Header.Set("X-Sbp-Trn-Num", transactionNumber)
	req.Header.Set("Content-Type", "application/xml")
	log.Debug().Interface("headers", req.Header).Msg("Set request headers")

	// Отправка запроса
	resp, err := a.Client.Do(req)
	if err != nil {
		log.Error().Err(err).Str("op", op).Str("url", url).Msg("HTTP request failed")
		return fmt.Errorf("%s: %w", op, err)
	}
	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusAccepted {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Error().Err(err).Str("op", op).Msg("Failed to read response body")
			return fmt.Errorf("%s: %w", op, err)
		}
		bodyString := string(bodyBytes)
		log.Error().
			Str("op", op).
			Str("response status", resp.Status).
			Str("response body", bodyString).
			Str("request url", req.URL.String()).
			Msg("Unexpected HTTP status code")
		return fmt.Errorf("%s: unexpected response status %d", op, resp.StatusCode)
	}

	log.Info().Str("op", op).Int("status_code", resp.StatusCode).Msg("sendRequest завершён успешно")
	return nil
}

// GetQRData — метод для получения данных QR-кода из внешней системы
func (a *IPSAdapder) GetQRData(ctx context.Context, uniqrID string) (string, error) {
	const op = "espp.GetQRData"
	log.Info().Str("op", op).Str("uniqrID", uniqrID).Msg("Начало GetQRData")

	// Формирование URL
	url := fmt.Sprintf("%s/payment/v1/universal-payment-link/paymentdata/%s", a.conf.CallbackAddress, uniqrID)
	log.Debug().Str("url", url).Msg("Constructed GET request URL")

	// Создание GET-запроса
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Error().Err(err).Str("op", op).Msg("Failed to create GET request")
		return "", fmt.Errorf("%s: failed to create request: %w", op, err)
	}

	// Установка заголовков
	payerID, ok := ctx.Value("payerID").(string)
	if !ok {
		log.Error().Str("op", op).Msg("Missing payerID in context")
		return "", fmt.Errorf("%s: missing payerID in context", op)
	}
	req.Header.Set("X-PAYER-ID", payerID)
	req.Header.Set("Accept", "application/json")
	log.Debug().
		Str("X-PAYER-ID", payerID).
		Str("Accept", "application/json").
		Msg("Set request headers")

	// Выполнение запроса
	resp, err := a.Client.Do(req)
	if err != nil {
		log.Error().Err(err).Str("op", op).Str("url", url).Msg("GET request failed")
		return "", fmt.Errorf("%s: failed to send request: %w", op, err)
	}
	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		log.Error().
			Str("op", op).
			Str("response status", resp.Status).
			Str("response body", bodyString).
			Str("request url", req.URL.String()).
			Msg("Unexpected response status")
		return "", fmt.Errorf("%s: unexpected response status %d", op, resp.StatusCode)
	}

	// Чтение тела ответа
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Str("op", op).Msg("Failed to read response body")
		return "", fmt.Errorf("%s: failed to read response body: %w", op, err)
	}

	log.Info().Str("op", op).Msg("GetQRData завершён успешно")
	return string(bodyBytes), nil
}

// SendBOperation — метод для отправки B-операций (B05, B06 и др.)
func (a *IPSAdapder) SendBOperation(ctx context.Context, operationType string, requestBody string) error {
	const op = "espp.SendBOperation"
	log.Info().Str("op", op).Str("operationType", operationType).Msg("Начало SendBOperation")

	transactionNumber, ok := ctx.Value("transactionNumber").(string)
	if !ok {
		return fmt.Errorf("%s: missing transactionNumber in context", op)
	}

	url := fmt.Sprintf("%s/v01/request/C2BQRD/1500020/%s", a.conf.CallbackAddress, operationType)
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(requestBody))
	if err != nil {
		return fmt.Errorf("%s: failed to create request: %w", op, err)
	}

	correlationID, ok := ctx.Value("correlationID").(string)
	if !ok {
		return fmt.Errorf("%s: missing correlationID in context", op)
	}
	req.Header.Set("X-Correlation-ID", correlationID)
	req.Header.Set("X-Sbp-Trn-Num", transactionNumber)
	req.Header.Set("Content-Type", "application/xml")

	resp, err := a.Client.Do(req)
	if err != nil {
		return fmt.Errorf("%s: failed to send request: %w", op, err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	if resp.StatusCode != http.StatusAccepted && resp.StatusCode != http.StatusOK {
		log.Error().
			Str("op", op).
			Str("response status", resp.Status).
			Str("response body", bodyString).
			Str("request url", req.URL.String()).
			Msg("Unexpected HTTP status in SendBOperation")
		return fmt.Errorf("%s: unexpected response status %d", op, resp.StatusCode)
	}

	// Логируем 200 OK отдельно, чтобы видеть, что оно прошло
	if resp.StatusCode == http.StatusOK {
		log.Info().
			Str("op", op).
			Str("response status", resp.Status).
			Str("response body", bodyString).
			Msg("SendBOperation успешно (200 OK)")
	} else {
		log.Info().
			Str("op", op).
			Str("response status", resp.Status).
			Msg("SendBOperation успешно (202 Accepted)")
	}

	return nil
}
