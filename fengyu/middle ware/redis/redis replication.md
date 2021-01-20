## redis repliation
redis replication -> 主从架构 -> 读写分离 -> 水平扩容支撑高并发



###### 注意事项
Master节点一定要做持久化。如果未做持久化，宕机后重启，内存是空的。由于主从架构的同步机制，会导致从节点数据也被清空。


#### 主从架构核心原理
1. slave节点启动后，发送PSYNC命令给Master node
2. 第一次连接，会进行full resynchronization全量复制
3. master把rdb发送给slave，slave存放到磁盘上
4. 

###### 断点续传
offset
###### 无磁盘化复制
`repl-diskless-sync yes` 开启无磁盘化数据同步
`repl-diskless-sync-delay 5` 

###### 过期Key的处理
master有过期key，但是slave不会有过期key。
master的过期key，或者LRU淘汰了一个key，会模拟一条del命令发送给slave。