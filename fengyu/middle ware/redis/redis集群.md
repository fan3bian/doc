
## Redis集群方案

redis-cluster高可用架构方案是目前最好的redis架构方案，redis的官方对redis-cluster架构是建议redis-master用于接收读写，而redis-slave则用于备份（备用），默认不建议读写分离。