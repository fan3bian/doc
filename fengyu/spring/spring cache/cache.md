#### Cache Annotation

+ `@EnableCaching`:缓存开关。
+ `Cacheable`:
+ `CacheaPut`:

当注解`EnableCaching`未使用时，标注Cacheable、CachePut等注解的程序运行时，不会产生错误(即便你使用的有问题)。


核心接口：
Cache: 对缓存抽象
CacheManager: 对缓存管理工具的抽象
CacheResolver


CacheAdviceParser

https://blog.csdn.net/hong10086/article/details/91811646

SPI:Service Provider Interface (SPI) is an API intended to be implemented or extended by a third party. It can be used to enable framework extension and replaceable components.


@NonNull可以标注在方法、字段、参数之上，表示对应的值不可以为空
@Nullable注解可以标注在方法、字段、参数之上，表示对应的值可以为空



CacheAutoConfiguration.CacheConfigurationImportSelector#selectImports -> 
ConfigurationClassParser#processConfigurationClass
ConfigurationClassParser#doProcessConfigurationClass


instanceof a.instanceof(A) a对象是否为A类的示例
isAssignableFrom A.instanceof(B) A类是否为B类的父类


`AbstractCacheResolver#resolveCaches`:从
`CacheOperation`:抽象类表示缓存的基本操作

`RedisCacheManager#loadCaches`:加载初始化的cache,需要实现类来完成

`AbstractFallbackCacheOperationSource#getCacheOperations`

- org.springframework.data.redis.cache.RedisCacheConfiguration
	- 


CacheAspectSupport#execute()
- doGet()
	- createCacheKey(key): 把key序列化
	- lookup(key): 获取缓存内容，
- invokeOperation: 获取真实数据
- doPut




1. Cache自动注入的问题：覆盖
2. CacheManager注入的问题
3. 序列化与反序列化的问题
	-3.1 key序列化与value的序列化
4. 