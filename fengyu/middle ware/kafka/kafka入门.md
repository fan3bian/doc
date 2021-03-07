## Kafka入门

#### 核心概念
Producer Broker Consumer Topic Partition Replica

#### Broker
<!-- Broker集群由一个Leader Broker和其他Follower Broker组成，连接同一个Zookeeper集群。Leader在kafka启动时选举出来
 -->
#### Consumer Group
Topic里的消息，一个Cunsumer Group只能消费一次，由其中某一个consumer进行消费

#### Topic 
消息订阅和发布的主题

#### Partition
分区，分区的作用是把消息分散到不同的broker上，提高并发能力。
1. 一个Topic可以分成多个partition，一个partition只能被一个consumer消费，一个consumer可以消费多个partition。多对一
2. 消费效率最高的情形是：一个partition由一个consumer消费
3. consumer数量大于partition时，会出现consumer闲置
4. Cousumer Group订阅者集群，每个consumer负责自己的消费分区

#### Replica
副本，副本的用途是备份。Leader Replica + Follower Replica

- AR: assigned replicas AR列表，每个partition存有最初分批的AR列表，不会改变
- PR: AR列表里第一个
- ISR: In Sync Replicas 同步副本

分区同步

#### Broker 
Broker启动时，会实例化一个KafkaController，并将broker的id注册到Zookeeper。集群启动过程中，选举一个broker做为leader

#### Leader Broker 控制器
分区leader的选举，选举partition的leader replica.

控制器选举(Leader选举)：
- 集群启动：
- 控制器所在代理故障：
- Zookeeper心跳机制感知控制器session过期

Zookeeper临时节点  controller_epoch 存储leader变更次数，初始为0，leader每变更一次，+1。
起到一个版本号的作用，保证集群控制器唯一。