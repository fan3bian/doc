（1）什么是散列表？
	通过关键码值直接访问的数据结果
（2）怎么实现一个散列表？
	数组+链表
（3）java中HashMap实现方式的演进？
	1.8增加了红黑树
（4）HashMap的容量有什么特点？
	数组的初始容量16，都是2的N次幂
（5）HashMap是怎么进行扩容的？
	每次扩容增加之前的一倍
（6）HashMap中的元素是否是有序的？
	无序的
（6.1）HashMap何时进行树化？何时进行反树化？
	HashMap在jdk1.8后引入红黑树结构，若桶中链表元素个数到达8的时候，链表转换为树结构，若桶中链表元素到达6的时候，树结果还原成链表。
 (7) HashMap是如何解决Hash碰撞的？
	发生Hash碰撞后，HashMap使用链表的方式存放hash值一致的元素，当链表的长度到达树化的阈值时，就会转换为红黑树。这个值是8
	理想状况下使用随机的哈希码，容器中节点分布在hash桶中的的频率遵循泊松分布。根据Poisson Distribution(泊松分布)，桶中出现8个元素的概率是约是千万分之6，原作者依据概率统计结果选择了8.
（8）HashMap是怎么进行缩容的？
	
（9）HashMap插入、删除、查询元素的时间复杂度各是多少？
	近似认为是O(1),操作数组里的数据，时间复杂度是O(1),链表里是O(N),红黑树里面是O(logN)
（10）HashMap中的红黑树实现部分可以用其它数据结构代替吗？

（11）LinkedHashMap是怎么实现的？

（12）LinkedHashMap是有序的吗？怎么个有序法？

（13）LinkedHashMap如何实现LRU缓存淘汰策略？

（14）WeakHashMap使用的数据结构？

（15）WeakHashMap具有什么特性？

（16）WeakHashMap通常用来做什么？

（17）WeakHashMap使用String作为key是需要注意些什么？为什么？

（18）什么是弱引用？

（19）红黑树具有哪些特性？

（20）TreeMap就有序的吗？怎么个有序法？

（21）TreeMap是否需要扩容？

（22）什么是左旋？什么是右旋？

（23）红黑树怎么插入元素？

（24）红黑树怎么删除元素？

（25）为什么要进行平衡？

（26）如何实现红黑树的遍历？

（27）TreeMap中是怎么遍历的？

（28）TreeMap插入、删除、查询元素的时间复杂度各是多少？

（29）HashMap在多线程环境中什么时候会出现问题？

（30）ConcurrentHashMap的存储结构？

（31）ConcurrentHashMap是怎么保证并发安全的？

（32）ConcurrentHashMap是怎么扩容的？

（33）ConcurrentHashMap的size()方法的实现知多少？

（34）ConcurrentHashMap是强一致性的吗？

（35）ConcurrentHashMap不能解决什么问题？

（36）ConcurrentHashMap中哪些地方运用到分段锁的思想？

（37）什么是伪共享？怎么避免伪共享？

（38）什么是跳表？

（40）ConcurrentSkipList是有序的吗？

（41）ConcurrentSkipList是如何保证线程安全的？

（42）ConcurrentSkipList插入、删除、查询元素的时间复杂度各是多少？

（43）ConcurrentSkipList的索引具有什么特性？

（44）为什么Redis选择使用跳表而不是红黑树来实现有序集合？