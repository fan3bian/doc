Collection与Map


主要又分为两大部分：List和Set

List接口通常表示一个列表（数组、队列、链表、栈等)，元素可以重复，有序
Set接口通常表示一个集合，其中的元素不允许重复（通过hashcode和equals函数保证),无序

AbstractCollection AbstractList AbstractSet

继承的优势：代码复用




ArrayList

动态数组，非线程安全
capacity:容量
initialCapacity：初始容量
size:数组大小
elementData：数组元素


扩容：1.5倍扩容 int newCapacity = (oldCapacity * 3)/2 + 1; 
大量使用了Arrays.copyOf作为数组赋值
