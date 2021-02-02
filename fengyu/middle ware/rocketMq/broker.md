为了保证Mq数据不丢失和具有一定的高可用，Broker部署成Master-Slave模式。

Master-Slave模式：

Slave Broker不停的发送请求到Master Broker去拉取消息

Master或Slave挂掉了会怎么样？

RocketMQ 4.5以后，Dledger机制(基于Raft协议实现),当Master Broker宕机后，通过Dleger技术和Raft协议进行leader进行选举，将Slave Broker选举为新的Master Broker。


#### Broker数据存储
Broker存储数据是MQ最核心的环节，它决定了生产者写入消息和消费者拉取消息的吞吐量，决定了消息不能丢失。


#### CommitLog消息顺序写入机制
CommitLog文件 末尾追加(顺序写入) 1GB固定大小 满了换新的
CommitLog的写入速度，影响到性能(吞吐量)。写入需要10ms,一个线程1s能执行100个消息，一百个线程才能处理10k个消息(TPS)。

提升写入性能机制：
1. PageCache：写入OS的PageCache内存缓存中，由OS后台线程异步刷盘
2. 顺序读写： 顺序写 > 随机写

异步刷盘有数据丢失风险，


#### ConsumseQueue
在Broker中，对Topic下每个MessageQueued都有一系列的ConsumeQueue(数量和MessageQueue一致)文件
`$HOME/store/consumequeue/${topic}/${queueId}/${fileName}`

文件里记录了每个消息在CommitLog文件中的物理位置(offset).



#### 
1. 为什么Broker数据存储机制是MQ最为核心的环节
2. CommitLog数据存储机制
3. MessageQueue对应的ConsumeQueue物理位置存储机制
4. 基于CommitLog顺序写+OS Cache + 异步刷盘的高吞吐消息写入机制
5. 同步刷盘和异步刷盘优缺点，同步数据不丢失，写入吞吐量下降；高吞吐量写入，有丢失数据风险。


### Broker网络通信机制

- Producer请求建立连接，Reactor主线程在端口上监听，建立长连接
- Reactor线程池并发的监听多个连接的请求是否到达
- Worker线程池对多个请求进行预处理(SSL、编码解码)
- SendMessage线程池对多个请求进行磁盘读写业务