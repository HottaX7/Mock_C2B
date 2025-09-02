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

func (h *Handler) HandleA01(w http.ResponseWriter, r *http.Request) {
	const op = "sbp.HandleA01"

	metrics.AddRequest("A01")

	correlationID := r.Header.Get("X-Correlation-ID")
	transactionNumber := "A32123420144" + strconv.Itoa(rand.Intn(9999999999-1000000000)+1000000000) + "IA00000001"

	ctx := context.WithValue(r.Context(), "correlationID", correlationID)
	ctx = context.WithValue(ctx, "transactionNumber", transactionNumber)

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error().Err(err).Str("op", op).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	a01 := &entity.A01{}
	if err := xml.Unmarshal(bodyBytes, a01); err != nil {
		log.Error().Err(err).Str("op", op).Str("body", string(bodyBytes)).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := h.SBPUsecase.A01(ctx, a01); err != nil {
		log.Error().Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("X-Sbp-Trn-Num", transactionNumber)
	w.Write([]byte("OK"))
}
