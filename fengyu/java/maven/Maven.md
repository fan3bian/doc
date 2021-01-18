## [Maven](https://maven.apache.org/index.html)
> 所有编程工具的开发都是为了减轻程序员的负担，在这个演进的过程中，也降低了程序员的门槛
#### Maven安装
maven是一个java工具，前提是JRE就位
```bash
mvn archetype:generate -DgroupId=com.mycompany.app -DartifactId=my-app -DarchetypeArtifactId=maven-archetype-quickstart -DarchetypeVersion=1.4 -DinteractiveMode=false
```
使用maven命令可以简单的生成java项目，这个项目的目录层级结构就是典型的maven风格（约定大于配置）
#### Maven依赖
maven通过比内地仓库(远程仓库以及中央仓库)简化了项目依赖JARs的问题
#### Maven POM
POM(Project Object Modal)：工程对象模型，为统一标准构建过程提供了支撑

###### Built-in Build Lifecycles
- clean : handles project cleaning
- default : handles your project deployment, 
- site : handles the creation of your project's site documentation

###### Phases
A Build Lifecycles is Made Up of Phases.Each of these build lifecycles is defined by a different list of build phases, wherein **a build phase represents a stage in the lifecycle**.

- compile - compile the source code of the project
- test - test the compiled source code using a suitable unit testing framework. These tests should not - require the code be packaged or deployed
- package - take the compiled code and package it in its distributable format, such as a JAR.
- verify - run any checks on results of integration tests to ensure quality criteria are met
- install - install the package into the local repository, for use as a dependency in other projects - locally
- deploy - done in the build environment, copies the final package to the remote repository for sharing with other developers and projects.

生命周期相互独立，周期内的阶段前后依赖

###### Plugin Goals
A Build Phase is Made Up of Plugin Goals

###### Features
Maven具有一套标准的骨架：
```java
mvn archetype:generate
```
###### Maven Coordinate
- groupId
- artifactId
- version
- type
- scope
- optional
- exclusions

#### Dependency
###### Dependency Scope（依赖范围）
Maven在编译/测试/运行时，对应三套classpath。Scope具有以下几种枚举：
- compile: 编译依赖范围(默认)
- test: 测试依赖范围，只对测试有效，例如JUnit
- provided: 已提供依赖范围，编译和测试有效，运行时无效，例如servlet-api
- runtime: 运行时依赖范围，测试和运行时有效，编译时无效。例如mysql-connector-java(jdbc驱动实现)
- system: 此类依赖不通过Maven仓库实现，往往与本机绑定。需使用systemPath显示指定依赖文件的路径。
- import：与dependencyManagement关联使用

###### 依赖的传递性
Maven会解析依赖的POM,将那些必要的间接依赖以传递性的形式引入当前的项目中。

依赖范围会影响传递性依赖，低级依赖范围不能超出高级依赖范围。

###### Dependency Mediation 依赖调解
- 第一原则：路径最近者优先
- 第二原则：第一声明者优先

###### 可选依赖
可选依赖不传递，使用时需显式声明

###### 归类依赖
Spring Framwork体系

###### 依赖分析
Maven会自动解析项目所有的直接依赖和传递性依赖(间接依赖)，根据规则判定依赖的范围。解决依赖冲突，确保每个构件有唯一的版本，以得到最终的已解析依赖(Resolved Dependency)
```bash
mvn dependency:list
mvn dependency:tree
mvn dependency:analyze
```
#### Maven Repository
- Loacl Repository
- Remote Repository
-- Central Repository
-- Private Repository (局域网)

###### 私服

私服是一种特殊的远程仓库，架设在局域网。当Maven需要下载构件时，从私服请求，私服不存在该构件，则从外部仓库下载，缓存在私服上。使用私服的有点是：节省外网带宽，加速Maven构建，稳定性高，可控

配置仓库信息
```xml
<repositories>
	<repository>
		<id>jboss</id>
		<name>JBoss Repository</name>
		<url>https://repository.jboss.com/maven2</url>
		<releases>
			<enabled>true</enabled>
		</releases>
		<snapshots>
			<enabled>false</enabled>
			<!-- 关闭该仓库快照版本的下载支持 -->
			<updatePolicy>daily</updatePolicy>
			<!-- 从远程仓库检查更新的频率 daily never always interval:5 times/min-->
		</snapshots>
		<layout>default</layout>
	</repository>
</repositories>
```
配置认证信息
```xml
<settings>
	<severs>
		<sever>
			<id>my-project</id>
			<username>repo-use</username>
			<password>repo-pwd</password>
		</sever>
	</severs>
</settings>
```
构建部署至远程仓库，修改POM
```xml
<distributionManagement>
	<repository>
		<id>proj-releases</id>
		<name>Proj Release Repository</name>
		<url>http://192.168.1.100/content/repositories/proj-releases</url>
	</repository>
	<snapshotRepository>
		<id>proj-snapshots</id>
		<name>Proj-snapshots</name>
		<ulr>http://192.168.1.100/content/repository/pro-snapshots</ulr>
	</snapshotRepository>
</distributionManagement>
```
id与<sever>id</sever>对应，执行 mvn clean deploy

