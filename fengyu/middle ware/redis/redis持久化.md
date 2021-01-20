使用redis做为缓存，redis基本上是基于内存完成的操作，数据都在内存里。一旦redis服务宕机了，内存中数据将全部丢失，怎么保证数据不会丢失？

## redis持久化
持久化的意义在于数据恢复。
redis持久化主要依靠两种机制：AOF日志和RDB快照。

#### AOF日志
Append Only File，类似于binlog，AOF日志是redis执行命令后，写入AOF。

###### AOF两个风险
1. 写入AOF日志不阻塞当前命令，AOF日志同样也是主线程执行，如果把日志文件写入磁盘，写入压力大，会阻塞下个操作
2. redis执行完命令后宕机，AOF日志没有写入磁盘，导致数据丢失

###### AOF写入磁盘策略
appendfsync用于配置AOF写入磁盘策略
- Always：同步写入
- EverySec：每秒写入，把日志写入AOF文件内存缓冲区，每秒把缓冲区内容写入磁盘。
- No：写入文件内存缓冲区，由操作系统控制写入磁盘。

trade-off: 性能和可靠性之间取舍，避免主线程阻塞



###### AOF文件重写机制
redis重写机制(rewrite)，可避免AOF文件过大
重写是通过bgwriteaof线程完成，不使用主线程，不会造成阻塞。
redis先进行内存拷贝，根据内存里数据最新状态，生成这些数据的插入命令，作为新日志。

`auto-aof-rewirte-percentage 100`: AOF执行条件,当前AOF文件大小超过上次100%
`auto-aof-rewrite-min-size 64m`: AOF执行条件,当前AOF文件大小超过64m

Q:如果执行命令前写入有什么问题？
1. 命令可能执行失败 2. 命令错误(语法) 


#### AOF持久化
数据写入AOF文件，数据通过AOF线程写入OS cache，OS线程周期性的把OS cache写入磁盘，刷盘。

AOF文件只有，会越来越大。redis内存是固定大小的，内存淘汰策略会，AOF rewrite操作会根据redis当前内存中的数据生成新的AOF

RDB和AOF都开启，会使用AOF来恢复。AOF是更细粒度的数据


#### 生产环境配置(企业级)
1. RDB适合做冷备，文件一旦生成，就不会发生变化。配置中，redis中RDB是默认是默认启用的
save `<seconds> <changes>` 时间范围key的变更数量达到目标值，进行RDB的备份

```
save 900 1
save 300 10
save 60 10000
```
2. 配置中，AOF默认是关闭的，生产环境一定要开启。fsync策略使用everysec
```
appendonly no
```

###### RDB 数据备份方案
(1) crontab定时调度脚本去做数据备份
(2) 按小时、天周期性的做备份