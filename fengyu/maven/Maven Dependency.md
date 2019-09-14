## Maven Dependency

#### Dependency Scope

Maven在编译项目主代码时需要使用一套classpath。通常存在三种classpath

- 编译classpath
- 测试classpath
- 运行classpath

Maven依赖范围就是用来控制依赖与三种classpath的关系，Maven有以下几种依赖:

###### Compile
	编译依赖范围，Maven默认的依赖范围，编译、测试、运行时都需要使用该依赖
###### Test
	测试依赖范围，只对测试classpath有效，再编译主代码或者运行项目时无法使用此类依赖。例如JUnit
###### Runtime
	运行时依赖范围，对于测试和运行都有效，但编译主代码无效。例如JDBC驱动mysql-connector-java，
###### Provided
已提供依赖范围。对于编译和测试classpath有效，但在运行时无效。例如(servlet-api)
###### System
###### Import

#### 传递性依赖
- 第一直接依赖(第一列)
- 第二直接依赖(第一行)

|	| Compile | Test | Provided | Runtime | 
| :-----| ----: | :----: |:----: |:----: |
| Compile | Compile | —— | —— |Runtime |
| Test | Test | —— | —— |Test |
| Provided | Provided | —— | Provided |Provided |
| Runtime | Runtime | —— | —— |Runtime |

基本上采用第一直接依赖

#### Dependency Mediation(依赖调解)
传递性依赖简化了依赖声明

- 第一原则：路径最优者优先
- 第二原则：第一声明者优先
