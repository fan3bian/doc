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


-Xms10m -Xmx10m -XX:+PrintGCDetails -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=路径