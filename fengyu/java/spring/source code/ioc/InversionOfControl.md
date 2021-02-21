## 控制反转

#### 依赖注入
1. 构造器注入
2. 

#### 依赖查找

1. 单一类型依赖查找 BeanFactory
- 根据名称查找：`beanFactory.getBean(name)`
- 根据类型查找：`beanFactory.getBean(class)`
2. 集合类型依赖查找 ListableBeanFactory 针对某一个类型去查找依赖，返回一个集合。一种是查询Bean名称，一种是查询Bean实例
- 获取同类型Bean名称列表：`listableBeanFactory.getBeanNamesForType(class)`
- 获取同类型Bean对象列表：`listableBeanFactory.getBeansOfType(class)`
3. 层级类型依赖查找：HierarchicalBeanFactory和子接口ConfigurableBeanFactory 
- 获取父BeanFactory：`HierarchicalBeanFactory.getParentBeanFactory`

核心接口: ConfigurableListableBeanFactory


依赖查找是在容器启动(加载)完成之后，来进行的初始化操作；BeanDefinition在Bean定义的注册阶段完成就有了。