## Chapter 9
This chapter is devoted to the nuts and bolts of the language. It discusses local variables, control structures, libraries, date types, and two extralinguistic facilites: *reflection* and *native methods*. Finally, it discusses optimization and naming conventions.

这章专门讨论语言的细节，本地变量，控制结构，类库，数据类型，和两个语言外的设施：*反射*和*本地方法*。最后，讨论一下优化和命名公约。

#### Itme 57: Minimize the scope fo the variables
最小化变量的范围

This item is similar in nature to Item 15. "Minimize the accessibility of classes and memebers". By minimize the scope of local variables, you increase the readability and maintainability of your code and reduce the likelihood of error.

> 这一条类目本质上有些像第15条，“最小化类和成员的可访问性”。通过最小化本地变量的范围，可以增加代码的可阅读性和可维护性，减少错误的可能性。

Older programming languages, such as C, mandated that local variables must be declared at the head of a block, and some programmers continue to do this out of habit. It's a habit worth breaking. As a gentle reminder, Java lets you declare variables anywhere a statement is legal( as does C, since C99).

>比较老的语言，比如C语言，规定本地变量必须写在代码块的头部，一些程序员出于习惯继续这样做。这是一个值得打破的习惯。温馨提示，java允许你在任何位置语句合法的位置上声明变量（像C99以后）

**The most powerful technique for minimizing the scope of a local variable is to declare it where it is first used**.
If a variable is declared before it is used, it's just clutter - one more thing to distract the reader who is trying to figure out what the program does. By the time the variable is used, the reader might not remember the variable's type or initial value.

> 使本地变量范围最小化最有效的方式在第一次使用的时候声明它。
如果一个变量在使用前声明，会变得混乱。另外，会转移正在试图理解程序代码读者的注意力。等到使用变量时，读者或许忘记了变量类型或者初始值

Declaring a local variable prematurely can cause its scope not only to begin too early but also to end too late. The scope of a local variable extends from the point where it is declared to the end of the enclosing block. If variable is declared outside the block in which it is used, it remains visible after the program exits that block. If a variable is used accidentally before or after its region of intended use, the consequences can be disastrous.

> 过早的声明本地变量会造成它的范围开始的太早，结束的过晚。局部变量的范围是从它开始声明的位置到代码块关闭的位置。如果变量在使用它代码块之外声明，程序退出该块后它仍然可见。如果一个变量在预期的范围之外被意外的使用，结果将是灾难。

Nearly every local variable declaration should contain an initializer. If you don't yet have enough information to initialize a variable sensibly, you should postpone the declaration until you do. One exception to do this rule concerns *try-catch* statements. If a variable is initialized to an expression whose evaluation can throw a checked exception, the variable must be initialized inside a try block(unless the enclosing method can propagate the exception). If the value must be used outside of the try block, then it must be declared before the try block, where it cannot yet be "sensibly initialized." For an example, see page 282.

> 几乎所有的本地变量的声明应该包含一个初始值，如果你还没有足够的信息，你应该推迟至有的时候再声明它。这条规则的一个例外是包含*try-catch*语句。如果一个变量由抛出受检异常的表达式返回值初始化，变量必须在try语句代码块内声明（除非封闭方法可以传播异常）。如果变量一定要在try模块外使用，那么它必须在try代码块之前声明，还不能“明智的初始化”。有关实例，请参见第282页。

Loops present a special opportunity to minimize the scope of variables. The for loop, in both its traditional and for-each forms, allows you to declare *loop variables*, limiting their scope to the exact region where they're needed. (This region consists of the body of the loop and the code in parentheses between the for keyword and the body.) Therefore, perfer for loop to while loops, assuming the contents of the loop variable aren't needed after the loop terminates.

For example, here is the perferred idiom for iterating over a ccollection(Item 58):
```java
//Preferred idiom for iterating over a collection or array
for (Element e: c) {
	...//Do Something with e	
}
```
If you need access to the iterator, perhaps to call its remove method, the perferred idiom uses a traditional for loop in place of the for-each loop:
```java
//Idiom for iterating when you need the iterator
for(Iterator<Element> i = c.iterator(); i.hasNext();){
	Element e = i.next();
	...// Do something with e and i
}
```
To see why these for loops are perferable to a while loop, consider the following code fragment, which contains two while loops and on bug:
```java
Iterator<Element> i = c.iterator();
while(i.hasNext()){
	doSomething(i.next());
}
...
Iterator<Element> i2 = i2.iterator();
while(i.hasNext()){    //BUG
	doSomething(i2.next());
}
```
The second loop contains a copy-and-paste error: it initialized a new loop variable i2, but use the old one, i which is, unfortunately, still in scope. The resulting code compiles without error and runs without throwing an exception, but it does the wrong thing. Instead of iterating over c2, the second loop terminates immediately, giving the false impression that c2 is empty. Because the program errs silently, the error can remain undetected for a long time.

If a similar copy-and-paste error were made in conjunction with either of the for loops(for-each or traditional), the resulting code wouldn't even compile. The element (or iterator) variable from the first loop would not be in scope in the second loop. Here's how it looks with the traditional for loop:
```java
for (Iterator<Element> i = c.iterator(); i.hasNext();){
	Element e = i.next();
	...// Do something wiht e and i
}
...
// Compile-time error -cannot find symbol i
for(Iterator<Element> i2 = c2.iterator(); i2.hasNext()){
	Element e2 = i2.next();
	...// Do something with e2 and i2
}

```
Moreover, if you use a for loop, it's much less likely that you'll make the copy-and-paste errros because there's no incentive to use different variable names in the two loops. The loops are completely independent, so there's no harm in reusing the element(or iterator) variable name. In fact, it's often stylish to do so.

