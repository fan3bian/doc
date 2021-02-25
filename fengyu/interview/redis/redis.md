#### redis线程模型
redis


#### redis持久化
1. redis持久化的目的？
- 数据恢复

2. 持久化方式
- RDB
- AOF
- 混合持久化

###### RDB

```sh
#save ponit
save 900 1 #900s 1个key发生了变化
save 300 20
save 60 10000
```
内存快照 保存内存中所有key和value生成rdb文件 二进制 压缩 

save basave
save 阻塞主进程
bgsave fork()创建子进程来生成rdb文件，在fork过程阻塞。

rdb优点：
1. 压缩 二进制 空间占用小 适合备份 多个save point多个版本
2. 灾难恢复 机器故障，传送其他数据中心
3. 性能较好 仅fork过程阻塞
4. 数据恢复速度要比aof快


###### AOF

```sh
appendonly yes #yes开启 no关闭 默认关闭

appendfsync everysec # 

# 当前AOF文件比上次重写后的AOF文件大小的增长比例超过100
auto-aof-rewrite-percentage 100 
# 当前AOF文件的文件大小大于64MB
auto-aof-rewrite-min-size 64mb
```
AOF分三个过程：
- 命令追加
  write命令完成后，追加写入aof_buf末尾。
- 文件写入 + 文件刷盘 将aof_buf写入Aof文件

AOF重写机制：AOF文件过大时，后台线程会自动重写(包含恢复当前数据集最小集合)

BGREWRITEAO ： 阻塞主进程
BGREWRITEAOF： fork创建子进程进行AOF重写，阻塞只发生在fork子进程

解决了主进程阻塞,引入了重写过程数据不一致问题，引入**重写缓冲区（aof_rewrite_buf_blocks）**。 fork子进程开始，redis主线程写完一个命令后，同时追加到AOF缓冲区和AOF重写缓冲区。 
serverCron检测到子进程完成AOF重写后：
1) 将AOF重写缓冲区重所有内容写入新的AOF
2）对AOF进行改名，原子性完成新旧文件替换。

appendfsync 参数
1）always：每处理一个命令都将 aof_buf 缓冲区中的所有内容写入并同步到AOF 文件，即每个命令都刷盘。
2）everysec：将 aof_buf 缓冲区中的所有内容写入到 AOF 文件，如果上次同步 AOF 文件的时间距离现在超过一秒钟， 那么再次对 AOF 文件进行同步， 并且这个同步操作是异步的，由一个后台线程专门负责执行，即每秒刷盘1次。
3）no：将 aof_buf 缓冲区中的所有内容写入到 AOF 文件， 但并不对 AOF 文件进行同步， 何时同步由操作系统来决定。即不执行刷盘，让操作系统自己执行刷盘


Aof优点：
1) AOF文件的里数据更完整，更可靠
2）AOF文件有序保存了所有的写入操作命令，以redis协议给保存，易读

Aof缺点：
1)保存相同的数据集，AOF文件要比rdb大
2）appendfsync的会占用性能

#### 混合持久化
redis4.0引入 仅发生在**AOF重写** 
格式：前半段是rdb格式全量数据，
```sh
aof-use-rdb-preamble no # yes 4.0默认关闭 6.0默认打开
```
混合持久化本质是通过 AOF 后台重写（bgrewriteaof 命令）完成的，不同的是当开启混合持久化时，fork 出的子进程先将当前全量数据以 RDB 方式写入新的 AOF 文件，然后再将 AOF 重写缓冲区（aof_rewrite_buf_blocks）的增量命令以 AOF 方式写入到文件，写入完成后通知主进程将新的含有 RDB 格式和 AOF 格式的 AOF 文件替换旧的的 AOF 文件。

优点：结合 RDB 和 AOF 的优点, 更快的重写和恢复。

#### redis事件

- 文件事件
常见的的文件事件：accept() read write close
- 时间事件(serverCron): serverCron 默认100ms执行一次，清理过期的key，Aof后台重写，rdb的save point检查，aof_buf刷盘(flushAppendOnlyFile)
