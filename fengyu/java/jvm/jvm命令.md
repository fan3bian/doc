##
jps
jstat -class pid
jstat -compiler pid
jstat -gc pid
jstat -heap pid


查看JVM默认的参数

```sh
# 查看默认成的线程栈内存 JVM
java -XX:+PrintFlagsFinal -version |grep ThreadStack

     intx CompilerThreadStackSize                   = 0                                   {pd product}
     intx ThreadStackSize                           = 1024                                {pd product}
     intx VMThreadStackSize                         = 1024                                {pd product}
```