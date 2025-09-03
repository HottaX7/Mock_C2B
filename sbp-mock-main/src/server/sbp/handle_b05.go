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

func (h *Handler) HandleB05(w http.ResponseWriter, r *http.Request) {
	const op = "sbp.HandleB05"
	logger := log.With().Str("op", op).Logger()

	metrics.AddRequest("B05")

	// Получаем заголовки
	correlationID := r.Header.Get("X-Correlation-ID")
	transactionNumber := "X32123420144" + strconv.Itoa(rand.Intn(9999999999-1000000000)+1000000000) + "IA00000001"

	// Устанавливаем контекст
	ctx := context.WithValue(r.Context(), "correlationID", correlationID)
	ctx = context.WithValue(ctx, "transactionNumber", transactionNumber)

	logger.Info().
		Str("correlationID", correlationID).
		Str("transactionNumber", transactionNumber).
		Str("method", r.Method).
		Str("url", r.URL.String()).
		Str("remote_addr", r.RemoteAddr).
		Msg("B05: входящий запрос получен")

	// Логируем все заголовки
	headers := make(map[string]string)
	for name, values := range r.Header {
		if len(values) > 0 {
			headers[name] = values[0]
		}
	}
	logger.Debug().
		Interface("headers", headers).
		Msg("B05: заголовки запроса")

	// Читаем тело
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error().Err(err).Msg("B05: ошибка чтения тела запроса")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Детальное логирование полученного XML
	logger.Debug().
		Str("raw_xml", string(bodyBytes)).
		Msg("B05: полученное сырое XML")

	// Парсим в существующую структуру B05
	var b05 entity.B05
	if err := xml.Unmarshal(bodyBytes, &b05); err != nil {
		logger.Error().
			Err(err).
			Str("body", string(bodyBytes)).
			Msg("B05: ошибка парсинга XML")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Логируем ключевые поля из B05
	logger.Info().
		Str("from_id", b05.AppHdr.Fr.FIId.FinInstnId.Othr.ID).
		Str("to_id", b05.AppHdr.To.FIId.FinInstnId.Othr.ID).
		Str("biz_msg_idr", b05.AppHdr.BizMsgIdr).
		Str("msg_def_idr", b05.AppHdr.MsgDefIdr).
		Str("biz_svc", b05.AppHdr.BizSvc).
		Str("cre_dt", b05.AppHdr.CreDt).
		Str("transactionNumber", transactionNumber).
		Msg("B05: сообщение успешно разобрано")

	// Логируем дополнительные детали из тела B05
	if b05.Document.FIToFICstmrCdtTrf.GrpHdr.MsgId != "" {
		logger.Debug().
			Str("grp_msg_id", b05.Document.FIToFICstmrCdtTrf.GrpHdr.MsgId).
			Str("cre_dt_tm", b05.Document.FIToFICstmrCdtTrf.GrpHdr.CreDtTm).
			Str("nb_of_txs", b05.Document.FIToFICstmrCdtTrf.GrpHdr.NbOfTxs).
			Msg("B05: детали GrpHdr")
	}

	// Логируем детали транзакции (CdtTrfTxInf - это не массив, а отдельная структура)
	tx := b05.Document.FIToFICstmrCdtTrf.CdtTrfTxInf
	logger.Debug().
		Str("end_to_end_id", tx.PmtId.EndToEndId).
		Str("tx_id", tx.PmtId.TxId).
		Str("amount", tx.IntrBkSttlmAmt.Text).
		Str("currency", tx.IntrBkSttlmAmt.Ccy).
		Msg("B05: детали транзакции")

	// Вызываем бизнес-логику
	logger.Info().
		Str("transactionNumber", transactionNumber).
		Msg("B05: вызов бизнес-логики")

	if err := h.SBPUsecase.B05(ctx, &b05); err != nil {
		logger.Error().
			Err(err).
			Str("transactionNumber", transactionNumber).
			Msg("B05: ошибка в бизнес-логике")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	logger.Info().
		Str("transactionNumber", transactionNumber).
		Msg("B05: бизнес-логика успешно выполнена")

	// Формируем B06 на основе данных из B05
	logger.Info().
		Str("transactionNumber", transactionNumber).
		Msg("B05: формирование ответа B06")

	responseXML := entity.NewB06(transactionNumber, &b05)

	// Логируем сформированный ответ
	logger.Debug().
		Str("transactionNumber", transactionNumber).
		Str("response_b06", responseXML).
		Msg("B05: сформированный ответ B06")

	// Отправляем клиенту
	w.Header().Set("Content-Type", "application/xml")
	w.Header().Set("X-Sbp-Trn-Num", transactionNumber)
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte(responseXML)); err != nil {
		logger.Error().
			Err(err).
			Str("transactionNumber", transactionNumber).
			Msg("B05: ошибка отправки ответа")
		return
	}

	logger.Info().
		Str("transactionNumber", transactionNumber).
		Msg("B05: ответ успешно отправлен клиенту")
}
