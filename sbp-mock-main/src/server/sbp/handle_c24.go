package sbp

import (
	"context"
	"encoding/xml"
	"espp-mock/entity"
	"espp-mock/metrics"
	"io"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

func (h *Handler) HandleC24(w http.ResponseWriter, r *http.Request) {
	const op = "sbp.HandleC24"

	metrics.AddRequest("C24")

	correlationID := r.Header.Get("X-Correlation-ID")
	address := r.URL.Path
	transactionNumber := strings.Split(address, "/")[5]

	ctx := context.WithValue(r.Context(), "correlationID", correlationID)
	ctx = context.WithValue(ctx, "transactionNumber", transactionNumber)

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error().Err(err).Str("op", op).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	c24 := &entity.C24{}
	if err := xml.Unmarshal(bodyBytes, c24); err != nil {
		log.Error().Err(err).Str("op", op).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := h.SBPUsecase.C24(ctx, c24); err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("X-Sbp-Trn-Num", transactionNumber)
	w.Write([]byte("OK"))
}
