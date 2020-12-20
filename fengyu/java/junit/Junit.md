#### 单元测试 AIR 原则
Automatic(自动化)
Independent(独立性)
Repeatable(可重复)


阿里 BCDE 原则
Border：边界值测试，包括循环边界、特殊取值、特殊时间点、数据顺序等；


Correct：正确的输入，并得到预期的结果；


Design：与设计文档相结合，来编写单元测试；


Error：强制错误信息输入（如:非法数据、异常流程、非业务允许输入等），并得到预期的结果；


#### Junit注解
@Test(excepted==xx.class,timeout=毫秒数)
修饰一个方法为测试方法，excepted参数可以忽略某些异常类


@Before
在每一个测试方法被运行前执行一次


@BeforeClass
在所有测试方法执行前执行


@After
在每一个测试方法运行后执行一次


@AfterClass
在所有测试方法执行后执行


@Ignore
修饰的类或方法会被测试运行器忽略


@RunWith
更改测试运行器

#### 如何给JUnit设置运行时内存