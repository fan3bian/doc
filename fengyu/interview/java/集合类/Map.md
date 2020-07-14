#### HashMap
1. HashMap底层实现原理
jdk1.7: 数组+链表
jdk1.8：数组+红黑树

1.1 数组的初始容量是多少？
1.2 数组类元素的类型是什么？

2. 把一个键值对存入map，也就是put方法，HashMap内部做了什么操作？
2.1 
HashMap扩容机制
负载因子: loadFactor 0.75
容量超过 Entry[] table的长度*loadFactor，就会左移


