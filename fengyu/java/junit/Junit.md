
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
