
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


CAP 理论指出对于一个分布式计算系统来说，不可能同时满足以下三点：
**一致性**：在分布式环境中，一致性是指数据在多个副本之间是否能够保持一致的特性，等同于所有节点访问同一份最新的数据副本。在一致性的需求下，当一个系统在数据一致的状态下执行更新操作后，应该保证系统的数据仍然处于一致的状态。
**可用性**：每次请求都能获取到正确的响应，但是不保证获取的数据为最新数据。

**分区容错性**：分布式系统在遇到任何网络分区故障的时候，仍然需要能够保证对外提供满足一致性和可用性的服务，除非是整个网络环境都发生了故障。

一个分布式系统最多只能同时满足一致性（Consistency）、可用性（Availability）和分区容错性（Partition tolerance）这三项中的两项。


BASE 理论
BASE 是 Basically Available(基本可用)、Soft-state(软状态) 和 Eventually Consistent(最终一致性) 三个短语的缩写。

基本可用：在分布式系统出现故障，允许损失部分可用性（服务降级、页面降级）。

软状态：允许分布式系统出现中间状态。而且中间状态不影响系统的可用性。这里的中间状态是指不同的 data replication（数据备份节点）之间的数据更新可以出现延时的最终一致性。

最终一致性：data replications 经过一段时间达到一致性。
BASE 理论是对 CAP 中的一致性和可用性进行一个权衡的结果，理论的核心思想就是：我们无法做到强一致，但每个应用都可以根据自身的业务特点，采用适当的方式来使系统达到最终一致性。