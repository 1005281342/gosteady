## 从系统架构谈起

### 传统分层架构与微服务架构

![](3_docker核心技术.assets/image-20220912173853718.png)

![](3_docker核心技术.assets/image-20220912174028824.png)

简单系统采用分层架构比较合适，复杂系统采用微服务架构比较合适。简单系统可能会随着业务发展迭代演进变得复杂。

### 微服务改造

#### 分离微服务方法建议

- 审视并发现可以分离的业务逻辑
- 寻找天生隔离的代码模块，可以借助于静态代码分析工具
- 不同并发规模，不同内存需求的模块都可以分离出不同的微服务，此方法可提高资源利用率，节省成本

#### 一些常用的可微服务化的模块: 

- 用户和账户管理
- 授权和会话管理
- 系统配置
- 通知和通讯服务
- 照片，多媒体，元数据等

#### 分解原则

- 基于微服务的大小
- 基于微服务的职责（工作范围）
- 基于微服务的能力

### 微服务间通讯

#### 点对点

![](3_docker核心技术.assets/image-20220912183219884.png)

- 多用于系统内部多模块之间通讯;
- 有大量的重复模块如认证授权;
- 缺少统一规范，如监控，审计等功能;
- 后期维护成本高，服务和服务的依赖关系错综复杂难以管理。

#### API网关

![](3_docker核心技术.assets/image-20220912183255218.png)

- 基于一个轻量级的 message gateway
- 新 API 通过注册至 Gateway 实现
- 整合实现 Common function （如：认证授权、审计日志等）

## Docker

- 基于 Linux 内核的 Cgroup，Namespace，以及 Union FS 等技术，对进程进行封装隔离，属于操作系统层面的虚拟化技术，由于隔离的进程独立于宿主和其它的隔离的进程，因此也称其为容器。
- 最初实现是基于 LXC，从 0.7 以后开始去除 LXC，转而使用自行开发的 Libcontainer，从 1.11 开始，则 进一步演进为使用 runC 和 Containerd。
- Docker 在容器的基础上，进行了进一步的封装，从文件系统、网络互联到进程隔离等等，极大的简化了容 器的创建和维护，使得 Docker 技术比虚拟机技术更为轻便、快捷。

### 使用Docker的好处

- 更高效的资源利用
- 更快速的启动时间
- 一致的运行环境（打包好的image就像一个箱子，它可以在不同的地方打开，而里面的内容是一样的）
- 持续交付和部署
- 更轻松地迁移、维护以及扩展

### 虚拟机和容器运行态的对比

![](3_docker核心技术.assets/image-20220912210641565.png)

![](3_docker核心技术.assets/image-20220912210710027.png)

### 性能对比

![](3_docker核心技术.assets/image-20220912210839303.png)



## 容器

### 容器标准（OCI）

**OCI 全称是Open Container Initiative**

#### 两个规范

规范1：Runtime Specification 文件系统包如何解压至硬盘，共运行时运行

规范2：Image Specification 如何通过构建系统打包，生成镜像清单（Manifest）、文件系统序列化文件、镜像配置。

### 容器主要特性

- 安全性
- 隔离性（基于namespace）
- 便携性
- 可配额（基于CGroup）

### namespace

Linux Namespace 是一种 Linux Kernel 提供的资源隔离方案

- 系统可以以进程分配不同的 Namespace
- 并保证不同的 Namespace 资源独立分配、进程彼此隔离，即不同的 Namespace 下的进程互不干扰

### cgroups

CGroups（Control Groups）是Linux下用于对一个或一组进程进行资源控制和监控的机制

- 可以对如 CPU 使用时间、内存、磁盘I/O 等进程所需资源进行限制
- 不同资源的具体管理工作由相应的 CGroup 子系统来实现
- 不同类型的资源限制，只要将限制策略在不同的子系统上进行关联
- CGroup 在不同的系统资源管理子系统中以层级树的方式来组织管理：每个 CGroup 都可以包含其他的子 CGroup。子 CGroup 能使用的资源除了受本 CGroup 配置的资源参数限制，还受到父 CGroup 设置的资源限制

#### 可配额/可度量 的资源

![](3_docker核心技术.assets/image-20220921000729013.png)

### CPU 子系统

cpu.shares：可出让的能获得 CPU 使用时间的相对值

cpu.cfs_period_us：配置时间周期长度，单位为 us （微秒）

cpu.cfs_quota_us：配置当前cgroup在 cfs_period_us 时间内最多能使用的 CPU 时间数，单位 us （微秒）

cpu.stat：cgroup 内的进程使用的 CPU 时间统计

nr_periods：经过 cpu.cfs_period_us 的时间周期数量

nr_throttled：在经过的周期内，有多少次因为进程在指定的时间周期内用光了配额时间而受到限制

throttled_time：cgroup 中的进程被限制使用 CPU 的总用时，单位是 ns （纳秒）

