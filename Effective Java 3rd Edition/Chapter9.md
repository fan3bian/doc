## Chapter 9
This chapter is devoted to the nuts and bolts of the language. It discusses local variables, control structures, libraries, date types, and two extralinguistic facilites: *reflection* and *native methods*. Finally, it discusses optimization and naming conventions.

这章专门讨论语言的细节，本地变量，控制结构，类库，数据类型，和两个语言外的设施：*反射*和*本地方法*。最后，讨论一下优化和命名公约。

#### Itme 57: Minimize the scope fo the variables
最小化变量的范围

This item is similar in nature to Item 15. "Minimize the accessibility of classes and memebers". By minimize the scope of local varialbes, you increase the readability and maintainability of your code and reduce the likelihood of error.

> 这一条类目本质上有些像第15条，“最小化类和成员的可访问性”。通过最小化本地变量的范围，可以增加代码的可阅读性和可维护性，减少错误的可能性。

Older programming languages, such as C, mandated that local variables must be declared at the head of a block, and some programmers continue to do this out of habit. It's a habit worth breaking. As a gentle reminder, Java lets you declare variables anywhere a statement is legel( as does C, since C99).

**The most powerful technique for minimizing the scope of a local varialbe is to declare it where it is first used**.
If a varialbe is declared before it is used, it's just clutter - one more thing to distract the reader who is trying to figure out what the program does. By the time the varialbe is used, the reader might not remember the varialbe's type or initial value.

Declaring a local varialbe prematurely can cause its scope not only to begin too early but also to end too late. The scope of a local varialbe extends from the point where it is declared to the end of the enclosing block. If varialbe is declared outside the block in which it is used, it remains visible after the program exits that book. If a varialbe is used accidentally before or after its region of intended use, the consequences can be disastrous.