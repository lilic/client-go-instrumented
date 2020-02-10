package metrics

import "github.com/prometheus/client_golang/prometheus"

// ListWatchMetrics stores the pointers of kube_state_metrics_[list|watch]_total metrics.
type ClientMetrics struct {
	ClientMetricErrors  *prometheus.CounterVec
	ClientMetricSuccess *prometheus.CounterVec
}

// NewClientMetrics
func NewClientMetrics(r *prometheus.Registry) *ClientMetrics {
	var m ClientMetrics
	m.ClientMetricErrors = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "k8sclient_resource_errors_total",
		Help: "Total count of all the resources error requests.",
	}, []string{"resource", "method", "name", "namespace", "error"})

	// ClientMetricSuccess
	m.ClientMetricSuccess = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "k8sclient_resource_success_total",
		Help: "Total count of all the resources success requests.",
	}, []string{"resource", "method", "name", "namespace"})
	if r != nil {
		r.MustRegister(
			m.ClientMetricErrors,
			m.ClientMetricSuccess,
		)
	}
	return &m
}
