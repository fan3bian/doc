## MAT(Memory Analyzer Tool)

#### 

内存泄漏分析，集合使用率，Hash性能分析，OQL快读定位空集合实


#### GC ROOT
jvm使用Reachability(可达性算法)来判定对象是否存活，以GC ROOT
为起点，进行链式搜索，不可达的对象被判定为被回收的对象。

GC ROOT可以是：

- 当前线程调用栈上的对象
- 线程自身或者系统类加载器加载的类
- native code保留的活动对象

#### 内存泄漏
当对象不会再使用，但仍然可达(未释放)，导致无法回收。

#### 对象应用
- strong reference
- soft reference
- weak reference
- phantom reference


##### Shallow Heap: 对象本身占用内存的大小，
shallow size: 对象本身占用内存的大小，由对象的成员变量和类型决定。数组还与数组长度相关。

```java
class A{
    public final class String {//8 Bytes header
    private char value[];// 4 Bytes
    private int offset;// 4 Bytes
    private int count; //4 Bytes
    private int hash = 0;// 4 Bytes
…}
//shallow size: 24 bytes
```
#### Retained Heap:当对象被回收时，所有将被GC回收的对象集合

#### Histogram(柱状图)
列出类的实例数
#### Dominator Tree 
如果从GC Root到达Y的的所有path都经过X，那么我们称X dominates Y,再Dominator Tree中，我们能看到内存中最大的对象，以及每个对的dominator。


http://www.lightskystreet.com/2015/09/01/mat_usage/