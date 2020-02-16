package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"

	v1 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	"k8s.io/client-go/metrics"
)

var (
	masterURL  string
	kubeconfig string
	name       string
	n, ns      bool
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	r := prometheus.NewRegistry()

	// Register default metrics onto new registry
	m := metrics.NewMetrics(r)
	// Create new metrics client with the default metrics we registered above
	// optionally disable name and namespace label values
	// to save on cardinality.
	clientMetrics := metrics.NewClientMetrics(m, n, ns)

	// Get autoscaling v1 client with client metrics
	clientv1, err := v1.NewWithMetrics(cfg, clientMetrics)
	if err != nil {
		klog.Fatalf("failed to create new client: %s", err.Error())
	}

	c := clientv1.HorizontalPodAutoscalers("default")
	// Perform a bunch of requests on the HPA
	_, err = c.Get("bla", metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}
	_, err = c.Get(name, metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}
	_, err = c.Get("blah", metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}
	existingHPA, err := c.Get(name, metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
		klog.Fatalf("failed to get hpa: %s", err.Error())
	}

	// update an existing hpa
	_, err = c.Update(existingHPA)
	if err != nil {
		fmt.Println(err)
	}

	// delete an existing hpa
	err = c.Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		fmt.Println(err)
	}

	// delete a failed hpa
	err = c.Delete("blah", &metav1.DeleteOptions{})
	if err != nil {
		fmt.Println(err)
	}

	// update a deleted hpa
	_, err = c.Update(existingHPA)
	if err != nil {
		fmt.Println(err)
	}

	// serve metrics
	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func init() {
	flag.StringVar(&name, "name", "php-apache", "Name of the hpa to get.")
	flag.BoolVar(&n, "disable-name", false, "Disable name label.")
	flag.BoolVar(&ns, "disable-namespace", false, "Disable namespace label.")
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}
