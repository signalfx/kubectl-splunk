module github.com/signalfx/kubectl-splunk

go 1.15

require (
	github.com/emicklei/go-restful v2.9.5+incompatible // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.4.0
	github.com/spf13/viper v1.7.1
	k8s.io/api v0.25.5
	k8s.io/apimachinery v0.25.5
	k8s.io/cli-runtime v0.25.5
	k8s.io/client-go v0.25.5
	k8s.io/kubectl v0.25.5
)

// security updates
replace github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
