

``` java 
//way 1
@RunWith(MockitoJUnitRunner.class)
// way 2
@Before
public void init() {
    MockitoAnnotations.initMocks(this);
}
//way 3
@Rule
public MockitoRule initRule = MockitoJUnit.rule();
...

```

#### annotations of the Mockito library

@Mock
@Spy
@Captor
@MockBean


