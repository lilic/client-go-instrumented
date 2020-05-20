module github.com/lilic/sample-instrumented-controller

go 1.13

require (
	github.com/lilic/client-go-instrumented v0.0.0-20200211204413-99e4cb00429f
	github.com/prometheus/client_golang v1.4.1
	k8s.io/apimachinery v0.17.3
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v1.0.0
)

replace (
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // pinned to release-branch.go1.13
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7 // pinned to release-branch.go1.13
	k8s.io/api => k8s.io/api v0.0.0-20200209065837-d5ab0ddd8b73
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20200209065656-fb45917aae8b
	k8s.io/client-go => github.com/lilic/client-go v0.0.0-20200216150340-4200a6fa8bd2
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20200209044353-396603da40fa
)
