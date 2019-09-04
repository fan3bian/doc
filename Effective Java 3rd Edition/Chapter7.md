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