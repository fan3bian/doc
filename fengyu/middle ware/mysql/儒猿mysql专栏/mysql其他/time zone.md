时区问题：
https://blog.csdn.net/zl1zl2zl3/article/details/90813681

```sql
SHOW VARIABLES LIKE "%time_zone%";
```

+------------------+--------+
| Variable_name    | Value  |
+------------------+--------+
| system_time_zone | CST    |
| time_zone        | SYSTEM |
+------------------+--------+