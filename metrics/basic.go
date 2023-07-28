package metrics

import "github.com/prometheus/client_golang/prometheus"

func NewLakeErrorGauge() *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "error_gauge",
		Help: "error gauge",
	}, []string{"chain", "task", "error"})
}

func NewStatusGauge() *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "status_gauge",
		Help: "status gauge",
	}, []string{"chain", "task", "status"})
}
