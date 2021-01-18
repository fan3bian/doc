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



- java语言如何进行内存管理的？
java的内存管理工作是JVM来完成的，内存的分配和回收都是JVM自动来完成的。
- JVM内存预取划分
java内区域划分：堆、方法区、程序计数器、虚拟机栈、本地方法栈。方法区在jdk1.8之后，称为元空间。
- java堆是如何划分的
新生代和老年代，新生代又分为Eden区和两个susvivor区。
- java代码是怎么运行起来
java代码(.java文件)编译成字节码(.class文件),JVM加载二进制的字节流到内存中，生成Klass元数据。从启动类的main线程开始运行。
- 类的加载器
- 类加载的过程
- 双亲委派模型
- JVM在什么情况下会加载一个类







