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

###### 可选依赖
```xml
<optional>true</optional>
```
基本上采用第一直接依赖

#### Dependency Mediation(依赖调解)
传递性依赖简化了依赖声明

- 第一原则：路径最优者优先
- 第二原则：第一声明者优先

#### Resolved Dependency
```
mvn dependency:list
mvn dependency:tree
mvn dependency:analyze
```

#### DependencyManagement
dependencyManagement 依赖管理是一种
```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/maven-v4_0_0.xsd">
    <modelVersion>4.0.0</modelVersion>
    <groupId>com.jd.clps</groupId>
    <artifactId>clps-dependency</artifactId>
    <version>0.0.1-SNAPSHOT</version>

    <packaging>pom</packaging>
    <name>clps-dependency</name>
    <description>统一管理依赖</description>

    <properties>
        <jdk.version>1.6</jdk.version>
        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
        <project.reporting.outputEncoding>UTF-8</project.reporting.outputEncoding>
        <maven.compiler.encoding>UTF-8</maven.compiler.encoding>

        <clps.version>0.0.1-SNAPSHOT</clps.version>
        <clps.new.version>0.0.2-SNAPSHOT</clps.new.version>

        <spring.version>3.2.6.RELEASE</spring.version>
        <aspectj.version>1.8.0</aspectj.version>
        <mybatis.version>3.0.6</mybatis.version>
        <mybatis-spring.version>1.0.1</mybatis-spring.version>
        <dbcp.version>1.4</dbcp.version>
        <mysql.version>5.1.18</mysql.version>
        <mysql-connector.version>5.1.18</mysql-connector.version>


        <slf4j.version>1.7.13</slf4j.version>
        <log4j2.version>2.3</log4j2.version><!--兼容jdk1.6使用-->
        <disruptor.version>3.3.4</disruptor.version>

        <gson.version>2.8.2</gson.version>
        <fastjson.version>1.2.60</fastjson.version>
        <jackson.version>1.9.6</jackson.version>
        <xstream.version>1.4.10</xstream.version>


        <jsf.version>1.6.9</jsf.version>
        <jfs.version>1.3.1-SNAPSHOT</jfs.version>
        <jimdb.version>1.4.5-SNAPSHOT</jimdb.version>
        <jmq.version>2.3.0</jmq.version>
        <ump.profiler.version>3.2.0</ump.profiler.version>
        <profiler_aop.version>1.0.0</profiler_aop.version>
        <jd.schedule.version>3.0.1-SNAPSHOT</jd.schedule.version>
        <jd.security.configsec.version>1.0.2.RELEASE</jd.security.configsec.version>


        <commons.lang.version>2.6</commons.lang.version>
        <commons.lang3.version>3.3.2</commons.lang3.version>
        <commons.collections.version>3.2.2</commons.collections.version>
        <commons.collections4.version>3.2.2</commons.collections4.version>
        <commons.beanutils.version>1.8.0</commons.beanutils.version>
        <camel.version>2.17.7</camel.version>
        <guava.version>14.0</guava.version>
        <okhttp.version>3.10.0</okhttp.version>
        <junit.version>4.10</junit.version>


    </properties>

    <dependencyManagement>
        <dependencies>
            <!--spring-->
            <dependency>
                <groupId>org.springframework</groupId>
                <artifactId>spring-beans</artifactId>
                <version>${spring.version}</version>
            </dependency>
            <dependency>
                <groupId>org.springframework</groupId>
                <artifactId>spring-context</artifactId>
                <version>${spring.version}</version>
            </dependency>
            <dependency>
                <groupId>org.springframework</groupId>
                <artifactId>spring-core</artifactId>
                <version>${spring.version}</version>
            </dependency>
            <dependency>
                <groupId>org.springframework</groupId>
                <artifactId>spring-expression</artifactId>
                <version>${spring.version}</version>
            </dependency>
            <dependency>
                <groupId>org.springframework</groupId>
                <artifactId>spring-aop</artifactId>
                <version>${spring.version}</version>
            </dependency>
            <dependency>
                <groupId>org.springframework</groupId>
                <artifactId>spring-jdbc</artifactId>
                <version>${spring.version}</version>
            </dependency>
            <dependency>
                <groupId>org.springframework</groupId>
                <artifactId>spring-test</artifactId>
                <version>${spring.version}</version>
                <scope>test</scope>
            </dependency>
            <!-- spring end-->

            <!-- jdbc -->
            <dependency>
                <groupId>commons-dbcp</groupId>
                <artifactId>commons-dbcp</artifactId>
                <version>${dbcp.version}</version>
            </dependency>
            <dependency>
                <groupId>org.mybatis</groupId>
                <artifactId>mybatis</artifactId>
                <version>${mybatis.version}</version>
            </dependency>
            <dependency>
                <groupId>org.mybatis</groupId>
                <artifactId>mybatis-spring</artifactId>
                <version>${mybatis-spring.version}</version>
                <exclusions>
                    <exclusion>
                        <groupId>org.mybatis</groupId>
                        <artifactId>mybatis</artifactId>
                    </exclusion>
                    <exclusion>
                        <groupId>org.springframework</groupId>
                        <artifactId>spring-core</artifactId>
                    </exclusion>
                    <exclusion>
                        <groupId>org.springframework</groupId>
                        <artifactId>spring-jdbc</artifactId>
                    </exclusion>
                    <exclusion>
                        <groupId>org.springframework</groupId>
                        <artifactId>spring-context</artifactId>
                    </exclusion>
                    <exclusion>
                        <groupId>org.springframework</groupId>
                        <artifactId>spring-tx</artifactId>
                    </exclusion>
                </exclusions>
            </dependency>
            <dependency>
                <groupId>mysql</groupId>
                <artifactId>mysql-connector-java</artifactId>
                <version>${mysql-connector.version}</version>
                <scope>runtime</scope>
            </dependency>
            <!-- jdbc end-->

            <!-- aop -->
            <dependency>
                <groupId>org.aspectj</groupId>
                <artifactId>aspectjweaver</artifactId>
                <version>${aspectj.version}</version>
            </dependency>

            <!-- Logging Start -->
            <dependency>
                <groupId>org.apache.logging.log4j</groupId>
                <artifactId>log4j-api</artifactId>
                <version>${log4j2.version}</version>
            </dependency>
            <dependency>
                <groupId>org.apache.logging.log4j</groupId>
                <artifactId>log4j-core</artifactId>
                <version>${log4j2.version}</version>
            </dependency>
            <dependency>
                <groupId>org.slf4j</groupId>
                <artifactId>slf4j-api</artifactId>
                <version>${slf4j.version}</version>
            </dependency>
            <dependency>
                <groupId>org.apache.logging.log4j</groupId>
                <artifactId>log4j-slf4j-impl</artifactId>
                <version>${log4j2.version}</version>
            </dependency>
            <dependency>
                <groupId>org.apache.logging.log4j</groupId>
                <artifactId>log4j-jcl</artifactId>
                <version>${log4j2.version}</version>
            </dependency>
            <!-- Logging End -->
            <!--京东中间件-->
            <dependency>
                <groupId>com.jd</groupId>
                <artifactId>jsf</artifactId>
                <version>${jsf.version}</version>
                <exclusions>
                    <exclusion>
                        <artifactId>slf4j-api</artifactId>
                        <groupId>org.slf4j</groupId>
                    </exclusion>
                </exclusions>
            </dependency>
            <dependency>
                <groupId>com.jd.jim.cli</groupId>
                <artifactId>jim-cli-api</artifactId>
                <version>${jimdb.version}</version>
                <exclusions>
                    <exclusion>
                        <groupId>log4j</groupId>
                        <artifactId>log4j</artifactId>
                    </exclusion>
                    <exclusion>
                        <groupId>org.slf4j</groupId>
                        <artifactId>slf4j-api</artifactId>
                    </exclusion>
                    <exclusion>
                        <groupId>org.slf4j</groupId>
                        <artifactId>slf4j-log4j12</artifactId>
                    </exclusion>
                    <exclusion>
                        <groupId>com.google.guava</groupId>
                        <artifactId>guava</artifactId>
                    </exclusion>
                    <exclusion>
                        <artifactId>slf4j-log4j12</artifactId>
                        <groupId>org.slf4j</groupId>
                    </exclusion>
                    <exclusion>
                        <artifactId>fastjson</artifactId>
                        <groupId>com.alibaba</groupId>
                    </exclusion>
                </exclusions>
            </dependency>
            <dependency>
                <groupId>com.jd.jfs</groupId>
                <artifactId>jfs-java-sdk</artifactId>
                <version>${jfs.version}</version>
                <exclusions>
                    <exclusion>
                        <groupId>org.slf4j</groupId>
                        <artifactId>slf4j-log4j12</artifactId>
                    </exclusion>
                    <exclusion>
                        <groupId>log4j</groupId>
                        <artifactId>log4j</artifactId>
                    </exclusion>
                </exclusions>
            </dependency>
            <dependency>
                <groupId>com.jd.ump</groupId>
                <artifactId>profiler</artifactId>
                <version>${ump.profiler.version}</version>
            </dependency>
            <dependency>
                <groupId>com.jd.ump</groupId>
                <artifactId>profiler_aop</artifactId>
                <version>${profiler_aop.version}</version>
            </dependency>
            <dependency>
                <groupId>com.jd.jmq</groupId>
                <artifactId>jmq-client-spring</artifactId>
                <version>${jmq.version}</version>
                <exclusions>
                    <exclusion>
                        <groupId>com.jd.jmq</groupId>
                        <artifactId>jmq-client-ump</artifactId>
                    </exclusion>
                    <exclusion>
                        <artifactId>fastjson</artifactId>
                        <groupId>com.alibaba</groupId>
                    </exclusion>
                    <exclusion>
                        <artifactId>slf4j-log4j12</artifactId>
                        <groupId>org.slf4j</groupId>
                    </exclusion>
                </exclusions>
            </dependency>
            <dependency>
                <groupId>com.jd.jmq</groupId>
                <artifactId>jmq-client-core</artifactId>
                <version>${jmq.version}</version>
                <exclusions>
                    <exclusion>
                        <artifactId>slf4j-api</artifactId>
                        <groupId>org.slf4j</groupId>
                    </exclusion>
                </exclusions>
            </dependency>
            <dependency>
                <groupId>com.jd.jmq</groupId>
                <artifactId>jmq-client-ump</artifactId>
                <version>${jmq.version}</version>
            </dependency>
            <dependency>
                <groupId>com.jd</groupId>
                <artifactId>joy-schedule-service</artifactId>
                <version>${jd.schedule.version}</version>
                <exclusions>
                    <exclusion>
                        <groupId>log4j</groupId>
                        <artifactId>log4j</artifactId>
                    </exclusion>
                    <exclusion>
                        <artifactId>guava</artifactId>
                        <groupId>com.google.guava</groupId>
                    </exclusion>
                    <exclusion>
                        <artifactId>jsf</artifactId>
                        <groupId>com.jd</groupId>
                    </exclusion>
                </exclusions>
            </dependency>
            <dependency>
                <groupId>com.jd.security.configsec</groupId>
                <artifactId>spring-configsec-sdk</artifactId>
                <version>${jd.security.configsec.version}</version>
            </dependency>

            <!-- json start -->
            <dependency>
                <groupId>com.google.code.gson</groupId>
                <artifactId>gson</artifactId>
                <version>${gson.version}</version>
            </dependency>
            <dependency>
                <groupId>com.alibaba</groupId>
                <artifactId>fastjson</artifactId>
                <version>${fastjson.version}</version>
            </dependency>
            <dependency>
                <groupId>org.codehaus.jackson</groupId>
                <artifactId>jackson-mapper-asl</artifactId>
                <version>${jackson.version}</version>
            </dependency>
            <dependency>
                <groupId>com.thoughtworks.xstream</groupId>
                <artifactId>xstream</artifactId>
                <version>${xstream.version}</version>
            </dependency>
            <!-- json end -->

            <!-- apache common -->
            <dependency>
                <groupId>org.apache.commons</groupId>
                <artifactId>commons-lang3</artifactId>
                <version>${commons.lang3.version}</version>
            </dependency>
            <dependency>
                <groupId>commons-lang</groupId>
                <artifactId>commons-lang</artifactId>
                <version>${commons.lang.version}</version>
            </dependency>
            <dependency>
                <groupId>commons-collections</groupId>
                <artifactId>commons-collections</artifactId>
                <version>${commons.collections.version}</version>
            </dependency>
            <dependency>
                <groupId>commons-beanutils</groupId>
                <artifactId>commons-beanutils</artifactId>
                <version>${commons.beanutils.version}</version>
            </dependency>
            <dependency>
                <groupId>com.lmax</groupId>
                <artifactId>disruptor</artifactId>
                <version>${disruptor.version}</version>
            </dependency>
            <dependency>
                <groupId>com.google.guava</groupId>
                <artifactId>guava</artifactId>
                <version>${guava.version}</version>
            </dependency>
            <dependency>
                <groupId>org.apache.camel</groupId>
                <artifactId>camel-core</artifactId>
                <version>${camel.version}</version>
            </dependency>
            <dependency>
                <groupId>com.squareup.okhttp3</groupId>
                <artifactId>okhttp</artifactId>
                <version>${okhttp.version}</version>
            </dependency>
            <dependency>
                <groupId>junit</groupId>
                <artifactId>junit</artifactId>
                <version>${junit.version}</version>
                <scope>test</scope>
            </dependency>
        </dependencies>
    </dependencyManagement>

    <distributionManagement>
        <repository>
            <id>jd-central</id>
            <name>libs-releases</name>
            <url>http://artifactory.jd.com/libs-releases-local</url>
        </repository>
        <snapshotRepository>
            <id>jd-snapshots</id>
            <name>libs-snapshots</name>
            <url>http://artifactory.jd.com/libs-snapshots-local</url>
        </snapshotRepository>
    </distributionManagement>
</project>

```