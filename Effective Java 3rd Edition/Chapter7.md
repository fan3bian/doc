## Chapter 7 Lambdas and Streams

In Java 8, functional interfaces, lambdas, and method reference were added to make it easier to create function objects. The streams API was added in tandem with these language changes to provide library support for processing sequences of data elements. In this chapter, we discuss how to make best use of these facilities.

#### Item 42: Perfer lambdas to anonymous classes

Historically, interfaces (or, rarely, abstract classes) with a single abstract method were used as fucntion types. Their instances, known as function objects, represent functions or actions. Since JDK 1.1 was released in 1997, the primary means of creating a funtion object was the annoymous class(Item 24). Here's a code snippet to sort a list of strings in order of length, using an anonymous class to create the sort's comparison function(which imposes the sort order):

```java
// Anonymous class instance as function object -obsolete
Collections.sort(words, new Comparator<String>() {
	public int compare(String s1, String s2) {
		return Integer.compare(s.length(), s2.length());
	}
});
```

Anonymous classes were adequate for the classic objected-oriented design patterns requiring function objects, notably the *Strategy* pattern [Gamma95]. The Comparator interface represents an *abstract strategy* for sorting; the anonymous class above is a *concrete strategy* for sorting; the anonymous classes, however, made functional programming in Java an unappealing prospect.

In java 8, the language formalized the notion that interfaces with a single abstract method are special and deserve special treatment. These interfaces are now known as *functional interfaces*, and the language allows you to create instances of these interfaces using lambda expressions, or lambdas for short.

Lambdas are similar in function to annoymous classes, but far more concise. Here's how the code snippet above looks with annoymous class replaced by a lambda. The boilerplate is gone, and behavior is clearly evident:

```java
//Lambda expression as function object (replaces annoymous class)
Collections.sort(words, (s1, s2) -> Integer.compare(s1.length(), s2.length()));
```
Note that the types of the lambda (Comparator<String>), of its parameters (s1 and s2, both String), and of its return value (int) are not present in the code. The compiler deduces these types from context, using a process known as *type interface*. In some cases, the compiler won't be able to determine the types, and you'll have to specify them. The rules for type interface are complex: they take up an entire chapter in the JLS[JLS,18]. Few programmers understand these rules in detail, but that's OK. **Omit the types of all lamdbas parameters unless their presence makes your program clear.**If the compiler generates an error telling you it can't infer the type of a lambda parameter, then specify it. Sometimes you may have to cast the return value or the entire lambda expression, but this is rare.

One caveat shoud be added concerning type interface. Item 26 tells you not to use raw types. Item 29 tells you to favor generic types, and Item 30 tells you to favor generic methods. This advice is doubly important when you're using lambdas, becasuse the compiler obtains most of the type information that allows it to perform type interface from generics. If you don't provide this information, the compiler will be unable to do type inference, and if you'll have to specify types manually in your lambdas, which will greatly increase their verbosity. By way of example, the code snippet above won't compile if the variable words is decleared to be of the raw type List instead of the parameterized type List<String>.

Incidentally, the comparator in the snippet can be made even more succinct if a *comparator constrction mehtod* is used in place of a lambda(Item 14.43):

```java 
 Collections.sort(words,comparingInt(String::length));
```
In fact, the snippet can be made still shorter by taking advantage of the sort method that was added the List interface in Java 8:
```java
words.sort(comparingInt(String::length));
```
The additon of labmdas to the language makes it practical to use funtion objects where it would not previously have made sense. For example, consider the Operation enum type in Item 34. Because each enum required different behavior for its apply mehtod, we used constant-specific class bodies and overrode the apply method in each enum constant. To refresh your memory, here is the code:
```java
// Enum type with constant-specific class bodies & data (Item 34)
public enum Operation {
PLUS("+") {
public double apply(double x, double y) { return x + y; }
},
MINUS("-") {
public double apply(double x, double y) { return x - y; }
},
TIMES("*") {
public double apply(double x, double y) { return x * y; }
},
DIVIDE("/") {
public double apply(double x, double y) { return x / y; }
};
private final String symbol;
Operation(String symbol) { this.symbol = symbol; }
@Override public String toString() { return symbol; }
public abstract double apply(double x, double y);
}
```

Item 34 says that enum instance fields are perfereable to constant-specific class bodies. Lambdas make it easy to implement constant-specific behavior using the former instead of the latter. Merely pass a lambda implementing each enum constant's behavior to its construtor. The construtor stores the lambda in an instance field, ant the apply method forwards invocations to the lambda. The resulting code is simpler and clearer than the original version:
```java
// Enum with function object fields & constant-specific behavior
public enum Operation {
PLUS ("+", (x, y) -> x + y),
MINUS ("-", (x, y) -> x - y),
TIMES ("*", (x, y) -> x * y),
DIVIDE("/", (x, y) -> x / y);
private final String symbol;
private final DoubleBinaryOperator op;
Operation(String symbol, DoubleBinaryOperator op) {
this.symbol = symbol;
this.op = op;
}
@Override public String toString() { return symbol; }
public double apply(double x, double y) {
return op.applyAsDouble(x, y);
}
}
```
Note that we're using the DoubleBinaryOperator insterface for the lambdas that represent the enum constant's behavior. This is one of the many perdefined funtional interfaces in java.util.funtion(Item 44). It represent a funtion that takes two double arguments and return a double result.

Looking at the lambda-based Operation enum, you might think constant-specific mehtod bodies have outlived their usefulness, but this is not the case. Unlike methods and classes, **Lambdas lack name and documentation; if a computation isn't self-explanatory, or exceeds a few lines, don't put it in a lambda.** One line is ideal for a lambda, and three lines is a reasonable maximum. If you violate this rule, it can cause serious harm to the readability of your programs. If a lambda is long or difficult to read, either find a way to simplify it or refactor your program to eliminate it. Also, the arguments passed to enum construtors are evaluated in a static context. Thus, lambdas in enum construtors can't access instance members of the enum. Constant-specfic class bodies are still the way to go if an enum types has constant-specific behavior that is difficult to understand, that can't be implemented in a few lines, or that requires access to instance fields or methods.

Likewise, you might think that annoymous classes are obsolete in the era of lambdas. This is closer to the truth, but there are a few things you can do with annoymous classes that you can't do with lambdas. Lambdas are limited to functional interfaces. If you want to create an instance of an abstract class, you can do with an annoymous class, but not a lambda. Similarly, you can use annoymous classes to create instances of interfaces with mutliple abstract methods. Finally, a lambda connot obtain a reference to itself. In a lambda, the this keyword refers to the enclosing instance, which is typically what you want. In an anonymous class, the this keyword refers to the annoymous class instance. If you need access to the funtion object from within its body, then you must use an annoymous class.

Lambdas share with annoymous classes the property that you can't reliably serialize and deserialize them across implementations. Therefore, you should rarely, if ever, serialize a lambda(or an annoymous class instance). If you have a funtion object that you want to make serialize, such as a Comparator, use an isntance fo private static nested class(Item 24).

In summary, as of java8, lambdas are by far the best way to repersent small funtion objects. Don't use annoymous classes for funtion objects unless you have to create instances of types that are't fucntional interfaces. Also, remember that lambdas make it so easy to represent small funtion objects that it opens the door to functional programming techniques that were not previously practical in Java.

Item 43: Perfer mehtod reference to lambdas

