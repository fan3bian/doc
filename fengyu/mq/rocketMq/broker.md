为了保证Mq数据不丢失和具有一定的高可用，Broker部署成Master-Slave模式。

Master-Slave模式：

Slave Broker不停的发送请求到Master Broker去拉取消息

Master或Slave挂掉了会怎么样？

RocketMQ 4.5以后，Dledger机制(基于Raft协议实现),当Master Broker宕机后，通过Dleger技术和Raft协议进行leader进行选举，将Slave Broker选举为新的Master Broker。