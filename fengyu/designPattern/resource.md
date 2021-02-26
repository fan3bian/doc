DDD(Domain Driven Design): 领域驱动设计
MDD(Model Driven Design): 模型驱动设计
MDA(Model Driven Architecture):领域驱动架构
Modle模型：模型是对现实世界的抽象

如何准确、合理的建立模型？

模型植根于领域、并精确反映出领域中的基础概念

紧密关联领域建模和设计。模型在构建时就考虑
到软件和设计。开发人员会加入到建模的过程中来。主要的想法
是选择一个能够恰当在软件中表现的模型，这样设计过程会很顺畅
并且基于模型。代码和其下的模型紧密关联会让代码更有意义并与
模型更相关。

当一个对象可以用其标识符而不是它的属性来区分时，可以将它作
为模型中的主要定义。保证类定义简洁并关注生命周期的延续性和
可标识性。对每个对象定义一个有意义的区分，而不管它的形式或
者历史。警惕要求使用属性匹配对象的需求。定义一个可以保证对
每一个对象产生一个唯一的结果的操作，这个过程可能需要某个符
号以保证唯一性。这意味着标识可以来自外部，或者它可以是由系
统产生、使用任意的标识符，但它必须符合模型中的身份差别。模
型必须定义哪些被看作同一事物。


值对象：用来描述领域的特殊方面、且没有标识符的一个对象，叫做值对象。

服务的特征：
1. 服务执行的操作涉及一个领域概念，这个领域概念通常不属于一
个实体或者值对象。
2. 被执行的操作涉及到领域中的其他的对象。
3. 操作是无状态的。