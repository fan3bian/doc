## Internal Topic

位移主题：`_consumer_offset`  消费者 偏移量

key-value格式 
key: `<groupId> <topic> <partition>`


#### 位移主题创建

创建时机：当kafka集群第一个consumer启动时，Kafka自动创建位移主题。

broker端参数
分区数：offsets.topic.nums.partitions  默认50
副本数：offsets.topic.replication.factor  默认：3

#### 写入主题
位移主题是为了记录kafka消费进度的，Kafka消费者提交消费位移时，向该主题写入消息。

Kafka提交位移的方式：
- 自动提交位移
- 手动提交位移

conusemr端参数：
enable.auto.commit 如果是true,kakfa会定期提交位移，时间间隔auto.commit.interval.ms控制；
一旦设置为false，consumer应用开发需要做提交位移，API: consumer.commitSync

删除消息：
kafka使用Compact策略来删除位移队列的中的过期消息，避免该主题无限膨胀。

使用Log Cleaner线程定期扫描，key(consuemr group topic partition)相同的不通消息，只保留最新的一条。
