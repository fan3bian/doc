 
#### lambda

lambda 包含类型推导(编译器自动推导类型信息)

```java
	x -> x.toString()
	(x, y) -> x + y

```

> The basic interfaces operate on object
reference types. The Operator interfaces represent functions whose result and
argument types are the same. The  Predicate interface represents a function that
takes an argument and returns a boolean . The  Function interface represents a
function whose argument and return types differ. The  Supplier interface
represents a function that takes no arguments and returns (or “supplies”) a value.
Finally, Consumer represents a function that takes an argument and returns
nothing, essentially consuming its argument


| Interface | Function Signature | Example |
| -- | -- | -- |
| UnaryOperator<T> | T apply(T t) | String::toLowerCase |
| BinaryOperator<T> | T apply(T t1, T t2) |  BigInteger::add |
| Predicate<T> | boolean test(T t) |  Collection::isEmpty |
| Function<T,R> |  R apply(T t) |  Arrays::asList |
| Supplier<T> |  T get() |  Instant::now |
| Consumer<T> |  void accept(T t) |  System.out::println |