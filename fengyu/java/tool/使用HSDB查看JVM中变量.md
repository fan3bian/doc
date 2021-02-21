
#### HSDB(Hotspot Debugger)
jdk自带的调试java进程的工具，位于lib\sa-jdi.jar。

启动HSDB
```java
java -classpath "%JAVA_HOME%/lib/sa.jdi.jar" sun.jvm.hostspot.HSDB
java -cp ".;%JAVA_HOME%/lib/sa-jdi.jar" sun.jvm.hotspot.HSDB
```