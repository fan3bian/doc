# 什么是MongoDB

MongoDB是一个文档数据库，提供好的性能，领先的非关系型数据库。采用BSON存储文档数据。2007年10月，MongoDB由[10gen](https://link.zhihu.com/?target=https%3A//zh.wikipedia.org/w/index.php%3Ftitle%3D10gen%26action%3Dedit%26redlink%3D1)团队所发展。2009年2月首度推出。MongoDB用c++编写的。

##### 优势：

- 面向文档的存储：以 JSON 格式的文档保存数据。
- 任何属性都可以建立索引。
- 复制以及高可扩展性。
- 自动分片。
- 丰富的查询功能。
- 快速的即时更新。
- 来自 MongoDB 的专业支持。

##### elasticsearch与MongoDB相同点与不同点

**相同点：**
1、都是以json格式管理数据的nosql数据库。
2、都支持CRUD操作。
3、都支持聚合和全文检索。
4、都支持分片和复制。
5、都支持阉割版的join操作。
6、都支持处理超大规模数据。

**不同点：**
1、es是java编写，通过RESTFul接口操作数据。mongodb是C++编写，通过driver操作数据。（es对java开发更有好，利于排查理解）
2、mongodb的分片有hash和range两种方式，es只有hash一种。
3、es是天生分布式，主副分片自动分配和复制，开箱即用。mongodb的分布式是由“前置查询路由+配置服务+shard集合”，需要手动配置集群服务。
4、内部存储ES是到排索引+docvalues+fielddata。mongodb暂时未知。
5、es全文检索有强大的分析器且可以灵活组合，查询时智能匹配。mongodb的全文检索字段个数有限制。
6、es所有字段自动索引，mongodb的字段需要手动索引。
7、MongoDB支持多文档事务


mongorestore -d taibai /opt/stocks：  导入数据命令

MongoDB特性：灵活性，可扩展性，强大的查询语言，优异的性能

# mongodb安装与启动

mongodb下载网址:    https://www.mongodb.com/download-center/community

选择版本、系统环境、包  。以下将以TGZ 为例

下载完成将得到mongodb-linux-x86_64-rhel70-4.2.3.tgz包，将其上传到linux（sentos7）某个目录

上传完成后解压并移动到/usr/local/mongodb目录

```
mv mongodb-linux-x86_64-rhel70-4.2.3 /usr/local/mongodb
```

创建专门的负责的用户并赋予权限

```
cd /usr/local/mongodb 
groupadd mongodb
useradd -s /sbin/nologin -g mongodb -M mongodb
mkdir data log run
chown -R mongodb:mongodb data log run
```

在/usr/local/mongodb 里面创建一个配置文件 mongodb.conf 

vim mongodb.conf  并写入下面的信息:

```
bind_ip=0.0.0.0
port=27017
dbpath=/usr/local/mongodb/data/
logpath=/usr/local/mongodb/log/mongodb.log
pidfilepath =/usr/local/mongodb/run/mongodb.pid
logappend=true
fork=true    
maxConns=500
noauth = true


配置解释：
fork=true   运行在后台
logappend=true  如果为 true，当 mongod/mongos 重启后，将在现有日志的尾部继续添加日志。否则，将会备份当前日志文件，然后创建一个新的日志文件；默认为 false。
noauth = true  不需要验证用户密码
maxConns=500 最大同时连接数 默认2000
```

以上是MongoDB的安装与启动的准备工作，可直接启动  启动命令：

```
/usr/local/mongodb/bin/mongod -f /usr/local/mongodb/mongodb.conf
```

**配置环境变量**

vim /etc/profile 

在/etc/profile文件末尾添加一行: 

```
 export PATH=/usr/local/mongodb/bin:$PATH
```

让其生效:

```
source /etc/profile
```

查看当前mongodb的版本:

```
mongod --version
```

# MongoDB的crud

执行mongo命令连接MongoDB

#### 数据库的操作

MongoDB自带的原始数据库：

admin：从权限角度来看，这是“root”数据库。如果将一个用户添加到这个数据库。这个用户自动继承所有数据库的权限。一些特定的服务器端命令也只能从这个数据库运行，比如列出所有的数据库或者关闭服务器

local：这个数据永远不会被复制，可以用来存储限于本地单台服务器的任意集合

config：当mongo用于分片设置时，config数据库在内部使用。用于保存分片的相关信息

**查看**：show dbs

**创建**：use 数据库名 

使用use时，如果数据库存在则会进入到相应的数 据库，如果不存在则会自动创建 – 一旦进入数据库，则可以使用db来引用当前库 ;如果是第一次创建，那个这个数据库只是存在于内存当中，直到这个数据库中创建了集合后才会持久化到磁盘上

**删除**：db.dropDatabase()

用于删除已经持久化的数据库，刚创建在内存中的数据库删除无效

#### 集合的操作

**创建：**db.createCollection("集合名称")

**查看：** show  tables  或者  show collections

**删除：**db.集合名称.drop() 

#### 文档的操作

**添加：**

```
db.xyj.insert({name:"猪八戒",age:28,gender:"男"});   

db.xyj.insertOne({_id:"workd",name:"猪八戒",age:28,gender:"男"});
```

xyj是集合名称，如果该集合还没有被创建，则会自动创建

**批量添加：**

```
db.xyj.insert([
     {name:"沙和尚",age:38,gender:"男"},
     {name:"白骨精",age:16,gender:"女"},
     {name:"蜘蛛精",age:14,gender:"女"}
]);

db.xyj.insertMany([
     {name:"沙和尚",age:38,gender:"男"},
     {name:"白骨精",age:16,gender:"女"},
     {name:"蜘蛛精",age:14,gender:"女"}
]);


db.collection.insertOne()     插入一个文档对象
db.collection.insertMany()    插入多个文档对象
```

 当我们向集合中插入文档时，如果没有给文档指定 _id属性，则数据库会自动为文档添加 _id该属性用来作为文档的唯一标识  _id我们可以自己指定，如果我们指定了数据库就不会在添加了，如果自己指定 _id 也必须确保它的唯一性

**额外小知识**

```
try{
	db.xyj.insert([
      {name:"沙和尚",age:38,gender:"男"},
      {name:"白骨精",age:16,gender:"女"},
      {name:"蜘蛛精",age:14,gender:"女"}
    ]);
}catch(e){
	print(e)
}

可以知道那条插入失败
```

**覆盖修改：**db.xyj.update({_id:"hello"},{age:NumberInt(30)})

执行效果：_id:hello这条数据只有age一个字段了

**局部修改：** db.xyj.update({_id:"workd"},{$set:{age:NumberInt(30)}})

执行效果：只会修改这条数据的某个字段

**批量修改：** db.xyj.update({name:"蜘蛛精"},{$set:{age:NumberInt(100)}},{multi:true})

在修改条数据时，必须要加上第三个参数，否则只会修改一条数据

**字段增加操作：** db.xyj.update({_id:"workd")},{$inc:{age:NumberInt(1)}})

