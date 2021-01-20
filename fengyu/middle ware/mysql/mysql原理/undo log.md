#### undo log
undo log:回滚日志。

我们对记录做了变更操作，就会产生undo log。事务发生回滚时，undo log作用是把数据恢复到事务执行前的状态。

undo log中存储的是数据的旧值(老版本数据)，当一个数据需要读旧数据时，可使用undo链表获取。

#### undo log类型和内容

1. insert语句的undo log日志类型是`TRX_UNDO_INSERT_REC`，包含

- 这条日志记录的起始位置
- 主键的各列长度和值
- 表id
- undo log日志编号
- undo log日志类型
- 这条日志记录结束位置

#### undo log 版本链
每条数据都包含两个隐藏字段`tax_id`和`roll_pointer`
- tax_id: 最近一次更新这条数据的事务id
- roll_pointer:指向更新这个数据之前的事务生成的undo log

#### ReadView
ReadView是基于undo log版本链实现的一套事务读机制，事务生成一个ReadView，事务能读到自己更新的数据和生成ReadView之前提交的事务修改的值。

执行一个事务，就会产生一个ReadView，里面有4个数据比较关键

- m_ids: 未提交的事务(包含当前事务)
- min_tax_id: m_ids里最小值
- max_tax_id: m_ids里最大值+1，下一个要生成的事务id
- creator_tax_id: 当前事务的id

当前事务只能读取，undo log版本链上小于或者等于本事务id的记录，来实现可重复读。

#### RC级别的实现原理
事务的每次查询生成一个新的ReadView，上一个ReadView中的事务如果已经提交了，就不在出现在新的ReadView的ms_ids列表中了。



