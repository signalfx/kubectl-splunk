module github.com/signalfx/kubectl-splunk

go 1.15

require (
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	k8s.io/api v0.27.2
	k8s.io/apimachinery v0.27.2
	k8s.io/cli-runtime v0.20.0
	k8s.io/client-go v0.27.2
	k8s.io/kubectl v0.20.0
)

// security updates
replace github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
