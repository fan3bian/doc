#### 主从复制原理

1-8 存量数据同步

1.开启主从复制 
```sh 
slavefo <masterIp> <masterPort> # 5.0 replicaof
```

2. 建立连接(Socket)
slave请求连接master，master在accept()后为套接字创建相应的客户端状态(多个客户端)

3. 发送Ping命令
slave -> ping -> master 
master -> pong -> slave 

4. 身份认证
密码验证

5.slave向master发送端口信息
slave向master发送自己监听的端口号，master收到后记录在slave客户端状态的slave_listening_port中

6. 发送IP地址
如果配置了slave_announce_ip，则slave向master发送slave_announce_ip配置的IP地址,master收到后记录在slave客户端状态的slave_ip中。解决主从

7. 发送capa
capa：capabilities,同步复制能力。
slave发送sapa告诉master自己具备同步复制能力，master收到后记录在slave客户端状态的salve_capa中。

8. 数据同步
slave向master发送PYSNC命令
- slave第一次执行，发送`PSYNC ? -1`，master返回`+FULLRESYNC replid offset`执行完整同步
- 如果不是第一次执行，向master发送`PSYNC <replid> <offset>` replid是复制id offest是复制偏移量 ; master 返回 `+CONTINUE <replid>`，仅执行部分同步

9. 命令传播阶段
存量数据同步完成之后，就会进入命令传播阶段。

salve默认1次/s频率向master发送命令`REPLCONF ACK <reploff>`，reploff是slave当前偏移量。

作用和意义：
1) 检测master和slave的网络连接状态
2) 汇报自己的偏移量，检测命令丢失。master会对比复制偏移量，如果发现slave偏移量小于自己，会向slave发送未发送的数据。
3) 辅助实现min-slaves，防止master在不安全的情况下执行写命令

```sh 
min-slaves-to-write 3 #
```


未完待续...