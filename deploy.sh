#!/bin/bash

export REGISTRY=192.168.122.190/library # Container registry + registry namespace. (ex 'ghcr.io/myuser').

# kubectl create namespace flows-example
kn broker create default -n flows-example

functions=("op/add1" "op/div2" "op/mul3" "filter/is-even" "filter/is-odd" "util/event-sender" "util/random-sender")
for function in "${functions[@]}"; do
    func deploy -n flows-example --path ./func/$function --registry $REGISTRY
done

resources=("resources/event-display.yaml" "resources/parallel.yaml" "resources/sequence.yaml" "resources/triggers.yaml")
for resource in "${resources[@]}"; do
    kubectl create -f $resource
done
