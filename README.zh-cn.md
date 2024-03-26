# Knative Function Flows 示例

[![en](https://img.shields.io/badge/lang-en-blue.svg)](./README.md)
[![zh](https://img.shields.io/badge/lang-zh--cn-red.svg)](./README.zh-cn.md)

这是一个 Knative Eventing（包括 Knative Eventing Flows）和 Knative Function 的使用示例。

这个工作流做了这样一件事（流程如下图所示）：
* `event-sender` 函数接受一个 HTTP POST，向 Broker 发送一个类型为 `com.example.hello` 的 CloudEvent。
* Broker 将收到的消息转发给各个 Trigger。`random-sender-trigger` 在过滤后接受 `com.example.hello` 事件，触发 `random-sender` 函数。
* `random-sender` 函数生成一个随机数，将消息类型改为 `com.example.collatz` 后发送给 Broker。
* Broker 将收到的消息转发给各个 Trigger。`parallel-trigger` 在过滤后接受 `com.example.collatz` 事件，触发 `Parallel` 工作流。
* `Parallel` 工作流分为两个分支，分别处理奇数和偶数的情况：
  * 过滤器 `is-odd` 筛选“奇数”事件，并触发 `Sequence` 工作流。
    * `Sequence` 工作流由 `mul3`、`add1` 两个函数组成，它们分别将数字乘 3、加 1，然后在修改事件类型为 `com.example.display` 后发送给 Broker。
  * 过滤器 `is-even` 筛选“偶数”事件，并触发 `div2` 函数。
    * `div2` 函数将数字除以 2，然后在修改事件类型为 `com.example.display` 后发送给 Broker。
* Broker 将收到的消息转发给各个 Trigger。`event-display-trigger` 在过滤后接受 `com.example.display` 事件，触发 `event-display` 服务，在日志中打印结果。

<p align="center">
  <img src="flows.svg" />
<p>

## 部署

### 前置条件

* 完整安装的 Kubernetes、Knative Serving、Knative Eventing 环境（包含网络插件）。
* InMemoryChannel、Magic DNS（sslip.io）组件。
* Knative Function 命令行工具。
* 如果 Kubernetes 安装于裸机，则需要 MetalLB 组件作为 Load-balancer 实现。

### 部署到集群

1. 设置 `deploy.sh` 中的环境变量 `REGISTRY`，指定镜像存储库（例：`docker.io/username`）。
2. 运行 `./deploy.sh` 命令。

### 测试

1. 运行 `kubectl get ksvc` 命令，找到 `event-sender` 函数的外部访问 URL（例：`http://event-sender.flows-example.192.168.0.0.sslip.io`）。
2. 运行 `curl -v "http://event-sender.default.192.168.0.0.sslip.io"`，其中，`{NUM}` 分别为奇数或偶数。
3. 运行 `kubectl get pods -n flows-example` 命令，找到 `event-display` 服务相关的 Pod，使用 `kubectl logs -n flows-example {POD_NAME}` 来查看日志。

我们应当看到该工作流正确地完成了“冰雹猜想”的运算，日志中呈现类似如下输出。

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
