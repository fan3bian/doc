#### Buffer Pool

```sql
SELECT @@innodb_buffer_pool_size;
SELECT @@innodb_buffer_pool_instances;
SELECT @@innodb_buffer_pool_chunk_size;
SELECT @@innodb_page_size;

innodb_buffer_pool_size:8589934592
innodb_buffer_pool_instances:4
innodb_buffer_pool_chrunk_size:16384
innodb_page_size=
```

常见配置：
8 188G
3 5G
10 20G



buffer pool将划分为多个实例，减少线程间读写缓存的争用，提高系统并发性。一般设置 buffer pool大小为总内存的 3/4 至 4/5

innodb_buffer_pool_size 大于 1GB 时, innodb_buffer_pool_instances 默认为 8。小于1G时，默认1个。



When increasing or decreasing innodb_buffer_pool_size, the operation is performed in chunks. Chunk size is defined by the innodb_buffer_pool_chunk_size configuration option, which has a default of 128M.