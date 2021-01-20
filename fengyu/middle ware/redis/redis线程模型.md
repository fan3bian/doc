1. 客户端与redis服务端进行网络连接, socket与serverSocket，会产生一个AE_READABLE事件(代表一个新的客户端来连接redis了)。
2. IO多路复用程序组件监听socket，把AR_REABABLE事件放入内存队列。
3. 文件事件分派器从队列中取出socket和事件，根据事件选择对应的事件处理器(连接应答处理器、命令请求处理器、命令回复处理器)
4. 