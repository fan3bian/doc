1. 常见的集合类有哪些？
Collection Set List Map
2. List和Set有什么特点、区别
List中元素允许重复,Set中元素无法重复;List是有序的,Set中HashSet是无序的。
3. List Set Map常用实现类？
ArrayList LinkedList Vector Stack
HashSet TreeSet LinkedHashSet
HashMap TreeMap HashTable ConCurrentHashMap
4. ArrayList
有序 允许元素重复 底层实现是数组 非线程安全 初始容量：10
原理：
arraylist中增加一个对象的时，会去检查的数组是否有足够的容量来存储这个新的对象。如果没有足够容量的话，那么就会新建一个长度是原来1.5倍的数组，旧的数组就会使用Arrays.copyOf方法被复制到新的数组中去，现有的数组引用指向了新的数组。
elementData = Arrays.copyOf(elementData, newCapacity);
每次数组扩容，增量是旧数组的0.5倍。

