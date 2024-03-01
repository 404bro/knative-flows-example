# Knative Function Flows Example

[![en](https://img.shields.io/badge/lang-en-blue.svg)](./README.md)
[![zh](https://img.shields.io/badge/lang-zh--cn-red.svg)](./README.zh-cn.md)

This is an example about how to use Knative Eventing(w/ Knative Eventing Flows) and Knative Function.

This workflow does the following (as shown in the following diagram):
* The `event-sender` function accepts an HTTP POST request with the request body being an integer string, and sends a CloudEvent with that integer string to the parallel resource.
* The parallel resource will execute two branches concurrently, with `is-odd` and `is-even` as filters. The former filters odd requests, while the latter filters even requests.
* The `is-odd` filter passes the CloudEvent to the `div2` function, which divides the integer corresponding to the event body by 2. The `is-even` filter passes the CloudEvent to a sequence (comprising the `mul3` and `add1` functions), which multiplies the integer corresponding to the event body by 3 and then adds 1.
* The results of these operations are sent as CloudEvents to the `event-display` service for displaying the received CloudEvents.

<p align="center">
  <img src="flows.svg" />
<p>

## Deploying

### Prerequisites

* Complete installation of Kubernetes, Knative Serving, Knative Eventing environment (including network plugins).
* InMemoryChannel, Magic DNS (sslip.io) components.
* Knative Function command-line tool.
* If Kubernetes is installed on bare metal, MetalLB component is required for Load-balancer implementation.

### Deploy to the Cluster

1. Set the environment variable `REGISTRY` in `deploy.sh` to specify the image repository (e.g., `docker.io/example`).
2. Run the `./deploy.sh` command.

### Test

1. Run the `kubectl get ksvc` command to find the external access URL for the `event-sender` function (e.g., `http://event-sender.default.192.168.10.0.sslip.io`).
2. Run `curl -v -d '{NUM}' "http://event-sender.default.192.168.10.0.sslip.io"`, where `{NUM}` is either an odd or even number.
3. Run the `kubectl get pods` command to find the Pod related to the `event-display` service, and use `kubectl logs {POD_NAME}` to view the logs.

We should see the workflow correctly completing calculations for the "Collatz Conjecture" with output similar to the following in the logs.

```plain
☁️  cloudevents.Event
Context Attributes,
  specversion: 1.0
  type: com.example.event
  source: event-sender
  id: 0
  time: 2024-03-01T14:01:19.215673814Z
  datacontenttype: text/plain
Data,
  62
```
