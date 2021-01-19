
## CAP理论
CAP理论是分布式系统的基石。
一个分布式系统最多只能同时满足一致性（Consistency）、可用性（Availability）和分区容错性（Partition tolerance）这三项中的两项

#### 一致性(Consistency)
all nodes see the same data at the same time. 所有节点在同一时间看到的数据完全一致

- 服务端：并发访问读取问题
- 客户端：数据最终分一致问题

三种一致性：

- 强一致性：例如关系型数据库，要求更新过的数据能被后续的访问都能看到，这是强一致性。
- 弱一致性：如果能容忍后续的部分或者全部访问不到，则是弱一致性
- 最终一致性：如果经过一段时间后要求能访问到更新后的数据，则是最终一致性。


#### 可用性(Availability):
Reads and writes always succeed，服务在正常响应时间内一直可用。


#### 分区容错性(Partition tolerance)
the system continues to operate despite arbitrary message loss or failure of part of the system.
即分布式系统在遇到某节点或网络分区故障的时候，仍然能够对外提供满足一致性或可用性的服务。