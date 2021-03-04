https://mp.weixin.qq.com/s/Y0jYbDhs5ARbDbkrnA7Rjg


Kafka是一种高吞吐量、持久性、分布式的发布订阅的消息队列系统。

支持负载均衡、数据持久化。

1. Kafka如何实现高吞吐量的？
2. Kafka持久化(消息存储机制)是怎么实现的？ 
3. kafka负载均衡是怎么实现的？
4. kafka如何实现故障转移
5. Kafka生产消息的过程
6. Kafka消费消息的过程
7. 消息队列设计中，为什么设计成consumer主动去拉去消息，而不是broker推送给consumer
8. Kafka如何实现状态管理

#### Kafka中的角色
- Producer:消息和数据的生产者，主要负责生产push消息到指定Broker的Topic中
- Broker：Kafka节点或Kafka进程称为Broker。Broker功能：创建Topic，存储Producer发布的消息，记录消费处理
- Topic: 同一个Topic消息可以分布在一个或多个Broker上，一个Topic包含一个或多个Partition分区，数据存储在partition中
- Partition：分区，Topic物理上的分组，一个Topic在Broker上分为一个或多个Partition,分区是真正存储数据的单元。
- Consumer：消息和数据的消费者，负责主动到已订阅的Topic中拉去消息并消费
- ZooKeeper：Zookeeper负责维护整个Kafka集群状管理，存储kafka各个节点的状态，协调kafka工作内容。

> Kafka 0.8以前，Producer从ZooKeeper获取集群状态；
0.8以后，在Producer中配置多个Broker节点发送消息，同时跟broker建立连接，从broker获取集群状态。

#### Topic和log
一个topic有多个分区，每个分区维护一个日志。
日志是有序的，通过追加的方式写入

消费偏移量offset

#### Partition
容错：分区使用副本的方案
Leader Partition：单节点，处理读写请求
Follower Partition：多节点，被动同步leader的数据

leader选举
#### 持久化
日志的保留策略

#### Geo Replication 地域复制
Mirrormarker为集群提供地域复制

异地容灾：数据备份和恢复

#### Producer
生产者 -> topic -> 分区

#### Consumer
Consumer Group：消费者组
