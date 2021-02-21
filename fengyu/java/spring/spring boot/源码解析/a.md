## 启动

```java
@SpringBootApplication
public class ElephantApplication {
    public static void main(String[] args) {
        SpringApplication.run(ElephantApplication.class, args);
    }
}
```
1. spring boot是如何判断工程类型(web应用还是普通的应用)?
`SpringApplication`里面有个属性叫`webApplicationType`，它有三种枚举(`NONE`,`SERVLET`,`REACTIVE`)用于表示是servlet-based的web应用、reactive web应用，还是普通的应用。使用`WebApplicationType.deduceFromClasspath()`扫描类路径，使用类加器加载指定类(比如`javax.servlet.Servlet`,`org.springframework.web.context.ConfigurableWebApplicationContext`,`org.springframework.web.reactive.DispatcherHandler`),如果加载到就推断出它的web application类型。

2. setInitializers 和 setListeners
#setInitializers
#setListeners
扫描类路径(classpath)下所有jar包的META-INFO/spring.factories，加载里面所有类缓存到`SpringFactoryLoader#cache`中。
创建`ApplicationContextInitializer`和`ApplicationListener`接口的实例。

3. 推断并加载应用主类
#deduceMainApplicationClass
通过RuntimeExcepiton.getStackTrace()获取运行时异常的栈信息，栈中搜索Main方法来确定Main类，进行加载。

Bean BeanDefinition BeanDefinitionRegistry beanFactoryPostProcessors
// Invoke factory processors registered as beans in the context.
org.springframework.context.support.AbstractApplicationContext#refresh
	org.springframework.context.support.AbstractApplicationContext#invokeBeanFactoryPostProcessors