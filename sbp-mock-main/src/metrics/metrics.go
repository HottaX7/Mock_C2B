package metrics

import (
	"context"
	"espp-mock/configs"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
	"github.com/rs/zerolog/log"
)

var (
	writeAPI api.WriteAPIBlocking
)

func InitInflux(conf configs.Influx) {
	client := influxdb2.NewClient(conf.Address, "")
	writeAPI = client.WriteAPIBlocking("", conf.Database)
}

func AddRequest(reqType string) {
	p := influxdb2.NewPointWithMeasurement("requests").
		AddTag("type", reqType).
		AddField("count", 1).
		SetTime(time.Now())
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		log.Error().Err(err).Send()
	}
}
