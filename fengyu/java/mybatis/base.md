
#### TypeAlias
*TypeAliasRegistry*
别名：系统别名和自定义别名
#### TypeHandler
类型转换器:处理javaType和jdbcType转换。
#### useGeneratedKeys和keyProperty
这两个属性用于id回填，自动生成id需要数据库的支持，mysql需使用anto_increament。

#### mysql中常用的表达式函数

```java
concat('%',userName,'%') //拼接，连接
now() //当前时间
```

#### mybatis一级缓存

session级别的缓存，相同线程，访问缓存，缓存不存在访问数据库。中间发生数据写操作，会删除缓存。

#### mybatias二级缓存

默认开启二级缓存，无需配置。

```xml
<setting name="cacheEnabled" value="true"></setting>
```
二级缓存特点：

1. 所有的select语句将会被缓存
2. 所有的更新语句(insert/update/delete)将会刷新缓存
3. 缓存采用LRU(Least Recently Used 最近最少使用)算法回收
4. 缓存会存储1024个对象引用

当一级缓存不存在时，若mapper中配置了二级缓存，会访问二级缓存。二级缓存的命中率：访问二级缓存次数/总的查询数

#### 三方EhCache
```xml
<!-- https://mvnrepository.com/artifact/org.mybatis.caches/mybatis-ehcache -->
<dependency>
    <groupId>org.mybatis.caches</groupId>
    <artifactId>mybatis-ehcache</artifactId>
    <version>1.1.0</version>
</dependency>

<cache type="org.mybatis.caches.enchache.EhcacheCache"></cache>
```
意义不大：因为有redis


export/Logs/jvm/jmap_140151_2019-10-06-14-12-38.txt