注意：`$inc`对应的字段必须是数字，而且递增或递减的值也必须是数字。

**删除文档：** db.xyj.remove({_id:"workd"})    

**删除文档字段：** db.xyj.update({"_id":123}, {"$unset": {"name":1}})

`$unset`指定字段的值只需是任意合法值即可。

**删除所有：**  db.xyj.remove({})  

```
插入测试数据

db.xyj.insertMany([
     {name:"沙和尚",age:38,gender:"男",hobby:["打篮球","吃喝"]},
     {name:"白骨精",age:16,gender:"女",hobby:["吃喝"]},
     {name:"蜘蛛精",age:14,gender:"女",hobby:["跑步","打乒乓球"]},
	 {name:"唐生",age:25,gender:"男",hobby:["坐禅","吃喝"]}
]);
```

数组添加：

```
db.xyj.update({"name": "白骨精"}, {"$push": {"hobby": "念佛"}})
```

删除元素（`$pop`）：

```
db.xyj.update({"_id": ObjectId("5e80950a2b2f5820a1958a43")}, {"$pop": {"hobby": 1}})  // 删除最后一个元素
db.xyj.update({"_id": ObjectId("5e80950a2b2f5820a1958a43")}, {"$pop": {"hobby": -1}})  // 删除第一个元素
```

删除特定元素（`$pull`）：

```
db.xyj.update({"_id": ObjectId("5e80950a2b2f5820a1958a44")}, {"$pull": {"hobby": "跑步" }})
```

更新嵌套数组的值：

```
db.xyj.insert({name:"猪八戒",age:28,gender:"男",address: [{place: "nanji", tel: 123}, {place: "dongbei", tel: 321}]});


db.xyj.update({"_id": ObjectId("5e809fc665a1965fdc792fc8")}, {"$set": {"address.0.tel": 213}})
```

数组查询：

```
db.xyj.find({"hobby":"跑步"})

多个元素的查询
db.xyj.find({"hobby":{"$all": ["跑步", "打乒乓球"]}})
只有hobby数组同时存在跑步和打乒乓球才会匹配

限制数组长度查询
db.xyj.find({"hobby": {"$size": 2}})
只有数组的长度是2才会匹配
```





**查看所有：**  

```
mongodb:   db.xyj.find()
sql:   select * from xyj
```

**投影查询：**   

```
mongodb:   db.xyj.find({name:"猪八戒"},{name:1,_id:0})   
sql:   select name from xyj where name= "猪八戒"
1表示显示 0表示强制隐藏
```

**按字段条件查询：**

```
mongodb:   db.xyj.find({name:"猪八戒"})
sql:   select * from xyj where name= "猪八戒"
```

**按字段条件查询并只返回一条：**

```
db.xyj.findOne({name:"猪八戒"})
```

**组合查询：** 

```
语法：db.xyj.find($and:[{},{},{}])
//查询年级大于20小于50的
db.xyj.find({$and:[{age:{$gt:NumberInt(20)}},{age:{$lt:NumberInt(50)}}]})  
//查询名字里有”精“的或者年纪大于30的
db.xyj.find({$or:[{age:{$gt:NumberInt(30)}},{name:/精/}]})  
```

**比较查询：** 

```
 db.xyj.find({age:{$gt:NumberInt(20)}})  //查询年级大于20岁的
 
$gt--》大于     $lt--》小于   $gte--》大于等于     $lte--》小于等于   $ne---》不等于(不等于不一定要用于数字)
```

**包含查询：**

```
 db.xyj.find({age:{$in:[28,38]}})
```

**不包含：**

```
 db.xyj.find({age:{$nin:[28,38]}})
```

**Like：**

```
db.xyj.find({"name": /精/})
```

**统计查询:**  

```
db.xyj.count()或者db.xyj.count({字段：条件})
```

**取模：** 

```
db.xyj.find({"age": {$mod: [5, 1]}})
比如我们要匹配 age % 5 == 1
```

**是否存在（$exists)**

```
db.xyj.find({"love": {"$exists": true}})  // 如果存在字段love，就返回
db.xyj.find({"love": {"$exists": false}}) // 如果不存在字段love，就返回
```



#### 分页查询

limit：显示几条记录

skip：跳过几条记录

第一次查询：db.xyj.find().limit(2)  

第一次查询：db.xyj.find().limit(2).skip(2)

结合排序：db.xyj.find().limit(2).skip(2).sort({age:1})   //  1代表升序，-1代表降序

执行顺序： sort > skip > limit



#### **聚合管道**

较常见的管道操作符以及他们的作用：

| 操作符         | 描述                                       | 语法                                       |
| ----------- | ---------------------------------------- | ---------------------------------------- |
| $project    | 数据投影，主要用于重命名，增加，删除字段                     | `db.article.aggregate({ $project : {title : 1 ,author : 1 ,}});` |
| $match      | 过滤，筛选符合条件的文档，作为下一阶段输入                    | `db.articles.aggregate( [{ $match : { score : { $gt : 70, $lte : 90 } } },{ $group: { _id: null, count: { $sum: 1 } } }] );` |
| $limit      | 限制经过管道的文档数量                              | `db.article.aggregate({ $limit : 5 });`  |
| $skip       | 待操作集合处理前跳过部分文档                           | `db.article.aggregate({ $skip : 5 });`   |
| $unwind     | 将数组拆分成独立字段                               | `db.article.aggregate({$project:{author:1,title:1,tags:1}},{$unwind:"$tags"})` |
| $group      | 对数据进行分组                                  | `db.article.aggregate({ $group : {_id : "$author",docsPerAuthor : { $sum : 1 },viewsPerAuthor : { $sum : "$pageViews" }}});` |
| $sort       | 对文档按照指定字段排序                              | `db.users.aggregate( { $sort : { age : -1, posts: 1 } });` |
| $sample     | 随机选择从其输入指定数量的文档。                         | `{ $sample: { size: <positive integer> } }` |
| $out        | 必须为pipeline最后一个阶段管道，因为是将最后计算结果写入到指定的collection中 |                                          |
| $indexStats | 返回数据集合的每个索引的使用情况                         | `{ $indexStats: { } }`                   |

