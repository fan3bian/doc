## Consumer Group
Consumer Group是kafka提供的可扩展的且具容错性的消费机制。每个Consumer Group在kafka集群有一个唯一的Group ID，Consumer Group内可以有多个Consumer实例，协调一起来消费订阅主题(Subscribed Topics)的所有分区(Partition)，每个分区只能由组内一个消费者消费。

Consumer Group之间相互独立，互不影响。

点对点消费队列模型：所有实例同属于一个Group
发布/订阅模型：所有实例分别属于不同的Group

#### 设置Consumer数量

理想情况下：Consumer 实例的数量应该等于该Group订阅主题的分区总数。如果实例数大于分区总数，只会造成资源浪费。

#### 消费者进度

Consumer Offset： Key-value key:partition value:consumer offset

旧版的Consumer Group把位移保存在ZooKeeper上，减少了Kakfa Broker端状态保存的开销。这样，Kafka broker服务器节点是无状态的，这样可以自由的扩容缩容，具有较高的伸缩性。缺点是Zookeeper不适合进行频繁的写更新，Consumer Group的位移更新是非常频繁，大吞吐量的操作会拖慢ZooKeeper集群性能。

新版本：Consumer Group 将位移保存到Broker端内部主题`_consumer_offsets`中

#### Rebalance机制

Rebalance本质是一个协议，规定一个Consumer Group下所有Consumer如何达成一致，来分配订阅Topic所有分区。平均分配
average = num of partition / num of consumer

Rebalance触发机制：
1. consumer group成员变更，新的实例加入或旧实例退出
2. 订阅主题数变更，comsumer group 支持表达式订阅主题`topic*`,在consuemr group运行中，创建了tpoicAAA。就会触发rebalance
3. 订阅主题的分区数发生变更，kafka目前只允许增加。

分配策略：kafka consuemr group 中 consumer实例分配partition的策略。

###### 缺点
Rebalance过程中，组内所有的consumer实例都会停止消费，等待rebalance完成。



#### 消费位移

消费位移记录了消费者消费消息的进度，Consuemr通过想Kafka提交汇报自己的位移数据(分区粒度)。
有多重位移提交的方式，灵活。Broker端并不对位移提供语义保障，完全由客户端自己保障。

以用户角度来说：自动提交和手动提交
以Consuemr角色来说：同步提交和异步提交

Consuemr端配置参数
`enable.auto.commit`:默认true
`auto.commit.interval.ms`: 自动提交频率，默认5s

手动提交`kafkaConsumer.commmitAsync`

自动提交：自动提交会造成重复消费。位移提交不够实时，如果发生了rebalance，当前具体上次提交的期间的消息会重新消费一次。
手动提交：调用`commitSync`时，consumer线程会阻塞，直到远程broker端返回提交结果。

#### CommitFailedException

场景1：消息处理时间超过预设的`max.poll.interval.ms`


#### 心跳线程
Consuemr里的心跳线程(Heartbeat Thread)负责定期给对应Broker机器发送心跳请求，以标识消费者的存活性(liveness).

主线程是单线程模型


#### 消费者建立连接
在调用`KafkaConsumer.poll`时，创建消费端的TCP连接。
1. 确定协调者和获取集群元数据：消费者向Broker(最小负载)发送findCoordinator请求
2. 连接协调者，令其执行组成员管理操作：通过`findCoordinator`响应结果，向Coordinator broker发起连接
3. 执行实际的消息获取：消费数据时，消费端和消费分区领导者副本所在的Broker建立连接

协调工作：加入组、等待组分配方案、心跳请求处理、位移获取、位移提交

#### 消费进度

消费滞后程度(Consumer Lag)：消费者滞后于生产者的程度。一般在主题级别上讨论Lag，实际上kafka监控lag在分区上。

一个正常的消费者，它的lag值很小接近于0，这表示消费者能及时消费生产者生产出来的消息，滞后程度很小。

如果一个消费者Lag很大，表示消费速度无法跟上生产者速度，Lag过大会拖慢下游生产的消息处理速度。最可怕是有可能导致消息从Broker操作系统的的内核缓冲区pageCache中，消费者需要从磁盘中的获取消息，马太效应，Lag值越来越大。
