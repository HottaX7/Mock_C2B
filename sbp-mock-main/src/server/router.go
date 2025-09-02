package server

import (
	"espp-mock/server/sbp"
	"net/http"

	//	"github.com/go-chi/chi"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

type RouterBuilder struct {
	SBPHandler *sbp.Handler
}

func (b *RouterBuilder) Build() *chi.Mux {
	r := chi.NewRouter()

	r.With(middleware.Logger).HandleFunc("/v01/request/C2CPush/{someID}/A01", b.SBPHandler.HandleA01)
	r.With(middleware.Logger).HandleFunc("/v01/request/C2CPush/1500020/A01", b.SBPHandler.HandleA01)
	r.With(middleware.Logger).HandleFunc("/v01/request/C2CPush/{transactionNumber}/C01", b.SBPHandler.HandleC01)
	r.With(middleware.Logger).HandleFunc("/v01/request/C2CPush/1500020/{transactionNumber}/C04", b.SBPHandler.HandleC04)
	r.With(middleware.Logger).HandleFunc("/v01/request/C2CPush/1500020/{transactionNumber}/C05", b.SBPHandler.HandleC05)
	r.With(middleware.Logger).HandleFunc("/v01/request/C2CPush/1500020/{transactionNumber}/C11", b.SBPHandler.HandleC11)
	r.With(middleware.Logger).HandleFunc("/v01/request/C2CPush/1500020/{transactionNumber}/C13", b.SBPHandler.HandleC13)
	r.With(middleware.Logger).HandleFunc("/v01/request/C2CPush/1500020/{transactionNumber}/C24", b.SBPHandler.HandleC24)
	r.With(middleware.Logger).HandleFunc("/v01/request/C2CPush/{transactionNumber}/{callID}", b.SBPHandler.Log)
	r.With(middleware.Logger).HandleFunc("/v01/request/C2CPush/1500020/{transactionNumber}/{callID}", b.SBPHandler.Log)
	r.With(middleware.Logger).HandleFunc("/payment/v1/universal-payment-link/paymentdata/{id}", b.SBPHandler.HandlePaymentData)
	r.With(middleware.Logger).HandleFunc("/api/v1/C2BQRD/150000000020/B05", b.SBPHandler.HandleB05)
	r.With(middleware.Logger).HandleFunc("/api/v1/C2BQRD/1500020/B05", b.SBPHandler.HandleB05)
	r.With(middleware.Logger).HandleFunc("/api/v01/request/C2BQRD/1500020/B05", b.SBPHandler.HandleB05)
	r.With(middleware.Logger).HandleFunc("/api/v1/C2BQRD/120000000020/B05", b.SBPHandler.HandleB05)
	r.With(middleware.Logger).HandleFunc("/v1/C2BQRD/120000000020/B21", b.SBPHandler.HandleB21)
	r.HandleFunc("/api/v1/C2BQRD/{memberId}/B22", b.SBPHandler.HandleB22)
	//	r.With(middleware.Logger).HandleFunc("/v01/request/C2CPush/120000000020/B05", b.SBPHandler.HandleB05)
	r.Post("/send-payment", b.SBPHandler.HandleSendPayment)

	// Логирование всех несуществующих путей
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		log.Warn().
			Str("method", r.Method).
			Str("url", r.URL.Path).
			Msg("Запрос на несуществующий путь")
		http.NotFound(w, r)
	})

	return r
}
func SetupRouter(handler *sbp.Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/payment/v1", func(r chi.Router) {
		r.Get("/universal-payment-link/paymentdata/{id}", handler.HandlePaymentData)
	})
	return r
}
