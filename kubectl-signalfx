#!/bin/bash


POD=""
AGENT="SignalFx Smart Agent"
TMP=$(mktemp -d)
SUPPORT=$TMP/signalfx-support
mkdir -p $SUPPORT

cleanup()
{
  rm -rf "$TMP"
  exit 1
}

trap cleanup 0 1 2 3 6 13 15

usage() {
    echo "
Usage: $0 [ support | status | config | logs | endpoints | configmap ]
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
"
}


get_signalfx_pods_status () {
	kubectl get pods --all-namespaces -l app=signalfx-agent -o wide
}

get_status () {
    namespace=$1
    pod_name=$2
    kubectl exec -t --namespace $namespace $pod_name signalfx-agent status
}

get_config () {
    namespace=$1
    pod_name=$2
    kubectl exec -t --namespace $namespace $pod_name signalfx-agent status config
}

get_logs () {
    namespace=$1
    pod_name=$2
    kubectl logs --namespace $namespace $pod_name
}

get_all_logs () {
    if [ -z "$POD" ]; then
        while IFS= read -r line; 
        do 
            echo "Getting logs from pod: $line"
            get_logs $line;
        done <<< "$(get_signalfx_agent_pods)"
    else
        echo "Getting logs from pod: $POD"
        get_logs "$(get_namespace)" $POD
    fi
}

get_all_status () {
    if [ -z "$POD" ]; then
        while IFS= read -r line; 
        do 
            echo "Getting status of pod: $line"
            get_status $line;
        done <<< "$(get_signalfx_agent_pods)"
    else
        echo "Getting status of pod: $PWD"
        get_status "$(get_namespace)" $POD
    fi
}

get_all_endpoints () {
    if [ -z "$POD" ]; then
        while IFS= read -r line; 
        do 
            echo "Getting endpoints discovered by agent: $line"
            get_endpoints $line;
        done <<< "$(get_signalfx_agent_pods)"
    else
        echo "Getting endpoints discovered by agent: $POD"
        get_endpoints "$(get_namespace)" $POD
    fi
}

get_endpoints () {
    namespace=$1
    pod_name=$2
    kubectl exec -t --namespace $namespace $pod_name signalfx-agent status endpoints
}

get_all_config () {
    if [ -z "$POD" ]; then
        while IFS= read -r line; 
        do 
            echo "Getting resolved config from agent: $line"
            get_config $line;
        done <<< "$(get_signalfx_agent_pods)"
    else
        echo "Getting resolved config from agent: $POD"
        get_config "$(get_namespace)" $POD
    fi
}

separator () {
    echo "--------------------------------------------------------------------------------------"
}

get_signalfx_agent_pods () {
    echo "$(kubectl get pods -l app=signalfx-agent --all-namespaces --no-headers | tr -s " " | cut -d     " " -f 1,2)"
}

get_namespace () {
    echo "$(kubectl get pods -l app=signalfx-agent --all-namespaces --no-headers | tr -s " " | cut -d     " " -f 1 | head -n 1)"
}

get_configmap() {
    echo "$(kubectl describe configmap signalfx-agent --namespace $(get_namespace))"
}

get_support() {
    # get config map
    echo "Getting config map"
    echo "$(get_configmap)" > $SUPPORT/configmap.yaml
    if [ -z "$POD" ]; then
        while IFS= read -r line; 
        do 
            get_support_for_pod $line;
        done <<< "$(get_signalfx_agent_pods)"
    else
        get_support_for_pod "$(get_namespace)" $POD
    fi
    zip_tmp
}

get_support_for_pod () {
    namespace=$1 
    pod=$2
    printf "Gathering config and logs for $pod\n"
    mkdir $SUPPORT/$pod
    get_logs $namespace $pod  > $SUPPORT/$pod/agent.log
    get_config $namespace $pod  > $SUPPORT/$pod/resolved_config.yaml
    get_status $namespace $pod  > $SUPPORT/$pod/agent_status.txt
    get_endpoints $namespace $pod  > $SUPPORT/$pod/agent_endpoints.txt
}

zip_tmp () {
    WORKING_DIR=$PWD
    cd $TMP && zip -r9 -q signalfx-agent-support.zip signalfx-support/
    cp $TMP/signalfx-agent-support.zip $WORKING_DIR
    cd $WORKING_DIR
}



# parse command line flags
POS_ARGS=()
while [[ $# -gt 0 ]]; do
    PARAM=`echo $1 | awk -F= '{print $1}'`
    VALUE=`echo $1 | awk -F= '{print $2}'`
    case $PARAM in
        -h|--help)
            usage
            exit 0
            ;;
        -p|--pod)
            POD=$VALUE
            ;;
        *)
            POS_ARGS+=("$1")
            ;;
    esac
    shift
done


# restore remaining positional arguments
if [[ ${#POS_ARGS[@]} -gt 0 ]]; then
    set -- "${POS_ARGS[@]}"
fi

if [[ "0" == $# ]]; then
   get_signalfx_pods_status
fi

for param in "$@"; do
    case $param in
        support)
            get_support
            ;;
        status)
            get_all_status
            ;;
        config)
            get_all_config
            ;;
        configmap)
            get_configmap
            ;;
        logs)
            get_all_logs
            ;;
        endpoints)
            get_all_endpoints
            ;;
        pods)
            get_signalfx_pods_status
            ;;
        *) # default
            usage
            ;;
    esac
done
