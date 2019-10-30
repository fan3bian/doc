#### jvm启动参数
https://blog.csdn.net/tterminator/article/details/54342666
```bash
export OPTS_MEMORY="-XX:+UseConcMarkSweepGC -XX:ParallelGCThreads=4 -XX:CMSInitiatingOccupancyFraction=70 -XX:+UseCMSInitiatingOccupancyOnly -XX:+CMSClassUnloadingEnabled -Xms4096m -Xmx4096m -Xmn1024m -XX:MaxPermSize=256m -XX:+UnlockExperimentalVMOptions -Xss512K -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=./java_pid<pid>.hprof $SGM_OPTS"
```


在UMP看到的应用参数
```
/export/servers/jdk1.7.0_71/bin/java -Djava.library.path=/usr/local/lib -server -Xms1024m -Xmx2048m -XX:MaxPermSize=256m -Djava.awt.headless=true -Dsun.net.client.defaultReadTimeout=60000 -Djmagick.systemclassloader=no -Dnetworkaddress.cache.ttl=300 -Dsun.net.inetaddr.ttl=300 -XX:+UseConcMarkSweepGC -XX:ParallelGCThreads=4 -XX:+UseConcMarkSweepGC -XX:CMSInitiatingOccupancyFraction=70 -XX:+UseCMSInitiatingOccupancyOnly -Xms4096m -Xmx4096m -Xmn1024m -XX:MaxPermSize=512m -XX:+UnlockExperimentalVMOptions -Xss512K -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=./java_pid<pid>.hprof -javaagent:/export/sgm-probe/sgm-probe-5.0.13-product/sgm-agent-5.0.13.jar -Xbootclasspath/a:/export/sgm-probe/sgm-probe-5.0.13-product/sgm-agent-5.0.13.jar -Dsgm.app.name=clps_goods -classpath /export/Instances/clps_goods/server1/runtime/conf:/export/Instances/clps_goods/server1/runtime/lib/* -Dbasedir=/export/Instances/clps_goods/server1/runtime -Dfile.encoding=UTF-8 com.jd.clps.master.goods.GoodsSafLauncher
```
#### CPU 

[java应用cpu使用率过高问题排查](https://blog.csdn.net/qq_34944965/article/details/81107419)
[java性能调优](https://www.jianshu.com/p/a75c1253fc11)
https://www.jianshu.com/p/3b2d71bf3eb5


Code Cache： JVM Code Cache is an area where JVM stores its bytecode compiled into native code. The just-in-time (JIT) compiler is the biggest consumer of the code cache area
