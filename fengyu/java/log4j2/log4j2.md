## Log4j2 

#### log4j2 configuration file

##### Level

All < Trace < Debug < Info < Warn < Error < Fatal < OFF



#### 关于nohup的使用
当关闭console日志appander时，当spring加载bean失败后，会触发自动销毁。但缺无法在日志appender中输出具体的原因。
 Destroying singletons in org.springframework.beans.factory.support.DefaultListableBeanFactory@3ee82600:


####  log4j2.component.properties

```java 
# 缓存时间戳,不会每个日志都去取system.curTime,而是缓存1毫秒或者1024条日志
log4j.Clock=CachedClock
# 异步缓存大小 512*1024
AsyncLogger.RingBufferSize=524288
# 写入磁盘io等待的时候,阻塞等待,较少锁竞争.Block在容器环境比较适合,cpu资源比日志吞吐重要的情况
AsyncLogger.WaitStrategy=Block
# 同上
AsyncLoggerConfig.RingBufferSize=524288
# 同上
AsyncLoggerConfig.WaitStrategy=Block
# 队列满了 丢弃
log4j2.AsyncQueueFullPolicy=Discard
# 丢弃日志级别
log4j2.DiscardThreshold=TRACE
```

#### MDC
MDC：Mapped Diagnostic Context(映射诊断上下文)

traceId : 调用链ID
spanId : 当前ID
parentId : 父ID


#### PatternLayout

<PatternLayout  charset="UTF-8" pattern="-5p - %d{yyyy-MM-dd HH:mm:ss} [%t] %l -- %m%n"/>
-5p - 2019-10-09 18:01:16 [JSF-BZ-22000-90-T-14] -- [transfer->CBS4418046561925]查询预占库存入参(4418046511713,385,CMG4418099962616,1,1,100)
