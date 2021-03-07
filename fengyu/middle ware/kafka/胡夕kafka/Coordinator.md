## 协调者Coordinator

每个Broker都有一个Coordinator，用于Consumer Group的元数据管理。

- 消费者组的注册
- 组成员管理记录
- 提交位移管理

Broker在启动时，会创建和开启响应的Coordinator。


#### 如何确定Consumer Group的Coordinator

1. 确定分区：`partitionId = Math.abs(groupId.hashCode % offsetsTopicPartitionCount)`  offsetsPartitionCount = 50(`_consumer_offset`)
2. 获取分区Leader副本所在Broker

#### Consuemr状态管理

心跳机制：每个Consuemr定期的向Coordinator发送心跳请求，表示存活。如果Consumer不能及时的发送心跳请求，Coordinator会认为Consuemr死了，将其从group中移除，然后开启新一轮的rebalance。

consumer端参数
+ `session.timeout.ms`：默认10s 决定了consumer存活性的时间间隔
+ `heartbeat.interval.ms`：这个值越小，consumer实例发送心跳请求的频率越高，频繁的心跳请求会消耗带宽资源，但也会快速的得知Coordinator是否启动rebalance。

Coordinator会在consumer心跳请求响应体内，设置是否开启rebalance，标志REBALANCE_NEEDED。

+ `max.poll.interval.ms`: 两个调用poll方法最大时间间隔，默认5分钟

#### 可避免Rebalance场景
1. 未能及时发送心跳导致Consumer实例被Coordinator被踢出group 
2. Consumer消费时间过长，触发`max.poll.interval.ms`

