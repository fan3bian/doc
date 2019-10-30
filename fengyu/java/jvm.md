b#### jvm启动参数

```bash
export OPTS_MEMORY="-XX:+UseConcMarkSweepGC -XX:ParallelGCThreads=4 -XX:CMSInitiatingOccupancyFraction=70 -XX:+UseCMSInitiatingOccupancyOnly -XX:+CMSClassUnloadingEnabled -Xms4096m -Xmx4096m -Xmn1024m -XX:MaxPermSize=256m -XX:+UnlockExperimentalVMOptions -Xss512K -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=./java_pid<pid>.hprof $SGM_OPTS"
```

#### CPU 

[java应用cpu使用率过高问题排查](https://blog.csdn.net/qq_34944965/article/details/81107419)
[java性能调优](https://www.jianshu.com/p/a75c1253fc11)
https://www.jianshu.com/p/3b2d71bf3eb5b