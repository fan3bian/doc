1. dubbo支持哪些协议？

默认是dubbo协议

- jdk标准的rmi协议
- WebService

2. dubbo集群的负载均衡策略
Dubbo内置4了四种负载均衡策略：

- RandomLoadBalance: 随机，默认策略
- RoundBinLoadBalance：轮询
- LeastActiveLoadBalance：最少活跃数
- ConsistentHashLoadBalance：一致性hash，相同请求参数落在同一台机器上


3. dubbo配置的优先级
方法 > 服务(接口) > 全局
客户端 > 服务端

4. dubbo服务启动时检查
check="true"
dubbo.reference.check=false，强制改变所有 reference 的 check 值，就算配置中有声明，也会被覆盖。
dubbo.consumer.check=false，是设置 check 的缺省值，如果配置中有显式的声明，如：`<dubbo:reference check="true"/>`，不会受影响。
dubbo.registry.check=false，前面两个都是指订阅成功，但提供者列表是否为空是否报错，如果注册订阅失败时，也允许启动，需使用此选项，将在后台定时重试。

5. dubbo集群容错
- Failover Cluster:失败转移，失败时重试其他服务器，通常用于读操作`retries=2`，默认重试次数
- FailFast Cluster:快速失败，发起一次调用，失败报错
- Failsafe Cluster:失败安全，出现异常时，忽略
- Failback Cluster:失败自动恢复，后台记录失败记录，定时重发
- Forking Cluster:并行访问多个服务器，只要有一个返回就成功。

这里的 Invoker 是 Provider 的一个可调用 Service 的抽象，Invoker 封装了 Provider 地址及 Service 接口信息
Directory 代表多个 Invoker，可以把它看成 List<Invoker> ，但与 List 不同的是，它的值可能是动态变化的，比如注册中心推送变更
Cluster 将 Directory 中的多个 Invoker 伪装成一个 Invoker，对上层透明，伪装过程包含了容错逻辑，调用失败后，重试另一个
Router 负责从多个 Invoker 中按路由规则选出子集，比如读写分离，应用隔离等
LoadBalance 负责从多个 Invoker 中选出具体的一个用于本次调用，选的过程包含了负载均衡算法，调用失败后，需要重选