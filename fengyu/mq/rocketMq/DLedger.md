## DLedger实现Broker高可用
Broker中，只有Leader Broker可以进行数据写入，Follower Broker只能接收(实际上是拉取)同步Leader的数据。
由DLedger实现管理CommitLog

#### Leader选举
基于Raft协议实现Leader选举：多数人投票一致
第一轮每个人投给自己，失败后各节点随机休眠，先苏醒的节点投给自己，后苏醒的节点会收到先苏醒的节点的投票信息，进行附议，超过一半(大多数)投票成功即可成为leader

``votes >= (num/2)+1``

#### 多副本同步
Leader向Follower同步数据分两个阶段

1. uncommitted
2. committed
