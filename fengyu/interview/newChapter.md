#### 基础面试题
1. ==和equals的区别
==是java里的操作符，equals是根类Object的成员方法。
对于基本类型，==用来比较两个值是否相等；对于引用类型，==用来比较两个对象的引用是否相同。Object类里的equals是用==比较两个对象的，如果没有重写equals方法，equals也是比较两个对象的引用是否相同。
```java
String x = "string";
String y = "string";
String z = new String("string");
System.out.println(x==y); // true 相同引用,常量池
System.out.println(x==z); // false new关键字会开辟新的内存空间
System.out.println(x.equals(y)); // true
System.out.println(x.equals(z)); // true
```
1.1 String是如何重写equals的？
先用==判断引用是否相等，第二判断对象是否是String类型，第三判断长度是否相等，第四片段每一个
1.2 hashCode和equals的关系？hashCode相同，equals一定相同吗？

2. final关键字 
final是最终的意思，可以用来修饰类、方法、变量。用来修饰类的时候，表示类不可以被扩展(继承)，当他用来修饰方法的时候，表示方法不可以被覆盖，修饰变量的时候，表示变量不能被赋值。

3. HashMap
3.1 HashMap的实现原理
HashMap底层结构是关联数组、哈希表，无序，线程不安全。1.8引入了红黑树来提高查询和插入的效率。

loadFactor(负载因子)：0.75
capacity: 默认初始容量 16
数组的扩容机制 threshold = capacity * loadFactor

树化的条件：
