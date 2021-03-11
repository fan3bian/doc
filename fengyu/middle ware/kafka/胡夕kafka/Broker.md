## Broker

#### Coordinator


#### 消息处理

1. Reactor模型

Dispatcher : SocketServer组件
Acceptor
工作线程池 : 网络线程池 
IO线程池：
Broker端参数
`num.network.threads`: 网络线程池的线程数。默认3，每台broker启动时会创建3个线程，专门处理客户端发送的网络请求。
`num.io.threads`：默认8，Broker启动后自动创建8个IO线程处理请求。

1. Acceptor线程采用轮训的方式将入站请求公平的发送给所有的网络线程中。
2. 网络线程接收到请求后，将请求放入共享队列中。
3. IO线程池负责从队列中获取请求，执行真正的处理。
4. IO线程处理完请求后，发送响应到网络线程池的响应队列中。
5. 网络线程将响应体发送给客户端

请求队列是网络线程共享，响应队列是网络线程专属。

请求分类
- 数据类请求：生产请求PRODUCE，将消息磁盘日志；消费请求FETCH，从pagecache获取磁盘中读取消息
- 控制类请求：LeaderAndIsr请求，负责更新leader、follower和ISR集合；StopReplica勒令副本下线。

Purgatory：缓存延时请求(Delayed Request),未满足条件不能立刻处理的请求。比如设置了acks=all的PRODUCE请求，需要等待ISR中所有副本都接受到了消息才能返回，此时处理该请求的IO线程必须等待其他Broker的写入结果。当请求不能立刻处理时，就会方式Purgatory