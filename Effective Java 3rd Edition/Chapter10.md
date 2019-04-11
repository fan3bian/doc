## Chapter 10
When used to best advantage, exceptions can improve a program's readablity, reliability, and maintainability. When used improperly, they can have the opposite effect. This chapter provides guidelines for using expcetions effectively.

> 异常的最大优势是能够提高程序的可读性、可靠性和可维护性。但是如果使用不当，会起到相反的效果。本章为有效的使用异常给出了指导意见。

#### Item 69 Use Exceptions only or exceptional conditions
Someday, if you are unlucky, you may stumble across a piece of code that looks something like this:
```
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

What does this code do ? It's not at all obvious from inspection, and that's reason enough not to use it (Item 67). It turns out be horrible ill-conceived idiom  for looping through the elements of an array. The infinite loop terminates by throwing, catching and ignoring an ArrayIndexOutOfBoundsEcpetion where it attempts to accessthe first array element outside the bounds fo the array. It's supposed to be equivalent to the standard idiom for looping throught an array, which is instantly recognizable to any Java programmer:
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