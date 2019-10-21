#### 

1. 加载
全xin

Constant Pool: 运行时常量池
本地方法栈：C语言实现的方法。
虚拟机栈：

	栈：
		当前栈帧
		局部变量表
		操作数栈：
		动态链接
		方法出口或返回值
 	程序计数器：






Methodref
Class
类模板信息，对应堆里的

静态链接：类加载阶段，找磁盘上的class文件。

加载 -> 验证 -> 准备 -> 解析 -> 

动态链接：运行时

#### 方法区(Method Area)

Mete-Space Non-Heap(非堆)

Young Generation新生代：1/3

Old Generation老年代：2/3

分配担保机制：
	1. 大的对象
	2. 年代久的对象
Survivor:
	Eden 
	From
	To

GC: Garabage Collention
Yonng GC(Minor GC)
Full GC(Major GC)：STW 

GC Root算法：可达性算法
程序技术法

堆的设计是为了方便垃圾回收。

#### 指标
吞吐量：
停顿时间：


-Xms10m -Xmx10m -XX:+PrintGCDetails -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=路径#### jvm启动参数

```bash
export OPTS_MEMORY="-XX:+UseConcMarkSweepGC -XX:ParallelGCThreads=4 -XX:CMSInitiatingOccupancyFraction=70 -XX:+UseCMSInitiatingOccupancyOnly -XX:+CMSClassUnloadingEnabled -Xms4096m -Xmx4096m -Xmn1024m -XX:MaxPermSize=256m -XX:+UnlockExperimentalVMOptions -Xss512K -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=./java_pid<pid>.hprof $SGM_OPTS"
```


在UMP看到的应用参数
```
/export/servers/jdk1.7.0_71/bin/java -Djava.library.path=/usr/local/lib -server -Xms1024m -Xmx2048m -XX:MaxPermSize=256m -Djava.awt.headless=true -Dsun.net.client.defaultReadTimeout=60000 -Djmagick.systemclassloader=no -Dnetworkaddress.cache.ttl=300 -Dsun.net.inetaddr.ttl=300 -XX:+UseConcMarkSweepGC -XX:ParallelGCThreads=4 -XX:+UseConcMarkSweepGC -XX:CMSInitiatingOccupancyFraction=70 -XX:+UseCMSInitiatingOccupancyOnly -Xms4096m -Xmx4096m -Xmn1024m -XX:MaxPermSize=512m -XX:+UnlockExperimentalVMOptions -Xss512K -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=./java_pid<pid>.hprof -javaagent:/export/sgm-probe/sgm-probe-5.0.13-product/sgm-agent-5.0.13.jar -Xbootclasspath/a:/export/sgm-probe/sgm-probe-5.0.13-product/sgm-agent-5.0.13.jar -Dsgm.app.name=clps_goods -classpath /export/Instances/clps_goods/server1/runtime/conf:/export/Instances/clps_goods/server1/runtime/lib/* -Dbasedir=/export/Instances/clps_goods/server1/runtime -Dfile.encoding=UTF-8 com.jd.clps.master.goods.GoodsSafLauncher
```
#### CPU 

[java应用cpu使用率过高问题排查](https://blog.csdn.net/qq_34944965/article/details/81107419)
[java性能调优](https://www.jianshu.com/p/a75c1253fc11)
https://www.jianshu.com/p/3b2d71bf3eb5
