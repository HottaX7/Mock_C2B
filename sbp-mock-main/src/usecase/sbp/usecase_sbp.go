package sbp

import (
	"espp-mock/configs"
	"espp-mock/espp"

	"github.com/rs/zerolog/log"
)

type SBPUsecase struct {
	conf *configs.Server

	ESPP *espp.IPSAdapder
}

func New(conf *configs.Server, espp *espp.IPSAdapder) *SBPUsecase {
	log.Info().
		Str("component", "SBPUsecase").
		Msg("Инициализация SBPUsecase и зависимостей")
	return &SBPUsecase{
		conf: conf,
		ESPP: espp,
	}
}
