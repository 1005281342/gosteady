# etcd 

基于 raft 协议开发的分布式 key-value 存储，可用于服务发现、共享配置以及一致性保障（如数据库选主、分布式锁等）

- 键值对存储：将数据存储在分层组织的目录中，如同在标准文件系统中
- 检测变更：通过 watch 检测 key 所对应的 value 变更
- 可靠：使用 raft 算法保证一致性
- 简单：curl 可访问用户的 API
- 安全：可选的 SSL 客户端证书认证

### 主要功能

- 基本的 key-value 存储
- 监听机制
- key 的过期及续约机制，用于服务发现
- 原子 `Compare And Swap` 和 `Compare And Delete`，用于分布式锁和 leader 选举

### 使用场景

- 可用于键值对存储，应用程序可以读取和写入 `etcd` 中的数据
- `etcd` 比较多的应用场景是用于服务注册与发现
- 基于监听机制的分布式异步系统

### 键值对存储

etcd 是一个`键值存储`的组件

- 采用 `key-value` 型数据存储，一般情况下比关系型数据库快
- 支持动态存储（内存）以及静态存储（磁盘）
- 分布式存储，可集成为多节点集群
- 存储方式，类似于目录结构

### 服务注册与发现

- 服务注册：服务提供者注册服务地址到服务注册中心
- 心跳保活：服务定期上报心跳
- 健康检查：服务注册中心定期检查服务状态，不可用时及时剔除
- 服务发现：服务消费者通过服务注册中心获取可用服务实例地址

![](5_kubernetes控制平面组件etcd.assets/image-20221018211142963.png)

### 消息发布和订阅

以使用 etcd 作为服务远程配置为例，在应用启动时，服务主动从 etcd 拉取一次配置信息，后续更新则通过 watcher 监听：

- etcd 服务作为消息中心
- 用户修改配置，发布消息
- 服务 watcher 订阅到数据变更，做出响应



## Raft 协议

Raft 协议基于 `quorum` 机制，即大多数同意原则，任何的变更都需要超过半数的成员确认

![](5_kubernetes控制平面组件etcd.assets/image-20221018234534570.png)

### 相关资料

理解分布式共识协议 Raft http://thesecretlivesofdata.com/raft/

https://raft.github.io/raft.pdf

https://raft.github.io/

https://github.com/maemual/raft-zh_cn/blob/master/raft-zh_cn.md

### 理解

#### 集群选主

**集群个数应该是素数个，避免出现多个候选节点拉票得票数一样的情况**