插入测试数据

```
document1=({name:'dogOne',age:1,tags:['animal','dog'],type:'dog',money:[{min:100},{norm:200},{big:300}]});

document2=({name:'catOne',age:3,tags:['animal','cat'],type:'cat',money:[{min:50},{norm:100},{big:200}]});

document3=({name:'catTwo',age:2,tags:['animal','cat'],type:'cat',money:[{min:20},{norm:50},{big:100}]});

document4=({name:'dogTwo',age:5,tags:['animal','dog'],type:'dog',money:[{min:300},{norm:500},{big:700}]});

document5=({name:'appleOne',age:0,tags:['fruit','apple'],type:'apple',money:[{min:10},{norm:12},{big:13}]});

document6=({name:'appleTwo',age:0,tags:['fruit','apple'],type:'apple',money:[{min:10},{norm:12},{big:13}]});

document7=({name:'pineapple',age:0,tags:['fruit','pineapple'],type:'pineapple',money:[{min:8},{norm:9},{big:10}]});

db.mycol.insert(document1)

db.mycol.insert(document2)

db.mycol.insert(document3)

db.mycol.insert(document4)

db.mycol.insert(document5)

db.mycol.insert(document6)

db.mycol.insert(document7)
```

假定我们想提取money中min为100的文档，并且只输出名称和money数组中的min那一项

```
db.mycol.aggregate(
   {$match:{'money.min':100}},
   {$project:{_id:0,name:'$name',minprice:'$money.min'}}
)
```

假定我们想提取money中min小于100的文档，并且限制3个文档，跳过一个文档再显示

```
db.mycol.aggregate(

    {$match:{'money.min':{$lt:100}}},

    {$limit:3},

    {$skip:1},

    {$project:{_id:0,name:'$name',minprice:'$money.min'}}

    )

```

通过type类型来对数据进行分类，并且同时统计他们的年龄age总和

```
db.mycol.aggregate(
    {$group:{_id:'$type',sumage:{$sum:'$age'}}}
)
```

按照年龄对数据进行排序

```
db.mycol.aggregate(
    {$group:{_id:'$type',sumage:{$sum:'$age'}}},
    {$sort:{sumage:1}}
)
```

# 文档模式设计的基本策略

- 其实很简单，我们一般建议的是先考虑内嵌， 直接按照你的对象模型来设计你的数据模型。如果你的对象模型数量不多，关系不是很复杂，那么恭喜你，可能直接一种对象对应一个集合就可以了。
- 内嵌是文档模型的特色，可以充分利用MongoDB的富文档功能来享受我们刚才谈到的一些文档模型的性能和扩展性等特性。一般的一对一、一对多关系，比如说一个人多个地址多个电话等等都可以放在一个文档里用内嵌来完成。
- 但是有一些时候，使用引用则难以避免。比如说， 一个明星的博客可能有几十万或者几百万的回复，这个时候如果把comments放到一个数组里，可能会超出16M的限制。这个时候你可以考虑使用引用的方式，在主表里存储一个id值，指向另一个表中的 id 值。使用引用要注意的就是：从性能上讲，一般我们可能需要两次以上才能把需要的数据取回来。更加重要的是：需要把数据存放到两个集合里，但是目前为止MongoDB并不支持跨表的事务性，所以对于强事务的应用场景要谨慎使用。

![MongoDB-模式设计进阶案例_页面_06](C:\Users\taibai\Desktop\MongoDB\MongoDB-模式设计进阶案例_页面_06.png)

**设计原则：**

为应用程序提供服务，而不是为了存储优化

为实现最佳性能而设计

**购物车场景设计：**

```
db.cart.insert({
  "userid": 1000,
  "last_activity": new Date(),
  "status": "active",
  "items": [
    {
    "itemid": 1000,
    "title": "mianbao",
    "price": 5.00,
    "quantity": 1,
    "img_rul":["mianbao.jpg"]
    }
  ]
});   
```

在这里我们把商品的一些主要信息放到购物车里了，比如说 name,price, quantity，为什么？ 读一次所有信息都拿到了：价格、数量等等，不需要再去查另一张表。这是一种比较常见的优化手段，用冗余的方式来提供读取性能。

向购物车添加数据：

比如说，如果我们想要往购物车里增加一个价值2元的面包，我们可以用下面的update语句。注意$push的用法。$push 类似于javascript的操作符，意思是往数组尾部增加一个元素。

```
db.cart.update({"_id": ObjectId("5e833a5acaea32e75f6f5b5e")},{"$push": {"items":{"itemid": 1001,"title": "niupai","price": 2.00,"quantity": 1,"img_url": "bread.jpg"}},$set:{"last_activity": new Date()}});
```

更新数量：

```
db.cart.update({"_id": ObjectId("5e833a5acaea32e75f6f5b5e"),"items.itemid":1001}, {$set:{"items.$.quantity":5,"last_activity": new Date()}})
```

统计所有商品的总价：

```
db.cart.aggregate(
   [
   {"$match":{"_id": ObjectId("5e833a5acaea32e75f6f5b5e")}},{"$unwind":"$items"},
      {
        $group : {
           _id : null,
           totalPrice: { $sum: { $multiply: [ "$items.price", "$items.quantity" ] } },
           count: { $sum: 1 }
        }
      }
   ]
)

$sum:1 :如果前面的情况出现一次,就加1, 如果后面$sum:2 那么每次前面条件满足一次就加2
使用了$group:_id ：强制必须存在。可以为 null。
 _id : null：则意为对所有行进行分组统计，
```

统计单个商品的总价：

```
db.cart.aggregate(
   [
   {"$match":{"_id": ObjectId("5e833a5acaea32e75f6f5b5e")}},{"$unwind":"$items"},
      { $project: {total_price: { $multiply: [ "$items.price", "$items.quantity" ] } } }
   ]
)
```

**社交网络案例：**

对于关系描述，使用文档模型的内嵌数组特性，我们可以很容易地把我关注的用户（following）和关注我的用户表示出来。下例表示TJ我的关注的用户是mandy和bert，而oscar和mandy则在关注我。这种模式是文档模型中最经典的。但是有一个潜在问题就是如果TJ我是一个明星，他们关注我的人可能有千万。一个千万级的数组会有两个问题：1） 有可能超出一个文档最大16M的硬性限制； 2） MongoDB数组太大会严重影响性能。

