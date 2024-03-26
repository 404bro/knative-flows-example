#!/bin/bash

functions=("add1" "div2" "mul3" "is-even" "is-odd" "event-sender" "random-sender")
for function in "${functions[@]}"; do
    func delete -n flows-example $function
done

resources=("resources/event-display.yaml" "resources/parallel.yaml" "resources/sequence.yaml" "resources/triggers.yaml")
for resource in "${resources[@]}"; do
    kubectl delete -f $resource --ignore-not-found=true
done

kn broker delete default -n flows-example
kubectl delete namespace flows-example
