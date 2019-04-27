## Chapter 10
When used to best advantage, exceptions can improve a program's readablity, reliability, and maintainability. When used improperly, they can have the opposite effect. This chapter provides guidelines for using exceptions effectively.

> 异常的最大优势是能够提高程序的可读性、可靠性和可维护性。但是如果使用不当，会起到相反的效果。本章为有效的使用异常给出了指导意见。

#### Item 69：Use Exceptions only or exceptional conditions
> 异常只用于异常情形

Someday, if you are unlucky, you may stumble across a piece of code that looks something like this:
```java
//Horrbile abuse of excetion. Don't ever do this！
//滥用异常，不要这样做。
try{
    int i = 0;
    while(true){
        range[i++].climb();
    }catch(ArrayIndexOutOfBoundsException e){
    
    }
}
```
> 有一天，你可能会遭遇像这样的一段代码。

What does this code do ? It's not at all obvious from inspection, and that's reason enough not to use it (Item 67). It turns out be horrible ill-conceived idiom  for looping through the elements of an array. The infinite loop terminates by throwing, catching and ignoring an ArrayIndexOutOfBoundsExcetion where it attempts to accessthe first array element outside the bounds fo the array. It's supposed to be equivalent to the standard idiom for looping throught an array, which is instantly recognizable to any Java programmer:
```
for (Moutain m : range)
	m.climb()
```
> 这段代码有什么用？从表面上看，一点也不明显。这让你有充足的理由不使用它。事实上，这是一段非常糟糕的用来遍历数组元素的代码。无限循环，当尝试获取数组边界外第一个元素时，通过抛出、捕获和忽略ArrayIndexOutOfBoundException来结束。它应该等同于
标准遍历数组的语句，JAVA程序员脑海里立即想到的：

So why would anyone use the excepiton-based loop in preference to the tried and true？ It's a misguided attempt to improve performance based on the faulty reasoning that, since the VM checks the bounds of all array accesses, the normal loop termination test - hidden by the compiler but still present in the for-each loop - is redundant and should be avoided. There are three things wrong with this reasoning:
- Because exceptionsare designed for excepitonal circumstances, there is little incentive for JVM implementors to make them as fast as explicit test.
- Placing code inside a try-catch bloc inhibits certain optimization that JVM implementations might otherwise them.
- The standard idiom for looping throught an array doesn't necessarily result in redundant checks. Many JVM implementations optimize them away

> 那为什么还有人愿意尝试基于异常的循环呢？这是利用失败去提升性能的一个错误尝试，由于Java虚拟机会检查数组访问的边界，正常循环的终止判断 - 被编译器隐藏，但仍存在于for-each循环中 - 是多余的应该避免。这个推理有三点错误：
- 因为异常是针对特殊情况设计的，没有激励JVM实现者把他们做的像显示判断那样迅速。
- 把代码放置在try-catch块会禁止JVM实现者执行特定的优化
- 标准数组的遍历语句并不会造成冗余的检查，很多JVM实现者优化了他们。

In fact, the excetion-base idiom is far slower than the standare one. On my machine, the exception-based idiom is about twice as slow as the standard one for arrays of one hundred elemetns.

> 实际上，基于异常的循环语句比标准的更慢，在我的机器上，100个元素的数组，基于异常的遍历语句比标准的慢一倍。

Not only does the exception-based loop obfuscate the purpose of the code and reduce its performance, but it's not guaranteed to work. If there is a bug in the loop, the use of exceptions for flow control can mask the bug, greatly complicating the debugging process. Suppose the computation in the body of the loop invokes a method that performs an out-of-bounds access to some unrelated array. If a reasonable loop idiom were used, the bug would generate an uncaught exception, resulting in immediate thread termination with a full stack trace. If the misguided excepiton-based loop were used, the bug-related exception would be caught and misinterpreted as a normal loop termination.