![MongoDB-模式设计进阶案例_页面_16-900x695](C:\Users\taibai\Desktop\MongoDB\MongoDB-模式设计进阶案例_页面_16-900x695.png)

怎么办？我们可以建立一个专门的集合来描述关注关系。这里就是一个内嵌和引用的经典选择。我们希望用内嵌，但是如果数组维度太大，就需要考虑用另外一个集合的方式来表示一对多的关系（用户 1–N 关注者）

![MongoDB-模式设计进阶案例_页面_17-900x695](C:\Users\taibai\Desktop\MongoDB\MongoDB-模式设计进阶案例_页面_17-900x695.png)

# 索引

#### 索引的类型

**单字段索引**

**复合索引**

**其他类型索引**

​	哈希索引：是指按照某个字段的hash值来建立索引，目前主要用于MongoDB Sharded Cluster的Hash分片，hash索引只能满足字段完全匹配的查询，不能满足范围查询等。

​	地理位置索引：能很好的解决O2O的应用场景，比如『查找附近的美食』、『查找某个区域内的车站』等。

​	文本索引：能解决快速文本查找的需求，比如有一个博客文章集合，需要根据博客的内容来快速查找，则可以针对博客内容建立文本索引。



#### 索引的操作

###### 创建匿名索引

准备测试数据：

```
for (var i = 1; i<= 100;i++) {
   db.xyj.insert({name: 'taibai'+i,num: i});
}
```

 查找`{num: 100}`文档的过程分析

```
 db.xyj.find({num: 100}).explain(true)
 
返回信息中：
 "stage" : "COLLSCAN",        意思为全表扫描
 executionTimeMillis" : 144   查询时间
```

创建匿名索引

```
db.xyj.createIndex({num: 1})
{
        "createdCollectionAutomatically" : false,
        "numIndexesBefore" : 1,  // 添加本次索引之前的索引数
        "numIndexesAfter" : 2,   // 添加本次索引之后的索引数
        "ok" : 1
}

```

建立索引后再次查找

```
 db.xyj.find({num: 100}).explain(true)
 
 "stage" : "IXSCAN", // 通过扫描索引进行查找
 "indexName" : "num_1",  // 如果没有指定索引名称的话， 默认格式为 ‘字段_1’ 或者 ‘字段_-1’ , 1 代表正序， -1 代表 倒序
 "executionTimeMillis" : 0,  // 本次执行时间的毫秒数为 0ms !!!   
```

创建命名索引

如果没有指定索引名称的话， 默认格式为 `fieldName_1` 或者 `fieldName_-1`，如果想要自定义索引名称的话，可以在创建的时候指定名称，如下：

```
db.xyj.createIndex({name: 1}, {name:' myIndexName'})
```

查询索引：db.xyj.getIndexes()

```
[
        {
                "v" : 2,
                "key" : {
                        "_id" : 1
                },
                "name" : "_id_",
                "ns" : "test.xyj"
        },
        {
                "v" : 2,
                "key" : {
                        "num" : 1
                },
                "name" : "num_1",
                "ns" : "test.xyj"
        },
        {
                "v" : 2,
                "key" : {
                        "name" : 1
                },
                "name" : " myIndexName",
                "ns" : "test.xyj"
        }
]

v:版本号
key： 在那个字段创建了索引，是升序还是降序的
name: 索引的名字
ns： 那个数据库的集合下
```

指定需要使用的索引

```
db.xyj.find({num: "100"}).hint({num:1}).explain(true)   //hint指定强制使用哪个索引
```

###### 创建复合索引

查询的条件不止一个，需要用复合索引

```
db.xyj.createIndex({name:1,num:1});
```

###### 创建ttl（过期）索引

------

mongodb也可以像redis一样给文档设置过期数据，建立ttl索引，这里设置了10秒钟的过期时间，10秒钟之后过期

```
db.classTwo.insert({time:new Date()});
db.classTwo.createIndex({time:1},{expireAfterSeconds:10});
```

- 索引字段的值必须Date对象，不能是其它类型比如时间戳
- 删除时间不精确，每60秒跑一次。删除也要时间，所以有误差。

###### 创建全文索引

准备测试数据：

```
db.xyj.insertMany([
 {"name" : "Lily", "content" : "I am a girl" },
 {"name" : "Tom", "content" : "I am a boy" },
 {"name" : "Carl", "content" : "I do not know boy girl"}
 ])
```

创建全文索引：

```
db.xyj.createIndex({content: 'text'})
```

```
// 查找包含 ‘girl’ 的文档
 db.xyj.find({$text: {$search: 'girl'}})
//查找包含 ‘girl’ 或者 ‘boy’ 的文档（多次查找是 与 的关系）
db.xyj.find({$text: {$search: 'boy girl'}})
// 查找仅包含‘girl’ 不包含‘boy’的文档
db.xyj.find({$text: {$search: 'girl -boy'}})
// 就是要查找包含 ‘girl boy’ 字符的文档（需要转义）
db.xyj.find({$text: {$search: '\"boy girl\"'}})
```

最大的问题,全文索引不支持中文。准确的来说，支持中文的能力没有想象中强大。

```
db.textIndexTest.insert({author:"杜甫",title:"绝句",article:"两个黄鹂鸣翠柳， 　　一行白鹭上青天。窗含西岭千秋雪，门泊东吴万里船。"})
db.textIndexTest.insert({author:"李白",title:"静夜思",article:"床前明月光，疑是地上霜。 举头望明月，低头思故乡。"})
db.textIndexTest.insert({author:"张 王",title:"你好",article:"测试数据"})

创建索引
db.textIndexTest.createIndex( { author: "text" } )

db.textIndexTest.find({$text:{$search:"李白"}}) 
db.textIndexTest.find({$text:{$search:"李"}}) 
db.textIndexTest.find({$text:{$search:"王"}}) 
```



###### 二维（地理位置）索引

准备测试数据：

```
db.dili.insertMany([
{ "gis" : [ 1, 1 ] },
{ "gis" : [ 1, 2 ] },
{ "gis" : [ 1, 3 ] },
{ "gis" : [ 2, 1 ] },
{ "gis" : [ 2, 2 ] },
{ "gis" : [ 2, 3 ] },
{ "gis" : [ 3, 1 ] },
{ "gis" : [ 3, 2 ] },
{ "gis" : [ 3, 3 ] }
]);
```

