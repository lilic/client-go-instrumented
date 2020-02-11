# examples

## Simple example

1. First run `./create-hpa.sh` which will create a deployment php-apache and autoscale, which in turn creates a Horizontal Pod Autoscaler resource.
2. `go build`
3.  `./simple --kubeconfig=<path-to-kubeconfig> --name=php-apache`
4. Go to http://localhost:8080/metrics to see metrics about the requests that the examples app created

### Example output of the /metrics

```
# HELP k8sclient_resource_errors_total Total count of all the resources error requests.
# TYPE k8sclient_resource_errors_total counter
k8sclient_resource_errors_total{error="404",method="delete",name="blah",namespace="default",resource="hpa"} 1
k8sclient_resource_errors_total{error="404",method="get",name="bla",namespace="default",resource="hpa"} 1
k8sclient_resource_errors_total{error="404",method="get",name="blah",namespace="default",resource="hpa"} 1
k8sclient_resource_errors_total{error="409",method="update",name="php-apache",namespace="default",resource="hpa"} 1
# HELP k8sclient_resource_success_total Total count of all the resources success requests.
# TYPE k8sclient_resource_success_total counter
k8sclient_resource_success_total{method="delete",name="php-apache",namespace="default",resource="hpa"} 1
k8sclient_resource_success_total{method="get",name="php-apache",namespace="default",resource="hpa"} 2
k8sclient_resource_success_total{method="update",name="php-apache",namespace="default",resource="hpa"} 1
```
