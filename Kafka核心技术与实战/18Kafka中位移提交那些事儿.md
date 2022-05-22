## 提交位移

consumer需要向Kafka汇报自己的位移数据，这个汇报过程被称为提交位移。

一个consumer能够同时消费多个分区，因此consumer需要为它所消费的每个分区提交各自的位移数据。

## 提交模式

### 用户角度

- 自动提交（此模式通过`auto.commit.interval.ms`设置自动提交间隔，默认为5000ms）
- 手动提交（可以通过设置`enable.auto.commit=false`关闭自动提交）

## consumer角度

- 同步提交
- 异步提交

## 自动提交造成消息重复消费案例

假设自动提交间隔为5秒，而再消费到第3秒的时候触发了rebalance，这个时候前3秒已经被处理的消息会在rebalance后重新消费。

