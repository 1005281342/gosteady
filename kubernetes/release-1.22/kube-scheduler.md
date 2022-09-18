## 处理过程

### 调度

```
sched.Run
  sched.scheduleOne
    sched.frameworkForPod（根据名字获取调度器）
    sched.skipPodSchedule（是否跳过本次调度）
    sched.Algorithm.Schedule（执行调度，默认使用通用调度器）
    sched.assume（pod和node的绑定信息更新缓存）
    fwk.RunReservePluginsReserve（预扣除资源，如扣除磁盘资源之类的）
    fwk.RunPermitPlugins（权限校验）
    sched.SchedulingQueue.Activate（推送pod到ActivateQ）
    go sched.bind（异步绑定）
```

### 通用调度器

```
sched.Algorithm.Schedule（genericScheduler通用调度器）
  g.snapshot (从缓存中拉取节点信息)
  g.findNodesThatFitPod（查找可用节点）
    fwk.RunPreFilterPlugins（过滤前执行的动作）
    g.findNodesThatPassFilters（过滤节点）
      g.numFeasibleNodesToFind（获取期望可行节点数，达到这个数量后就不会继续判断处理后面的节点了）
      fwk.RunFilterPluginsWithNominatedPods（执行过滤插件）
    findNodesThatPassExtenders（执行扩展过滤插件）
  prioritizeNodes（处理节点优先级）
    fwk.RunPreScorePlugins（打分前执行的动作）
    fwk.RunScorePlugins（打分）
  g.selectHost（选择主机）
```



## 参考

[Kubernetes Scheduler 设计与实现](https://www.bilibili.com/video/BV1N7411w7M9?vd_source=261fdf90969b104ee0e32522eed85cba)