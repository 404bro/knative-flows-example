# Knative Function Flows 示例

[![en](https://img.shields.io/badge/lang-en-blue.svg)](./README.md)
[![zh](https://img.shields.io/badge/lang-zh--cn-red.svg)](./README.zh-cn.md)

这是一个 Knative Eventing（包括 Knative Eventing Flows）和 Knative Function 的使用示例。

这个工作流做了这样一件事（流程如下图所示）：
* `event-sender` 函数接受一个请求体为整数字符串的 HTTP POST，向 parallel 资源发送带有该整数字符串的 CloudEvent。
* parallel 将以 `is-odd` 和 `is-even` 两个函数作为 filter，同时执行两个分支，其中前者筛选奇数请求，后者筛选偶数请求。
* `is-odd` 过滤器将 CloudEvent 传入函数 `div2`，将事件体对应的整数除以2；`is-even` 过滤器将 CloudEvent 传入 sequence（依次包含函数 `mul3` 和 `add1`），将事件体对应的整数乘 3 后再加 1。
* 上述的运算结果作为 CloudEvent 传入 `event-display` 服务，对收到的 CloudEvent 进行展示。

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

1. 设置 `deploy.sh` 中的环境变量 `REGISTRY`，指定镜像存储库（例：`docker.io/example`）。
2. 运行 `./deploy.sh` 命令。

### 测试

1. 运行 `kubectl get ksvc` 命令，找到 `event-sender` 函数的外部访问 URL（例：`http://event-sender.default.192.168.10.0.sslip.io`）。
2. 运行 `curl -v -d '{NUM}' "http://event-sender.default.192.168.10.0.sslip.io"`，其中，`{NUM}` 分别为奇数或偶数。
3. 运行 `kubectl get pods` 命令，找到 `event-display` 服务相关的 Pod，使用 `kubectl logs {POD_NAME}` 来查看日志。

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