The for loop has one more advantage over the socpe of local variables:
```java
for (int i = 0; n = expensiveComputation(); i < n; i++){
	...//Do something with i;
}
```
The important thing to notice about this idiom is that it has two loop variales, i and n, both of which have exactly the right scope. The second varialbe, n, is used to store the limit of the first, thus avoiding the cost of a redundant computation in every iteration. As a rule, you should use this idiom if the loop test involves a method invocation that is guaranteed to return the same result on each iteration.

> 关于这个语句需要注意的事情是，它有两个变量i和n，都有完全正确的
范围。第二个变量，n，用来存储第一个的上限，从而避免了每次迭代中冗余的计算成本。作为一条规则，如果循环的判断调用了一个每次都返回结果相同的方法，你应该使用这个语句。

A final technique to minimize the scope of local variables is to keep methods small and focused. If you combine two activities in the same method, local variables relevant to one activity may be in th scope of the code performing the other activity. To prevent this from happening, simply separate the method into two: one for each activity.


#### Item 58: Perfer for-each method loops to traditional for loops

As discussed in Itme 45, some tasks are best accomplished with streams, others with iteration. Here is a traditonal for loop to literate over a collection:
```java
//Not the best way to iterate over a collection!
for (Iterator<Element> i = c.iterator(); i.hasNext(); ) {
	Element e = i.next();
	... //Dos something with e
}
```
and here is a traditional for loop to iterate over an array:
```java
//Not the best way to iterate over an array!
for (int i = 0; i < a.length(); i++) {
	...//Do somthing with a[i]
}
```
These idioms are better than while loops(Item 77), but they aren't perfect. The iterator and the index variables are both just clutter - all you need are the elements. Furthermore, they represent opportunities for error. The iterator occurs three times in each loop and the index variable four, which gives you many chances to use the wrong variable. If you do, there is no guarantee that the compiler will catch the problem. Finally, the two loops are quite different, drawing unnecessary attention to the type of the container and adding a (minor) hassle to changing that type.

The for-each loop(offically known as the "enhanced for statement") solves all these problems. It gets rid of the clutter and the opportunity for error by hiding the iterator or index variable. The resulting idiom applies equally to collections and arrays, easing the process of switching the implementation type of a container from one to the other:
```java
// The perferred idiom for iterating over collections and arrays
for (Element e : elements){
	...// Do something whih e
}
```
When you see the colon (:), read it as "in". Thus, the loop above reads as "for each element *e* in *elements*." There in no performance penalty for using for-each loops, even for arrays: the code they generate is essentially identical to the code you would wirte by hand.

The advantages fo the for-each loop over the tranditional for loop are even greater when it comes to nested iteration. Here is a common mistake that people make when doing nested iteration:

```java 
//Can you spot the bug?
enum Suit { CLUB, DIAMOND, HEART, SPADE}
enum Rank { ACE, DEUCE, THREE, FOUR, FIVE ,SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING}
...
static Collections<Suit> suits = Arrays.asList(Suit.values());
static Collections<Rank> ranks = Arrasy.asList(Rank,values());

List<Card> deck = new ArrayList<>();
for(Iterator<Suit> i = suits.iterator(); i.hasNext(); )
	for(Iterator<Rank> j =ranks.iterator(); j.hasNext(); )
		deck.add(new Card(i.next(), j.next()));
```
Don't feel bad if you didn't spot the bug. Many expert programmers have made this mistake at one time or another.
The problem is that the *next* method is called too many times on the iterator for the outer collection(suits). It should be called from the inner loop, so it is called once per card. After you run out of suits, the loop thorws a NoSuchElementException.

If you're really unlukcy and the size fo the outer collection is a multiple of the size of the inner collection - perhaps because they're the same collection - the loop will terminate normally, but it won't do what you want. For example, consider this ill-conceived attempt to print all the possible rolls of a pair of dice:
```java
//Same bug, different symptom!
enum Face {ONE, TWO, THREE, FOUR, FIVE, SIX}
...
Collection<Face> faces = EnumSet.allOf(Face.class);

for (Iterator<Face> i = faces.iterator(); i.hasNext();)
	for(Iterator<Face> j = faces.iterator(); i.hasNext();)
		System.out.println(i.next() + " " +j.next());
```
The program doesn't throw an exception, but it prints only the six "doubles"(from "ONE ONE" to "SIX SXI"), instead of the expected thirty-six combinations.

To fix the bugs in these examples, you must add a variable in the scope of the outer loop to hold the outer element:
```java
// Fixed, but ugly - you can do better!
for (Iterator<Suit> i = suits.iterator(); i.hasNext(); ) {
	Suit suit = i.next();
	for (Iterator<Rank> j = ranks.iterator(); j.hasNext(); )
		deck.add(new Card(suit, j.next()));
}
```
If instead you use a nested for-each loop, the problem simply disappears. The resulting code is as succinct as you could wish for:
```java
//Perferred idiom for nested iteration on collections and arrays
for (Suit suit : suits)
	for (Rank rank : ranks)
		deck.add(new Card(suit, rank));
```
Unfortunately, there are three common situations where you can't use for-each:

- Destructive filtering - If you need to traverse a collection removing selected elements, then you need to use an explicit iterator so that you can its remove method. You can often avoid explicit traversal by using Collections's removeIf method, added in Java8.

- Transforming - If you need to traverse a list or array and replace some or all of the values of the elements, then you need the list iterator or array index in order to replace the value of the element.

- Parallel iteration - If you need to traverse multiple collections in parallel, then you need explicit control over the iteration or index variable so that all iterators or index variables can be advanced in lockstep(as demonstrated unintentionally int the buggy and dice examples above).

If you find yourself in any of these situations, use an ordinary for loop and be wary of the traps mentioned in this item.

