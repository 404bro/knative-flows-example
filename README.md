# Knative Function Flows Example

[![en](https://img.shields.io/badge/lang-en-blue.svg)](./README.md)
[![zh](https://img.shields.io/badge/lang-zh--cn-red.svg)](./README.zh-cn.md)

This is an example about how to use Knative Eventing(w/ Knative Eventing Flows) and Knative Function.

This workflow accomplishes the following (as shown in the diagram):

- The `event-sender` function receives an HTTP POST request with an integer string in the request body and sends a CloudEvent with that integer string to the Broker resource.
- The Broker forwards the received message to the `parallel-trigger` trigger and the `event-display-trigger` trigger.
  - The `parallel-trigger` trigger filters events of type `com.example.collatz` and sends them to `parallel`.
  - The `event-display-trigger` trigger filters events of type `com.example.display` and sends them to `event-display`.
- `parallel` uses two functions, `is-odd` and `is-even`, as filters, executing two branches simultaneously, where the former filters odd requests and the latter filters even requests.
- The `is-odd` filter passes the CloudEvent to the `div2` function, which divides the integer corresponding to the event body by 2; the `is-even` filter passes the CloudEvent to a sequence (comprising the `mul3` and `add1` functions in sequence), which multiplies the integer corresponding to the event body by 3 and then adds 1.
- The results of the above computations are passed as CloudEvents to the Broker (with the type modified to `com.example.display`), which then forwards them to `event-display` for display.

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

1. Run the `kubectl get ksvc` command to find the external access URL for the `event-sender` function (e.g., `http://event-sender.flows-example.192.168.0.0.sslip.io`).
2. Run the `curl -v -d '{NUM}' "http://event-sender.default.192.168.0.0.sslip.io"` command, where `{NUM}` is either an odd or even number.
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
