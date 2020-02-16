# Example of using client-go with metrics

## How to use:
1. `./create-hpa.sh` # creates 1 deployment and its HPA
2. `go build`
3. `./sample-instrumented-controller --kubeconfig=/Users/lili/.kube/config`
4.  go to `http://localhost:8080/metrics`

/metrics example:
```
# HELP kube_client_resource_request_errors_total Total count of all the resource error requests.
# TYPE kube_client_resource_request_errors_total counter
kube_client_resource_request_errors_total{error="404",method="delete",name="",namespace="default",resource="hpa"} 1
kube_client_resource_request_errors_total{error="404",method="get",name="bla",namespace="default",resource="hpa"} 1
kube_client_resource_request_errors_total{error="404",method="get",name="blah",namespace="default",resource="hpa"} 1
kube_client_resource_request_errors_total{error="409",method="update",name="php-apache",namespace="default",resource="hpa"} 1
# HELP kube_client_resource_requests_total Total count of all the resource requests.
# TYPE kube_client_resource_requests_total counter
kube_client_resource_requests_total{method="delete",name="",namespace="default",resource="hpa"} 2
kube_client_resource_requests_total{method="get",name="bla",namespace="default",resource="hpa"} 1
kube_client_resource_requests_total{method="get",name="blah",namespace="default",resource="hpa"} 1
kube_client_resource_requests_total{method="get",name="php-apache",namespace="default",resource="hpa"} 2
kube_client_resource_requests_total{method="update",name="php-apache",namespace="default",resource="hpa"} 2
```

## Optional args:
```
--disable-name false If true does not include name of the resource in the metrics to prevent high cardinality.
--disable-namespace false If True does not include namespace of the resource in the metrics to prevent high cardinality.
```
