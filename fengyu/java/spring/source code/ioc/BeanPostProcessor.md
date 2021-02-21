

#### BeanPostProcessor
```java
@Nullable
default Object postProcessBeforeInitialization(Object bean, String beanName) throws BeansException {
	return bean;
}

@Nullable
default Object postProcessAfterInitialization(Object bean, String beanName) throws BeansException {
	return bean;
}
```
后置处理器，在初始化Bean前后的进行