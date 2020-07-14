1. ==和equals方法的区别
2. String里面的equals是如何实现的？
2.1.1 String StringBuffer StringBuilder的区别是什么？
2.1 fianl关键字的用法
2.2 抑制序列化 transient
3. 常见的集合类
4. ArrayList的实现原理和扩容机制
4.1 如何在循环中删除一个元素？在for循环中删除一个元素会出现什么情况 fail-fast
5. HashMap的实现原理和扩容机制？
	5.1 put方法是如何获得value的
6. java内存模型，那一部分不会发生内存溢出
7. 垃圾回收?垃圾回收算法？
标记 -清除算法
复制算法
标记-压缩算法
分代收集算法
8. 判断对象是否存活一般有
引用计数：每个对象有一个引用计数属性，新增一个引用时计数加1，引用释放时计数减1，计数为0时可以回收。此方法简单，无法解决对象相互循环引用的问题。
可达性分析（Reachability Analysis）：从GC Roots开始向下搜索，搜索所走过的路径称为引用链。当一个对象到GC Roots没有任何引用链相连时，则证明此对象是不可用的。不可达对象。
9. 垃圾回收器
Serial收集器
ParNew收集器
Parallel收集器
Parallel Old 收集器
CMS收集器
10. 线程的声明周期：
新建状态 就绪状态 运行状态 阻塞状态 死亡状态
11. 异常的分类有哪些？
12. 一个class文件，如何判断是由哪个版本jdk编译的
13. 泛型的作用？泛型擦除


#### Mysql
1. 事务有哪些特性？原子性 一致性 隔离性 持久性
2. mysql提供的四种隔离级别：
　① Serializable (串行化)：可避免脏读、不可重复读、幻读的发生。
　② Repeatable read (可重复读)：可避免脏读、不可重复读的发生。Mysql默认
　③ Read committed (读已提交)：可避免脏读的发生。
　④ Read uncommitted (读未提交)：最低级别，任何情况都无法保证。
3.

#### Spring 
1. Srping的核心概念是什么？
2. AOP的用途


## 缓存
#### redis

#### Mybatis
Mybatis的一级缓存和二次缓存

#### 
