#### Spring boot搭建rest服务

###### 注解

spring boot 默认情况下，没有spring-main.xml配置文件，因项目需要使用jsf协议，需要引入jsf的配置文件，
```java
@ImportResource("classpath:spring-jsf.xml")
```
jsf使用spring aop，在spring启动的时候，建立动态代理，需开启切面。
```java
@EnableAspectJAutoProxy
```



#### Restful风格http服务

@RequestMapping：请求地址,命名空间映射,注于类或方法
- value(path)
- method(POST,GET)
- consumes（Content-Type）application/json, text/html;
- produces
- headers
- params 
@RequestBody：以json格式接收，注于方法参数
@RequestParam：接收key=value格式的url参数，注于方法参数
@RestController：相当于@Controller加@ResponseBody，注于类
@PathVariable：接收路径上命名空ian

