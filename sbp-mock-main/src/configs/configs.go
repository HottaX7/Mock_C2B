package configs

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Influx struct {
	Address  string `mapstructure:"address"`
	Database string `mapstructure:"database"`
}

type Server struct {
	ListenAddress   string        `mapstructure:"listen_address"`
	CallbackDelay   time.Duration `mapstructure:"callback_delay"`
	CallbackAddress string        `mapstructure:"callback_address"`
	Influx          Influx        `mapstructure:"influx"`
}

func New(cfgFile string) (*Server, error) {
	const op = "configs.New"

	vp, err := load(cfgFile)
	if err != nil {
		return nil, err
	}

	config := &Server{}

	err = vp.Unmarshal(config)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return config, nil
}

func load(cfgFile string) (*viper.Viper, error) {
	const op = "configs.load"

	vp := newViper()

	if cfgFile != "" {
		vp.SetConfigFile(cfgFile)
	} else {
		vp.AddConfigPath(".")
		vp.SetConfigType("yaml")
		vp.SetConfigName("config")
	}

	if err := vp.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			log.Debug().Err(err).Msgf("%s: using default config", op)
		case *os.PathError:
			log.Debug().Err(err).Msgf("%s: using default config", op)
		default:
			return nil, err
		}
	}

	return vp, nil
}

func newViper() *viper.Viper {
	vp := viper.New()

	vp.AutomaticEnv()
	vp.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	vp.SetDefault("listen_address", ":9201")
	vp.SetDefault("callback_delay", 1*time.Second)
	vp.SetDefault("callback_address", "http://10.21.58.53:11502")

	vp.SetDefault("influx.address", "http://10.21.25.119:8086")
	vp.SetDefault("influx.database", "ESPP_SBP_MOCK/autogen")

	return vp
}
