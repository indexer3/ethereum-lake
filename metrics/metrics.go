package metrics

import (
	"time"

	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/common/utils"
	"github.com/indexer3/ethereum-lake/constant/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

var metrics []prometheus.Collector

func CollectMetrics(toCollect []prometheus.Collector) {
	metrics = append(metrics, toCollect...)
}

// SinceInSeconds gets the time since the specified start in seconds.
func SinceInSeconds(start time.Time) float64 {
	return time.Since(start).Seconds()
}

// SinceInMilliseconds gets the time since the specified start in milliseconds.
func SinceInMilliseconds(start time.Time) float64 {
	return float64(time.Since(start).Milliseconds())
}

func RunPusher(chain, name string) {
	log.Info("run pusher", zap.String("chain", chain), zap.String("name", name))

	pusher := push.New(viper.GetString(config.MetricsCollectorURL), name)

	for _, c := range metrics {
		pusher = pusher.Collector(c)
	}

	go utils.Recoverable(func() {
		for range time.Tick(time.Duration(viper.GetInt(config.MetricsCollectInterval)) * time.Second) {
			err := pusher.Push()
			if err != nil {
				log.Error("failed to push metrics", zap.Error(err))
			}
		}
	})
}
