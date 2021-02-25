## JVM参数

#### 标准参数
标准参数不会变
```java
-version
-help
```
#### -X参数(非标准)
- Xint

#### -XX

- `-ea`: enable assertion，启用断言。JUnit测试一般带上这个参数
- `-Xms`: the initial memory allocation pool。 初始堆大小 -XX:MaxHeapSize
- `-Xmx`: Size of young generation heap space。 最大堆大小 -XX:InitialHeapSize
- `-Xss`: the memory of each stack 栈内存
- `-MaxPermSize`: Size of the permanent space(before 1.8)
- `-PermSize`: intial size of the permanent space(before 1.8)
- `-XX`:G1NewSizePercent: 新生代占比，默认5%
- `-XX`:MaxG1NewSizePercent: 新生代最大占比，默认60%
- `-XX`:+UseParNewGC 启用ParNew垃圾回收器
- `-XX`:ParallelGCThreads ParNew垃圾回收线程数量


## G1参数配置
```
-Xms4096M -Xmx4096M -Xss=1M -XX:MaxPermSize=256M -XX:UseG1GC

```
Memory = 4096M
RegionSize = Memory / 2048 = 2M
NewGenerationSize = 4096 * 5% = 200M(近似)

## 参数模板
```java 1.7
-Xms4096M -Xmx4096M -Xmn3072M -Xss1M  -XX:PermSize=256M -XX:MaxPermSize=256M -XX:+UseParNewGC -XX:+UseConcMarkSweepGC -XX:CMSInitiatingOccupancyFaction=92 -XX:+UseCMSCompactAtFullCollection -XX:CMSFullGCsBeforeCompaction=0 -XX:+CMSParallelInitialMarkEnabled -XX:+CMSScavengeBeforeRemark
```

```java 1.8
-Xms4096M -Xmx4096M -Xmn2048M -Xss1M  -XX:MetaspaceSiz=256M -XX:MaxMetaspaceSizee=256M -XX:+UseParNewGC -XX:+UseConcMarkSweepGC -XX:CMSInitiatingOccupancyFaction=92 -XX:+UseCMSCompactAtFullCollection -XX:CMSFullGCsBeforeCompaction=0 -XX:+CMSParallelInitialMarkEnabled -XX:+CMSScavengeBeforeRemark
```
1.8 后
PermSize = MaxMetaspaceSize


https://www.oracle.com/java/technologies/javase/vmoptions-jsp.html
https://www.cnblogs.com/redcreen/archive/2011/05/04/2037057.html