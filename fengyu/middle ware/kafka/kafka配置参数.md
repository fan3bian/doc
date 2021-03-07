## 生产配置

#### Broker配置

Broker需要配置存储信息

log.dirs：/home/kafka1,/home/kafka2,/home/kafka3
最好挂在不同物理磁盘上：1.多块磁盘同时读写，拥有更高的吞吐量 2.Failover故障转移

Borker连接
listeners:监听器
advertised.listeners: Broker对外发布的监听器

日记留存
log.retention.{hour|mintues|ms}
log.retention.hour=168
log.retention.bytes: 指定Broker为消息保存的总容量大小 默认-1
message.max.bytes: 控制broker能够接受到的最大消息大小 默认1m 过小
#### ZooKeeper配置
分布式协调框架，负责协调和管理Kafka集群元数据信息 broker状态、topic信息、partition信息

zookeeper.connect: zk1:2181,zk2:2181,zk3:2181

管理两套Kafka集群：
zk1:2181,zk2:2181,zk3:2181/kafka1
zk1:2181,zk2:2181,zk3:2181/kafka2

#### Topic参数

auto.create.topics.enable: 是否允许自动创建Topic 建议false
unclean.leader.election.enable: 是否允许Unclean Leader选举 建议false 不允许落后太多的副本参与选举，会丢失数据
auto.leader.rebalance.enable：是否允许定期进行leader选举 建议false 不重新选举

日志留存
retention.ms 消息被保存时间，默认7天，设置后会覆盖Broker端的值
retention.bytes 为Topic预留磁盘空间，默认-1 表示无限制使用空间
max.message.bytes Broker能正常接收该Topic最大消息大小

#### JVM相关
KAFKA_HEAP_OPTS：指定堆大小
KAFKA_JVM_PERFORMANCE_OPTS：指定 GC 参数。
``
$> export KAFKA_HEAP_OPTS=--Xms6g --Xmx6g
$> export KAFKA_JVM_PERFORMANCE_OPTS= -server -XX:+UseG1GC
$> bin/kafka-server-start.sh config/server.properties
``

#### 操作系统
1. ulimit -n 1000000 文件描述符设置一个较大的值 
2. 文件系统使用XFS ZFS
3. swap 设置为较小的值1
4. 异步刷盘时间 flush 向kafka写入数据，写入操作系统页缓存(Page Cache)，操作系统根据LRU算法定期将刷到物理磁盘上，默认是5S。可适当的增加提交时间间隔降低物理磁盘的写操作。


ulimit -n 如果不设置，单机在Centos7上几百的并发就报“Too many open files”了



