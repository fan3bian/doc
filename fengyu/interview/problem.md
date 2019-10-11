#### 

- 如何保证redis和数据库的数据一致性
	Nosql数据库，key-value形式存储，数据存储在内存里。数据库IO和并发量是一个薄弱环节。所以对数据库进行优化，分库分表/索引。但数据库是有瓶颈的。避免数据库的并发压力，在应用和数据库中间增加缓存，如果缓存没有命中，去数据库读取，再把数据填充到缓存中。

-- Cache Aside:先更新数据库，再更新缓存
-- Read/Write 
-- 

数据库 binlog 三种形态 canal中间件订阅binlog，拿到binlog数据的操作。消息队列，


- 缓存数据预热
- 缓存穿透
- Redis里的LRU机制
	LRU：Least Recently Used ked的淘汰机制
	maxmemory-samples:随机选择最n个Key,把时间最大的删除。

- AP模型和CP模型
- 幂等性