> 基于异常的循环不仅会模糊代码的意图和降低性能，而且它不保证总是正常运行的。假如循环里有一个bug，对流程使用的异常会掩盖掉错误，使调试工作变的更为复杂。比如，循环里和执行的计算调用了其他数组越界访问的操作。如果使用的循环语句，这个错误将生成未捕获的异常，造成线程的立即终止附带完整的堆栈信息。如果使用了误导性基于异常的循环，由于错误导致的异常将会被捕捉到并且被当作正常的循环退出。

The moral of this story is simple: **Exceptions are, as their name impplies, to be used only for exceptional conditions; they should never be used for ordinary control flow.** More generally, use standard, easily recongnizable idioms in preference to overly clever techniques that purport to offer better performance. Even if the performance advantage is real, it may not remain in the face of steadily improving platform implementations. The subtle bugs and maintenance headaches that come from overly clever techniques, however, are sure to remian.

> 这个故事的寓意很简单：正如他们名字所暗示的，异常只能用于异常情形。他们永远不应该用于典型的流程控制。更常用的做法是，优先使用简单的标准的易于识别的语句而不是
旨在更好的性能而聪明过头的。即使真的有性能优势，面对平台实现稳定的改进，可能会当然无存。然而，聪明过头技术带来的微小bug和维护难题，一定会留下来。

This principle also has implications for API design. **A well-designed API must not force its clients to use exceptions for ordinary control flow.**A class with a "state-dependent" method that can be invoked only under cretain unpredictable conditions should generally have a separate "state-testing" method whether it is appropriate to invoke the state-dependent method. For example, the *Iterator* interface has the state-dependent method *next* and the corresponding state-testing method *hasNext*. This enables the standard idiom for iterating over a collection with traditional for loop(as well as the for-each loop, where the *hasNext* method is used internally):
```java
 	for (Iterator<Foo> i = collection.iterator(); i.hasNext();) {
 		Foo foo = i.next();
 		...
 	}
```

>这个原则也可以应用于API设计。精心设计的API不会强制客户端为了典型流程控制使用异常。具有状态依赖的方的类在某些不可预测的条件下调用，通常具有一个独立的“状态检测”的方法判断是否应该调用状态依赖的方法。比如说，*Iterator*接口拥有状态依赖方法*next*和与之对应的状态检测的方法*hasNext*。这使得标准语句能用于迭代传统的for循环的集合(for-each循环中，*hasNext*方法在内部使用)。

If Iterator lacked the hasNext method, clients would forced to do this instead:
```java
	// Do not use this hideous code for iteration over a collection
	try{
		Iterator<Foo> it = collection.iterator();
		while(true){
			Foo fool = i.next();
			...
		}
	} catch (NoSuchElementException e){

	}
```

This should look very familiar after the array iteration example that began this item. In addition to being wordy and misleading, the excetpion-based loop is likely to perform poorly and can mask bugs in unrelated parts of the system.

An alternative to providing a separate state-testing method is to have the state-dependent method return an empty optional(Item 55) or a distinguished value such as *null* if it cannot perform the desired computation.

> 这看上去与开篇的迭代数组的例子非常相似，除了冗长和容易误解外，基于异常的循环性能不佳，会掩盖系统中不相关部分的错误。

提供独立状态检测方法的替代办法是，当状态依赖的方法不能执行预期的计算时，返回一个空选项或者一个可区别的值例如“null”。

Here are some guidelines to help you to choose between a state-testing method and an optional or distinguished return value. If an object is to be accessed concurrently without exteranl synchronization or is subject to externally induced state transitions, you must use an optional or distinguished return value, as the object's state could change in the interval between the invocation of a state-testing method and its state-dependent method. Performance concerns may dictate that an optional or distinguished return value be used if a separate state-testing method would duplicate the work of the state-dependent method. All other things being equal, a state-testing method is mildly perferable to distinguished return value. It offers slightly better readability, and incorrect use may be easier to detect: if you forget to call a state-testing method, the state-dependent method will throw an exception, making the bug obvious; if you forget to check for a distinguished return value, the bug may subtle. This is not an issue for optional values.

