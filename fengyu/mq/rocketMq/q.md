Q: NameServer宕机后MQ还能继续运行吗？生产者还能发送消息到Broker吗？消费者还能从Broker里拉取消息吗？

A: 注册中心nameserver宕机后通过producer本地缓存的broker来进行数据的发送，如果producer重启了则不能发送数据 跟dubbo讲提供者信息注册到zk上消费者通过zk获取也是消费者本地缓存提供者的信息即zk宕机后可用