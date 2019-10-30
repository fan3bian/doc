## 双十一值班关注点

#### 单据导出开关

双十一期间关闭商家销售订单、采购单、退货单、退供应商单的导出功能。excel导出容易在内存中产生大对象，为保证业务的稳定的性能，双十一期间导出功能关闭。
通过UCC(统一配置中心)


订单详情:

{"orderCode":"CSL0001","buyer":"zhangshuyi"}

        jQuery.ajax({
            "type": "post",
            "url": "/order/add.do?rand=" + Math.random(),
            "data": {"orderCode":"CSL0001","buyer":"zhangshuyi"},
            success: function (resp) {}
            });