>这里为选择状态判断方法还是可选或可区分返回值方法提供了一些指导。如果在没有外部同步或者外部引发的状态转换的情况下并发的访问对象，你一定要使用可选或可区别的返回值，因为对象的状态在调用状态判断方法和状态依赖方法间隔中可能发生改变。如果单独的状态检查方法会重复状态依赖方法的工作，出于性能需求要使用可选的或者可区分返回值方法。其他情况都是相同的，状态判断方法要比可选或可区分返回值方法稍稍好一点。它具有稍好的可读性，并且错误使用会被轻易的检测到。如果你忘记调用状态检测方法，状态依赖方法会抛出异常，从而使错误更明显。如果你忘记去检查可区分返回值，错误会变的隐晦。对于可选值，这不是问题。

In Summary, excepitons are designed for exceptional conditions. Don't use them for ordinary control flow, and write APIs that force others to do so.

> 总结来说，异常是为了异常情形而设计的。不要把他们用于普通的流程控制，编写强制别人这么做的API。


#### Item 70：Use checked exceptions for recoverable conditions and runtime exceptions for programming errors

>对可恢复的条件使用受检的异常，对程序错误使用运行时异常

Java provides three kinds of throwables: checked exceptions, runtime exceptions, and errors. There is some confusion among programmers as to when it is appropriate to use each kind of throwable. While the decision is not always clear-cut, there are some general rules that provide strong guidance.

> Java提供了三种可抛出错误，受检的异常，运行时异常和错误。程序员对每一种可爆出错误的适合情形会感到困惑。尽管结论并不总是清晰的，但也有一些强有力指导的通用规则。

The cardinal rule in deciding whether to use a checked or an unchecked exception is this : use checked excepitons for conditions from which the caller can reasonably by expected to recover. By throwing a checked exception, you force the caller to handle the excepiton in a catch clause or to propagate it outward. Each checked excepiton that a method is declared to throw is therefore a potent indication to the API user that the associated condition is a possible outcome of invoking the method.

>决定使用受检的异常还是非受检异常的基本准则是：使用受检的异常是为了调用者合理地期望从中恢复。通过抛出受检的异常，你可以强制调用方再catch子句中处理它或者向外传播出去。方法声明上抛出的每个受检的异常，对于API使用者是强有力指示：异常是调用方法可能得到的结果。

By confronting the use with a checked exception, the API designer presents a mandate to recover from the condition. The user can disregard the mandate by catching the exception and ignoring it, but this is usually a bad idea(Item 77).

> 当使用受检的异常时，API设计者提出了在改条件下恢复的指令。使用者可以通过捕捉异常忽略异常来漠视指令，但这同通常是个坏主意(见第77条)

There are two kinds of unchecked throwables: runtime exceptions and errors. They are identical in their behavior: both are throwables that needn't, and generally shouldn't, be caught. If a program throws an unchecked exception or an error, it is generally the case that recovery is possible and continued exception would do more harm than good. If a program does not catch such a throwable, it will cause the current thread to halt with an appropriate error message.

>有两种受检的可抛出项:运行时异常和错误。他们在行为上一致：他们是不需要，通常也是不应该被捕获的抛出项。如果程序抛出未受检的异常或者错误，如果程序抛出了受检的异常或者错误，通常是这样的情形：恢复是可能的并且继续执行异常弊大于利。如果程序不捕获这样的可抛出项，会导致当前线程停止并打印对应的错误信息。

Use runtime exception to indicate programming erros. The great majority of runtime exceptions indicate *precondition violations*. A precondition violation is simply a failure by the client of an API to adhere to the contract established by the API specification. For example, the contract for array access specifies that the array index must be between zero and the array length minus one, inclusive. ArrayIndexOutOfBoundsException indicates that this precondition was violated.

> 使用运行时异常提示程序错误。绝大多数的运行时异用来提示“先决条件违规”。违反先决条件失败仅仅是API客户端未能遵循API规范所建立的契约。例如，数据访问契约指定数组下标必须在0和数组长度减1之间，包含临界点。ArrayIndexOutOfBoundsException表示违反此先决条件。

