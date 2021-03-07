
SHOW GLOBAL VARIABLES LIKE 'innodb_flush_log%'

redo log 刷盘策略:
innodb_flush_log_at_trx_commit


提交事务时，redo log buffer数据刷新到磁盘文件的策略
0：提交事务时，不把redo log buffer里面数据刷入磁盘
1：提交事务时，就必须把redo log从内存中刷入到磁盘文件中
2：提交事务时，就必须把redo log从内存中刷入到磁盘文件对应的OS cache中

默认是1，通常也设置为1，可以避免mysql突然宕机后数据丢失。一旦宕机，可以通过redo log进行数据恢复。