package main

import (
	"context"
	"fmt"

	"github.com/lilic/client-go-instrumented/metrics"
	scale "github.com/lilic/client-go-instrumented/typed/autoscaling/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type HPA struct {
	// kubeclientset is a standard kubernetes clientset
	kubeclientset kubernetes.Interface
	auto          scale.HorizontalPodAutoscalerInterface
}

func NewHPA(client kubernetes.Interface, m *metrics.ClientMetrics) *HPA {
	return &HPA{
		kubeclientset: client,
		auto:          scale.NewHorizontalPodAutoscalers("default", client.AutoscalingV1(), m),
	}

}

func (hpa *HPA) Get(name string) error {
	h, err := hpa.auto.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		// error metric
		return err
	}
	fmt.Println(h)
	return nil
}
