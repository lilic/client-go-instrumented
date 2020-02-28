package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

type ClientMetrics struct {
	metrics          *metrics
	disableName      bool
	disableNamespace bool
}

type metrics struct {
	errors *prometheus.CounterVec
	total  *prometheus.CounterVec
}

func NewClientMetrics(m *metrics, disableName, disableNamespace bool) *ClientMetrics {
	return &ClientMetrics{
		metrics:          m,
		disableName:      disableName,
		disableNamespace: disableNamespace,
	}
}

func NewMetrics(r *prometheus.Registry) *metrics {
	var m metrics
	m.errors = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "kube_client_resource_request_errors_total",
		Help: "Total count of all the resource error requests.",
	}, []string{"resource", "method", "name", "namespace", "error"})

	m.total = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "kube_client_resource_requests_total",
		Help: "Total count of all the resource requests.",
	}, []string{"resource", "method", "name", "namespace"})
	if r != nil {
		r.MustRegister(
			m.errors,
			m.total,
		)
	}
	return &m
}

func (m *ClientMetrics) Inc(resource, action, name, namespace string, err error) {
	n := m.getName(name)
	ns := m.getNs(namespace)
	if err != nil {
		m.incErrors(resource, action, n, ns, err)
	}
	m.incTotal(resource, action, n, ns)
}

func (m *ClientMetrics) incErrors(resource, action, name, ns string, err error) {
	code := getStatusCode(err)
	m.metrics.errors.WithLabelValues(resource, action, name, ns, code).Inc()
}

func (m *ClientMetrics) incTotal(resource, action, name, ns string) {
	m.metrics.total.WithLabelValues(resource, action, name, ns).Inc()
}

func (m *ClientMetrics) getNs(ns string) string {
	if m.disableNamespace {
		return ""
	}
	return ns
}

func (m *ClientMetrics) getName(name string) string {
	if m.disableName {
		return ""
	}
	return name
}

func getStatusCode(err error) string {
	code := ""
	switch t := err.(type) {
	case apierrors.APIStatus:
		c := t.Status().Code
		code = fmt.Sprintf("%d", c)
	}
	return code
}
