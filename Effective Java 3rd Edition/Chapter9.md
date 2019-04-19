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