One problem with this advice is that it is not always clear whether you're dealing with a recoverable conditions or a programming error. For example, consider the case of resource exhaustion, which can be caused by a programming error such as allocating an unreasonably large array, or by a genuine shortage of resources. If resource exhaustion is caused by a temporary shortage or by temporarily heightened demand, the condition may well be recoverable. It is a matter of judgement on the part of the API designer whether a given instance of resource exhaustion is likely to allow for recovery. If you believe a condition is likely to allow for recovery, use a checked exception; if not, use a runtime exception. If it isn't clear whether recovery is possible, you're probably better off using a unchecked exception, for reasons discussed in Item 71.

> 这条建议的一个问题是，并不总是清楚你正在处理的是可恢复条件还是编程错误
。例如，资源耗尽的情形，可能是编程错误引起的比如说为一个不合理的大数组分配内存或者真正的资源短缺。如果资源枯竭是由短暂短缺或是由暂时需求增加造成，这种情况可能被恢复。一个给定的资源耗尽实例是否允许恢复，这取决于API设计者的判断。如果你认为这个情况能够恢复，使用受检的异常；如果不能，使用运行时异常。如果不清楚能否恢复，最好使用未受检的异常，原因在第71条中讨论。

While the Java Language Specification does not require it , there is a strong convention that errors are reserved for use by the JVM to indicate resource deficiencies, invariant failures, or other conditions that make it impossible to continue execution. Given the almost universal acceptance of this convention, it's best not to implement any new Error Subclasses. Therefore, **all of the unchecked throwables you implement should subclass RuntimeException**(directly or indirectly). Not only shouldn't you define *Error* subclasses, but with the excepiton of *AssertionError*, you shouldn't thorw them either.

> 虽然java预言规范并不要求，但有一个强烈的约定，JVM保留error作为资源不足，不变的错误，或者其他不可继续执行的情形。考虑到几乎普遍接受这种约定，最好不要实现Error创建任何新的子类。因此，所有你实现的非受检可抛出项应该继承*RuntimeException*直接或间接的。你除了不应该定义Error子类，也不应该定义AssertionError子类，你也不应该抛出他们。

It is possible to define a throwable that is not a subclass of Exception, RuntimeException, or Error. The JLS doesn't address such throwables directly but specifies implicitly that they behave as ordinary checked exceptions(which are subclasses of Exception but not RuntimeException). So when should you use such a beast? In a word, never. They have no benefits over ordinary checked exceptions and would serve merely to confuse the use of your API.

> 定义一个非*Exception, RuntimeException或者Error*的子类。Java语言规范不直接处理这种可抛出项，但隐式指定他们作为普通受检的异常（是Exception的子类，但不是RuntimeExcetpion的子类）。所以你什么时候应该使用这个样怪物？一句话，永远不要。他们相比普通受检的异常，没有任何益处，仅仅是混淆API的使用。

API designers often forget that exception are full-fledged objects on which arbitrary methods can be defined. The primary use of such methods is to provide code that cachees the exception with additional information concerning the condition that caused the exception to be thorwn. In the absence of such methods, programmers have been known to parse the string representation of an excepiton to ferret out additional information. This is extremely bad practice(Itme 12). Throwable classes seldom specify the details of their string representations, so string representations can differ from implementation to implementation and release to release. Therefore, code that parses the string representation of an excepiton is likely to be nonportable and fragile.

>API的设计者往往忘记exception是完全成熟的对象，可以在定义在任何方法上。
这种方法的主要用途是提供缓存异常的代码，以及有关导致异常被激活的条件的其他信息。在没有这种方法的情况下，程序员会解析异常的字符串形式来获取附件信息。这是非常糟糕的行为(第十二条)。可抛出类并不指定其字符串表现形式的细节，所以字符串表示因实现和版本而异
因此，解析异常的字符串表现形式的代码很有可能是不可移植和脆弱的。

