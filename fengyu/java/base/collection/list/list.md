（1）ArrayList和LinkedList有什么区别？
	ArrayList和LinkedList都是List接口的实现类,ArrayList是基于动态数组实现的，LinkedList是基于链表实现的。对于随机访问get和set，ArrayList性能优于LinkedList，对于添加和删除操作，更适合用LinkedList
（2）ArrayList是怎么扩容的？
	ArrayList内部结构是动态数组，当元素的个数比数组长度长的时候需要扩容。扩容的方式是原长度的一半与新增加元素个数的最大值。当然这个值要小于数组能支持的最大长度。如果初始化时未指定长度，添加第一个元素时是，自动扩容至10个。
（3）ArrayList插入、删除、查询元素的时间复杂度各是多少？
	add方法在尾部进行插入，复杂度是O(1),在指定位置插入需要移动后面的元素，复杂度是O(N),使用get方法通过下标查询复杂度是O(1),删除元素复杂度是O(1)
（4）怎么求两个集合的并集、交集、差集？
	addAll retainAll removeAll
（5）ArrayList是怎么实现序列化和反序列化的？
	writeObject，readObject方法。ArrayList使用transient抑制elementData数组的序列化，只把序列化元素。
（6）集合的方法toArray()有什么问题？

（7）什么是fail-fast？

（8）LinkedList是单链表还是双链表实现的？

（9）LinkedList除了作为List还有什么用处？

（10）LinkedList插入、删除、查询元素的时间复杂度各是多少？

（11）什么是随机访问？

（12）哪些集合支持随机访问？他们都有哪些共性？

（13）CopyOnWriteArrayList是怎么保证并发安全的？

（14）CopyOnWriteArrayList的实现采用了什么思想？

（15）CopyOnWriteArrayList是不是强一致性的？

（16）CopyOnWriteArrayList适用于什么样的场景？

（17）CopyOnWriteArrayList插入、删除、查询元素的时间复杂度各是多少？

（18）CopyOnWriteArrayList为什么没有size属性？

（19）比较古老的集合Vector和Stack有什么缺陷？