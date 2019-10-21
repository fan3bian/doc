## Mybatis cache

#### 一级缓存
MyBatis执行SQL语句之后，这条语句就是被缓存，以后再执行这条语句的时候，会直接从缓存中拿结果，而不是再次执行SQL。

一级缓存的作用域scope是SqlSession
```java
@Test
public void test() {
    SqlSession sqlSession = sqlSessionFactory.openSession();
    try {
        User user = (User)sqlSession.selectOne("org.format.mybatis.cache.UserMapper.getById", 1);
        log.debug(user);
        User user2 = (User)sqlSession.selectOne("org.format.mybatis.cache.UserMapper.getById", 1);
        log.debug(user2);
    } finally {
        sqlSession.close();
    }
}
```
