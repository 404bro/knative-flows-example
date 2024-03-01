# Knative Function Flows Example | Knative Function Flows 示例
This is an example about how to use Knative Eventing(w/ Knative Eventing Flows) and Knative Function.

这是一个 Knative Eventing（包括 Knative Eventing Flows）和 Knative Function 的使用示例。

This workflow does the following (as shown in the following diagram):
* The `event-sender` function accepts an HTTP POST request with the request body being an integer string, and sends a CloudEvent with that integer string to the parallel resource.
* The parallel resource will execute two branches concurrently, with `is-odd` and `is-even` as filters. The former filters odd requests, while the latter filters even requests.
* The `is-odd` filter passes the CloudEvent to the `div2` function, which divides the integer corresponding to the event body by 2. The `is-even` filter passes the CloudEvent to a sequence (comprising the `mul3` and `add1` functions), which multiplies the integer corresponding to the event body by 3 and then adds 1.
* The results of these operations are sent as CloudEvents to the `event-display` service for displaying the received CloudEvents.

这个工作流做了这样一件事（流程如下图所示）：
* `event-sender` 函数接受一个请求体为整数字符串的 HTTP POST，向 parallel 资源发送带有该整数字符串的 CloudEvent。
* parallel 将以 `is-odd` 和 `is-even` 两个函数作为 filter，同时执行两个分支，其中前者筛选奇数请求，后者筛选偶数请求。
* `is-odd` 过滤器将 CloudEvent 传入函数`div2`，将事件体对应的整数除以2；`is-even`过滤器将 CloudEvent 传入 sequence（依次包含函数 `mul3` 和 `add1`），将事件体对应的整数乘 3 后再加 1。
* 上述的运算结果作为 CloudEvent 传入 `event-display` 服务，对收到的 CloudEvent 进行展示。

![](flows.svg)

## Deploying | 部署

### Prerequisites | 前置条件

* Complete installation of Kubernetes, Knative Serving, Knative Eventing environment (including network plugins).
* InMemoryChannel, Magic DNS (sslip.io) components.
* Knative Function command-line tool.
* If Kubernetes is installed on bare metal, MetalLB component is required for Load-balancer implementation.

* 完整安装的 Kubernetes、Knative Serving、Knative Eventing 环境（包含网络插件）。
* InMemoryChannel、Magic DNS（sslip.io）组件。
* Knative Function 命令行工具。
* 如果 Kubernetes 安装于裸机，则需要 MetalLB 组件作为 Load-balancer 实现。

### Deploy to the Cluster | 部署到集群
