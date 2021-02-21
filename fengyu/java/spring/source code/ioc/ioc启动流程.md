
## IOC启动

1. 创建Spring容器 new AnnotationConfigApplicationContext(MainConfig.class)
2. Spring容器刷新 AbstractApplicationContext#refresh()
3. 获取BeanFactory ConfigurableListableBeanFactory beanFactory = obtainFreshBeanFactory();
4. 注册bean的后处理器 registerBeanPostProcessors(beanFactory);
4. (1) 创建Bean的实例，popluateBean给bean各属性赋值
5. 初始化bean AbstractAutowireCapableBeanFactory#initializeBean BeanPostProcess在bean初始化前后工作的
5. (1) invokeAwareMethods 处理Aware接口的赋值
5. (2) applyBeanPostProcessorsBeforeInitialization 应用后置处理的BeforeInitialization方法
5. (3) invokeInitMethods 执行自定义初始化方法(init-method destroy-method)
5. (4) applyBeanPostProcessorsFfterInitialization

