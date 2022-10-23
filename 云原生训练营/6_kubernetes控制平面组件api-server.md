# API-Server

kube-apiserver 是 Kubernetes 最重要的核心组件之一，主要提供以下的功能：

- 提供集群管理的 rest-API 接口，包括认证授权、参数校验以及集群状态变更等
- 提供其他模块之间的数据交互和通信的枢纽（其他模块通过 API Server 查询或修改数据，只有 API Server 才直接操作 etcd）

**访问控制概览**

Kubernetes API 的请求都会经过多阶段的访问控制之后才会被接受，这包括认证、授权以及准入控制等

![](6_kubernetes控制平面组件api-server.assets/image-20221023135526848.png)



**访问控制细节**

![](6_kubernetes控制平面组件api-server.assets/image-20221023135741466.png)



## 认证

开启 TLS 时，所有请求都需要先进行认证。

Kubernetes 支持多种认证机制，并支持同时开启多个认证插件（只要有一个认证通过即可）。

如果认证成功，则用户的 username 会传入授权模块做进一步授权验证，认证失败的请求则返回 HTTP 401。

### 认证插件

- x509 证书
- 静态 token 文件
- 引导 token
- 静态密码文件
- Service Account
  - 是 Kubernetes 自动生成的，并会自动挂载到容器  `/run/secrets/kubernetes.io/serviceaccount` 目录中
- Open ID
  - OAuth 2.0 的认证机制
- Webhook 令牌认证
- 匿名请求
  - 如果使用 AlwaysAllow 的认证模式，则匿名请求默认开启，但可用 `--anonymous-auth=false` 禁止匿名请求。

### 基于 webhook 的认证服务集成

- 开发认证服务
  - 解码认证请求
  - 转发认证请求至认证服务器
  - 认证结果返回给 APIServer
- 配置认证服务
- 配置 APIServer