#!/bin/bash

export REGISTRY=YOUR_REGISTRY # Container registry + registry namespace. (ex 'ghcr.io/myuser').

functions=("op/add1" "op/div2" "op/mul3" "filter/is-even" "filter/is-odd" "util/event-sender")
for function in "${functions[@]}"; do
    func deploy --path $function
done

resources=("resources/event-display.yaml" "resources/parallel.yaml" "resources/sequence.yaml")
for resource in "${resources[@]}"; do
    kubectl create -f $resource
done