###### 快照版本
mvn clean compile -U：强制更新，忽略<updatePolicy>配置
快照版本只应该在组织内部或模块间依赖使用。

#### 插件目标
为了方便构建，Maven在核心的生命周期阶段绑定了插件目标
| Phases | Plugin Goals | Task Content|
|--------|--------|--------|
|clean  |maven-clean-plugin:clean|清空文件
|process-resources | maven-resource-plguin:resources|复制主资源文件至主输出目录|
|compile | maven-compiler-plugin:compile|编译主代码至主输出目录|
|process-test-resources | maven-resources-plugin:testResources|复制测试资源文件至测试输出目录|
|test-compile | maven-compiler-plugin:testCompile|编译测试代码至测试输出目录|
|test | maven-surefire-plugin:test |执行测试用例|
|package | maven-jar-plugin:jar | 创建项目jar包|
|install | maven-install-plugin:install|将项目输出构建安装到本地仓库|
|deploy | maven-deploy-plugin:deploy|将项目输出构建安装到远程仓库|
|site | maven-site-plugin:site| |
|site-deploy | maven-site-plugin:deploy| |

###### 自定义绑定
```xml
<build>
	<plugin>
		<groupId>org.apache.maven.plugins</groupId>
		<artifactId>maven-source-plugin</artifactId>
		<version>2.1.1</version>
		<executions>
			<execution>
				<id>attach-sources</id>
				<phase>verify</phase>
				<!-- 绑定到verify阶段 -->
				<goals>
					<goal>jar-no-for</goal>
					<!-- 执行目标 -->
				</goals>
			</execution>
		</executions>
	</plugin>
</build>
```
查看插件默认绑定的阶段：
```bash
mvn help:describe -Dplugin=org.apache.maven.plugins:maven-source-plugin:2.1.2 -Ddetail
```
Bound to phase: package

插件配置
```
mvn install -Dmaven.test.skip=true
```
使用命令行参数可以方便的更改插件的行为

pom中插件的全局配置
<plugin>
	<gorupId>org.apache.maven.plugins</gorupId>
	<artifactId>maven-compiler-plugin</artifactId>
	<version>2.1</version>
	<configuration>
		<source>1.7</source>
		<target>1.7</target>
		<!-- 插件的全局配置，基于插件的任务，会使用该配置 -->
	</configuration>
</plugin>

maven-help-plugin
```bash
mvn help:describe -Dplugin=org.apache.maven.plguins:maven-compiler-plugin:2.1.7
mvn help:describe -Dplugin=org.apache.maven.plguins:maven-compiler-plugin

mvn help:describe -Dplugin=compiler -Dgoal=compile

mvn help:describe -Dplugin=compiler -Ddetail

mvn org.apache.maven.plugins:maven-help-plugin:2.1:describe -Dplugin=compiler
mvn [plugin]:goal
```
可以使用前缀(Goal Prefix)替代插件
| Goal Prefix | Plugin |
|------|------|
| help | maven-help-plugin|
| denpendency | maven-dependency-plugin|
| compiler | maven-compiler-plugin|

###### 插件仓库
插件与构件同样基于坐标存储在Maven仓库中
```xml
<pluginRepositories>
	<pluginRepository>
		<id>jboss</id>
		<name>JBoss Repository</name>
		<url>https://repository.jboss.com/maven2</url>
		<releases>
			<enabled>true</enabled>
			<updatePolicy>never</updatePolicy>
		</releases>
		<snapshots>
			<enabled>false</enabled>
			<!-- 关闭该仓库快照版本的下载支持 -->
			<!-- 从远程仓库检查更新的频率 daily never always interval:5 times/min-->
		</snapshots>
		<layout>default</layout>
	</pluginRepository>
</pluginRepositories>
```
<groupId>org.apache.maven.plugins</groupId>
Maven官方插件可以省略groupId

解析插件版本：

Maven在Super POM中为所有的核心插件设定了版本，其他插件版本maven会检查仓库内的
maven-metadata.xml,将该路径的所以后的仓库元数据归并后，计算出latest和release的值。maven以后对于插件，不再解析lastest。
#### Maven的聚合和继承
聚合使用多模块的方式，


常用Maven配合
https://www.cnblogs.com/zincredible/p/9661243.html