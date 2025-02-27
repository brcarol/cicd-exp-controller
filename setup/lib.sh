#!/usr/bin/env bash

cur_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" > /dev/null && pwd )"

export APPS_DIR=${cur_dir}/../clusters/common/apps

# Host system's port mapped to the local container registry's port.
export REGISTRY_PORT='5000'

# Using localhost instead of kind.local to push images to the local kind-registry.
export KO_DOCKER_REPO=localhost:$REGISTRY_PORT

export CLUSTERS_DIR=${cur_dir}/../clusters

list_clusters() {
    ls $CLUSTERS_DIR | grep -v common
}

list_kind_clusters() {
    clusters=$(kind get clusters | grep -Ei 'c[0-9]+')
    echo "${clusters[@]}"
}

get_interface() {
    ip route get 8.8.8.8 | head -n1 | awk '{print $5}'
}

get_ip() {
    ifconfig $(get_interface) |
        grep -i mask |
        awk '{print $2}'|
        cut -f2 -d:
}

## Parallelism

# Array to store the pids of the background processes
pids=()

# Function to trap signals that interrupt a script
interrupt_handler() {
    echo -e "\nScript interrupted."
    for pid in "${pids[@]}"; do
        kill $pid
    done
    exit 1
}

# Function to wait for all background processes to finish
wait_pids() {
    for pid in "${pids[@]}"; do
        wait $pid
    done
}

# Trap the interrupt signal (Ctrl + C) to call the interrupt_handler function
trap interrupt_handler SIGINT