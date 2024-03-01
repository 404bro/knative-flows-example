#!/bin/bash

export REGISTRY=YOUR_REGISTRY # Container registry + registry namespace. (ex 'ghcr.io/myuser').

functions=("add1" "div2" "mul3" "is-even" "is-odd" "event-sender")
for function in "${functions[@]}"; do
    func delete $function
done

resources=("resources/event-display.yaml" "resources/parallel.yaml" "resources/sequence.yaml")
for resource in "${resources[@]}"; do
    kubectl delete -f $resource --ignore-not-found=true
done