Not only does the for-each loop let you iterate over collections and arrays, it lets you iterate over any object that implements the *Iterable* interface, which consists of a single method. Here is how the interface looks:
```java
public interface Iterable<E> {
	// Returns an iterator over the elements in this iterable
	Iterator<E> iterator();
}
```
It is a bit tricky to implement Iterable if you have to write your own Iterator implementation from scratch, but if you are writing a type that represents a group of elements, you should strongly consider having it implement Iterable, even if you choose not to have it implement Collection. This will allow your users to iterate over your type using the for-each loop, and they will be forever grateful.

In summary, the for-each loop provides compelling advantages over the traditonal for loop in clarity, flexiblity, and bug prevention, with no performance. Use for-each loops in perference to for loops wherever you can.


#### Item 59: Know and use libraries

Suppose you want to generate randon integers between zero and some upper bound. Faced with this common task, many programmers would write a little method that looks something like this:
```java
// Common but deeply flawd!
static Random rand = new Random();

static int random (int n) {
	return Maths.abs(rand.nextInt()) % n;
}
```
This method may look good, but it has three flaws. The first is that if n is a small power of two, the sequence of random numbers will repeat itself after a fairly short period. The second flaw is that if n is not a power of two, some numbers will, on average, be returned more frequently than others. If n is large, this effect can be quite pronounced. This is powerful demonstrated by the following program, which generates a million random numbers in a carefully chosen range and then prints out how many of the numbers fell in the lower half the range:
```java
public static void main(String[] args) {
	int n = 2 * (Integer.MAX_VALUE  / 3);
	int low = 0;
	for (int i = 0; i < 1000000; i++)
		if (random(n) < n/2 ) {
			low++;
		}
	System.out.println(low);
}
```

> 这个方法看起来不错，但是它有三个缺点。首先，如果n是2的小次幂，随机序列会在相当短的时间里重复，第二个缺点是如果n不是2的次幂，平均而言，一些数字将比其他数字更频繁的返回。如果n很大，这种效果会很明显。下面的程序可以很好的展示这一点。在精心选择的范围内生成100万个随机数，然后打印出有多少落在了下板部分。

If the random method worked properly, the program would print a number close to half a million, but if you run it, you'll find that it prints a number close to 666,666. Two-thirds of the numbers generated by the random mehtod fall in the lower half of its range!

The third flaw in the random method is that it can, on rare occassions, fail catastrophically, returning a number outside the specified range. This is so becasuce the method attempts to map the value return by rnd.nextInt() to a non-negative int by calling Math.abs. If nextInt() returns Integer.MIN_VALUE, Math.abs will also return Integer.MIN_VALUE, and the remainder operator(%) will return a negative numver, assuming n is not a power of two. This will almost certainly cause your program to fail, and the failure may be difficult to reproduce.

