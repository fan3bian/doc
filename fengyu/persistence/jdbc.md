#### JDBC

数据持久化：把内存中的数据存入磁盘(断电可恢复)的过程叫数据持久化。数据是真实世界运行的抽象，在应用服务其中，数据在内存中被加工、处理、运算。数据在运行在内存中，存储在磁盘上。应用程序中，将内存数据写入磁盘的部分，通常被称为持久层(Persistence Layer).

数据库：数据库是存放数据的仓库。数据库管理系统是为管理数据库而设计的电脑软件系统，一般具有存储、截取、安全保障、备份等基础功能。应用程序(业务系统)需要通过数据库管理系统来操作数据库文件，来实现数据的落盘。通常我们说的数据库，指的是关系型数据库，例如Mysql。

JDBC(Java DataBase Connectivity)：用于执行SQL的Java API,为多种关系数据库提供统一访问，它由一组用Java语言编写的类和接口组成，位于java.sql包之下。JDBC提供了基准，数据库供应商各自实现JDBC接口。

```java
Connection con = DriverManager.getConnection("jdbc:odbc:wombat","login",
"password");
Statement stmt = con.createStatement();
ResultSet rs = stmt.executeQuery("SELECT a, b, c FROM Table1");
while (rs.next()) {
int x = rs.getInt("a");
String s = rs.getString("b");
float f = rs.getFloat("c");
}
```