Because checked exceptions generally indicate recoverable conditions, it's especially important for them to provide methods that furnish information to help the caller recover from the exceptional condition. For example, suppose a checked exception is thrown when an attempt to make a purchase with a gift card fails due to insufficient funds. The excepiton should provide an accessor method to query the amount of the shortfall. This enable the caller to relay the amount to the shopper. See Item 75 for more this topic.

> 由于受检的异常通常指示可恢复的情形，所以通过方法给调用者提供进一步的信息帮助调用者从异常情形中恢复出来尤为重要。例如，假设当是试图使用礼物卡购物由于资金不足抛出受检的异常。这个异常应该提供一个访问方法图查询短缺的数量。这使得调用者把金额传递给购物者。有关这个而主题更多的信息，请参见第75条。

To summarize, throw checked exceptions for recoverable conditions and unchecked exceptions for programming errors. When in doubt, throw unchecked excepitons. Don't define any thorwables that are neither checked exceptions nor runtime exceptions. Provide methods on your checked exceptions to aid in recovery.

>总而言之，可恢复的情形抛受检的异常，编程错误抛未受检的异常。当有疑问时，抛未受检的异常。不要定义任何非受检的和运行时异常的抛出项。为受检的提供方法来帮助恢复。


#### Item 71: Avoid unnecessary use of checked excepitons

Many Java programmers dislike checked excepitons, but used properly, they can improve the APIs and programs. Unlike return codes and unchecked excepitons, they force programmers to deal with problems, enhancing reliability. That said, overuse of checked excepitons in APIs can make them far less pleasant to use. If a methed throws checked exceptions, the code that invokes it must handle thme in or more *catch* blocks, or delare that it throws them and let them propagate outward. Either way, it places a burden on the user of the API. The burden increased in java 8, as methods throwing checked exceptions can't be used directly in streams(Item 45-48).

This burden may be justified if the exceptional condition cannot be prevented by proper use of th API and the programer using the API can take some useful action once confronted with the exception.
Unless both of there conditions are met, an uncheked exception is appropriate. As a litmus test, ask yourself how the programmer will handle the exception. Is this the best that can be done?
```java
} catch (TheCheckedException e) {
	throw new AssertionError(); //Can't happen!
}
```
Or this?
```java
} catch (TheCheckedException e) {
	e.printStackTrace(); //Oh well, we lose
	System.exit()
}
```
If the programmer can do no better, an unchecked exception is called for.

The additional burden on the programmer caused by a checked exception is substantially higher if it is the sole checked exception thrown by a method. If there are others, the method must already appear in a try block, and this exception requires, at most, another catch block. If a method throws a single checked exception, this exception is the sole reason the method must appear in try block and can't be used directly in streams. Under these circumstances, it pays to ask yourself if there is a way to avoid the ckecked excepiton.

>如果方法只抛出唯一的受检异常，那么给程序员带来的额外负担就会大很多。如果还有其他的异常，这方法一定出现在try语句块中，并且此异常最多需要另一个catch块。如果方法只抛出一个受检的异常，并且这个异常是方法必须出现在try块中唯一的原因并且造成不能直接在stream中使用。在这种情况下，值得问问自己是否有办法去避免这个受检的异常。

The easiest way to eliminate a checked exception is to return an optional of the desired result type(Item 55). Instead of throwing a ckecked exception, the method simply returns an empty optional. The disadvantage of this technique is that the method can't return any additional infomation detailing its inability to perform the desired computation. Exceptions, by constrast, have descriptive types, and can export methods to provide additional information(Item 70).

> 消除受检的异常最简单的方法是返回一个结果所需类型的可选项。方法只返回一个空选项，而不是抛出一个受检的异常。这个技术的缺点是方法不能返回描述不能执行预期计算细节的额外信息。异常，相反，有描述性类型并且能输出方法来提供附加信息(见第70条)

