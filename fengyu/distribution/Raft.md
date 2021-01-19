
Raft协议：
Raft协议将一致性协议的核心内容分拆成为几个关键阶段，以简化流程，提高协议的可理解性。

Leader election

Raft协议的每个副本都会处于三种状态之一：Leader、Follower、Candidate。

Leader：所有请求的处理者，Leader副本接受client的更新请求，本地处理后再同步至多个其他副本；
Follower：请求的被动更新者，从Leader接受更新请求，然后写入本地日志文件
Candidate：如果Follower副本在一段时间内没有收到Leader副本的心跳，则判断Leader可能已经故障，此时启动选主过程，此时副本会变成Candidate状态，直到选主结束。