To write a version of the random method that corrects these flaws, you'd have to know a fair amount about pseudorandom number generators, number theory, and two's complement arithmetic. Luckily, you don't have to do this - it's been done for you. It's called Random.nextInt(int). You needn't concern yourself with the details of how it does job (although you can stud the documentation or the source code if you're curious). A senior engineer with a background in algorithms spent a good deal of time designing, implementing, and testing this method and then showed it to several experts in the field to make sure it was right. Then the library was beta tested, released, and used extensively by millions of programmers for almost two decades. No flaws have yet been found in the mothod, but if a flaw were to be discovered, it would be fixed in the next release. **By using a standard library, you take advantage of the knowledge of the experts who wrote it and the experience of those who used it before you.**

As of Java 7, you should no longer use Random. For most uses, **the random number generator of choice is now ThreadLocalRandom.** It produces higer quality random numbers, and it's very fast. On my machine, it is 3.6 times faster than Random. For fork join pools and parallel streams, use *SplittableRandom*.

A second advantage of using the libraries is that you don't have to waste your time ad hoc solutions to problems that only marginally related to your work. If you are like most programmers, you'd rather spend your time woking on your application that on the underlying plumbing.

A third advantage of using the libraries is that their performance tends to improve over time, with no effort on your part. Because many people use them and because the're used in industry-standard benchmarks, the organizations that supply these libraries have a strong incentive to maek them run faster. Many of the Java platform libraries have been rewritten over the years, sometimes repeately, resulting in dramatic performance improvements.

A fourth advantage of using libraries is that they tend to gain functionality over time. If a library is missing something, the developer community will make it known, and the missing functionality may get added in a subsequent release.

A final advantage of using libraries is that you place your code in the mainstream. Such code is more easily readable, maintainable, and reusable by the multitude of developers.

Given all these advantages, it seems only logical to use library facilites in perference to ad hoc implementations, yet many programmers don't. Why not? Perhaps they don't know the library facilities exist. **Numberous features are added to the libraries in every major release, and it pays to keep abreast of these additions.** Each time there is a major release of the Java platform, a web page is published describing its new features. These pages are well worth reading [Java8-feat,Java9-feat]. To reinforce this point, suppose you wanted to write a program to print the contents of a URL specified on the command line(which is roughly what the Linux curl command does). Prior to Java 9, this code was a bit tedious, but in java 9 the transferTo method was added to InputStream. Here is a complete program to perform this task using the new method:
```java
//Printing the contents fo a URL with transferTo, added in java 9
public static void main(String[] args) throws IOException{
	try (InputStream in = new URL(arge[0]).openStream()){
		in.transferTo(System.out)
	}
}
```
The libraries are too big to study all the documentation [Java9-api], but **every programmer should be familiar with the basics of java.lang.util and java.io and their subpacakges.** Knowledeg of their libaries can be acquired on an-needed basis. It is beyond the scope of this item to summarize the facilities in the libraries, which have grown immense over the years.

Several libraries bear special mention. The collections framework and the stream libraries(Item 45-48) should be part fo every programmer's basic tookkit, as should parts of the concurrency utlities in java.util.concurrent. This package contains both high-level utilities to simplify the task of multithreaded programming and low-level primitives to allow experts to write their own higher-level concurrent abstractions. The high-level parts of java.util.concurrent are discussed in Item 80 and 81.

Occasionally, a library facility can fail to meet your needs.The more specialized your needs, the more likely this is to happen. While your first impulse should be to use the libraries, if you've looked at what they have to offer in some area and it doesn't meet your needs, the use an alternate implementation. There will always be holes in the functionality provided by any finite set of libraries. If you can't find what you need in Java platform libraries, your next choice should be to look in high-quality third-party libraries, such as Googles's excellent, open source Guava library[Guava].If you can't find the functionality that you need in any appropriate library, you may hava no choice but to implement it yourself.

To summarize, don't reinvent the wheel. If you need to do something that seems like it should be reasonably common, there may already be a facility in the libraries that does what you want. If there is, use it;
if you don't know, check. Generally speaking, library code is likely to be better that code that you'd write yourself and is likely to improve over time. This is no reflection on your abilities as a programmer. Economies of scale dictate that library code receive far more attention than most developers cound afford to devote the same functionality.


#### Item 60: Avoid float and double if exact answers are required

The float and double types are designed primarily for scientific and engineering calculations. They perform binary floating-point arithmetic, which was carefully designed to furnish accurate approximations quickly over a broad range of magnitudes. They do not, however, provide exact results and should not be used where exact results are required. **The float and double types are particularly ill-suited for monetary calculations** because it is imporssible to represent 0.1 (or any other negative power of ten) as a float or double exactly.

For example, suppose you have $1.03 in your pocket, and you spend 42￠.How much money do you have left? Here's a naive program fargment that attempts to answer this question:
```java
System.out.println(1.03 - 0.42);
```
Unfortunately, it prints out 0.6100000000000001. This is not an isolated case. Suppose you have a dollar in your pocket, and you buy nine washers priced at tencents each. How much change do you get?
```java
System.out.println(1.00 - 9 * 0.10);
```
According to this program fragment, you get $ 0.09999999999999998.

You might think that the problem could be solved merely by rounding results prior to printing, but unfortunately this does not always work. For example, suppose you have a dollar in your pocket, and you see a shelf with row of delicious candies priced at 10￠,20￠,30￠,and so forth, up to a dolllar. You buy one of each candy, starting with the one that costs 10￠, until you can't afford to by the next candy on the shelf. How many candies do you buy, and how much change do you get? Here's a naive program designed to solve this problem:
```java
// Broken - uses floating point for monetary calculation!
public static void main(String[] args) {
	double funds = 1.00;
	int itemsBought = 0;
	for (double price = 0.10; funds >= price; price += 0.10) {
		funds -= price;
		itemsBought++;
	}
	System.out.println(itemsBought + "items bought.");
	System.out.println("Change: $" + funds);
}
```
If you run the program, you'll find that you can afford three pieces of candy, and you have $0.3999999999999999 left. This is the wrong answer! The right way to solve the problem is to **use BigDecimal, int , or long for monetary calculations.**

Here's straightforward transformation of previous program to use the BigDecimal type in place of double. Note that BigDecimal's String constructor is used rather than its double constructor. This is required
in order to avoid introducing inaccurate values into the computation[Bloch05,Puzzle 2]:
```java
public static void main(String[] args) {
	final BigDecimal TEN_CENTS = new BigDecimal(".10");

	int itemsBought = 0;
	BigDecimal funds = new BigDecimal("1.00");
	for (BigDecimal price = TEN_CENTS : funds.compareTo(price) >=0;price = price.add(TEN_CENTS) ) {
		funds = funds.subtract(price);
		itemsBought++;
	}
	System.out.println(itemsBought + "items bought.");
	System.out.println("Money left over: $" + funds);
}
```
If you run the revised program, you'll find that you can afford four pieces of candy, with $0.00 left over. This is the correct answer.

There are, however, two diadvantages to using BigDecimal: it's a lot less convenient than using a primitive arithmetic type, and it's a lot slower. The latter disavantage is irrelevant if you're solving a single short problem, but the former may annoy you.

An alternative to using BigDecimal is to use int or long, depending on the amounts involved, and to keep track of the decimal point yourself. In this example, the obvious approach is to do all copmutation in cents instead of dollars. Here's a straightforward transformation that takes this approach:
```java
public static void main(String[] args) {
	int itemsBought = 0;
	int funds = 100;
	for (int price = 10; funds >= price ;price +=10) {
		funds -= price;
		itemsBought ++;
	}
	System.out.println(itemsBought + "items bought.");
	System.out.println("Cash left over: " + funds + "cents");
}
```
In summary, don't use float or double for any calculations the require an exact answer. Use BigDecimal if you want the system to keep track of the decimal poit and you don't mind the inconvenience and cost of not using a primitive type. Using BigDecimal has the added advantage that it gives you full control over rounding, letting you select from eight rounding modes whenever an operation that entails rounding is performed. This comes in handy if you're performing business calculations wiht legally mandated rounding behaivor. If performance is of the essense, you don't mind keeping track of the decimal point yourself, and the quantities aren't too big, use int or long. If the quantities don't exceed nine decimal digits, you can use it: if they don't exceed eighteen digits, you can use long. If the quantities might exceed nine decimal digits, you can use it; if the don't exceed eighteen digits, you can use long. If the quantities might exceed eighteen digits, use BigDecimal.


#### Item 61: Perfer primitive types to boxed primitives

Java has a two-part type system, consisiting of primitives, such as int, double, and boolean, and reference types, such as String and List. Every primitive type has a corresponding to int, double, and boolean are Integer, Double, and Boolean.

As mentioned in Item 6, autoboxing and auto-unboxing blur but do not erase the distinction betwwen the primitive and boxed primitive types. There are real difference between the two, and it's important that you remain aware of which you are using and that you choose carefully between them.

There are three major differences between primitives and boxed primitives. First primitives have only their values, whereas boxed primitive instances can have the same value and differnet identites. In other words, two boxed primitive instances can have the same value and different identities. Second, primitive types have only fully funtional values, whereas each boxed primitive type have one nonfuntional value, which is null, in addition to all the functional values of the corresponding primitive type. Last, primitives are more time-and space-efficient than boxed primitives. All three of these differences can get you into real trouble if you aren't careful.

Consider the following comparator, which is designed to represent ascending numerical order on Integer values.(Recall that a comparator's compare mehtod returns a number that is negative, zero, or positive, depending on whether its first arguemnt is less than, equal to, or greater than its second.) You wouldn't need to write this comparator in practice because it implements the natural ordering on Integer, but it makes for an interesting example:
```java
// Broken comparator -can you spot the flaw?
Comparator<Integer> naturalOrder = (i,j) -> (i<j) ? -1: (i==j ? 0 :1);
```
This comparator looks like it ought to work, and it will pass many tests. For example, it can be used with Collections.sort to correctly sort a million-element list, whether or not the list contains duplicate elements. But the comparator is deeply flawed. To convince yourself of this , merely print the value of *naturalOrder.compare(new Integer(42), new Integer(42)).* Both Integer instances represents the same value(42), so the value of this expression should be 0, but it's 1, which indicates that the first Integer value is greater than the second!

So what's the problem? The first test in naturalOrder works fine. Evaluating the expression i < j causes the Integer instances referred to by i and j to be anto-unboxed; that is, it extracts their primitive values. The evaluation proceeds to check if the first of the resulting int values is less than the second. But suppose it is not. Then the next test evaluates the expression i==j, which performs an identiy comparison on the two object references. If i and j refer to distinct Integer instances that represent the same int value, this comparison will return false, and the comparator will incorrectlyu return 1, indicating that the first Integer value is greater than the second. **Applying the == operator to boxed primitives is almost always wrong.**

In practice, if you need a comparator to describe a type's naturl order, you should simply call Comparartor.naturalOrder(), and if you wirte a comparator yourself, you should use the comparator construction methods, or the static compare methods on primitive types(Item 14). That said, you could fix the problem in the broken comparator by adding two local variables to store the primitve int values corresponding to the boxed Integer parameters, and performing all of the comparisons on these variables. This avoids the erroneous identity comparison:
```java
Comparator<Integer> naturalOrder = (iBoxed, jBoxed) -> {
	int i = iBoxed, j = jBoxed; //Anto-unboxing
	return i < j ? -1 : (i == j ? 0 : 1);
}
```
Next, consider the delightful little problem:
```java
public class Unbelievalbe{
	static Integer i;
	public static void main(String[] args) {
		if (i = 42)
			System.out.println("Unbelievalbe");
	}
}
```
No, it doesn't print Unbelievable - but what it does is almost as strange. It throws a NullPointerException when evaluating the expression i == 42. The problem is that i is an Integer, not an int, and like all nonconstant object reference fields, its initial value is null. When the program evaluates the expression i == 42, it is comparing an Integer to an int. In nearly every case **when you mix primitives and boxed primitives in an operation, the boxed primitive is auto-unboxed.** If a null object reference is auto-unboxed, you get a NullPointerExcepiton. As this program demonstrates, it can happen almost anywhere. Fixing the problem is as simple as declaring i to be an int instead of an Integer.

Finally, consider the program from page 24 and Item 26:
```java
//Hideously slow program! Can you spot the object creation?
public static void main(String[] args) {
	Long sum = 0L;
	for (long i = 0; i < Integer.MAX_VALUE; i++) {
		sum += i;
	}
	System.out.println(sum);
}
```
This program is much slower than it should be because it accidentally declares a local variable (sum) to be of the boxed primitive type Long instead of the primitive type long. The program compiles without error or warning, and the variable is repeatelly boxed and unboxed, causing the observed performance degradation.

In all three of the programs discussed in this item, the problem was the same: the programmer ignored the distinction between primitives and boxed primitives and suffered the consequences. In the first two programs, the consequences were outright failure; in the third, severe performance problems.

So when should you use boxed primitives? They have several legitimate
uses. The first is an elements, keys, and values in collections. You can't put primitives in collections, so you're forced to use boxed primitives.This is special case of a more general one. You must use boxed primitives as type parameters in parameterized types and methods(Chapter 5), because the language does not permit you to use primitives. For example, you cannot declare a varialbe to be of type ThreadLocal<int>, so you must use ThreadLocal<Ineger> instead. Finally, you must use boxed primitives where making relective method invocations(Item 65).

In summary, use primitives in preference to boxed primitives whenever you have the choice. Primitive types are simpler and faster. If you must use boxed primitives, be careful!.**Antoboxing reduces the verbosity, but not the danger, of using boxed primitives.** When your program compares two boxed primitives with the == operator, it does an identity comparison, which is almost certainly not what you want. When your program does mixed-type computations involving boxed and unboxed primitives, it does unboxing, and **when your program does unboxing, it can throw a NullPointerException.** Finally, when your program boxes primitive values, it can result in costly and unnecessary object creations.

#### Item 62: Avoid strings where other types are more appropriate

Strings are design to represent text, and they do a fine job of it. Because strings are so common and so well supported by the language, there is a natural tendency to use strings for purposes other than those for which they ware designed. This item discusses a few things that you shouldn't do with strings.

**Strings are poor substitutes for other value types.** When a piece of data comes into a program from a file, from the network, or from keyboard input, it is often in string form. There is a natural tendency to leave it that way, but this tendency is justified only if the data really textual in nature. If it's numeric, it should be translated into the appropriate numeric type, such as int, float, or BigIneger. If it's the answer to a yes-or-no question, it should be translated into an appropriate enum type or a boolean. More generally, if there's an appropriate vlaue type, whether primitive or object reference, you should use it; if there isn't, you should write one. While this advice my seem obvious, it is often violated.

**Strings are poor substitudes for enum types**. As discussed in Item 34, enums make for better enumerated types constants than strings.

**Strings are poor substitudes for aggregate types.** If an entity has mulitple components, it is usually a bad idea to represent it as a single string. For example, here's a line of code that comes from 
a real system - identifier names have been changed to protect the guilty:
```java
// Inappropriate use of using as aggregate type
String compoundKey = className + "#" + i.next();
```

This approach has many disadvantages. If the character used to separate fields occurs in one of the fields, chaos may result. To access individual fields, you have to parse the string, which is slow, tedious, and error-prone. You can't provide equals, to String, or compareTo methods but are forced to accept the behavior that String provides. A better approach is simply to write a class to represent the aggregate, often a private static member class(Item 24).

Strings are poor substitutes for capabilities. Occasionally, strings are used to grant access to some functionality. For example, consider the design of a thread-local variable facility. Such a facility provides varialbes for which each thread has its own value. The java libraries have had a thread-local varialbe facility since release 1.2, but prior to that, programmers had to roll their own. When confronted with the task of designing such a facility many years ago, several people independently came up with the same design, in which client-provided string keys are used to identify each thread-local varialbes:
```java
//Broken - inappropriate use of string as capability
public class ThreadLocal{
	private ThreadLocal)() {} // Noninstantiable

	//Sets the current thead's value for named variable.
	public static void set (String key, Object value)；

	//Returns the current threads vlaue for the named variable.
	public static Object get (String key);
}
```
The problem with this approach is that the string keys represent a shared global namespaces for thread-local varialbes. In order for the approach to work, the client-provided string keys have to be unique: if two clients independently decide to use the same name for their thread-local variable, they unintentionally share a single varialbe, which will generally causes both clients to fail. Also, the security is poor. A malicious client could intentionally use the same string key as another client to gain illicit access to other clients data.

This API can be fixed by replacing the strings with unforgable key (sometiems called a capability):
```java
public class ThreadLocal() {} 
	private ThreadLocal() {} //Noninstantiable
	public static class Key { //(Capability)
		Key(){}
	}
	// Generates a unique, unforgeable
	// 生成唯一的不可伪造的
	public static Key getKey(){
		return new Key();
	}
	public static void set(Key key, Object value);
	public static Object get(Key key);
```
While this solves both the problems with string-based API, you can do much better. You don't really need the static mehtods anymore. They can instead become instance methods on the key, at which point the key is no longer a key for a thread-local varialbe: it is a thread-local variable. At this ponit, the top-level class isn't doing anything for you anymore, so you might as well get rid of it and rename the nested class to ThreadLocal:
```java
public final class ThreadLocal {
	public ThreadLocal ();
	public void set(Object value);
	pulbic Object get();
}
```
This API isn't typesafe, because you have to cast the value from Object to its actual type when you retrive it from a thread-loval variable. It is impossible to make the original String-based API typesafe and difficult to make the Key-based API typesafe, but it is a simple matter to make this API typesafe by making ThreadLocal a parameterized class(Item 29):
```java
pulbic final class ThreadLocal<T> {
	public ThreadLocal();
	public void set(T value);
	pulbic T get();
}
```
This is, roughly speaking, the API that java.lang.ThreadLocal provides. In addition to solving the problems with the string-based API, it is faster and more elegant than either of the key-based APIs.

To summarize, avoid the natural tendency to represnet objects as strings when better data types exist or can be written. Used inappropriately, strings are more cumbersome, less flexible, slower, and more error-prone than other types. Types for wihch string are commonly miused include primitive types, enums, and aggregate types.

#### Item 63: Beware the performance of string concatenation

The string concatenation operator(+) is a convenient way to combine a few strings into one. It is fine for generating a single line of output or constructing the string representation of a small, fiexd-size object, but it does not scale. **Using the string concatenation opreator repeatedly to concatenate n strings requires time quadratic in n.** This is an unfortunate consequence of the fact that strings are immutable(Item 17). When two strins are concatenated, the contents of both are copied.

For example, consider this method, which constructs the string representation of a billing statement by repeatedly concatenating a line for each item:
```java
// Inappropriate use of strng concatenation - Performs poorly
public String statement () {
	String result = "";
	for (int i =0 ;i < intItems(); i++)
		result += lineForItem(i);// String concatenantion
	return result;
}
```
The method performs abysamll if the number of items is  large. **To achieve acceptable performance, use a StringBuilder in place of a String** to store the statement under construction:
```
public String statement(){
	StringBuilder b = new StringBuilder(numItems() * LINE_WIDTH);
	for (int i = 0; i < numItems(); i++)
		b.append(lineForItem(i))
	return b.toString();
}
```

A lot of work has gone into making string concatenation faster since Java 6, but the difference in the performance of two methods is still dramatic: If numItems returns 100 and lineForItems return an 80-character string, the second method run 6.5 times faster than the first on my machine. Because the first method is quardratic in the number of items and the second is liner, the performance difference gets much larger as the number of items grwows. Note that the second method prealloates a StringBuilder large enough to hold the entire result, eliminating the need for automatic growth. Even if it is detuned to use a default-size StringBuilder, it is still 5.5 times faster than the first method.

The moral is simple: **Don't use the string concatenation operator to combine more than a few strings** unless performance is irrelevant. Use StringBuilder's append method instead. Alternatively, use a character array, or process the strings one at a time instead of combining them.

#### Item 64: Refer to objects by their interfaces

Item 51 says that you should use interfaces rathar than classes as parameter types. More generally, you should favor the use of interfaces over classes to refer to objects. **If appropriate interface types exist, then parameters, return values, variables, and fields should all be declared using interface types.** The only time you really need to refer to an object's class is when you're creating it with a constructor. To make this concrete, consider the case of LinkedHashSet, which is an implementation of the Set interface. Get in the habit of typing this:
```java
//Good -uses interface as type
Set<Son> sonSet = new LinkedHashSet<>();
```
not this
```java
//Bad
LinkedHashSet<Son> sonSet = new LinkedHashSet<>();
```

**If you get into the habit of using interfaces as types, your program will be much more flexible.** If you decide that you want to switch implementations, all you have to do is change the class name in the constructor( or use a different static factory.) For example, the first declaration could be changed to read:
```java
Set<Son> sonSet = new HashSet<>();
```
and all of the surrounding code would continue to work. The surrounding code was unaware to old implemntation type, so it would be oblivious to the change.

There is one caveat: if the original implementation offered some special functionality not required by the general contract of the interface and the code depended on the same functionality, then it is critical of the interface and the code depneded on that functionality. For example, if the code surrounding the first declaration depended on LinkedHashSet's ordering policy, then it would be incorrect to subsitude HashSet for LinkedHashSet in the declaration, because HashSet makes no guarantee concerning iteration order.

So why would you want to change an implementation type? Beacuse the second implementation offers better performance than the original, or because it offers desirable functionality that the orginal implementation lacks. For example, suppose a field contains a HashMap instance. Changing it to an EnumMap will provide better performance and iteration order consisitent with the natural order of the keys, but you can only use an EnumMap if the key type is an enum type.

You might think it's OK to declare a variable using its implementation type, because you can change the declarationn type and the implementation type at the same time, but there is no guarantee that this change will result in a program that compiles. If the client code used methods on the original implementation type; then the code will no longer complie after making this change. Declaring the variable with the interface keeps you honest.

**It is entirely appropriate to refer to an object by a class rather than an interface if no appropriate interface exists.** For example, consider value class, such as String and BigInteger. Value classes are rarely written with multiple implementations in mind. They are often final and rarely have corresponding interfaces. It is perfectly appropriate to use such a value class as a parameter, variable, field, or return type.

A second case in which there is no appropriate interface type is taht of objects belonging to a framework whose fundamental type are classes rather than interfaces. If an object belongs to such a class-based framework, it is preferable to refer to it by the relevant base class, which is often abstarct, rather than by its implementation class. Many java.io classes such as OutputStream fall into this category.

A final case in which there is no appropriate interface type is that of classes that implement an interface but also provide extra methods not found in the interface - for example, PriorityQueue has a comparator method that is not present on the Queue interface. Such a class should be used to refer to ists instances only if the program relies on the extra methods, and this should be very rare.

There three cases are not meant to be exhaustive but merely to convery the flavor of situations where it is appropriate to refer to an object by its class. In practice, it should be apparent whether a given object has an appropriate interface.If it does, your program will be more flexible and stylish if you use the interface to refer to the object. **If there is no appropriate interface, just use the least specific class in the class hierarchy that provides the required functionality.**

#### Item 65: Perfer interfaces to reflection

The core reflection facility, java.lang.reflect, offers programming access to arbitrary classes. Given a Class object, you can obtain Constructor, Method, and Field instances representing the constructors, methods, and fields of the class repersented by the Class instance. These objects provide programmatic access to the class's member names, field types, method signatures, and so on.

Moreover, Constructor, Method, and Field instances let you manipulate theire underlying counterparts relection: you can construct instances reflectively: you can construct instances, invoke methods, and access fields of the underlying class by invoking methods on the Constructor, Methods, and Field instances. For example, Method.invoke lets you invoke any method on any object of any class(subject to the usual security constraints). Reflection allows one class to use another, even if the latter class dit not exist when the former was compiled.This power, however, comes at a price:

- **You lose all the benefits of compile-time type checking,** including exception checking. If a program attempts to invoke a nonexistent or inaccessible method reflectively, it will fail at runtime unless you've taken special precautions.

- **The code required to perform reflective access is clumsy and verbose.** It is tedious to write and difficult to read.

- **Performance suffers.**Reflective method invocation is much slower than normal method invocation. Exactly how much slower is hard to say, as there are many factors at work. On my machine, invoking a method with no input parameters and an int return was eleven times slower when done reflectively.

There are a few sophisticated applications that require relfection. Examples include code analysis tools and depnedency injection frameworks. Even such tools have been moving away from relection of late, as its disadvantages become clearer. If you have any doubts as to whether your application requires relfection, it probably doesn't.

**You can obtain many of the benefits of reflection while incurring few of its costs by using it only in a very limited form.** For many programs that must use a class that is unavailable at compile time, there exists at compile time an appropriate interface or superclass by which to refer to the class. If this is the case, you can create instances reflectively and access them normally via their interface or superclass.

For example, here is a program that creates a Set<String> instance whose class is specified by the first command line argument. The program inserts the remaining command line arguments into the set and prints it. Regardless of the frist argument, the program prints the remaining arguments with duplicates eliminated. The order in which these arguemnts are printed, however, depends on the class specified in the first argument. If you specify java.util.HashSet, they're printed in apparently random order; if you specify java.util.TreeSet, they're printed in alphabetical order because the elemetns in a TreeSet are sorted:
```java
// Reflective instantiation with interface access
public static void main(String[] args) {
	//Translate the class name into a Class object
	Class<? extends Set<String>> cl = null;
	try {
		cl = (Class<? extends Set<String>>) Class.forName(args[0]);//Unchecked cast!
	} catch (ClassNotFoundException e) {
		fatalError("Class not found.");
	}
	//Get the constructor
	Constructor<? extends Set<String>> cons = null;
	try {
		cons = cl.getDeclaredConstructor();
	} catch (NoSuchElementException e) {
		fatalError("Class not found.");
	}
	//Instantiate the set
	String<Set> s = null;
	try {
		s = cons.newInstance();
	} catch (IllegalAccessException e) {
		fatalError("Constructor not accessible");
	} catch (InstantiationException e) {
		fatalError("Class not instantible.");
	} catch (InvocationTargetException e) {
		fatalError("Constructor threw " + e.getCause());
	} catch (ClassCastException e) {
		fatalError("Class doestn't implemtn Set");
	}
	// Exercise the set
	s.addAll(Arrays.asList(args).subList(1,args.length));
	System.out.println(s);
}
private static void fatalError(String msg) {
	System.err.println(msg);
	System.exit(1);
}
```
While this program is just a toy, the technique it demonstrates is quite powerful. The toy program could easily be returned into a generic set tester that validates the specified Set implementation by aggressively manipulating one or more instances and checking that they obey the Set constract. Similarly, it could be turned into a generic set performance analysis tool. In fact, this technique is sufficiently powerful to implement a full-blown *service provider framework*(Item 1). Usually, this technique is all that you need in the way of reflection.

This example demonstrates two disadvantages of relfection. First, the example can generate six different exceptions at runtime, all of which would have been compile-time
errors if reflective instantiation were not used. (For fun, you can cause the program to generate each of the six exceptions by passing in appropriate command line arguments.) The second disavantage is that it takes twenty-five lines of tedious code to generate an instance of the class from its name, whereas a constructor invocation would fit neatly on a single line. The length of the program could be reduced by catching ReflectiveOperationException, a superclass of the various reflective exceptions that was introduced in Java 7. Both disavantages are restricted to the part of the program that instantiates the object. Once instantiated, the set is indistinguishable from any other Set instance. In a real program, the great bukl of the ocde thus unaffected by this limited use of relection.

If you compile this program, you'll get an unckecked cast warning. This warning is legitimate, in that the cast to <Class? extends Set<String>> will succeed even if the named class is not Set implementation, in which case the program with throw a ClassCastException when it instantiates the class. To learn about suppressing the warning, read Item 27.

A legitimate, if rare, use of relection is to manage a class's dependencies on other classes, methods, or fields that may be absent at runtime. This can be useful if you are writing a package that must be absent at runtime. This can be useful if you are writing a package that must run against multiple versions of some other package. The technique is to compile your pakcage against the minimal environment required to support it, typically the oldest version, and to access any newer classes or methods reflectively. To make this work, you have to take appropirate action if a newer class or method that you are attempting to access does not exist at runtime. Appropriate action might consist of using some alternate means to accomplish the same goal or operating with reduced functionality.

In summary, relection is a powerful facility that is required for certian sophisticated system programming tasks, but it has many disadvantages. If you are writing a program that has to work with classes unknown at compile time, you should, if at all possilbe, use reflection only to instantiate objects, and access the objects using some interface or superclass that is known at compile time.

#### Item 66： Use native methods judiciously

The Java Native Interface(JNI) allows Java programs to call native methods, which are methods written in native programming languages such as C or C++. Historically, native methods have had three main uses. They provide access to platform-specific facilities such as registries. They provide access to existing libraries of native code, including legacy libraries that provide access to legacy data. Finally, native methods are used to write performance-critical parts of applications in native languages for improved performance.

It is legitimate to use native methods to access platform-specific facilites, but it is seldom necessary: as the Java platform matured, it provided access to many features perviously found only in host platform. For example, the process API, add in java 9, provides access to OS processes. It is also legitimate to use native methods to use native libraries when no equivalent libraries are available in java.

**It is rarely adivsable to use native methods for improved performance.** In early releases(perior to java 3), it was necessary, but jVMs have gotten much faster since then. For most tasks, it is now possible to obtain comparable performance in Java. For example, when java.math was added in release 1.1,BigInteger relied on a then-fast multiprecision arithmetic libarary in C. In java 3, BigInteger was reimplemented in Java, and carefully tuned to the point where it ran faster than the original native implementation.

A sad coda to this story is that BigInteger has changed little since then, with the excepiton of faster multiplication for large numbers in java 8. In that time, work continued apace on native libraries, notably GNU Multiple Precision arithmetic library(GMP). Java programmers in need of truly high-performance multiprecision artimetic are now justified in using GMP via native methods[Blum14].

The use of native methods has serious disadvantages. Beacuse native languages are not safe(Item 50), applications using native methods are no longer immune to memory corruption errors. Beacuse native languages are more platform-dependent than Java, programs using native methods can decrease performance because the garbage collector can't automate, or even track, native memory usage(Item 8), and there is a cost associated with going into and out of native code. Finally, native methods require "glue code" that is difficult to read and tedious to wirte.

In summary, think twice before using native methods. It is rare that you need to use them for imporved performance. If you must use native methods to access low-level resources or native libraries, use as little nativve code as possible and test in thoroughly. A single bug in the native code can corrupt your entire application.


#### Optimize judiciously

There are three aphorisms concerning optimization that everyone should know:

> More computing sins are commmited in the name of efficiency(without necessary achieving it) than for any other single reason - including blind stupidity.

> We should forget about small efficiencies, say about 97% of the time: premature
optimization is the root of all evil.

> We follow two rules in the matter of optimization:
Rule 1. Don't do it.
Rule 2. (for experts only). Don't do it yet - that is , not until you have a perfect clear and unoptimized solution.

All of these aphorisms predate the Java programming language by two decades. They tell a deep truth about optimization: it is easy to do more harm than goods, especially of you optimize prematurely. In the process, you may produce software that is neither fast or correct and cannot easily be fixed.

Don't sacrifice sound architectural principles for performance.**Strive to write good programs rather than fast ones.** If a good program is not fast enough, its architecture will allow it to be optimized. Good programs embody the principle of information hiding: where possible, they localize design decisions within individual components, so individual decisions can be changed without affecting the reminder of the system(Item 15).

This does not mean that you can ignore performance concerns until your program is complete. Implementation problems can be fixed by later optimizaiton, but pervasive architectual flaws that limit performance can be impossible to fix without rewriting the system. Changing a fundamental facet of your design after the fact can result in an ill-strutured system that is difficult to maintain and evole. Therefore you must think about performance during the design process.
