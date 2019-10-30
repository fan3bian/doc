#### 2019/09/24
昨天升级了WMS的最后两台JDK.16的容器。步骤

- 申请新的容器，先进行扩容。[弹性云计算平台](http://cap.jd.com)，需重新编译提测上线。新机器需配置日志接收器。
- 扩容完成后对1.6的容器进行下线。jone执行【停止】，cap执行销毁。logbook中查到日志停止，jsf接口无心跳，未从实例里移除，logbook中一段时间还有两台机器订单信息。

其实有些担心jsf接口和jmq出现问题，不过最后也没有发现什么问题。

rts修复重复下发导致退供单取消的问题，添加了redis防重，但是我的注解最后好像没有生效。

整理了一下isv和gateway接口调用次数，需要恶补一下这方面知识。
#### 2019/09/25 
定位一下防重切面不生效的原因，结果发现生效了。

今天统一配置一下JVM启动参数

```xml
    <if test="po.orgId > 0">
        <![CDATA[    and org_id = #{po.orgId,jdbcType=BIGINT}   ]]>
    </if>
```
这段代码为什么不抛NullPointerExcepiton?

这是双十一我负责的应用
 应用				JVM配置	 JMQ/JSF升级 taskfunnul-api(测试)
clps_so_jproxy	不必         完成		
clps_so			不必         完成		不必
clps_sop		完成         完成		完成
clps_wms		完成         完成		完成
clps-biz-pk		完成         完成		不必
clps-biz-so		完成         完成		完成
clps-base-so	完成         完成		完成
clps_jddl_proxy  不必         完成		不必
clps-compensate 完成         完成		不必
clps_rts	完成         完成			完成
clps_po		完成         完成			完成
clps_rtw	完成         完成			完成
clps_ib		完成         完成			完成
clps-ao		完成         完成			完成
clps_biz_oso	完成         完成		完成
clps_base_oso   完成         完成		完成
clps-biz-po  	完成         完成		不必



jtms状态回传与kdn状态回传优化 50%
运单流程节点handler异常处理 50%
27 sclient接单失效以后无法及时恢复BUG修复 70%
29 20%

#### 2019/09/26
jsf接口下线：
g.jsf.jd.local/com.jd.jsf.deploy.app.InstOperateForDeployService

继续做昨天的jvm参数的配置，发现 po没有接入sgm

1. po模块慢SQL seller上线
2. po接入sgm po上线 店铺商品查询接口优化
3. 

#### 2019/09/27

1. jdk1.6应用容器升级至1.8
2. JVM配置完成
3. 修复退供单重复下发导致退供单取消的BUG
4. taskfunnul-api 测试环境升级，50%
5. 解决采购单慢SQL完成
6. 需求：青龙退货标 测试中
7. 临时任务：ISV接口监控数据收集

taskfunnul-api 测试环境升级 完成

采购单越库 完成

#### 2019/09/29

整理紧急预案： 关闭4种单据的导出
机器扩容申请：

1. 为扩容的机器建立单独的分组，便于缩容。注意配置文件的抽取
2. 注意扩容的机器IP段，出现与原机器不同IP段，需检查该IP段数据库授权
3. 使用IntefaceLog的扩容机器，注意添加logbook的CLPS日志收集器
4. 扩容优先选择jdk1.8_60机器

#### 2019/10/10 预售订单

1. 采购单下发:加可用库存
2. 采购单上架回传：加实物库存
3. 采购单关单：

#### 2019/10/14 预售订单
clps_po dbsql
clps_wms
clps_isv
clps_master

#### 2019/10/16 UCC配置
ucc:线上地址
http://coocmc.jd.com/
需要在UCC上建立部门，应用。应用的负责人为其他参与者授权。
1. 新建配置变量
2. 

#### 2019/10/16 rtw异常

java.lang.NoSuchMethodError: com.jd.ump.profiler.proxy.Profiler.registerInfo(Ljava/lang/String;Ljava/lang/String;ZZ)Lcom/jd/ump/profiler/CallerInfo;

#### 2019/10/21 
蒙牛预入库
clps_isv clps_gateway clps_po clps_wms clps_seller 

#### 2019/10/22 Tuesday
clps-task-funnel 分布式任务调度系统

根据业务模块划分为order stock other三个集群，集群的划分避免了某一个模块任务量激增，通道阻塞，从而影响其他模块的任务(库存重算的任务往往导致挤压)。
同一集群内，正常任务与错误任务(重试任务)使用不同的mq队列，这样又避免了重试任务过多导致通道阻塞，影响到正常任务的运行。

#### 2019/10/28 Monday
```sql
select warehouse_name,warehouse_no from warehouse where warehouse_name in('邢台南宫百川云仓1号库','威海中外运云仓1号库','北京悠酒国际云仓1号库','定州隆德体育云仓1号库','北京配思云仓1号库','石家庄亮道云仓1号库','廊坊霸州波菲特云仓1号库','北京盛世嘉祥云仓1号库','呼和浩特天马云仓1号库','聊城鑫火云仓1号库','滨州林源云仓1号库','保定奥黛莉云仓1号库','保定顺捷云仓1号库','北京凤鸟云仓2号库','衡水九号仓河北供应链云仓1号库','日照河川云仓1号库','青岛电商通云仓1号库','张家口卿安云仓1号库','廊坊霸州祥其瑞云仓1号库','石家庄微客云仓1号库','沧州任丘鼎盛云仓1号库') and yn = 1;

select a.seller_no,a.seller_name,group_concat(b.dept_no) as dept_nos
from seller a inner join dept b on a.id = b.seller_id where a.seller_name in ('河北御泰汽车用品有限公司','山东连京商业管理有限公司','北京中通春华物流有限公司','定州市隆德体育用品有限公司','北京圆邦富成物流有限公司','石家庄花辰熙月商贸有限公司','天津波菲特家具有限公司','北京华艺拓海文化传媒有限公司','内蒙古天马物流服务有限公司','阳谷县鑫火新能源有限公司','山东林源供应链管理有限公司','保定奥黛莉皮具制造有限公司','高阳县顺捷物流有限公司','北京中和致远科技发展有限公司','九号仓河北供应链管理有限公司','山东河川经贸有限公司','青岛电商通国际物流有限公司','上海卿安信息科技有限公司','福州顺道电子商务有限公司','石家庄微客物流有限公司','任丘市鼎盛采暖设备有限公司') and a.yn = 1 and b.yn = 1 
group by a.seller_name,a.seller_no ;
```
#### 2019/10/30 Wednesday
com.jd.clps.base.so.service.SoQueryService