You can also turn a checked exception into a unchecked exception by breaking the method that throws the exception into two methods, the first of which returns a boolean indicating whether the exception would be thrown. This API refactoring transforms the calling sequence from this:
```java
// Invocation with checked exception
try{
	obj.action(args);
} catch (TheCheckedException e) {
	...// Handle exceptional condition
}
```
into this:
```java
// Invocation with state-testing method and unchecked excepiton
if (obj.actionPermitted(args)) {
	obj.action(args);
} else {
	...//Handle excepitonal condition
}
```
This refactoring is not always appropriate, but where it is, it can make an API more pleasent to use. While the latter calling sequence is no prettier that the former, the refactored API is more flexible. If the programmer knows the call will succeed, or is content to let the thread terminate if it fails, the refactoring also allows this trivial calling sequence:
```java
obj.action(args);
```
If you suspect the trivial calling sequence will be the norm, then the API refactoring may be appropriate. The resulting API is essentially the state-testing mehtod API in Item 69 and the same caveats apply: if an object is to be accessed concurrently without external synchronization or it is subject to externally induced state transistions, this refactoring is inappropriate because the object's state may change between the calls to actionPermitted and action. If a separate actionPermitted method would duplicate the work of the action mehtod, the refactoring may be ruled out on the performance grounds.

In summary, when used sparing, checked excepitons can increase the reliability of programs: when overused, they make APIs painful to use. If callers won't be able to recover from failures, throw unchecked excepitons. If recovery may be possible and you want to force callers to handle excepitonal conditions, first consider returning an optional. Only if this would provide insufficient information in the case of failure should you throw a checked excepiton.

#### Item 72: Favor the use of standard exception

An attribute the distinguishes expert programmers from less experienced ones is that expert strive for and usually achieve a higher digree of code reuse. Exceptions are no excepiton to the rule that code reuse is a goods thing. The java libraries provide a set of excepitons that covers most of the exception-throwing needs of most APIs.

Reusing standard exceptions has several benefits. Chief among them is that it makes your API easier to learn and use because it matches the established conventions that programmers are already familiar with. A close second is that programs useing your API are easier to read because they aren't cluttered with unfamiliar excepitons. Last(and lesat), fewer excepiton classes means a smaller memory footprint and less time spent loading classes.

The most commonly reused exception type is *IllegalArgumentExcetpion*(Item 49). This is generally the exception to throw when the caller passes in an argument whose value is inappropriate. For examlpe, this would be the exception to throw if the caller passed the negative number in a parameter representing the number of times some action was to be repeated.

Another commonly used exception is *IllegalStateException*. This is generally the exception to throw if the invocation is illegal because of the state of the reveiving object. For example, this would be exception to throw is the caller attempted to use some object before it had been properly initialized.

Arguably, every erroneous method invocation boils down to an illegal argument or state, but other exceptions are standardly used for certain kinds of illegal arguments and states. If a caller passes null in some parameters for which null values are prohibited, convention dictates that *NullPointerException* be thrown rather that *IllegalArgumentExcetpion*. Similarly, if a caller passes an out-of-range value in a parameter representing an index into a sequence, *IndexOutOfBoundException* should be thrown rather that *IllegalArgumentExcetpion*.

Another reuseable exception is *ConcurrentModificationException*. It should be thrown if an object that was designed for use by a single thread(or with external synchronization) detects that it is being modified concurrently. This exception is at best a hint because it is impossible to relialby detect concurrent modification.

At last standard exception of note is *UnsupportedOperationException*. This is the exception to throw if an object does not supprot an attempted operation. Its use is rare because most objcets supprot all of their methods. This exception is used by clasess that fail to implement one or more *optional operations* defined by an interface that they implement.For example, an append-only List implementation would throw this exception if someone tried to delete an element from the list.

**Do not reuse Exception, RuntimeException, Throwable, or Error  directly.** Treat these classes as if they ware abstract. You can't reliably test for these exceptions because they are superclasses of other exceptions that a method may thorw.

This table summarizes the most commonly reused exceptions:
| Exception| Occasion for Use |
|--|--|
| IllegalArgumentExcetpion | Non-null parameter value is inappropriate |
| IllegalStateException | Object state is inappropriate for method to invovation |
| NullPointerException | Parameter value is null where prohibited |
| IndexOutOfBoundException | Index parameter value is out of range |
| ConcurrentModificationException | Concurrent modification of an object has been detected where it is prohibited |
| UnsupportedOperationException | Object does not support method |

