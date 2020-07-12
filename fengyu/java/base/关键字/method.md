The two most prevalent modes of passing arguments to methods are “passing-by-value” and “passing-by-reference”. Different programming languages use these concepts in different ways. As far as Java is concerned, everything is strictly Pass-by-Value.

https://www.baeldung.com/java-pass-by-value-or-pass-by-reference

modes of passing arguemnts to method:
- passing-by-value : parameters passed to the callee method will be clones of original parameters
- passing-by-reference：the unique identifier of the object is sent to the method

In java, everything is strictly Pass-by-Value.Primitive variables store the actual values, whereas Non-Primitives store the reference variables which point to the addresses of the objects they're referring to.Both values and references are stored in the stack memory.

Java是值传递还是引用传递？
值传递
为什么是值传递？
值传递还是引用传递描述的是调用方法时参数的传递方式。值传递指的是传给方法的是原始参数的一个拷贝，在方法内部的对进行修改，原参数不会改变。