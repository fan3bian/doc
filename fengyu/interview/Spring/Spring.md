## Spring面试题

- 依赖注入和依赖查找的区别？
> 
- 什么是Spring Bean?
> Bean是能够在Spring容器中创建、组装被Spring容器管理的对象
- BeanFactory和FactroyBean有什么区别？

- BeanFactory.getBean是否是线程安全的?
> 是线程安全的，回家互斥锁。
- singleton Bean是否在一个应用中是唯一的?
> singleton Bean仅在当前Spring Ioc容器(BeanFactory)中是单例对象，但一个应用可能包含多个上下文。
- BeanPostProcessor使用场景有哪些?
> BeanPostProcessor提供Spring Bean初始化前和初始化后生命周期回调，分别对应postProcessBeforeInitialization和postProcessAfterInitialization方法，允许对指定的bean(根据类型、名称拦截)进行扩展甚至替换

- Spring是如何解决循环依赖？
- Spring Aop的是如何实现的？
- JDK动态代理和Cglib动态代理区别是什么？
- JDK动态代理为什么只能代理接口实现类？
- 



#### 谈谈对Spring AOP的理解
AOP的意思是面向切面编程，是一种设计思想，它是对面向对象编程的一种补充，用于处理系统中的分布在各个模块的横切关注点。Spring的AOP是基于动态代理实现的，支持两种方式的动态代理。一种是JDK动态代理，一种是Cglib实现的动态代理。JDK动态代理是利用反射机制生成一个实现目标接口的代理类，调用具体方法转换成调用InvocationHandler#invoke方法，所以要求目标类一定是接口的实现类；Cglib是通过asm修改字节码的方式，去创建目标类的子类，通过子类覆盖父类方法的方式来实现动态代理。相比较而言，Cglib支持的场景更多。默认情况下，如果是目标对象是接口实现类，Spring会使用JDK动态代理；非接口类使用Cglib来实现。

#### Spring的Bean的生命周期

1. 实例化Bean对象，默认是单例
2. 设置对象属性，依赖注入
3. 检查Aware相关接口，并设置相关依赖。如果实现了BeanNameAware，就调用setBeanName()方法；如果实现了BeanFactoryAware，就调用setBeanFactory()接口，如果实现了BeanClassLoaderAware接口，就调用setBeanClassLoader()方法
4. BeanPostProcesser前置处理。 如果bean实现了BeanPostProcessor接口，就会用调用postProcessBeforeInitialization()方法
5. 检查是否是InitializingBean接口，如果实现了调用afterPorpertiesSet()方法
6. 检查是否有配置自定义的init-method方法，有的话执行。
7. 执行BeanPostProcessor后置处理
8. 检查是否有实现Destruction相关回调接口，是否实现DisposableBean接口，注册销毁函数。


#### Spring容器是如何管理注册Bean的？
Spring通过加载配置文件或者扫描注解的方式，获取bean的元信息,通过Bean的元信息生成BeanDefinition并向注册到BeanDefinitionRegistry中。根据BeanDefinition创建并初始化Spring Bean。

#### Spring是如何解决循环依赖的？
构造器的构造参数出现循环依赖，是无法解决的。Spring能支持下循环依赖指的循环引用，在AbstactAutowireCapableBeanFactory#setAllowCircularReferences有个开关。Spring是用三级缓存在执行createBeanInstance创建Bean后，执行addSingltonFactory()
singletonObjects.put() earlySingletonObjects.remove registeredSingletons.add()


earlySingletonObjects bean在实例化之后往earlySingletonObjects放入

Spring Ioc容器解析bean的定义，生成并注册BeanDefinition。Spring在bean的创建阶段，会把                                                                                                                              