While there are by far the most commonly reused exception, other may be reused where circumstances warrant. For example, it appropriate to reuse *ArithmeticException* and *NumberFormatException* if you were implementing arithmetic objects such as complex numbers or rational numbers. If an exception fits your needs, go ahead and use it, but only if the conditions under which you would throw it are consistent with exception's documentation: reuse must be based on documented semantics, not jsut on name. Alos, feel free to subclass a stardard exception if you want to add more details(Item 75), but remember the exceptions are serializable(Chapter 12). That alone is reason to write your own exception class without good reason.

Choosing which exception to reuse can be tricky because the "occasions for use" in the tale above do not appear to be mutually exclusive. Consider the case of an object represening a deck of cards, and suppose there were a method to deal a hand from the deck that took as an argument the size of the hand. If the caller passed a vlaue larger than the number of cards remaining in the deck, it could be construed as an *IllegalArgumentExcetpion*(the handSize parameter value is too high) or an *IllegalStateException*(the deck contains too few cards). Under these circumstances, the rule is to **throw IllegalStateException if no argument values would have worked, otherwise throw IllegalArgumentExcetpion.**

#### Item 73: Throw excepitons appropriate to the abstraction

It is disconcerting when a method throws an exception that has no apparent connection to the task it performs. This often happens when a method propagates an exception thrown by a low-level abstraction. Not only is it disconcerting, but it polluets the API of the higher layer with implementation details. If the implementation of the higher layer changes in a later release, the exceptions it throws will change too, potentially breaking existing client programs.

> 当方法抛出一个与它执行的任务无明显关系的异常时，这令人不安。
这通常发生在方法传播由低级抽象抛出的异常时。这不仅令人不安，还会通过实习细节污染高等级的API。如果更高层的实现在之后的版本发生了变化，它抛出的异常也会改变，从而可能破坏现有的客户端程序。

