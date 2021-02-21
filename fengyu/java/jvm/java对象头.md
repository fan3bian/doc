https://www.jianshu.com/p/9d729c9c94c4


## Java对象格式

- Instance Header(对象头)：java对象hashCode、GC年龄、锁标识、class指针、数据长度等
- Instance Data(实例数据)：
- Padding(补齐)：


#### 对象头


#### 实例数据
jvm中，对象的类型分为基本数据类型和引用数据类型。
byte、bolean = 1 
char、short = 2
int、float = 4
long、double = 8
rel = 4 or 8

rel：引用类型，实际上是地址指针。32bit机器，占4个字节；64bit机器，压缩指针占4个字节(-XX:UseCompressedOops)，不压缩8个字节，默认压缩。

#### 对齐填充
JVM要求对象大小是8的整数倍，padding
```C++
class oopDesc {
  friend class VMStructs;
  private:
  volatile markOop  _mark;// mark word
  //union 联合
  union _metadata {
    wideKlassOop    _klass;//class 指针
    narrowOop       _compressed_klass;//压缩的class 指针
  } _metadata;
  ...
}
```
Oop模型：`Ordinary Object Pointer`，即普通对象指针。JVM层用于定义Java对象模型及一些元数据格式的模型就是：Oop，可以认为是JVM层中的“类”。

HotSpot是基于c++语言实现的，它最核心的地方是设计了两种模型,分别是`OOP`和`Klass`，称之为`OOP-Klass Model`.  其中`OOP`用来将指针对象化，比C++底层使用的"`*`"更好用，**每一个类型的OOP都代表一个在JVM内部使用的特定对象的类型**。而`Klass`则用来描述JVM层面中对象实例的具体类型，它是java实现语言层面类型的基础，或者说是**对java语言层类型的VM层描述**。所以看到openJDK源码中的定义基本都以Oop或Klass结尾，如图所示：

Oop就是JVM内部对象类型，而Klass就是java类在JVM中的映射。JVM中把我们上层可见的Java对象在底层实际上表示为两部分，分别是oop和`klass`，