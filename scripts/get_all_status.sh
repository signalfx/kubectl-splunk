#!/bin/bash

TMP=$(mktemp -d)

get_status () {
    namespace=$1
    pod_name=$2
    echo "Pod: $pod_name" $'\t' "Namespace: $namespace"
    separator
    kubectl exec -t --namespace $namespace $pod_name signalfx-agent status config
    echo ""
    echo ""
}

get_logs () {
    namespace=$1
    pod_name=$2
    kubectl logs --namespace $namespace $pod_name
}

get_all_logs () {
    while IFS= read -r line; 
    do 
        get_logs $line;
    done <<< "$(get_signalfx_agent_pods)"
}

get_all_status () {
    while IFS= read -r line; 
    do 
        get_status $line;
    done <<< "$(get_signalfx_agent_pods)"
}

separator () {
    echo "----------------------------------------------------------"
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
    echo "$(get_configmap)" > $TMP/configmap.yaml
    while IFS= read -r line; 
    do 
        get_support_for_pod $line;
    done <<< "$(get_signalfx_agent_pods)"
}

get_support_for_pod () {
    namespace=$1 
    pod=$2
    printf "%s\r" "Getting logs and status for $pod"
    mkdir $TMP/$pod
    get_logs $namespace $pod  > $TMP/$pod/agent.log
    get_status $namespace $pod  > $TMP/$pod/agent_status.txt
}

trap cleanup 0 1 2 3 6 13 15

cleanup()
{
  echo "Cleaning up temp files:"
  rm -rf "$TMP"
  exit 1
}

get_support
