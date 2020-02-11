/*
Copyright 2017 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"

	"github.com/lilic/client-go-instrumented/metrics"
	scale "github.com/lilic/client-go-instrumented/typed/autoscaling/v1"
)

var (
	masterURL  string
	kubeconfig string
	name       string
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	r := prometheus.NewRegistry()
	m := metrics.NewClientMetrics(r)

	// autoscalingV1InstrumentedClient for the default namespace, with disabled name and enabled namespace label values
	c := scale.NewHorizontalPodAutoscalers("default", kubeClient.AutoscalingV1(), m, false, false)

	// do a bunch of get requests on the instrumented client
	_, err = c.Get(context.TODO(), "bla", metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}
	_, err = c.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}
	_, err = c.Get(context.TODO(), "blah", metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}
	existingHPA, err := c.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
		klog.Fatalf("failed to get hpa: %s", err.Error())
	}

	// update an existing hpa
	_, err = c.Update(context.TODO(), existingHPA, metav1.UpdateOptions{})
	if err != nil {
		fmt.Println(err)
	}

	// delete an existing hpa
	err = c.Delete(context.TODO(), name, &metav1.DeleteOptions{})
	if err != nil {
		fmt.Println(err)
	}

	// delete a failed hpa
	err = c.Delete(context.TODO(), "blah", &metav1.DeleteOptions{})
	if err != nil {
		fmt.Println(err)
	}

	// update a deleted hpa
	_, err = c.Update(context.TODO(), existingHPA, metav1.UpdateOptions{})
	if err != nil {
		fmt.Println(err)
	}

	// serve metrics
	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func init() {

	flag.StringVar(&name, "name", "php-apache", "Name of the hpa to get.")
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}
