# examples

## Simple example

1. First run `./create-hpa.sh` which will create a deployment php-apache and autoscale, which in turn creates a Horizontal Pod Autoscaler resource.
2. `go build`
3.  `./simple --kubeconfig=<path-to-kubeconfig> --name=php-apache`
4. Go to http://localhost:8080/metrics to see metrics about the requests that the examples app created

