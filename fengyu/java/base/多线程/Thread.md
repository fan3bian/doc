#### 线程的状态
- NEW
- TIMED_WAITING
- BLOCKED
- DESTROY

#### Jconsole

#### 反编译
javap -v 
```sh
javap -v SynDemo.class > a.txt
```
指令：

monitorenter 
monitorexit

synchronized 修饰符会生成monitor，monitorenter与monitorexit之间的指令部分，即是加锁(Lock)的部分。一个monitorenter可以对应两个monitorexit，分别对应正常出口和异常出口。

```java
  public static synchronized void access();
    descriptor: ()V
    flags: ACC_PUBLIC, ACC_STATIC, ACC_SYNCHRONIZED
    Code:
      stack=3, locals=1, args_size=0
         0: getstatic     #2                  // Field java/util/concurrent/TimeUnit.SECONDS:Ljava/util/concurrent/TimeUnit;
         3: lconst_1
         4: invokevirtual #3                  // Method java/util/concurrent/TimeUnit.sleep:(J)V
         7: getstatic     #4                  // Field java/lang/System.out:Ljava/io/PrintStream;
        10: new           #5                  // class java/lang/StringBuilder
```
ACC_SYNCHRONIZED：标记是同步方法