## 处理过程

### 调度

```
sched.Run
  sched.scheduleOne
    sched.frameworkForPod
    sched.skipPodSchedule
    sched.Algorithm.Schedule
    sched.assume
    fwk.RunReservePluginsReserve
    fwk.RunPermitPlugins
    sched.SchedulingQueue.Activate
    go sched.bind
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