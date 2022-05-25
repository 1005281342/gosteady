## CommitFailedException异常

指的是Consumer客户端在提交位移的时候出现了错误或异常，而且还是那种不可恢复的严重异常。



### 案例

消费者组开启了ReBalance过程，并且将要提交位移的分区分配给了另一个消费者实例。使得Consumer实例连续两次调用poll方法的时间间隔查过了期望的`max.poll.interval.ms`参数值。

**解决方法有两个**：

1. 增加期望的时间间隔`max.poll.interval.ms`参数值
2. 通过调整`max.poll.records`参数值，减少poll方法一次性返回的消息数量



## 手动提交时抛出CommitFailedException异常的场景

### 消息处理总时间超过预设的` max.poll.interval.ms`参数值

消息处理之后进行手动提交，但是间隔已经操作了consumer端`max.poll.interval.ms`参数值

解决方法：

1. 缩短单条消息处理的时间
2. 增加`max.poll.interval.ms`的时间
3. 减少下游系统一次性消费的消息总数
4. 下游系统使用多线程来加速消费

