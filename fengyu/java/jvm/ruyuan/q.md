1. 什么时候触发Minor GC
对象一般都是在新生代(Eden区)分配，Eden区快满的时候就会执行Minor GC
2. 垃圾回收器用的什么？
新生代用ParNew,老年代用CMS(Current )
3.什么是Stop the word？
在进行垃圾回收的时候，会把系统程序所有的工作线程都停掉，仅垃圾回收线程在运行。
4.ParNew和Serial区别？
ParNew使用多线程，Serial使用单线程。
5.JVM调优的目的？
1.为了更少的停顿时间(stop the world) 2为了更大吞吐量
