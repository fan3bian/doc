
#### Lambda Expression

匿名类的出现简化了类的声明，lambda本质上是匿名方法的调用。
lambda 表达式由两部分构成，参数列表 -> 方法实现。
```
(params) -> expression
(params) -> statement
(params) -> { statements }
```
#### Method Reference
java 8新增了函数式接口的注解：**FunctionalInterface**，声明在只有一个接口上，表明是函数式接口(不声明也能运行，为了良好的阅读性)。

#### Standard Functional interfaces

- The Operator interfaces represent functions whose result and argument types are the same.
- The Predicate interface represents a function that takes an argument and returns a  boolean .
- The Function interface represents a function whose argument and return types differ.
- The Supplier interface represents a function that takes no arguments and returns (or “supplies”) a value.
- Finally,  Consumer represents a function that takes an argument and returns nothing, essentially consuming its argument. 

 The six basic functional interfacesare summarized below:

|Interface |Function Signature| Example|
|--|-- | -- |
|UnaryOperator<T> |T apply(T t)| String::toLowerCase|
|BinaryOperator<T> |T apply(T t1, T t2)| BigInteger::add|
|Predicate<T> |boolean test(T t) |Collection::isEmpty|
|Function<T,R> |R apply(T t) |Arrays::asList|
|Supplier<T> |T get() |Instant::now|
|Consumer<T> |void accept(T t)| System.out::println|



	count_韵铭云仓_newCombination.submitNewCombinationRule.do,10000