module github.com/signalfx/kubectl-splunk

go 1.15

require (
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
	k8s.io/api v0.23.6
	k8s.io/apimachinery v0.23.6
	k8s.io/cli-runtime v0.23.6
	k8s.io/client-go v0.23.6
	k8s.io/kubectl v0.23.6
)

// security updates
replace github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