创建二维索引

```
db.dili.ensureIndex({gis:'2d'})
```

查询距离`[1,1]`最近的四个点

```
 db.dili.find({gis:{$near: [1,1]}}).limit(4)
```

查询以点`[1,2]`和点`[3,3]`为对角线的正方形中的所有的点

```
db.dili.find({gis: {$within:{$box: [[1,2],[3,3]]}}})
```

 查出以`[2,2]`为圆心，以`1` 为半径为规则下圆心面积中的点

```
db.dili.find({gis: {$within:{$center: [[2,2],1]}}})
```

删除索引： db.集合名称.dropIndex( 索引名称或者 创建时的条件 )    //如果是文本索引，只能通过名字删

 db.xyj.dropIndex({age:1})  //条件

 db.xyj.dropIndex("age_1")//名称

 db.xyj.dropIndexes()   //删除所有索引   自带的这个id的索引是不会删除的

# 副本

讲了多么多nosql数据库，一些概念的东西就不写不讲了

### 搭建： 一台机器启动多个实例

创建三个文件夹做为三个MongoDB实例的工作空间

```
mkdir /myMongoDb/mongodb1  -p 
mkdir /myMongoDb/mongodb2  -p 
mkdir /myMongoDb/mongodb3  -p 
```

在三个工作空间下分别创建数据，日志，进程存放等目录或者文件

```
mkdir /myMongoDb/mongodb1/data  -p 
mkdir /myMongoDb/mongodb1/log  -p 
mkdir /myMongoDb/mongodb1/pid  -p 

mkdir /myMongoDb/mongodb2/data  -p 
mkdir /myMongoDb/mongodb2/log  -p 
mkdir /myMongoDb/mongodb2/pid  -p 

mkdir /myMongoDb/mongodb3/data  -p 
mkdir /myMongoDb/mongodb3/log  -p 
mkdir /myMongoDb/mongodb3/pid  -p 

touch /myMongoDb/mongodb1/data/mongodb.log
........
.......
```

在每个工作空间下创建配置文件

```
vim /myMongoDb/mongodb1/mongodb.conf
写入
systemLog:
  destination: file
  logAppend: true
  path: /myMongoDb/mongodb1/log/mongodb.log

storage:
  dbPath: /myMongoDb/mongodb1/data
  journal:
    enabled: true

processManagement:
  fork: true
  pidFilePath: /myMongoDb/mongodb1/pid/mongodb.pid

net:
  port: 7000
  bindIp: 192.168.204.199

replication:
  replSetName: rs0
```

```
vim /myMongoDb/mongodb2/mongodb.conf
写入
systemLog:
  destination: file
  logAppend: true
  path: /myMongoDb/mongodb2/log/mongodb.log

storage:
  dbPath: /myMongoDb/mongodb2/data
  journal:
    enabled: true

processManagement:
  fork: true
  pidFilePath: /myMongoDb/mongodb2/pid/mongodb.pid

net:
  port: 7001
  bindIp: 192.168.204.199

replication:
  replSetName: rs0
```

```
vim /myMongoDb/mongodb3/mongodb.conf
写入
systemLog:
  destination: file
  logAppend: true
  path: /myMongoDb/mongodb3/log/mongodb.log

storage:
  dbPath: /myMongoDb/mongodb3/data
  journal:
    enabled: true

processManagement:
  fork: true
  pidFilePath: /myMongoDb/mongodb3/pid/mongodb.pid

net:
  port: 7002
  bindIp: 192.168.204.199

replication:
  replSetName: rs0
```

启动：

```
/usr/local/mongodb/bin/mongod -f /myMongoDb/mongodb1/mongodb.conf
/usr/local/mongodb/bin/mongod -f /myMongoDb/mongodb2/mongodb.conf
/usr/local/mongodb/bin/mongod -f /myMongoDb/mongodb3/mongodb.conf
```

客户端连接，初始化副本：

```
mongo --host=192.168.204.199 --port=7000

> rs.initiate()
{
        "info2" : "no configuration specified. Using a default configuration for the set",
        "me" : "192.168.204.199:7000",
        "ok" : 1,
        "$clusterTime" : {
                "clusterTime" : Timestamp(1584452776, 1),
                "signature" : {
                        "hash" : BinData(0,"AAAAAAAAAAAAAAAAAAAAAAAAAAA="),
                        "keyId" : NumberLong(0)
                }
        },
        "operationTime" : Timestamp(1584452776, 1)
}
rs0:OTHER> 
rs0:PRIMARY> 

OK为1的话 则初始化成功
rs.conf()  查看副本信息


```

加入副本和仲裁节点

```
rs.add("192.168.204.199:7001")  加入副本节点
rs.addArb("192.168.204.199:7002") 
```

测试：

现在连接的是主节点：

```
use taibai
db.xyj.insert({name:"猪八戒",age:28,gender:"男"});  
```

连接副本节点：

```
mongo --host=192.168.204.199 --port=7001
show dbs-----》  会报错

执行一下：rs.slaveOk() 命令
show dbs-----》  OK

rs.slaveOk()  默认  rs.slaveOk(true)   不想复制的时候可以执行rs.slaveOk(false)

```

连接仲裁节点：

```
mongo --host=192.168.204.199 --port=7002
执行一下：rs.slaveOk() 命令
show dbs
返回
local  0.000GB

仲裁节点不会复制数据

```

### 故障测试：

副本挂掉的情况：

​	主节点和仲裁节点不受影响，如果副本挂掉的时候主节点还写入数据，当副本重启时主会同步数据给副本

主节点挂掉的情况：

​	执行故障转移，从节点成为主节点，仲裁节点没有选举权，当挂掉的主节点回来了。将成为副本

主节点和仲裁节点挂掉的情况：

​	副本还是副本，但是此时随便一个成员加入：

​	仲裁节点加入：  副本成为主节点

​	主节点加入：看谁的数据新

副本节点和仲裁节点挂掉的情况：

​	主节点降级为副本

### 数据是如何复制的？

当一个修改（增删改）操作到达主节点时，它对数据的操作经过特定的转化之后将被记录下来，这些称为oplog，

从节点通过在主节点上打开一个tailable游标不断获取新进入主节点的oplog，并在自己的数据上回放。

### 故障转移

