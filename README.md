# Knative Function Flows Example

[![en](https://img.shields.io/badge/lang-en-blue.svg)](./README.md)
[![zh](https://img.shields.io/badge/lang-zh--cn-red.svg)](./README.zh-cn.md)

This is an example about how to use Knative Eventing(w/ Knative Eventing Flows) and Knative Function.

This workflow does the following (as shown in the diagram):

- The `event-sender` function receives an HTTP POST and sends a CloudEvent of type `com.example.hello` to the Broker.
- The Broker forwards the received message to various Triggers. The `random-sender-trigger` accepts the `com.example.hello` event after filtering and triggers the `random-sender` function.
- The `random-sender` function generates a random number, changes the message type to `com.example.collatz`, and sends it to the Broker.
- The Broker forwards the received message to various Triggers. The `parallel-trigger` accepts the `com.example.collatz` event after filtering and triggers the `Parallel` workflow.
- The `Parallel` workflow branches into two parts, handling odd and even cases separately:
  - The `is-odd` filter selects "odd" events and triggers the `Sequence` workflow.
    - The `Sequence` workflow consists of two functions, `mul3` and `add1`, which multiply the number by 3 and add 1, respectively. They then change the event type to `com.example.display` and send it to the Broker.
  - The `is-even` filter selects "even" events and triggers the `div2` function.
    - The `div2` function divides the number by 2 and changes the event type to `com.example.display` before sending it to the Broker.
- The Broker forwards the received message to various Triggers. The `event-display-trigger` accepts the `com.example.display` event after filtering and triggers the `event-display` service, which logs the result.

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

1. Set the environment variable `REGISTRY` in `deploy.sh` to specify the image repository (e.g., `docker.io/username`).
2. Run the `./deploy.sh` command.

### Test

1. Run the `kubectl get ksvc` command to find the external access URL for the `event-sender` function (e.g., `http://event-sender.flows-example.192.168.0.0.sslip.io`).
2. Run the `curl -v "http://event-sender.default.192.168.0.0.sslip.io"` command, where `{NUM}` is either an odd or even number.
3. Run the `kubectl get pods -n flows-example` command to find the Pod related to the `event-display` service, and use `kubectl logs -n flows-example {POD_NAME}` to view the logs.

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
