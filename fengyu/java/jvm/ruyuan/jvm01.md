#### 编译
A.java -> A.class

编译：编译器把java源代码编译成(JVM能够识别的)字节码文件的过程。

#### 运行阶段

JVM通过类加载器把.class加载到内存中。

## 类加载

#### 类加载过程

加载 -> 验证 -> 准备 -> 初始化 -> 使用 -> 卸载

Q: JVM什么时候会去加载一个类？
A: 代码用到这个类的时候，才会去加载。(lazy)

验证：jvm根据虚拟机规范，来校验加载进来的.class文件是否符合规范（保证JVM的安全）

准备：给加载进来的类分配内存空间，给类变量分配内存空间，赋默认初始值。

解析：把**符号引用**替换为**直接引用**

初始化：执行static静态变量和静态代码块。(静态方法只有在被调用的时候才会执行)

#### 类加载器

Bootstrap ClassLoader:
Extension ClassLoader:
Application ClassLoader:
Customized ClassLoader:

##### 双亲委派模型
为了避免多层级的类加载器重复某些类，使用双亲委派模型，先由父类加载去，如果不在所负责目录范围内，再由自类去加载。