具有投票权的节点之间两两互相发送心跳

当5次心跳未收到则判断为节点失联

如果是主节点，从节点会发起选举，选出主节点

如果失联的是从节点则不会发生新的选举

选举基于RAFT一致性算法，选举成功的必要条件是大多数投票节点存活

复制集中最多可以有50个节点，但最多只能有7个投票节点

**影响选举的因素：**

被选举为主节点的节点必须：

能与多数节点建立连接

具有较新的oplog

根据配置文件的优先级



# 分片

搭建两套副本

创建shard文件夹，作为分片集群工作空间

在shard创建replication1，作为第一个副本工作空间，然后创建三个文件夹做为三个MongoDB实例的工作空间

```
mkdir /shard/replication1/mongodb1  -p 
mkdir /shard/replication1/mongodb2  -p 
mkdir /shard/replication1/mongodb3  -p 
```

在三个工作空间下分别创建数据，日志，进程存放等目录或者文件

```
mkdir /shard/replication1/mongodb1/data  -p 
mkdir /shard/replication1/mongodb1/log  -p 
mkdir /shard/replication1/mongodb1/pid  -p 

mkdir /shard/replication1/mongodb2/data  -p 
mkdir /shard/replication1/mongodb2/log  -p 
mkdir /shard/replication1/mongodb2/pid  -p 

mkdir /shard/replication1/mongodb3/data  -p 
mkdir /shard/replication1/mongodb3/log  -p 
mkdir /shard/replication1/mongodb3/pid  -p 

touch /shard/replication1/mongodb1/log/mongodb.log
touch /shard/replication1/mongodb2/log/mongodb.log
touch /shard/replication1/mongodb3/log/mongodb.log

touch /shard/replication1/mongodb1/pid/mongodb.pid
touch /shard/replication1/mongodb2/pid/mongodb.pid
touch /shard/replication1/mongodb3/pid/mongodb.pid
```

在每个工作空间下创建配置文件

```
vim /shard/replication1/mongodb1/mongodb.conf
写入
systemLog:
  destination: file
  logAppend: true
  path: /shard/replication1/mongodb1/log/mongodb.log

storage:
  dbPath: /shard/replication1/mongodb1/data
  journal:
    enabled: true

processManagement:
  fork: true
  pidFilePath: /shard/replication1/mongodb1/pid/mongodb.pid

net:
  port: 7000
  bindIp: 192.168.204.199

replication:
  replSetName: shard0
sharding:
  clusterRole: shardsvr
```

```
vim /shard/replication1/mongodb2/mongodb.conf
写入
systemLog:
  destination: file
  logAppend: true
  path: /shard/replication1/mongodb2/log/mongodb.log

storage:
  dbPath: /shard/replication1/mongodb2/data
  journal:
    enabled: true

processManagement:
  fork: true
  pidFilePath: /shard/replication1/mongodb2/pid/mongodb.pid

net:
  port: 7001
  bindIp: 192.168.204.199

replication:
  replSetName: shard0
sharding:
  clusterRole: shardsvr
```

```
vim /shard/replication1/mongodb3/mongodb.conf
写入
systemLog:
  destination: file
  logAppend: true
  path: /shard/replication1/mongodb3/log/mongodb.log

storage:
  dbPath: /shard/replication1/mongodb3/data
  journal:
    enabled: true

processManagement:
  fork: true
  pidFilePath: /shard/replication1/mongodb3/pid/mongodb.pid

net:
  port: 7002
  bindIp: 192.168.204.199

replication:
  replSetName: shard0
sharding:
  clusterRole: shardsvr
```

启动：

```
/usr/local/mongodb/bin/mongod -f /shard/replication1/mongodb1/mongodb.conf
/usr/local/mongodb/bin/mongod -f /shard/replication1/mongodb2/mongodb.conf
/usr/local/mongodb/bin/mongod -f /shard/replication1/mongodb3/mongodb.conf
```



在shard创建replication2，作为第二个副本工作空间，然后创建三个文件夹做为三个MongoDB实例的工作空间

```
mkdir /shard/replication2/mongodb1  -p 
mkdir /shard/replication2/mongodb2  -p 
mkdir /shard/replication2/mongodb3  -p 
```

在三个工作空间下分别创建数据，日志，进程存放等目录或者文件

```
mkdir /shard/replication2/mongodb1/data  -p 
mkdir /shard/replication2/mongodb1/log  -p 
mkdir /shard/replication2/mongodb1/pid  -p 

mkdir /shard/replication2/mongodb2/data  -p 
mkdir /shard/replication2/mongodb2/log  -p 
mkdir /shard/replication2/mongodb2/pid  -p 

mkdir /shard/replication2/mongodb3/data  -p 
mkdir /shard/replication2/mongodb3/log  -p 
mkdir /shard/replication2/mongodb3/pid  -p 

touch /shard/replication2/mongodb1/log/mongodb.log
touch /shard/replication2/mongodb2/log/mongodb.log
touch /shard/replication2/mongodb3/log/mongodb.log

touch /shard/replication2/mongodb1/pid/mongodb.pid
touch /shard/replication2/mongodb2/pid/mongodb.pid
touch /shard/replication2/mongodb3/pid/mongodb.pid
```

在每个工作空间下创建配置文件

```
vim /shard/replication2/mongodb1/mongodb.conf
写入
systemLog:
  destination: file
  logAppend: true
  path: /shard/replication2/mongodb1/log/mongodb.log

storage:
  dbPath: /shard/replication2/mongodb1/data
  journal:
    enabled: true

processManagement:
  fork: true
  pidFilePath: /shard/replication2/mongodb1/pid/mongodb.pid

net:
  port: 8000
  bindIp: 192.168.204.199

replication:
  replSetName: shard1
sharding:
  clusterRole: shardsvr
```

```
vim /shard/replication2/mongodb2/mongodb.conf
写入
systemLog:
  destination: file
  logAppend: true
  path: /shard/replication2/mongodb2/log/mongodb.log

storage:
  dbPath: /shard/replication2/mongodb2/data
  journal:
    enabled: true

processManagement:
  fork: true
  pidFilePath: /shard/replication2/mongodb2/pid/mongodb.pid

net:
  port: 8001
  bindIp: 192.168.204.199

replication:
  replSetName: shard1
sharding:
  clusterRole: shardsvr
```

