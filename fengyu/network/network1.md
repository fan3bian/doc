## Network


#### TCP/IP 四层网络模型
应用层、传输层、网络层、链路层

#### OSI七层网络模型
OSI:Open Systems Interconnection model 开放式系统互联模型

应用层、表示层、会话层、传输层、网络层、数据链路层、物理层

应用层协议：HTTP、FTP(文件传输协议)、SNMP(简单网络管理协议)、

传输层(Transport Layer): TCP协议和UDP协议

网络层(Network Layer): IP协议、ICMP协议、ARP和RARP协议

#### TCP和UDP的区别


#### 数据链路层协议

帧(Frame)：数据链路层会把一组电信号封装成数据包(帧)，分为标头和数据。标头里包含了源和目的Mac地址。
以太网协议(Ethernet): 以太网规定接入网络设备必须有网卡，网卡地址就是Mac地址，Mac地址是惟一的。以太网协议以广播的形式发送数据包，每个终端设备铜鼓Mac地址进行匹配与过滤。 

#### 网络层协议
IP协议(Intenet Potocal): 32位2进制，前24位代表网络，厚8位代表主机。
子网掩码: 比如255.255.255.0 两个IP是否是相同子网，分别于子网掩码进行与运算，网络部分一致，则是同一个子网。