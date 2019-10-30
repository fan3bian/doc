#### 调优经验
export OPTS_MEMORY="-XX:+UseConcMarkSweepGC -XX:ParallelGCThreads=4 -XX:CMSInitiatingOccupancyFraction=70 -XX:+UseCMSInitiatingOccupancyOnly -XX:+CMSClassUnloadingEnabled -Xms4096m -Xmx4096m -Xmn1024m -XX:MaxPermSize=256m -XX:+UnlockExperimentalVMOptions -Xss512K -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=./java_pid<pid>.hprof -XX:+PrintGCDetails -XX:+PrintHeapAtGC -XX:+PrintGCDateStamps  -XX:+PrintTenuringDistribution -verbose:gc  -Xloggc:/export/Logs/po.clps.jd.com/gc.log $SGM_OPTS"

Problem: 从sgm上发现clps_po应用的young gc的频率有些高。

> 最近1小时58分钟共发生714垃圾回收，平均9秒发生一次，每次耗时12毫秒。 最近10分钟发生了72次垃圾回收，耗时791毫秒。

每九秒一次young gc,实际上频率的最高的时间段，大约每4秒发生一次young GC。认为频率略微高一些，其他应用大概十几秒，30几秒甚至有的1分钟发生一次。

应用使用doceker部署，4C 8G jdk1.7 
| 参数 | 含义 |
| -- | -- | 
|-Xms4096m | 初始堆4096m |
|-Xmx4096m | 最大堆4096m |
|-Xmn1024m | 新生代1024m |
|-XX:+UseConcMarkSweepGC |使用CMS垃圾回收器 |
|-XX:CMSInitiatingOccupancyFraction=70 | 老年代容量到达70%触发CMS回收|
|-XX:MaxPermSize=256m | 持久代256m |

未指定-XX:SurvivorRatio(Eden区与survivor区的比例)，jvm使用默认值8

这样可得出:
Young Gen: 1024
	Eden: 819.2
	Survivor: 102.4
Old Gen:3096
这与sgm得到的数据基本一致：
 heap: 3993m
	 Old Gen: 3072m
 non-heap: 304m
 	Perm Gen: 256m
	Code Cache: 48m

新生代占堆的比例为：0.25，官方推荐是3/8，小于官方推荐。

