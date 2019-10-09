#### 线上问题分析

CSL8816149692680 

1. 未生成订单
2. 订单预占了库存

9月18日扣减库存接口超时，库存预占成功。服务端成功，客户端失败。

checkTimeoutSoHandler 


#### 订单模块外部门系统接口
//gis地址解析服务
com.jd.lbs.consgaddress.jsf.IAddressMapService
//台账
com.jd.orderbank.export.rest.OrderQueryResource
//大件 TC
com.jd.las.delivery.service.DeliveryRouterRpcService
//骑行路径规划服务 众包
com.jd.lbs.jdlbsapi.search.DirectionService
//地理编码服务 众包
com.jd.lbs.geocode.api.GeocodingService
//京东订单中间件 pop
com.jd.ioms.jsf.export.OrderMiddlewareJSFService
//青龙tms
com.jd.tms.basic.ws.BasicQueryWS
//短信
com.jd.mobilePhoneMsg.sender.client.service.SmsMessageRpcService
//青龙运单轨迹
com.jd.etms.waybill.api.WaybillTraceApi
//jwms取消接口
com.jd.jwms.front.service.JWmsOpenService
//青龙拦截接口
com.jd.etms.receive.api.saf.ExceptionOrderSaf
//青龙包裹接口
com.jd.etms.receive.api.saf.ExceptionOrderSaf
//京东支付二维码接口
com.jd.jr.pospay.mobile.gateway.export.inf.QrOrderService


