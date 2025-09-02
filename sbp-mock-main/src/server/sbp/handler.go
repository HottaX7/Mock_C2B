package sbp

import (
	"espp-mock/metrics"
	"espp-mock/usecase/sbp"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	// "context" //new
	// "encoding/json" //new
)

type Handler struct {
	SBPUsecase *sbp.SBPUsecase
}

func NewHandler(usecase *sbp.SBPUsecase) *Handler {
	return &Handler{
		SBPUsecase: usecase,
	}
}

func (h *Handler) Log(w http.ResponseWriter, r *http.Request) {
	metrics.AddRequest("Undefined")

	bodyBytes, _ := io.ReadAll(r.Body)
	log.Info().Msgf("body: %s", string(bodyBytes))
	log.Info().Msgf("headers: %v", r.Header)

	w.WriteHeader(http.StatusOK)
}

