## Rocket Mq

#### NameServer
负责管理集群里所有的Broker信息
Producer和Consumer通过NameServer知道和哪个具体的Broker进行通信，
Producer和Consumer定时发送请求到NamServer去拉取最新的Broker集群信息。

NameServer需使用集群部署(高可用),每个Broker启动时向所有NameServer进行注册，所以每个NameServer都会有集群中所有Broker信息。

NameServer通过Broker的心跳机制判定存活

担当路由中心角色

#### Broker
Broker通过多台机器部署集群，使用主从架构实现多副本存储和高可用

#### 生产者

#### Consumer消费者
Consumer每次从Broker里拉取消息

首先启动注册中心nameserver，每个nameserver之间互不通信，启动broker时，会把自己的信息注册到每一个nameserver，broker每30s发送心跳包给注册中心，注册中心更新broker的最后更新时间。nameserver会定时10秒检测更新时间是否超过120s，超过则将这个broker路由原信息剔除。生产者和消费者定时去获取broker的路由信息，根据轮询生产消息和消费消息的负载。当注册中心挂了，本地还会有缓存信息能够继续通信。