To avoid this problem, **Higher layers should catch lower-level excepptions and in their place, throw exceptions that can be explained in terms of the higher-level abstraction**.This idiom is known as *exception translation*
```java
// Exception Translation
try {
	...// Use lower-level abstraction to do our bidding
} catch (LowerLevelException e) {
	throw new HigherLevelExcepiton(...);
}
```
Here is an example of excepiton translation taken from the *AbstractSequentialList* class, whick is a *skeletal implementation*(Item 20) of the List interface. In this example, excepption translation is mandated by the specification of the get method in the List<E> interface:
```java
/**
 * Returns the element at the specified position in this list.
 *@throws IndexOutOfBoundsException if the index is out of range({@code index <0 || index >= size()}).
 *
 **/
public E get(int index) {
	ListIterator<E> i = listIterator(index);
	try {
		return i.next();
	} catch (NuSucnElementException e) {
		throw new IndexOutOfBoundException("Index: " + index);
	}
}
```
A special form of exception translation called *exception chaining* is called for in cases where the lower-level exception might be helpful to someone debugging the problem that caused the higher-level excepiton. The lower-level exception(the cause) is passed to the higher-level exception, which provides an accessor method(Throwable's getCause method) to retrieve the lower-level exception:
```java
// Exception Chaining
try {
	...// Use lower-level abstraction to do our bidding
} catch (LowerLevelException cause) {
	throw new HigherLevelExcepiton(cause);
}
```
The higher-level exception's constructor passes the cause to a chaining-aware superclass constructor, so it is ultimately passed to one of Throwable's chaining-aware constructors, such as Throwable(Throwable):
```java
//Exception with chaining-aware constructor
calss HigherLevelExcepiton extends Exception {
	HigherLevelExcepiton(Throwable cause) {
		super(cause);
	}
}
```
Most standard exceptions have chaining-aware constructors. For exceptions that don't, you can set the cause using Throwable's initCause method。 Not only does exception chaining let you access the cause programmatically(with getCause), but it integrates the cause's stack trace into that of the higher-level exception.

**While exception translation is superior to mindless propagation of exceptions from lower layers, it should not be overused.** Where possible, the best way to deal with excepitons from lower layers is to avoid them, by ensuring that lower-level methods succeed. Sometimes you can do this by checking the validity of higher-level method's parameters before passing them on to lower layers.

If it is impossible to pervent exceptions from lower layers, the next best thing is have the higher layer silently work around these exceptions, insulating the caller of the higher-level method from lower-level problems. Under thees circumstances, it may be appropriate to log the exception using some appropriate logging facility such java.util.logging. This allows programmers to investigate the problem, while insulating client code and the users from it.

In summary, if it isn't feasible to prevent or handle exceptions from lower layers, use exception translation, unless the lower-level method happens to guarantee that all of its exceptions are appropriate to the higher level. Chaining provides the best of both worlds: it allows you to throw an appropriate higher-level exception, while capturing the underlying cause for failure analysis(Item 75).


#### Item 74: Document all exceptions thrown by each method

A descprition of the exceptions thrown by a method is an important part of the documentation required to use the metohd properly. Therefore, it is critically important that you take the time to carefully document all of the excepitons thrown by each method(Item 56).

**Always declare checked exceptions individually, and document precisely the conditions under which each one is thrown** using the Javadoc @throws tag. Don't take the shortcut of declaring that a method throws some superclass of multiple exception classes that it can throw. As an extreme example, don't declare that a public method throws Excepiton or, worse, throws Throwable. In addition to denying an guidance to the method's user concerning the exceptions it is capable of throwing, such a declaration greatly hinders the use of the method because it effectively obscures any other exception that may be thrown in the same context. One exception to this advice is the main method, which can safely be declared to throw Exception because it is called only by VM.

While the language does not require programmers to declare the unchecked exceptions taht a method is capable of throwing, it is wise to document them as carefully as the checked exceptions. Unchecked exceptions generally represent programming errors(Item 70), and familiarizing programmers with all of the errors they can make helps them avoid making these errors. A well-documented list of the unchecked excepitons that a method can throw effectively describes the preconditions for its successful execution. It is essential that every pulbic method's documentation describe its preconditios(Item 56), and documenting its unchecked exceptions is the best way to satisfy this requirement.

It is particularly important that methods in interfaces document the unchecked excepitons the may throw. This documentation forms a part of the interface's  general contract and enables common behavior among multiple implementations of the interface.

**Use the Javadoc @throw tag to document each exception that a method can throw, but do not use the throws keyword on unchecked exceptions.** It is important that programmers using your API are aware of which excepitons are checked and whick are unchecked because the programmers responsibilities differ in these two cases. The documentation generated by the Javadoc @throws tag without a corresponding throws clause in the method declaration provides a strong visual cue to the programmer that an excepiton is unchecked.

It should be noted that documenting all of the unchecked exceptions that each method can thorw is an ideal, not always achievable in the real world. when a class undergoes revision, it is not a violation of source or binary compatibility if an exported method is modified to throw additional unchecked exceptions. Suppose a class invokes a method from another, independently written class. The authors of the former class may carefully document all of the unchecked exceptions that each method throws, but if the latter class is revised to throw additional unchecked exceptions, it is quite likely that the formar class (whick has not undergoes revision) will propagate the new unchecked exception even though it does not document them.

**If an exception is thrown by many methods in a class for the same reason, you can document the exception in the class's documentation comment** rather than documenting it individually for each method. A common example is NullPointerException. It is fine for a class's documentation comment to say,
"All methods in this class throw a NullPointerException if a null object reference is passed in any parameter," or words to that effect.

In summary, document every exception that can be thrown by each method that you write. This is true for unchecked as well as checked exceptions, and for abstract as well as concrete methods. This documentation should take the form of @throws tags in doc documents. Declaring each checked exceptions individually in method's throw clause, but do not declare unchecked excepitons. If you fail to document the exceptions that your method can throw, it will be difficult or impossible for others to make effective use of your classes and interfaces.


