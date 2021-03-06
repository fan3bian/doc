#### 团队协作如何高效的开发

有些问题是早晚都需要考虑的，即便当下思考它没有不会带来实质的收益，可磨刀终究不误砍柴功。

现实工作中，团队总是会不断的接到项目。正常的项目按照需求迭代的模式，产品出方案文档，研发排期开发，随后测试跟上，接着就是发版。对于一个成熟的团队，这些工作都会有条不紊的推进着，时间 * 人力 = 结果。

有时候，我们会遇到紧急的项目。项目的截至时间是确定的，任务量远大于平时的工作量。通常决策人会选择增加项目参与人员来缩短时间，从程序开发来说这个方法是必要的，但存在误区的一点是，决策者对于新增人员能够填补工作量预期往往过于乐观，实际上他们也并不关注项目的细节。

对于紧急项目，我的观点是：良好的开发规范，优秀的开发者，熟悉的开发框架优于堆砌开发者。

项目搭建：负责代码工程搭建的人极为关键，技术选型是整个开发中很关键的一环，这一环如果做的好，能够避免很多的小问题。但是如果一开始选型就错了，后期则会麻烦不断。优秀的框架都育有相似的有点：简单、易用、稳健。团队中的做法往往是，在一个已有的项目上，改改名字换成一个新项目。我们来推演一下，这里面要做什么工作：

1. 从原有的代码库拉一份代码推送到新的代码库上
2. 把原有的项目名、文件名、配置替换为新的
3. 把原来的工具类功复制过来

然后就在上堆叠业务代码。实际上，如果不是痛下决心另起炉灶，大多数情况下都会这么干。当前的社会环境中，往往会有这样的情况，你手里维护的代码是你从别的家伙手里接过来的。可能是一朵花，也可能是一坨屎。按照既有的工程去拉新工程，不过是把一坨屎从过一个茅坑换到了另一个茅坑，依然是臭不可闻。

批量更改文件名称的工具：everyThing
批量更改文件内容的工具：sublime或idea