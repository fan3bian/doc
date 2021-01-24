## AOP

#### EnableAspectJAutoProxy
开启切面: EnableAspectJAutoProxy

`EnableAspectJAutoProxy`里`@Import(AspectJAutoProxyRegistrar.class)`,通过实现`ImportBeanDefinitionRegistrar`覆盖registerBeanDefinitions，向spring容器进行beanName=`org.springframework.aop.config.internalAutoProxyCreator`的AnnotationAwareAspectJAutoProxyCreator.class。


所有以Enable开头的注解，都向Spring容器注册了组件(Bean)，而注解的功能也就是由这些组件决定的。

#### AnnotationAwareAspectJAutoProxyCreator.class

`advisedBeans`: 保存了所有增强的bean

#### Advisor
增强器：切面中的通知方法 @Before @AfterReturning 


通知方法之间的相对顺序是如何保证的?

Spring容器中保存了组件(Component),也就是被代理类的的代理对象，代理对象里保存了增强器和目标对象。

#### 被代理方法的执行过程(切面)

CglibAopProxy.DynamicUnadvisedInterceptor#intercept