```
vim /shard/replication2/mongodb3/mongodb.conf
写入
systemLog:
  destination: file
  logAppend: true
  path: /shard/replication2/mongodb3/log/mongodb.log

storage:
  dbPath: /shard/replication2/mongodb3/data
  journal:
    enabled: true

processManagement:
  fork: true
  pidFilePath: /shard/replication2/mongodb3/pid/mongodb.pid

net:
  port: 8002
  bindIp: 192.168.204.199

replication:
  replSetName: shard1
sharding:
  clusterRole: shardsvr
```

启动：

```
/usr/local/mongodb/bin/mongod -f /shard/replication2/mongodb1/mongodb.conf
/usr/local/mongodb/bin/mongod -f /shard/replication2/mongodb2/mongodb.conf
/usr/local/mongodb/bin/mongod -f /shard/replication2/mongodb3/mongodb.conf
```



在shard创建configservice，作为配置节点副本工作空间，然后创建三个文件夹做为三个MongoDB实例的工作空间

```
mkdir /shard/configservice/mongodb1  -p 
mkdir /shard/configservice/mongodb2  -p 
mkdir /shard/configservice/mongodb3  -p 
```

在三个工作空间下分别创建数据，日志，进程存放等目录或者文件

```
mkdir /shard/configservice/mongodb1/data  -p 
mkdir /shard/configservice/mongodb1/log  -p 
mkdir /shard/configservice/mongodb1/pid  -p 

mkdir /shard/configservice/mongodb2/data  -p 
mkdir /shard/configservice/mongodb2/log  -p 
mkdir /shard/configservice/mongodb2/pid  -p 

mkdir /shard/configservice/mongodb3/data  -p 
mkdir /shard/configservice/mongodb3/log  -p 
mkdir /shard/configservice/mongodb3/pid  -p 

touch /shard/configservice/mongodb1/log/mongodb.log
touch /shard/configservice/mongodb2/log/mongodb.log
touch /shard/configservice/mongodb3/log/mongodb.log

touch /shard/configservice/mongodb1/pid/mongodb.pid
touch /shard/configservice/mongodb2/pid/mongodb.pid
touch /shard/configservice/mongodb3/pid/mongodb.pid
```

在每个工作空间下创建配置文件

```
vim /shard/configservice/mongodb1/mongodb.conf
写入
systemLog:
  destination: file
  logAppend: true
  path: /shard/configservice/mongodb1/log/mongodb.log

storage:
  dbPath: /shard/configservice/mongodb1/data
  journal:
    enabled: true

processManagement:
  fork: true
  pidFilePath: /shard/configservice/mongodb1/pid/mongodb.pid

net:
  port: 9000
  bindIp: 192.168.204.199

replication:
  replSetName: shard2
sharding:
  clusterRole: configsvr
```

```
vim /shard/configservice/mongodb2/mongodb.conf
写入
systemLog:
  destination: file
  logAppend: true
  path: /shard/configservice/mongodb2/log/mongodb.log

storage:
  dbPath: /shard/configservice/mongodb2/data
  journal:
    enabled: true

processManagement:
  fork: true
  pidFilePath: /shard/configservice/mongodb2/pid/mongodb.pid

net:
  port: 9001
  bindIp: 192.168.204.199

replication:
  replSetName: shard2
sharding:
  clusterRole: configsvr
```

```
vim /shard/configservice/mongodb3/mongodb.conf
写入
systemLog:
  destination: file
  logAppend: true
  path: /shard/configservice/mongodb3/log/mongodb.log

storage:
  dbPath: /shard/configservice/mongodb3/data
  journal:
    enabled: true

processManagement:
  fork: true
  pidFilePath: /shard/configservice/mongodb3/pid/mongodb.pid

net:
  port: 9002
  bindIp: 192.168.204.199

replication:
  replSetName: shard2
sharding:
  clusterRole: configsvr
```

启动：

```
/usr/local/mongodb/bin/mongod -f /shard/configservice/mongodb1/mongodb.conf
/usr/local/mongodb/bin/mongod -f /shard/configservice/mongodb2/mongodb.conf
/usr/local/mongodb/bin/mongod -f /shard/configservice/mongodb3/mongodb.conf
```

客户端连接，初始化第一套副本：

```
mongo --host=192.168.204.199 --port=7000

> rs.initiate()
```

加入副本和仲裁节点

```
rs.add("192.168.204.199:7001")  加入副本节点
rs.addArb("192.168.204.199:7002") 
```



客户端连接，初始化第二套副本：

```
mongo --host=192.168.204.199 --port=8000

> rs.initiate()
```

加入副本和仲裁节点

```
rs.add("192.168.204.199:8001")  加入副本节点
rs.addArb("192.168.204.199:8002") 
```



客户端连接，初始化配置节点副本：

```
mongo --host=192.168.204.199 --port=9000

> rs.initiate()
```

加入副本和仲裁节点

```
rs.add("192.168.204.199:9001")  加入副本节点
rs.add("192.168.204.199:9002") 
```



创建仲裁节点

在shard下创建route文件夹，然后创建两个文件夹作为路由节点实例工作空间

```
mkdir /shard/route/mongodb1/log -p
mkdir /shard/route/mongodb1/pid -p

touch /shard/route/mongodb1/log/mongodb.log
touch /shard/route/mongodb1/pid/mongodb.pid
```

```
 vim /shard/route/mongodb1/mongos.conf
 写入
 systemLog:
  destination: file
  logAppend: true
  path: /shard/route/mongodb1/log/mongodb.log

processManagement:
  fork: true
  pidFilePath: /shard/route/mongodb1/pid/mongodb.pid

net:
  port: 6000
  bindIp: 192.168.204.199

sharding:
  configDB: shard2/192.168.204.199:9000,192.168.204.199:9001,192.168.204.199:9002   #连接配置节点
```

启动

```
mongos -f /shard/route/mongodb1/mongos.conf
```

连接路由节点，将副本加入进来

```
mongos --host=192.168.204.199 --port=6000

sh.addShard(ip:port)   单机
sh.addShard(副本名:/ip:端口,ip:端口....)   副本

sh.addShard("shard0/192.168.204.199:7000,192.168.204.199:7001,192.168.204.199:7002")
sh.addShard("shard1/192.168.204.199:8000,192.168.204.199:8001,192.168.204.199:8002")

sh.status()  查看
```

初始化分片

