1. redis和memchached的区别？
redis和memcached都是内存数据库，memcached是早期互联网公司的技术选型，现在几乎所有的主流公司都在用redis。相对于memchached，redis具有几个显著的优点。
(1) redis支持更多的数据类型和更丰富的的操作,,memcached只支持简单的字符串类型
(2) redis是单线程模型，memcached使用的是多线程模型
(3) redis原生支持集群模式(1.3以后支持cluster模式)，memcahched需要依靠客户端来实现往集群里分片写入数据
2. 缓存的优点
高性能： 读多写少的场景，从数据库读取的代价要比从缓存中读取的代价要高
高并发： 数据库并发过高单机(比如4000/s),可能会出现宕机的，高峰期会出现问题。数据库建议不超过2000个请求每秒

