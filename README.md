# kubectl-signalfx
A kubectl plugin for interacting with signalfx-agent deployments


## Installation
Clone this repository and copy the `kubectl-signalfx` binary to `/usr/local/bin`
1. `git clone https://github.com/signalfx/kubectl-signalfx-agent.git`
2. `chmod +x kubectl-signalfx`
3. `cp kubectl-signalfx /usr/local/bin`

## Usage
```bash

kubectl-signalfx [ support | status | config | logs | endpoints | configmap ]
   pods:        Display status of $AGENT pods, 
                this is the default command when none is passed.
   support:     Zip up the configmap, logs, status, and resolved config.
   status:      print the $AGENT statuses.
   config:      print the $AGENT resolved config.
   configmap:   print the $AGENT config map.
   endpoints:   print the endpoints discovered by the $AGENT.
   logs:        print the $AGENT logs.

   Optional args:

   (-p|--pod)=signalfx-agent-pod-name. If no pod is passed, all
              $AGENT pods are used.
```


## License
[Apache 2.0](./LICENSE)