```
sh.enableSharding("数据库名")   //开启此数据库分片
sh.shardCollection("数据库名.集合名",{"根据此集合中那个字段进行分片":"分片规则是什么"})

sh.enableSharding("taibai")  
sh.shardCollection("taibai.xyj",{"likename":"hashed"})   hash
sh.shardCollection("taibai.sg",{"age":1})   范围
```

测试分片：

注意：必要要在路由节点上执行

```
测试hash
use taibai
for(var i=1;i<=1000;i++){db.comment.insert({_id:i+"";likename:"taibai"+i})}  插入1000条数据
登录7000和8000两个副本查看
7000  ---》  502
8000  ---》  498

登录7000的副本7001，执行rs.slaveOk()命令之后，可能看到数据也同步过来了

测试范围
在做范围测试之前，由于范围默认是写入某个库，直到这个库当中这个集合的数据达到64M之后才会进行分片，然而我们没有这么大的数据量，所有将64M设置为1M
use config
db.settings.save({_id:"chunksize",value: 1})

use taibai
for(var i=1;i<=20000;i++){db.sg.save({"name":"fwefaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"+i,"age":NumberInt(i%2000)})}
登录7000和8000两个副本查看
7000  ---》  7940
8000  ---》  16782
总数据：24722  
```

搭建多个路由节点：

步骤和上面完全一样。。。。。

# 事务相关

MongoDB的事务，可以实现和关系型数据库类似的事务场景

必须使用与MongoDB4.2兼容的驱动

事务默认必须在60s（可调）内完成，否则将被取消

涉及事务的分片不能使用仲裁节点

事务会影响chunk迁移效率。正在迁移的chunk也可能造成事务提交失败，此时重试即可

多文档事务中的读操作必须使用主节点读

readConcern只应该在事务级别设置，不能设置在每次读写操作上

### 写

writeConcern决定一个写操作落到多少个节点上才算成功。writeConcern的取值包括：

0：发起写操作，不关心是否成功

1~集群最大数据节点数：写操作需要被复制到节点执定节点数才算成功

majority: 写操作需要被复制到大多数节点才算成功。

发起写操作的程序将阻塞到写操作到达指定的节点数为止

journal则定义如何才算成功。

true：写操作落到journal文件中才算成功；

false：写操作到达内存即算做成功。

使用：当指定了超过节点数量时，会提示报错，但是数据还是会写入成功

```
db.test.insert({count:2},{writeConcern:{w:3}})
```

**writeConcern实验：**

查看节点设置：

```
var conf=rs.conf()
conf.members
```

设置延迟节点：

```
conf.members[1].slaveDelay=10  设置延迟
conf.members[1].priority=0     设置不参与选举

rs.reconfig(conf)  使配置生效
rs.reconfig()命令执行后，会强制当前的主节点下线，然后进行新的主节点选择。主节点下线时，会关闭所有客户端的连接，这个过程会持续10-20s，因此该操作应该在维护时间执行，减少对系统的影响。
```

设置超时时间：

```
db.test.insert({count:10},{writeConcern:{w:3}})
db.test.insert({count:10},{writeConcern:{w:3,wtimeout:3000}})
```

### 读

readPreference决定读哪个节点

primary：只选择主节点（默认）

primaryPreferred：优先选择主节点，如果不可用则选择从节点

secondary： 只选择从节点

secondaryPreferred：优先选择从节点，如果从节点不可用则选择主节点

nearest：选择最近的节点；

**readPreference选择从节点读实验：**

```
db.test.remove({})
db.test.insert({count:100})
在两个从节点分别执行db.fsyncLock()    禁止同步（写入）
db.test.insert({count:200})
db.test.find().readPref("secondary").readConcern("majority")      //4.0版本中的所有对从节点的读取都将来自快照，无需等待副本数据写入完成。
db.fsyncUnlock()     解锁
```

什么是readConcern？

在readPreference选择了指定的节点后，readConcern决定这个节点上的数据那些是可读的，类似关系型数据库的隔离级别。可选值包括：

available：读取所有可用的数据

local：读取所有可用且属于当前分片的数据

majority：读取在大多数节点上提交完成的数据

linearizable：可线性化读取文档

snapshot：读取最近快照中的数据

**实现安全的读写分离**

写使用writeConcern   majority

读使用readConcern   majority



### Change Stream

 什么是Change Stream

```
Change Stream 是MongoDB用于实现变更追踪的解决方案，类似于关系数据库的触发器，但原理不完全相同：
    
    |              |  Change Stream  |  触发器        |
    |--------------|-----------------|---------------|
    |   触发方式    |  异步           |  同步（事务保证） |
    |   触发位置    |  应用回调事件    |  数据库触发器   |
    |   触发次数    |  每个订阅事件的客户端  |  1次（触发器） |
    |   故障恢复    |  从上次断点重新触发  |  事务回滚    |
    
```

MongoDB 从3.6版本开始支持了 Change Stream 能力（4.0、4.2 版本在能力上做了很多增强），用于订阅 MongoDB 内部的修改操作，change stream 可用于 MongoDB 之间的增量数据迁移、同步，也可以将 MongoDB 的增量订阅应用到其他的关联系统；比如电商场景里，MongoDB 里存储新的订单信息，业务需要根据新增的订单信息去通知库存管理系统发货。

使用 Change Stream 非常简单，mongo shell 封装了针对整个实例、DB、Collection 级别的订阅操作。

注意：集群环境一定要开启enableMajorityReadConcern: true    配置文件中开启

```
db.getMongo().watch()    订阅整个实例的修改
db.watch()               订阅指定DB的修改
db.collection.watch()    订阅指定Collection的修改
```

新建连接1发起订阅操作

```
db.test.watch([], {maxAwaitTimeMS: 60000})    // 最多阻塞等待 1分钟
```

新建连接2写入新数据

```
db.test.insert({x: 100})
db.test.insert({x: 200})
db.test.insert({x: 300})
db.test.insert({x: 400})

上述 ChangeStream 结果里，_id 字段标识着 oplog 的某个位置，如果想从某个位置继续订阅，在 watch 时，通过 resumeAfter 指定即可。比如每个应用订阅了上述3条修改，但只有第一条已经成功消费了，下次订阅时指定第一条的 resume token 即可再次订阅到接下来的2条。

db.test.watch([], {maxAwaitTimeMS: 60000,resumeAfter:{ "_data" : "825E8DB54F000000012B022C0100296E5A1004BDB3105CB8B44265BB704E58E8EA8B0446645F696400645E8DB54FF368D1BC67ABE7AE0004" }}) 
```

