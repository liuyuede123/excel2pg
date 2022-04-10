package config

var CreateSql = `
--105张表的建表语句
CREATE TABLE IF NOT EXISTS bi_master_test.public."万相台_创意_货品加速" 
(
    日期 varchar(200) null
    ,效果模型 varchar(200) null
    ,统计口径 varchar(200) null
    ,转化周期 varchar(200) null
    ,创意id varchar(200) null
    ,商品id varchar(200) null
    ,店铺url varchar(200) null
    ,计划id varchar(200) null
    ,计划名称 varchar(200) null
    ,消耗 varchar(200) null
    ,曝光量 varchar(200) null
    ,点击量 varchar(200) null
    ,点击率 varchar(200) null
    ,点击成本 varchar(200) null
    ,直接加购数 varchar(200) null
    ,间接加购数 varchar(200) null
    ,总加购量 varchar(200) null
    ,收藏量 varchar(200) null
    ,加购收藏成本 varchar(200) null
    ,预售成交笔数 varchar(200) null
    ,预售成交金额 varchar(200) null
    ,直接成交笔数 varchar(200) null
    ,间接成交笔数 varchar(200) null
    ,总成交笔数 varchar(200) null
    ,直接成交金额 varchar(200) null
    ,间接成交金额 varchar(200) null
    ,总成交金额 varchar(200) null
    ,成交转化率 varchar(200) null
    ,roi varchar(200) null
    ,新客成交金额贡献占比 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."万相台_账户" 
(
    日期 varchar(200) null
    ,优化目标 varchar(200) null
    ,效果模型 varchar(200) null
    ,统计口径 varchar(200) null
    ,转化周期 varchar(200) null
    ,消耗 varchar(200) null
    ,曝光量 varchar(200) null
    ,点击量 varchar(200) null
    ,点击率 varchar(200) null
    ,点击成本 varchar(200) null
    ,收藏量 varchar(200) null
    ,加购收藏量 varchar(200) null
    ,加购收藏成本 varchar(200) null
    ,成交转化率 varchar(200) null
    ,roi varchar(200) null
    ,预售成交笔数 varchar(200) null
    ,预售成交金额 varchar(200) null
    ,总加购量 varchar(200) null
    ,直接加购数 varchar(200) null
    ,间接加购数 varchar(200) null
    ,总成交笔数 varchar(200) null
    ,直接成交笔数 varchar(200) null
    ,间接成交笔数 varchar(200) null
    ,总成交金额 varchar(200) null
    ,直接成交金额 varchar(200) null
    ,间接成交金额 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_品牌专区报表_账户" 
(
    日期 varchar(200) null
    ,搜索量 varchar(200) null
    ,搜索访客数 varchar(200) null
    ,访客触达率 varchar(200) null
    ,展现量 varchar(200) null
    ,自然流量增量曝光 varchar(200) null
    ,触达访客数 varchar(200) null
    ,回搜展现量 varchar(200) null
    ,回搜触达访客数 varchar(200) null
    ,点击量 varchar(200) null
    ,跳转点击量 varchar(200) null
    ,点击访客数 varchar(200) null
    ,点击率 varchar(200) null
    ,进店访客数 varchar(200) null
    ,互动点击量 varchar(200) null
    ,跳转点击率 varchar(200) null
    ,回搜点击量 varchar(200) null
    ,回搜点击访客数 varchar(200) null
    ,店铺收藏数 varchar(200) null
    ,宝贝收藏数 varchar(200) null
    ,宝贝加购数 varchar(200) null
    ,宝贝浏览数 varchar(200) null
    ,店铺浏览数 varchar(200) null
    ,访问时长 varchar(200) null
    ,店铺收藏访客数 varchar(200) null
    ,宝贝收藏访客数 varchar(200) null
    ,宝贝加购访客数 varchar(200) null
    ,宝贝浏览访客数 varchar(200) null
    ,成交笔数 varchar(200) null
    ,成交金额 varchar(200) null
    ,转化率 varchar(200) null
    ,预售成交笔数 varchar(200) null
    ,成交访客数 varchar(200) null
    ,自然流量增量成交 varchar(200) null
    ,预售成交金额 varchar(200) null
    ,搜索进店率 varchar(200) null
    ,进店行动率 varchar(200) null
    ,行动成交率 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_明星店铺报表_创意" 
(
    日期 varchar(200) null
    ,计划名称 varchar(200) null
    ,单元名称 varchar(200) null
    ,创意名称 varchar(200) null
    ,展现量 varchar(200) null
    ,自然流量增量曝光 varchar(200) null
    ,消耗 varchar(200) null
    ,触达访客数 varchar(200) null
    ,千次展现成本 varchar(200) null
    ,点击单价 varchar(200) null
    ,跳转点击单价 varchar(200) null
    ,点击量 varchar(200) null
    ,跳转点击量 varchar(200) null
    ,点击访客数 varchar(200) null
    ,互动点击量 varchar(200) null
    ,点击率 varchar(200) null
    ,跳转点击率 varchar(200) null
    ,进店访客数 varchar(200) null
    ,店铺收藏数 varchar(200) null
    ,宝贝收藏数 varchar(200) null
    ,宝贝加购数 varchar(200) null
    ,行动访客数 varchar(200) null
    ,店铺收藏访客数 varchar(200) null
    ,宝贝收藏访客数 varchar(200) null
    ,宝贝加购访客数 varchar(200) null
    ,成交笔数 varchar(200) null
    ,成交金额 varchar(200) null
    ,回报率 varchar(200) null
    ,转化率 varchar(200) null
    ,自然流量增量成交 varchar(200) null
    ,预售成交笔数 varchar(200) null
    ,预售成交金额 varchar(200) null
    ,成交访客数 varchar(200) null
    ,搜索进店率 varchar(200) null
    ,进店行动率 varchar(200) null
    ,行动成交率 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_明星店铺报表_品牌流量包" 
(
    日期 varchar(200) null
    ,计划名称 varchar(200) null
    ,单元名称 varchar(200) null
    ,词包名称 varchar(200) null
    ,展现量 varchar(200) null
    ,自然流量增量曝光 varchar(200) null
    ,消耗 varchar(200) null
    ,触达访客数 varchar(200) null
    ,千次展现成本 varchar(200) null
    ,点击单价 varchar(200) null
    ,跳转点击单价 varchar(200) null
    ,点击量 varchar(200) null
    ,跳转点击量 varchar(200) null
    ,点击访客数 varchar(200) null
    ,互动点击量 varchar(200) null
    ,点击率 varchar(200) null
    ,跳转点击率 varchar(200) null
    ,进店访客数 varchar(200) null
    ,店铺收藏数 varchar(200) null
    ,宝贝收藏数 varchar(200) null
    ,宝贝加购数 varchar(200) null
    ,行动访客数 varchar(200) null
    ,店铺收藏访客数 varchar(200) null
    ,宝贝收藏访客数 varchar(200) null
    ,宝贝加购访客数 varchar(200) null
    ,成交笔数 varchar(200) null
    ,成交金额 varchar(200) null
    ,回报率 varchar(200) null
    ,转化率 varchar(200) null
    ,自然流量增量成交 varchar(200) null
    ,预售成交笔数 varchar(200) null
    ,预售成交金额 varchar(200) null
    ,成交访客数 varchar(200) null
    ,搜索进店率 varchar(200) null
    ,进店行动率 varchar(200) null
    ,行动成交率 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_明星店铺报表_定向人群" 
(
    日期 varchar(200) null
    ,计划名称 varchar(200) null
    ,单元名称 varchar(200) null
    ,创意名称 varchar(200) null
    ,定向人群名称 varchar(200) null
    ,展现量 varchar(200) null
    ,自然流量增量曝光 varchar(200) null
    ,消耗 varchar(200) null
    ,触达访客数 varchar(200) null
    ,千次展现成本 varchar(200) null
    ,点击单价 varchar(200) null
    ,跳转点击单价 varchar(200) null
    ,点击量 varchar(200) null
    ,跳转点击量 varchar(200) null
    ,点击访客数 varchar(200) null
    ,互动点击量 varchar(200) null
    ,点击率 varchar(200) null
    ,跳转点击率 varchar(200) null
    ,进店访客数 varchar(200) null
    ,店铺收藏数 varchar(200) null
    ,宝贝收藏数 varchar(200) null
    ,宝贝加购数 varchar(200) null
    ,行动访客数 varchar(200) null
    ,店铺收藏访客数 varchar(200) null
    ,宝贝收藏访客数 varchar(200) null
    ,宝贝加购访客数 varchar(200) null
    ,成交笔数 varchar(200) null
    ,成交金额 varchar(200) null
    ,回报率 varchar(200) null
    ,转化率 varchar(200) null
    ,自然流量增量成交 varchar(200) null
    ,预售成交笔数 varchar(200) null
    ,预售成交金额 varchar(200) null
    ,成交访客数 varchar(200) null
    ,搜索进店率 varchar(200) null
    ,进店行动率 varchar(200) null
    ,行动成交率 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_明星店铺报表_推广单元" 
(
    日期 varchar(200) null
    ,计划名称 varchar(200) null
    ,单元名称 varchar(200) null
    ,展现量 varchar(200) null
    ,自然流量增量曝光 varchar(200) null
    ,消耗 varchar(200) null
    ,触达访客数 varchar(200) null
    ,千次展现成本 varchar(200) null
    ,点击单价 varchar(200) null
    ,跳转点击单价 varchar(200) null
    ,点击量 varchar(200) null
    ,跳转点击量 varchar(200) null
    ,点击访客数 varchar(200) null
    ,互动点击量 varchar(200) null
    ,点击率 varchar(200) null
    ,跳转点击率 varchar(200) null
    ,进店访客数 varchar(200) null
    ,店铺收藏数 varchar(200) null
    ,宝贝收藏数 varchar(200) null
    ,宝贝加购数 varchar(200) null
    ,行动访客数 varchar(200) null
    ,店铺收藏访客数 varchar(200) null
    ,宝贝收藏访客数 varchar(200) null
    ,宝贝加购访客数 varchar(200) null
    ,成交笔数 varchar(200) null
    ,成交金额 varchar(200) null
    ,回报率 varchar(200) null
    ,转化率 varchar(200) null
    ,自然流量增量成交 varchar(200) null
    ,预售成交笔数 varchar(200) null
    ,预售成交金额 varchar(200) null
    ,成交访客数 varchar(200) null
    ,搜索进店率 varchar(200) null
    ,进店行动率 varchar(200) null
    ,行动成交率 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_明星店铺报表_推广计划" 
(
    日期 varchar(200) null
    ,计划名称 varchar(200) null
    ,展现量 varchar(200) null
    ,自然流量增量曝光 varchar(200) null
    ,消耗 varchar(200) null
    ,触达访客数 varchar(200) null
    ,千次展现成本 varchar(200) null
    ,点击单价 varchar(200) null
    ,跳转点击单价 varchar(200) null
    ,点击量 varchar(200) null
    ,跳转点击量 varchar(200) null
    ,点击访客数 varchar(200) null
    ,互动点击量 varchar(200) null
    ,点击率 varchar(200) null
    ,跳转点击率 varchar(200) null
    ,进店访客数 varchar(200) null
    ,店铺收藏数 varchar(200) null
    ,宝贝收藏数 varchar(200) null
    ,宝贝加购数 varchar(200) null
    ,行动访客数 varchar(200) null
    ,店铺收藏访客数 varchar(200) null
    ,宝贝收藏访客数 varchar(200) null
    ,宝贝加购访客数 varchar(200) null
    ,成交笔数 varchar(200) null
    ,成交金额 varchar(200) null
    ,回报率 varchar(200) null
    ,转化率 varchar(200) null
    ,自然流量增量成交 varchar(200) null
    ,预售成交笔数 varchar(200) null
    ,预售成交金额 varchar(200) null
    ,成交访客数 varchar(200) null
    ,搜索进店率 varchar(200) null
    ,进店行动率 varchar(200) null
    ,行动成交率 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_明星店铺报表_账户" 
(
    日期 varchar(200) null
    ,搜索量 varchar(200) null
    ,搜索访客数 varchar(200) null
    ,访客触达率 varchar(200) null
    ,展现量 varchar(200) null
    ,自然流量增量曝光 varchar(200) null
    ,消耗 varchar(200) null
    ,触达访客数 varchar(200) null
    ,千次展现成本 varchar(200) null
    ,点击单价 varchar(200) null
    ,跳转点击单价 varchar(200) null
    ,点击量 varchar(200) null
    ,跳转点击量 varchar(200) null
    ,点击访客数 varchar(200) null
    ,互动点击量 varchar(200) null
    ,点击率 varchar(200) null
    ,跳转点击率 varchar(200) null
    ,进店访客数 varchar(200) null
    ,店铺收藏数 varchar(200) null
    ,宝贝收藏数 varchar(200) null
    ,宝贝加购数 varchar(200) null
    ,行动访客数 varchar(200) null
    ,店铺收藏访客数 varchar(200) null
    ,宝贝收藏访客数 varchar(200) null
    ,宝贝加购访客数 varchar(200) null
    ,成交笔数 varchar(200) null
    ,成交金额 varchar(200) null
    ,回报率 varchar(200) null
    ,转化率 varchar(200) null
    ,自然流量增量成交 varchar(200) null
    ,预售成交笔数 varchar(200) null
    ,预售成交金额 varchar(200) null
    ,成交访客数 varchar(200) null
    ,搜索进店率 varchar(200) null
    ,进店行动率 varchar(200) null
    ,行动成交率 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."引力魔方_推荐_展现效果_创意" 
(
    日期 varchar(200) null
    ,计划组 varchar(200) null
    ,计划 varchar(200) null
    ,主体类型 varchar(200) null
    ,主体名称 varchar(200) null
    ,商品id varchar(200) null
    ,创意名称 varchar(200) null
    ,创意id varchar(200) null
    ,消耗 varchar(200) null
    ,展现量 varchar(200) null
    ,点击量 varchar(200) null
    ,点击率 varchar(200) null
    ,千次展现成本 varchar(200) null
    ,点击成本 varchar(200) null
    ,引导访问量 varchar(200) null
    ,引导访问人数 varchar(200) null
    ,引导访问潜客数 varchar(200) null
    ,引导访问潜客占比 varchar(200) null
    ,引导访问率 varchar(200) null
    ,深度访问量 varchar(200) null
    ,平均访问页面数 varchar(200) null
    ,关注店铺量 varchar(200) null
    ,回搜量 varchar(200) null
    ,回访量 varchar(200) null
    ,回搜成本 varchar(200) null
    ,回访成本 varchar(200) null
    ,平均深度访问时长 varchar(200) null
    ,新客获取量 varchar(200) null
    ,拉新率 varchar(200) null
    ,拉新消耗 varchar(200) null
    ,拉新成本 varchar(200) null
    ,收藏宝贝量 varchar(200) null
    ,添加购物车量 varchar(200) null
    ,拍下订单量 varchar(200) null
    ,拍下订单金额 varchar(200) null
    ,预售订单量 varchar(200) null
    ,预售金额 varchar(200) null
    ,加购收藏量 varchar(200) null
    ,加购收藏成本 varchar(200) null
    ,成交订单量 varchar(200) null
    ,成交订单金额 varchar(200) null
    ,展现转化率 varchar(200) null
    ,点击转化率 varchar(200) null
    ,投资回报率 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."引力魔方_推荐_展现效果_定向人群" 
(
    日期 varchar(200) null
    ,计划组 varchar(200) null
    ,计划 varchar(200) null
    ,主体类型 varchar(200) null
    ,主体名称 varchar(200) null
    ,商品id varchar(200) null
    ,定向人群 varchar(200) null
    ,消耗 varchar(200) null
    ,展现量 varchar(200) null
    ,点击量 varchar(200) null
    ,点击率 varchar(200) null
    ,千次展现成本 varchar(200) null
    ,点击成本 varchar(200) null
    ,引导访问量 varchar(200) null
    ,引导访问人数 varchar(200) null
    ,引导访问潜客数 varchar(200) null
    ,引导访问潜客占比 varchar(200) null
    ,引导访问率 varchar(200) null
    ,深度访问量 varchar(200) null
    ,平均访问页面数 varchar(200) null
    ,关注店铺量 varchar(200) null
    ,回搜量 varchar(200) null
    ,回访量 varchar(200) null
    ,回搜成本 varchar(200) null
    ,回访成本 varchar(200) null
    ,平均深度访问时长 varchar(200) null
    ,新客获取量 varchar(200) null
    ,拉新率 varchar(200) null
    ,拉新消耗 varchar(200) null
    ,拉新成本 varchar(200) null
    ,收藏宝贝量 varchar(200) null
    ,添加购物车量 varchar(200) null
    ,拍下订单量 varchar(200) null
    ,拍下订单金额 varchar(200) null
    ,预售订单量 varchar(200) null
    ,预售金额 varchar(200) null
    ,加购收藏量 varchar(200) null
    ,加购收藏成本 varchar(200) null
    ,成交订单量 varchar(200) null
    ,成交订单金额 varchar(200) null
    ,展现转化率 varchar(200) null
    ,点击转化率 varchar(200) null
    ,投资回报率 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."引力魔方_推荐_展现效果_资源位" 
(
    日期 varchar(200) null
    ,计划组 varchar(200) null
    ,计划 varchar(200) null
    ,主体类型 varchar(200) null
    ,主体名称 varchar(200) null
    ,商品id varchar(200) null
    ,资源位 varchar(200) null
    ,消耗 varchar(200) null
    ,展现量 varchar(200) null
    ,点击量 varchar(200) null
    ,点击率 varchar(200) null
    ,千次展现成本 varchar(200) null
    ,点击成本 varchar(200) null
    ,引导访问量 varchar(200) null
    ,引导访问人数 varchar(200) null
    ,引导访问潜客数 varchar(200) null
    ,引导访问潜客占比 varchar(200) null
    ,引导访问率 varchar(200) null
    ,深度访问量 varchar(200) null
    ,平均访问页面数 varchar(200) null
    ,关注店铺量 varchar(200) null
    ,回搜量 varchar(200) null
    ,回访量 varchar(200) null
    ,回搜成本 varchar(200) null
    ,回访成本 varchar(200) null
    ,平均深度访问时长 varchar(200) null
    ,新客获取量 varchar(200) null
    ,拉新率 varchar(200) null
    ,拉新消耗 varchar(200) null
    ,拉新成本 varchar(200) null
    ,收藏宝贝量 varchar(200) null
    ,添加购物车量 varchar(200) null
    ,拍下订单量 varchar(200) null
    ,拍下订单金额 varchar(200) null
    ,预售订单量 varchar(200) null
    ,预售金额 varchar(200) null
    ,加购收藏量 varchar(200) null
    ,加购收藏成本 varchar(200) null
    ,成交订单量 varchar(200) null
    ,成交订单金额 varchar(200) null
    ,展现转化率 varchar(200) null
    ,点击转化率 varchar(200) null
    ,投资回报率 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."引力魔方_推荐_点击效果_创意" 
(
    日期 varchar(200) null
    ,计划组 varchar(200) null
    ,计划 varchar(200) null
    ,主体类型 varchar(200) null
    ,主体名称 varchar(200) null
    ,商品id varchar(200) null
    ,创意名称 varchar(200) null
    ,创意id varchar(200) null
    ,消耗 varchar(200) null
    ,展现量 varchar(200) null
    ,点击量 varchar(200) null
    ,点击率 varchar(200) null
    ,千次展现成本 varchar(200) null
    ,点击成本 varchar(200) null
    ,引导访问量 varchar(200) null
    ,引导访问人数 varchar(200) null
    ,引导访问潜客数 varchar(200) null
    ,引导访问潜客占比 varchar(200) null
    ,引导访问率 varchar(200) null
    ,深度访问量 varchar(200) null
    ,平均访问页面数 varchar(200) null
    ,关注店铺量 varchar(200) null
    ,回搜量 varchar(200) null
    ,回访量 varchar(200) null
    ,回搜成本 varchar(200) null
    ,回访成本 varchar(200) null
    ,平均深度访问时长 varchar(200) null
    ,新客获取量 varchar(200) null
    ,拉新率 varchar(200) null
    ,拉新消耗 varchar(200) null
    ,拉新成本 varchar(200) null
    ,收藏宝贝量 varchar(200) null
    ,添加购物车量 varchar(200) null
    ,拍下订单量 varchar(200) null
    ,拍下订单金额 varchar(200) null
    ,预售订单量 varchar(200) null
    ,预售金额 varchar(200) null
    ,加购收藏量 varchar(200) null
    ,加购收藏成本 varchar(200) null
    ,成交订单量 varchar(200) null
    ,成交订单金额 varchar(200) null
    ,展现转化率 varchar(200) null
    ,点击转化率 varchar(200) null
    ,投资回报率 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."引力魔方_推荐_点击效果_定向人群" 
(
    日期 varchar(200) null
    ,计划组 varchar(200) null
    ,计划 varchar(200) null
    ,主体类型 varchar(200) null
    ,主体名称 varchar(200) null
    ,商品id varchar(200) null
    ,定向人群 varchar(200) null
    ,消耗 varchar(200) null
    ,展现量 varchar(200) null
    ,点击量 varchar(200) null
    ,点击率 varchar(200) null
    ,千次展现成本 varchar(200) null
    ,点击成本 varchar(200) null
    ,引导访问量 varchar(200) null
    ,引导访问人数 varchar(200) null
    ,引导访问潜客数 varchar(200) null
    ,引导访问潜客占比 varchar(200) null
    ,引导访问率 varchar(200) null
    ,深度访问量 varchar(200) null
    ,平均访问页面数 varchar(200) null
    ,关注店铺量 varchar(200) null
    ,回搜量 varchar(200) null
    ,回访量 varchar(200) null
    ,回搜成本 varchar(200) null
    ,回访成本 varchar(200) null
    ,平均深度访问时长 varchar(200) null
    ,新客获取量 varchar(200) null
    ,拉新率 varchar(200) null
    ,拉新消耗 varchar(200) null
    ,拉新成本 varchar(200) null
    ,收藏宝贝量 varchar(200) null
    ,添加购物车量 varchar(200) null
    ,拍下订单量 varchar(200) null
    ,拍下订单金额 varchar(200) null
    ,预售订单量 varchar(200) null
    ,预售金额 varchar(200) null
    ,加购收藏量 varchar(200) null
    ,加购收藏成本 varchar(200) null
    ,成交订单量 varchar(200) null
    ,成交订单金额 varchar(200) null
    ,展现转化率 varchar(200) null
    ,点击转化率 varchar(200) null
    ,投资回报率 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."引力魔方_推荐_点击效果_资源位" 
(
    日期 varchar(200) null
    ,计划组 varchar(200) null
    ,计划 varchar(200) null
    ,主体类型 varchar(200) null
    ,主体名称 varchar(200) null
    ,商品id varchar(200) null
    ,资源位 varchar(200) null
    ,消耗 varchar(200) null
    ,展现量 varchar(200) null
    ,点击量 varchar(200) null
    ,点击率 varchar(200) null
    ,千次展现成本 varchar(200) null
    ,点击成本 varchar(200) null
    ,引导访问量 varchar(200) null
    ,引导访问人数 varchar(200) null
    ,引导访问潜客数 varchar(200) null
    ,引导访问潜客占比 varchar(200) null
    ,引导访问率 varchar(200) null
    ,深度访问量 varchar(200) null
    ,平均访问页面数 varchar(200) null
    ,关注店铺量 varchar(200) null
    ,回搜量 varchar(200) null
    ,回访量 varchar(200) null
    ,回搜成本 varchar(200) null
    ,回访成本 varchar(200) null
    ,平均深度访问时长 varchar(200) null
    ,新客获取量 varchar(200) null
    ,拉新率 varchar(200) null
    ,拉新消耗 varchar(200) null
    ,拉新成本 varchar(200) null
    ,收藏宝贝量 varchar(200) null
    ,添加购物车量 varchar(200) null
    ,拍下订单量 varchar(200) null
    ,拍下订单金额 varchar(200) null
    ,预售订单量 varchar(200) null
    ,预售金额 varchar(200) null
    ,加购收藏量 varchar(200) null
    ,加购收藏成本 varchar(200) null
    ,成交订单量 varchar(200) null
    ,成交订单金额 varchar(200) null
    ,展现转化率 varchar(200) null
    ,点击转化率 varchar(200) null
    ,投资回报率 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."淘宝联盟_商品分析_商品" 
(
    商品名称 varchar(200) null
    ,商品id varchar(200) null
    ,商品链接 varchar(200) null
    ,"付款佣金支出(元)" varchar(200) null
    ,付款佣金率 varchar(200) null
    ,"付款金额(元)" varchar(200) null
    ,付款笔数 varchar(200) null
    ,商品折扣率 varchar(200) null
    ,进店量 varchar(200) null
    ,收藏宝贝量 varchar(200) null
    ,添加购物车量 varchar(200) null
    ,预售定金笔数 varchar(200) null
    ,"预售定金金额(元)" varchar(200) null
    ,"预估预售尾款金额(元)" varchar(200) null
    ,"预估预售整单金额(元)" varchar(200) null
    ,"结算佣金支出(元)" varchar(200) null
    ,结算佣金率 varchar(200) null
    ,预估预售整单佣金 varchar(200) null
    ,"付款服务费支出(元)" varchar(200) null
    ,付款服务费率 varchar(200) null
    ,"结算服务费支出(元)" varchar(200) null
    ,结算服务费率 varchar(200) null
    ,"单件商品付款支出费用(元)" varchar(200) null
    ,"付款支出费用(元)" varchar(200) null
    ,"结算支出费用(元)" varchar(200) null
    ,预估预售整单佣金率 varchar(200) null
    ,进店人数 varchar(200) null
    ,"平均优惠券面额(元)" varchar(200) null
    ,付款人数 varchar(200) null
    ,统计日期 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."淘宝联盟_成交订单_订单结算" 
(
    确认收货时间 varchar(200) null
    ,账户支出时间 varchar(200) null
    ,淘客结算时间 varchar(200) null
    ,创建时间 varchar(200) null
    ,计划名称 varchar(200) null
    ,商品id varchar(200) null
    ,商品名称 varchar(200) null
    ,实际成交价格 varchar(200) null
    ,成交商品数 varchar(200) null
    ,佣金比例 varchar(200) null
    ,佣金 varchar(200) null
    ,服务费率 varchar(200) null
    ,服务费金额 varchar(200) null
    ,淘宝父订单编号 varchar(200) null
    ,淘宝子订单编号 varchar(200) null
    ,来源或淘客昵称 varchar(200) null
    ,团长名称 varchar(200) null
    ,是否预售未付尾款订单 varchar(200) null
    ,备注 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."淘宝联盟_类目分析_商品" 
(
    叶子类目 varchar(200) null
    ,叶子名称 varchar(200) null
    ,"付款佣金支出(元)" varchar(200) null
    ,付款佣金率 varchar(200) null
    ,"付款金额(元)" varchar(200) null
    ,付款笔数 varchar(200) null
    ,商品折扣率 varchar(200) null
    ,进店量 varchar(200) null
    ,收藏宝贝量 varchar(200) null
    ,添加购物车量 varchar(200) null
    ,预售定金笔数 varchar(200) null
    ,"预售定金金额(元)" varchar(200) null
    ,"预估预售尾款金额(元)" varchar(200) null
    ,"预估预售整单金额(元)" varchar(200) null
    ,"结算佣金支出(元)" varchar(200) null
    ,结算佣金率 varchar(200) null
    ,预估预售整单佣金 varchar(200) null
    ,"付款服务费支出(元)" varchar(200) null
    ,付款服务费率 varchar(200) null
    ,"结算服务费支出(元)" varchar(200) null
    ,结算服务费率 varchar(200) null
    ,"单件商品付款支出费用(元)" varchar(200) null
    ,"付款支出费用(元)" varchar(200) null
    ,"结算支出费用(元)" varchar(200) null
    ,预估预售整单佣金率 varchar(200) null
    ,进店人数 varchar(200) null
    ,"平均优惠券面额(元)" varchar(200) null
    ,付款人数 varchar(200) null
    ,统计日期 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."淘宝联盟_维权退款订单" 
(
    淘宝订单编号 varchar(200) null
    ,淘宝子订单编号 varchar(200) null
    ,商品名称 varchar(200) null
    ,维权退款金额 varchar(200) null
    ,应退回服务费 varchar(200) null
    ,应退回佣金 varchar(200) null
    ,维权状态 varchar(200) null
    ,订单结算时间 varchar(200) null
    ,维权创建时间 varchar(200) null
    ,维权完成时间 varchar(200) null
    ,备注 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_v任务_v任务分析_图文任务" 
(
    商品id varchar(200) null
    ,商品名称 varchar(200) null
    ,商品商品点击次数 varchar(200) null
    ,商品商品点击人数 varchar(200) null
    ,商品种草成交金额 varchar(200) null
    ,商品种草成交人数 varchar(200) null
    ,商品种草成交笔数 varchar(200) null
    ,商品种草成交件数 varchar(200) null
    ,作者类型 varchar(200) null
    ,发布时间 varchar(200) null
    ,任务id varchar(200) null
    ,内容id varchar(200) null
    ,作者id varchar(200) null
    ,标题 varchar(200) null
    ,作者 varchar(200) null
    ,内容类型 varchar(200) null
    ,内容商品点击次数 varchar(200) null
    ,内容商品点击人数 varchar(200) null
    ,内容种草成交金额 varchar(200) null
    ,内容种草成交人数 varchar(200) null
    ,内容种草成交笔数 varchar(200) null
    ,内容种草成交件数 varchar(200) null
    ,下单角色 varchar(200) null
    ,服务方 varchar(200) null
    ,拍下时间 varchar(200) null
    ,任务金额 varchar(200) null
    ,交付时间 varchar(200) null
    ,任务名称 varchar(200) null
    ,统计日期 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_内容专题_订阅分析_单条效果" 
(
    内容标题 varchar(200) null
    ,内容id varchar(200) null
    ,作者 varchar(200) null
    ,发布时间 varchar(200) null
    ,内容曝光人数 varchar(200) null
    ,内容曝光次数 varchar(200) null
    ,内容查看人数 varchar(200) null
    ,内容查看次数 varchar(200) null
    ,商品点击人数 varchar(200) null
    ,商品点击次数 varchar(200) null
    ,引导进店人数 varchar(200) null
    ,引导进店次数 varchar(200) null
    ,引导加购人数 varchar(200) null
    ,引导加购件数 varchar(200) null
    ,订阅直接引导成交金额 varchar(200) null
    ,直接引导成交人数 varchar(200) null
    ,直接引导成交笔数 varchar(200) null
    ,订阅玩法名称 varchar(200) null
    ,统计日期 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_内容专题_订阅分析_整体概况" 
(
    内容曝光次数 varchar(200) null
    ,订阅直接引导成交金额 varchar(200) null
    ,引导加购件数 varchar(200) null
    ,券核销引导成交金额 varchar(200) null
    ,引导加购人数 varchar(200) null
    ,内容查看人数 varchar(200) null
    ,内容曝光人数 varchar(200) null
    ,内容查看次数 varchar(200) null
    ,订阅引导总成交金额 varchar(200) null
    ,直接引导成交人数 varchar(200) null
    ,引导进店次数 varchar(200) null
    ,直接引导成交笔数 varchar(200) null
    ,商品点击人数 varchar(200) null
    ,引导进店人数 varchar(200) null
    ,商品点击次数 varchar(200) null
    ,引导直播种草成交金额 varchar(200) null
    ,日期 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_内容专题_逛逛_单条效果_其他内容" 
(
    统计日期 varchar(200) null
    ,内容标题 varchar(200) null
    ,内容id varchar(200) null
    ,作者 varchar(200) null
    ,发布时间 varchar(200) null
    ,内容曝光人数 varchar(200) null
    ,内容曝光次数 varchar(200) null
    ,内容查看人数 varchar(200) null
    ,内容查看次数 varchar(200) null
    ,"人均停留时长(秒)" varchar(200) null
    ,内容互动人数 varchar(200) null
    ,内容互动人次数 varchar(200) null
    ,内容点赞人数 varchar(200) null
    ,内容点赞次数 varchar(200) null
    ,内容评论人数 varchar(200) null
    ,内容评论次数 varchar(200) null
    ,内容分享人数 varchar(200) null
    ,内容分享次数 varchar(200) null
    ,内容收藏人数 varchar(200) null
    ,内容收藏次数 varchar(200) null
    ,商品点击人数 varchar(200) null
    ,商品点击次数 varchar(200) null
    ,逛逛种草成交金额 varchar(200) null
    ,逛逛种草成交人数 varchar(200) null
    ,逛逛种草成交笔数 varchar(200) null
    ,逛逛种草成交件数 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_内容专题_逛逛_商品分析" 
(
    内容标题 varchar(200) null
    ,内容id varchar(200) null
    ,商品名称 varchar(200) null
    ,商品id varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_商品客群" 
(
    日期 varchar(200) null
    ,属性 varchar(200) null
    ,占比 varchar(200) null
    ,访客数 varchar(200) null
    ,商品id varchar(200) null
    ,标签 varchar(200) null
    ,角色 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_商品流量_流量来源_pc端" 
(
    流量来源 varchar(200) null
    ,访客数 varchar(200) null
    ,浏览量 varchar(200) null
    ,支付金额 varchar(200) null
    ,浏览量占比 varchar(200) null
    ,店内跳转人数 varchar(200) null
    ,跳出本店人数 varchar(200) null
    ,收藏人数 varchar(200) null
    ,加购人数 varchar(200) null
    ,下单买家数 varchar(200) null
    ,下单转化率 varchar(200) null
    ,支付件数 varchar(200) null
    ,支付买家数 varchar(200) null
    ,支付转化率 varchar(200) null
    ,直接支付买家数 varchar(200) null
    ,"收藏商品-支付买家数" varchar(200) null
    ,粉丝支付买家数 varchar(200) null
    ,"加购商品-支付买家数" varchar(200) null
    ,访问来源 varchar(200) null
    ,统计日期 varchar(200) null
    ,商品id varchar(200) null
    ,来源明细 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_商品流量_流量来源_无线端" 
(
    流量来源 varchar(200) null
    ,来源明细 varchar(200) null
    ,访客数 varchar(200) null
    ,浏览量 varchar(200) null
    ,支付金额 varchar(200) null
    ,浏览量占比 varchar(200) null
    ,店内跳转人数 varchar(200) null
    ,跳出本店人数 varchar(200) null
    ,收藏人数 varchar(200) null
    ,加购人数 varchar(200) null
    ,下单买家数 varchar(200) null
    ,下单转化率 varchar(200) null
    ,支付件数 varchar(200) null
    ,支付买家数 varchar(200) null
    ,支付转化率 varchar(200) null
    ,直接支付买家数 varchar(200) null
    ,"收藏商品-支付买家数" varchar(200) null
    ,粉丝支付买家数 varchar(200) null
    ,"加购商品-支付买家数" varchar(200) null
    ,访问来源 varchar(200) null
    ,统计日期 varchar(200) null
    ,商品id varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_商品评价_评价" 
(
    评论 varchar(2000) null
    ,昵称 varchar(200) null
    ,追评内容 varchar(2000) null
    ,追评时间 varchar(200) null
    ,评价时间 varchar(200) null
    ,商品id varchar(200) null
    ,颜色分类 varchar(200) null
    ,评价分类 varchar(200) null
	,skuid varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_商品销售_类目" 
(
    父级类目id varchar(200) null
    ,叶子类目id varchar(200) null
    ,叶子类目名称 varchar(200) null
    ,类目层级 varchar(200) null
    ,是否叶子类目 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_商品销售_销售分析sku" 
(
    加购件数 varchar(200) null
    ,支付件数 varchar(200) null
    ,支付买家数 varchar(200) null
    ,支付金额 varchar(200) null
    ,商品id varchar(200) null
    ,sku信息 varchar(200) null
	,skuid   varchar(200) null
    ,统计日期 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_商品销售_销售分析价格" 
(
    件单价 varchar(200) null
    ,商品id varchar(200) null
    ,sku信息 varchar(200) null
	,skuid   varchar(200) null
    ,统计日期 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_商品销售_销售分析商品1" 
(
    时间 varchar(200) null
    ,商品访客数 varchar(200) null
    ,商品浏览量 varchar(200) null
    ,商品平均停留时长 varchar(200) null
    ,商品详情页跳出率 varchar(200) null
    ,商品加购人数 varchar(200) null
    ,商品加购件数 varchar(200) null
    ,商品收藏人数 varchar(200) null
    ,下单买家数 varchar(200) null
    ,下单金额 varchar(200) null
    ,下单转换率 varchar(200) null
    ,支付买家数 varchar(200) null
    ,下单件数 varchar(200) null
    ,支付转换率 varchar(200) null
    ,支付金额 varchar(200) null
    ,支付件数 varchar(200) null
    ,聚划算支付金额 varchar(200) null
    ,支付新买家数 varchar(200) null
    ,支付老买家数 varchar(200) null
    ,老买家支付金额 varchar(200) null
    ,月累计支付金额 varchar(200) null
    ,月累计支付件数 varchar(200) null
    ,年累计支付金额 varchar(200) null
    ,访客平均价值 varchar(200) null
    ,成功退款金额 varchar(200) null
    ,商品id varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_商品销售_销售分析商品2" 
(
    商品名称 varchar(200) null
    ,商品url varchar(200) null
    ,商品图片url varchar(200) null
    ,货号 varchar(200) null
    ,月支付金额进度 varchar(200) null
    ,月支付件数进度 varchar(200) null
    ,近30天本店排名 varchar(200) null
    ,属性 varchar(200) null
    ,标签 varchar(200) null
    ,商品类目id varchar(200) null
    ,商品类目 varchar(200) null
    ,商品id varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_商品销售_销售分析属性" 
(
    支付买家数 varchar(200) null
    ,支付件数 varchar(200) null
    ,支付金额占比 varchar(200) null
    ,属性名称 varchar(200) null
    ,支付金额 varchar(200) null
    ,支付买家数占比 varchar(200) null
    ,商品id varchar(200) null
    ,属性类型 varchar(200) null
    ,统计日期 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_图文_单条效果_其他" 
(
    图文标题 varchar(200) null
    ,作者 varchar(200) null
    ,发布时间 varchar(200) null
    ,互动人数 varchar(200) null
    ,互动次数 varchar(200) null
    ,新增粉丝数 varchar(200) null
    ,内容引导访客数 varchar(200) null
    ,商品点击人数 varchar(200) null
    ,商品点击次数 varchar(200) null
    ,引导加购人数 varchar(200) null
    ,引导加构次数 varchar(200) null
    ,内容种草人数 varchar(200) null
    ,种草粉丝人数 varchar(200) null
    ,种草成交金额 varchar(200) null
    ,种草成交人数 varchar(200) null
    ,统计日期 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_图文_单条效果_达人带货视频" 
(
    图文标题 varchar(200) null
    ,作者 varchar(200) null
    ,发布时间 varchar(200) null
    ,互动人数 varchar(200) null
    ,互动次数 varchar(200) null
    ,新增粉丝数 varchar(200) null
    ,内容引导访客数 varchar(200) null
    ,商品点击人数 varchar(200) null
    ,商品点击次数 varchar(200) null
    ,引导加购人数 varchar(200) null
    ,引导加构次数 varchar(200) null
    ,内容种草人数 varchar(200) null
    ,种草粉丝人数 varchar(200) null
    ,种草成交金额 varchar(200) null
    ,种草成交人数 varchar(200) null
    ,统计日期 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_图文_商品分析" 
(
    发布时间 varchar(200) null
    ,创作者 varchar(200) null
    ,图文id varchar(200) null
    ,图文名称 varchar(200) null
    ,内容商品点击人数 varchar(200) null
    ,内容种草成交人数 varchar(200) null
    ,内容引导加购人数 varchar(200) null
    ,内容引导加购件数 varchar(200) null
    ,内容种草成交金额 varchar(200) null
    ,内容商品点击次数 varchar(200) null
    ,商品总引导加购人数 varchar(200) null
    ,商品id varchar(200) null
    ,商品总引导加购件数 varchar(200) null
    ,商品总种草成交金额 varchar(200) null
    ,商品名称 varchar(200) null
    ,商品总点击次数 varchar(200) null
    ,商品总点击人数 varchar(200) null
    ,商品总种草成交人数 varchar(200) null
    ,统计日期 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_短视频_全屏视频_单条效果_其他" 
(
    视频id varchar(200) null
    ,视频标题 varchar(200) null
    ,作者 varchar(200) null
    ,发布时间 varchar(200) null
    ,活动标签 varchar(200) null
    ,商品点击人数 varchar(200) null
    ,商品点击次数 varchar(200) null
    ,引导加购人数 varchar(200) null
    ,引导加购件数 varchar(200) null
    ,点击种草成交金额 varchar(200) null
    ,点击种草成交人数 varchar(200) null
    ,点击种草成交笔数 varchar(200) null
    ,点击种草成交件数 varchar(200) null
    ,新客种草成交人数 varchar(200) null
    ,老客种草成交人数 varchar(200) null
    ,新客种草成交金额 varchar(200) null
    ,老客种草成交金额 varchar(200) null
    ,作品标签 varchar(200) null
    ,内容引导访客数 varchar(200) null
    ,种草成交金额 varchar(200) null
    ,种草成交人数 varchar(200) null
    ,统计日期 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_短视频_全屏视频_单条效果_商家自制视频" 
(
    视频id varchar(200) null
    ,视频标题 varchar(200) null
    ,作者 varchar(200) null
    ,发布时间 varchar(200) null
    ,活动标签 varchar(200) null
    ,曝光人数 varchar(200) null
    ,曝光次数 varchar(200) null
    ,曝光点击率 varchar(200) null
    ,播放次数 varchar(200) null
    ,播放人数 varchar(200) null
    ,有效播放次数 varchar(200) null
    ,有效播放人数 varchar(200) null
    ,完整播放次数 varchar(200) null
    ,完整播放人数 varchar(200) null
    ,完播率 varchar(200) null
    ,"人均观看时长(秒)" varchar(200) null
    ,视频互动人数 varchar(200) null
    ,视频互动次数 varchar(200) null
    ,新增粉丝数 varchar(200) null
    ,商品点击人数 varchar(200) null
    ,商品点击次数 varchar(200) null
    ,商品点击率 varchar(200) null
    ,引导加购人数 varchar(200) null
    ,引导加购件数 varchar(200) null
    ,点击种草成交金额 varchar(200) null
    ,点击种草成交人数 varchar(200) null
    ,点击种草成交笔数 varchar(200) null
    ,点击种草成交件数 varchar(200) null
    ,新客播放人数 varchar(200) null
    ,老客播放人数 varchar(200) null
    ,新客种草成交人数 varchar(200) null
    ,老客种草成交人数 varchar(200) null
    ,新客种草成交金额 varchar(200) null
    ,老客种草成交金额 varchar(200) null
    ,直播间曝光人数 varchar(200) null
    ,直播间曝光次数 varchar(200) null
    ,引导进直播间人数 varchar(200) null
    ,引导进直播间次数 varchar(200) null
    ,引导直播种草成交金额 varchar(200) null
    ,引导直播种草成交人数 varchar(200) null
    ,引导直播种草成交笔数 varchar(200) null
    ,作品标签 varchar(200) null
    ,内容引导访客数 varchar(200) null
    ,内容种草人数 varchar(200) null
    ,种草粉丝人数 varchar(200) null
    ,种草成交金额 varchar(200) null
    ,种草成交人数 varchar(200) null
    ,统计日期 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_短视频_全屏视频_单条效果_达人带货视频" 
(
    视频id varchar(200) null
    ,视频标题 varchar(200) null
    ,作者 varchar(200) null
    ,发布时间 varchar(200) null
    ,活动标签 varchar(200) null
    ,商品点击人数 varchar(200) null
    ,商品点击次数 varchar(200) null
    ,引导加购人数 varchar(200) null
    ,引导加购件数 varchar(200) null
    ,点击种草成交金额 varchar(200) null
    ,点击种草成交人数 varchar(200) null
    ,点击种草成交笔数 varchar(200) null
    ,点击种草成交件数 varchar(200) null
    ,新客种草成交人数 varchar(200) null
    ,老客种草成交人数 varchar(200) null
    ,新客种草成交金额 varchar(200) null
    ,老客种草成交金额 varchar(200) null
    ,作品标签 varchar(200) null
    ,内容引导访客数 varchar(200) null
    ,种草成交金额 varchar(200) null
    ,种草成交人数 varchar(200) null
    ,统计日期 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_短视频_全屏视频_商品分析" 
(
    视频收藏次数 varchar(200) null
    ,视频点击次数 varchar(200) null
    ,视频id varchar(200) null
    ,视频点击种草成交金额 varchar(200) null
    ,视频名称 varchar(200) null
    ,视频种草成交人数 varchar(200) null
    ,视频点击成交件数 varchar(200) null
    ,视频点金种草成交人数 varchar(200) null
    ,视频加购件数 varchar(200) null
    ,视频种草成交金额 varchar(200) null
    ,视频点击种草成交笔数 varchar(200) null
    ,视频粉丝点击次数 varchar(200) null
    ,统计日期 varchar(200) null
    ,商品收藏次数 varchar(200) null
    ,商品名称 varchar(200) null
    ,商品点击次数 varchar(200) null
    ,商品点击种草成交金额 varchar(200) null
    ,商品种草成交人数 varchar(200) null
    ,商品点击成交件数 varchar(200) null
    ,商品点金种草成交人数 varchar(200) null
    ,商品加购件数 varchar(200) null
    ,商品种草成交金额 varchar(200) null
    ,商品点击种草成交笔数 varchar(200) null
    ,商品粉丝点击次数 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_短视频_全屏视频_整体概况_其他" 
(
    统计日期 varchar(200) null
    ,商品点击人数 varchar(200) null
    ,商品点击次数 varchar(200) null
    ,引导加购人数 varchar(200) null
    ,引导加购件数 varchar(200) null
    ,点击种草成交金额 varchar(200) null
    ,点击种草成交人数 varchar(200) null
    ,点击种草成交笔数 varchar(200) null
    ,点击种草成交件数 varchar(200) null
    ,新客种草成交人数 varchar(200) null
    ,老客种草成交人数 varchar(200) null
    ,新客种草成交金额 varchar(200) null
    ,老客种草成交金额 varchar(200) null
    ,被播放视频数 varchar(200) null
    ,覆盖商品数 varchar(200) null
    ,内容引导访客数 varchar(200) null
    ,占比全店访客数 varchar(200) null
    ,种草成交金额 varchar(200) null
    ,占比全店支付金额 varchar(200) null
    ,种草成交人数 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_短视频_全屏视频_整体概况_商家自制视频" 
(
    统计日期 varchar(200) null
    ,曝光人数 varchar(200) null
    ,曝光次数 varchar(200) null
    ,曝光点击率 varchar(200) null
    ,播放人数 varchar(200) null
    ,播放次数 varchar(200) null
    ,"有效播放人数(短视频访客数)" varchar(200) null
    ,有效播放次数 varchar(200) null
    ,完整播放人数 varchar(200) null
    ,完整播放次数 varchar(200) null
    ,完播率 varchar(200) null
    ,"人均观看时长(秒)" varchar(200) null
    ,视频互动人数 varchar(200) null
    ,视频互动次数 varchar(200) null
    ,新增粉丝数 varchar(200) null
    ,商品点击人数 varchar(200) null
    ,商品点击次数 varchar(200) null
    ,商品点击率 varchar(200) null
    ,引导加购人数 varchar(200) null
    ,引导加购件数 varchar(200) null
    ,点击种草成交金额 varchar(200) null
    ,点击种草成交人数 varchar(200) null
    ,点击种草成交笔数 varchar(200) null
    ,点击种草成交件数 varchar(200) null
    ,新客播放人数 varchar(200) null
    ,老客播放人数 varchar(200) null
    ,新客种草成交人数 varchar(200) null
    ,老客种草成交人数 varchar(200) null
    ,新客种草成交金额 varchar(200) null
    ,老客种草成交金额 varchar(200) null
    ,直播间曝光人数 varchar(200) null
    ,直播间曝光次数 varchar(200) null
    ,引导进直播间人数 varchar(200) null
    ,引导进直播间次数 varchar(200) null
    ,引导直播种草成交金额 varchar(200) null
    ,引导直播种草成交人数 varchar(200) null
    ,引导直播种草成交笔数 varchar(200) null
    ,被播放视频数 varchar(200) null
    ,猜你喜欢审核通过视频数 varchar(200) null
    ,猜你喜欢视频发布数 varchar(200) null
    ,猜你喜欢视频覆盖公域商品数 varchar(200) null
    ,猜你喜欢公域商品视频覆盖率 varchar(200) null
    ,内容引导访客数 varchar(200) null
    ,占比全店访客数 varchar(200) null
    ,内容种草人数 varchar(200) null
    ,种草粉丝人数 varchar(200) null
    ,种草成交金额 varchar(200) null
    ,占比全店支付金额 varchar(200) null
    ,种草成交人数 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_短视频_全屏视频_整体概况_达人带货视频" 
(
    统计日期 varchar(200) null
    ,商品点击人数 varchar(200) null
    ,商品点击次数 varchar(200) null
    ,引导加购人数 varchar(200) null
    ,引导加购件数 varchar(200) null
    ,点击种草成交金额 varchar(200) null
    ,点击种草成交人数 varchar(200) null
    ,点击种草成交笔数 varchar(200) null
    ,点击种草成交件数 varchar(200) null
    ,新客种草成交人数 varchar(200) null
    ,老客种草成交人数 varchar(200) null
    ,新客种草成交金额 varchar(200) null
    ,老客种草成交金额 varchar(200) null
    ,被播放视频数 varchar(200) null
    ,覆盖商品数 varchar(200) null
    ,内容引导访客数 varchar(200) null
    ,占比全店访客数 varchar(200) null
    ,种草成交金额 varchar(200) null
    ,占比全店支付金额 varchar(200) null
    ,种草成交人数 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_短视频_微详情视频_单条效果" 
(
    统计日期 varchar(200) null
    ,视频id varchar(200) null
    ,视频标题 varchar(200) null
    ,作者 varchar(200) null
    ,作品标签 varchar(200) null
    ,发布时间 varchar(200) null
    ,播放人数 varchar(200) null
    ,播放次数 varchar(200) null
    ,有效播放人数 varchar(200) null
    ,有效播放次数 varchar(200) null
    ,有效播放率 varchar(200) null
    ,"人均观看时长(秒)" varchar(200) null
    ,卡片点击率 varchar(200) null
    ,视频覆盖商品数 varchar(200) null
    ,商品点击人数 varchar(200) null
    ,商品点击次数 varchar(200) null
    ,内容引导访客数 varchar(200) null
    ,内容种草人数 varchar(200) null
    ,种草粉丝人数 varchar(200) null
    ,种草成交金额 varchar(200) null
    ,种草成交人数 varchar(200) null
    ,完整播放人数 varchar(200) null
    ,完整播放次数 varchar(200) null
    ,完播率 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_短视频_微详情视频_商品分析" 
(
    视频播放次数 varchar(200) null
    ,"视频人均观看时长(秒)" varchar(200) null
    ,视频播放人数 varchar(200) null
    ,视频种草成交人数 varchar(200) null
    ,视频完整播放次数 varchar(200) null
    ,视频有效播放次数 varchar(200) null
    ,视频有效播放人数 varchar(200) null
    ,视频完播率 varchar(200) null
    ,视频有效播放率 varchar(200) null
    ,视频商品点击人数 varchar(200) null
    ,视频商品点击次数 varchar(200) null
    ,视频卡片点击率 varchar(200) null
    ,视频完整播放人数 varchar(200) null
    ,视频种草成交金额 varchar(200) null
    ,视频标题 varchar(200) null
    ,创作者名称 varchar(200) null
    ,创作者角色 varchar(200) null
    ,创作者id varchar(200) null
    ,视频id varchar(200) null
    ,统计日期 varchar(200) null
    ,商品id varchar(200) null
    ,商品名称 varchar(200) null
    ,商品播放人数 varchar(200) null
    ,商品播放次数 varchar(200) null
    ,商品有效播放人数 varchar(200) null
    ,商品有效播放次数 varchar(200) null
    ,商品有效播放率 varchar(200) null
    ,"商品人均观看时长(秒)" varchar(200) null
    ,商品卡片点击率 varchar(200) null
    ,商品商品点击人数 varchar(200) null
    ,商品商品点击次数 varchar(200) null
    ,商品种草成交金额 varchar(200) null
    ,商品种草成交人数 varchar(200) null
    ,商品完整播放人数 varchar(200) null
    ,商品完整播放次数 varchar(200) null
    ,商品完播率 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_短视频_微详情视频_整体概况" 
(
    统计日期 varchar(200) null
    ,播放人数 varchar(200) null
    ,播放次数 varchar(200) null
    ,有效播放人数 varchar(200) null
    ,有效播放次数 varchar(200) null
    ,有效播放率 varchar(200) null
    ,"人均观看时长(秒)" varchar(200) null
    ,卡片点击率 varchar(200) null
    ,商品点击人数 varchar(200) null
    ,商品点击次数 varchar(200) null
    ,被播放视频数 varchar(200) null
    ,微详情渠道发布视频数 varchar(200) null
    ,审核通过视频数 varchar(200) null
    ,审核通过率 varchar(200) null
    ,视频覆盖商品数 varchar(200) null
    ,内容种草人数 varchar(200) null
    ,种草粉丝人数 varchar(200) null
    ,种草成交金额 varchar(200) null
    ,种草成交人数 varchar(200) null
    ,"（种草成交金额）占比全店支付金额" varchar(200) null
    ,内容引导访客数 varchar(200) null
    ,"（内容访客数）占比全店访客数" varchar(200) null
    ,完整播放人数 varchar(200) null
    ,完整播放次数 varchar(200) null
    ,完播率 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_营销玩法和工具_单品宝" 
(
    支付件数 varchar(200) null
    ,支付买家数 varchar(200) null
    ,支付金额 varchar(200) null
    ,客单价 varchar(200) null
    ,统计日期 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_营销玩法和工具_店铺券" 
(
    领取张数 varchar(200) null
    ,使用张数 varchar(200) null
    ,支付金额 varchar(200) null
    ,支付买家数 varchar(200) null
    ,支付件数 varchar(200) null
    ,客单价 varchar(200) null
    ,统计日期 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_营销玩法和工具_商品券" 
(
    领取张数 varchar(200) null
    ,使用张数 varchar(200) null
    ,支付金额 varchar(200) null
    ,支付买家数 varchar(200) null
    ,支付件数 varchar(200) null
    ,客单价 varchar(200) null
    ,统计日期 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_营销玩法和工具_店铺宝" 
(
    支付父定单 varchar(200) null
    ,支付买家数 varchar(200) null
    ,支付金额 varchar(200) null
    ,客单价 varchar(200) null
    ,人均支付件数 varchar(200) null
    ,支付件数 varchar(200) null
    ,统计日期 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."直通车_账户报表_单元" 
(
    日期 varchar(200) null
    ,计划id varchar(200) null
    ,计划名称 varchar(200) null
    ,推广类型 varchar(200) null
    ,单元id varchar(200) null
    ,单元名称 varchar(200) null
    ,商品id varchar(200) null
    ,展现量 varchar(200) null
    ,点击量 varchar(200) null
    ,花费 varchar(200) null
    ,点击率 varchar(200) null
    ,平均点击花费 varchar(200) null
    ,千次展现花费 varchar(200) null
    ,总收藏数 varchar(200) null
    ,宝贝收藏数 varchar(200) null
    ,店铺收藏数 varchar(200) null
    ,总购物车数 varchar(200) null
    ,直接购物车数 varchar(200) null
    ,间接购物车数 varchar(200) null
    ,加购成本 varchar(200) null
    ,宝贝收藏成本 varchar(200) null
    ,宝贝收藏转化率 varchar(200) null
    ,加购率 varchar(200) null
    ,总预售成交金额 varchar(200) null
    ,总预售成交笔数 varchar(200) null
    ,直接预售成交金额 varchar(200) null
    ,直接预售成交笔数 varchar(200) null
    ,间接预售成交金额 varchar(200) null
    ,间接预售成交笔数 varchar(200) null
    ,总成交金额 varchar(200) null
    ,直接成交金额 varchar(200) null
    ,间接成交金额 varchar(200) null
    ,总成交笔数 varchar(200) null
    ,直接成交笔数 varchar(200) null
    ,间接成交笔数 varchar(200) null
    ,投入产出比 varchar(200) null
    ,点击转化率 varchar(200) null
    ,直接点击转化率 varchar(200) null
    ,购物金充值笔数 varchar(200) null
    ,购物金充值金额 varchar(200) null
    ,自然流量曝光 varchar(200) null
    ,自然流量转化金额 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."直通车_账户报表_计划" 
(
    日期 varchar(200) null
    ,计划id varchar(200) null
    ,计划名称 varchar(200) null
    ,推广类型 varchar(200) null
    ,展现量 varchar(200) null
    ,点击量 varchar(200) null
    ,花费 varchar(200) null
    ,点击率 varchar(200) null
    ,平均点击花费 varchar(200) null
    ,千次展现花费 varchar(200) null
    ,总收藏数 varchar(200) null
    ,宝贝收藏数 varchar(200) null
    ,店铺收藏数 varchar(200) null
    ,总购物车数 varchar(200) null
    ,直接购物车数 varchar(200) null
    ,间接购物车数 varchar(200) null
    ,加购成本 varchar(200) null
    ,宝贝收藏成本 varchar(200) null
    ,宝贝收藏转化率 varchar(200) null
    ,加购率 varchar(200) null
    ,总预售成交金额 varchar(200) null
    ,总预售成交笔数 varchar(200) null
    ,直接预售成交金额 varchar(200) null
    ,直接预售成交笔数 varchar(200) null
    ,间接预售成交金额 varchar(200) null
    ,间接预售成交笔数 varchar(200) null
    ,总成交金额 varchar(200) null
    ,直接成交金额 varchar(200) null
    ,间接成交金额 varchar(200) null
    ,总成交笔数 varchar(200) null
    ,直接成交笔数 varchar(200) null
    ,间接成交笔数 varchar(200) null
    ,投入产出比 varchar(200) null
    ,点击转化率 varchar(200) null
    ,直接点击转化率 varchar(200) null
    ,购物金充值笔数 varchar(200) null
    ,购物金充值金额 varchar(200) null
    ,自然流量曝光 varchar(200) null
    ,自然流量转化金额 varchar(200) null
    ,特权订金金额 varchar(200) null
    ,尾款金额 varchar(200) null
    ,一口价金额 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."直通车_账户报表_账户" 
(
    日期 varchar(200) null
    ,展现量 varchar(200) null
    ,点击量 varchar(200) null
    ,花费 varchar(200) null
    ,点击率 varchar(200) null
    ,平均点击花费 varchar(200) null
    ,千次展现花费 varchar(200) null
    ,总收藏数 varchar(200) null
    ,宝贝收藏数 varchar(200) null
    ,店铺收藏数 varchar(200) null
    ,总购物车数 varchar(200) null
    ,直接购物车数 varchar(200) null
    ,间接购物车数 varchar(200) null
    ,加购成本 varchar(200) null
    ,宝贝收藏成本 varchar(200) null
    ,宝贝收藏转化率 varchar(200) null
    ,加购率 varchar(200) null
    ,总预售成交金额 varchar(200) null
    ,总预售成交笔数 varchar(200) null
    ,直接预售成交金额 varchar(200) null
    ,直接预售成交笔数 varchar(200) null
    ,间接预售成交金额 varchar(200) null
    ,间接预售成交笔数 varchar(200) null
    ,总成交金额 varchar(200) null
    ,直接成交金额 varchar(200) null
    ,间接成交金额 varchar(200) null
    ,总成交笔数 varchar(200) null
    ,直接成交笔数 varchar(200) null
    ,间接成交笔数 varchar(200) null
    ,投入产出比 varchar(200) null
    ,点击转化率 varchar(200) null
    ,直接点击转化率 varchar(200) null
    ,购物金充值笔数 varchar(200) null
    ,购物金充值金额 varchar(200) null
    ,自然流量曝光 varchar(200) null
    ,自然流量转化金额 varchar(200) null
    ,特权订金金额 varchar(200) null
    ,尾款金额 varchar(200) null
    ,一口价金额 varchar(200) null
    ,新成交用户数 varchar(200) null
    ,新成交用户占比 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."营销活动-行业活动" 
(
    活动名称 varchar(200) null
    ,预热时间 varchar(200) null
    ,活动时间 varchar(200) null
    ,商品名称 varchar(200) null
    ,商品id varchar(200) null
    ,营销id varchar(200) null
    ,商品状态 varchar(200) null
    ,一口价 varchar(200) null
    ,活动价格 varchar(200) null
    ,库存 varchar(200) null
    ,报名时间 varchar(200) null
    ,活动状态 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."顺手买一件_活动信息" 
(
    活动id varchar(200) null
    ,活动名称 varchar(200) null
    ,活动开始时间 varchar(200) null
    ,活动结束时间 varchar(200) null
    ,创建时间 varchar(200) null
    ,活动门槛 varchar(200) null
    ,活动状态 varchar(200) null
    ,换购销量 varchar(200) null
    ,换购剩余库存 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."顺手买一件_活动效果" 
(
    活动点击pv varchar(200) null
    ,优惠时间 varchar(200) null
    ,活动曝光pv varchar(200) null
    ,支付金额 varchar(200) null
    ,活动id varchar(200) null
    ,活动名称 varchar(200) null
    ,支付件数 varchar(200) null
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."万相台_创意_上新快" 
(
    日期 varchar(255)
    ,效果模型 varchar(255)
    ,统计口径 varchar(255)
    ,转化周期 varchar(255)
    ,创意id varchar(255)
    ,商品id varchar(255)
    ,店铺url varchar(255)
    ,计划id varchar(255)
    ,计划名称 varchar(255)
    ,消耗 varchar(255)
    ,曝光量 varchar(255)
    ,点击量 varchar(255)
    ,点击率 varchar(255)
    ,点击成本 varchar(255)
    ,直接加购数 varchar(255)
    ,间接加购数 varchar(255)
    ,总加购量 varchar(255)
    ,收藏量 varchar(255)
    ,加购收藏成本 varchar(255)
    ,预售成交笔数 varchar(255)
    ,预售成交金额 varchar(255)
    ,直接成交笔数 varchar(255)
    ,间接成交笔数 varchar(255)
    ,总成交笔数 varchar(255)
    ,直接成交金额 varchar(255)
    ,间接成交金额 varchar(255)
    ,总成交金额 varchar(255)
    ,成交转化率 varchar(255)
    ,roi varchar(255)
    ,新客成交金额贡献占比 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."万相台_创意_会员快" 
(
    日期 varchar(255)
    ,效果模型 varchar(255)
    ,统计口径 varchar(255)
    ,转化周期 varchar(255)
    ,创意id varchar(255)
    ,商品id varchar(255)
    ,店铺url varchar(255)
    ,计划id varchar(255)
    ,计划名称 varchar(255)
    ,消耗 varchar(255)
    ,曝光量 varchar(255)
    ,点击量 varchar(255)
    ,点击率 varchar(255)
    ,点击成本 varchar(255)
    ,总加购量 varchar(255)
    ,收藏量 varchar(255)
    ,加购收藏成本 varchar(255)
    ,预售成交笔数 varchar(255)
    ,预售成交金额 varchar(255)
    ,总成交笔数 varchar(255)
    ,总成交金额 varchar(255)
    ,入会量 varchar(255)
    ,会员成交金额 varchar(255)
    ,成交转化率 varchar(255)
    ,roi varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."万相台_创意_拉新快" 
(
    日期 varchar(255)
    ,效果模型 varchar(255)
    ,统计口径 varchar(255)
    ,转化周期 varchar(255)
    ,创意id varchar(255)
    ,商品id varchar(255)
    ,店铺url varchar(255)
    ,计划id varchar(255)
    ,计划名称 varchar(255)
    ,消耗 varchar(255)
    ,曝光量 varchar(255)
    ,点击量 varchar(255)
    ,点击率 varchar(255)
    ,点击成本 varchar(255)
    ,总加购量 varchar(255)
    ,收藏量 varchar(255)
    ,加购收藏成本 varchar(255)
    ,预售成交笔数 varchar(255)
    ,预售成交金额 varchar(255)
    ,总成交笔数 varchar(255)
    ,总成交金额 varchar(255)
    ,成交转化率 varchar(255)
    ,新客roi varchar(255)
    ,派样量 varchar(255)
    ,派样金额 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."万相台_创意_活动加速" 
(
    日期 varchar(255)
    ,效果模型 varchar(255)
    ,统计口径 varchar(255)
    ,转化周期 varchar(255)
    ,创意id varchar(255)
    ,商品id varchar(255)
    ,店铺url varchar(255)
    ,计划id varchar(255)
    ,计划名称 varchar(255)
    ,消耗 varchar(255)
    ,曝光量 varchar(255)
    ,点击量 varchar(255)
    ,点击率 varchar(255)
    ,点击成本 varchar(255)
    ,直接加购量 varchar(255)
    ,间接加购量 varchar(255)
    ,总加购量 varchar(255)
    ,收藏量 varchar(255)
    ,加购收藏成本 varchar(255)
    ,预售成交笔数 varchar(255)
    ,预售成交金额 varchar(255)
    ,直接成交笔数 varchar(255)
    ,间接成交笔数 varchar(255)
    ,总成交笔数 varchar(255)
    ,直接成交金额 varchar(255)
    ,间接成交金额 varchar(255)
    ,总成交金额 varchar(255)
    ,成交转化率 varchar(255)
    ,roi varchar(255)
    ,千次展现成本 varchar(255)
    ,加购收藏量 varchar(255)
    ,店铺订阅量 varchar(255)
    ,店铺订阅成本 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."万相台_创意_测款快" 
(
    日期 varchar(255)
    ,效果模型 varchar(255)
    ,统计口径 varchar(255)
    ,转化周期 varchar(255)
    ,创意id varchar(255)
    ,商品id varchar(255)
    ,店铺url varchar(255)
    ,计划id varchar(255)
    ,计划名称 varchar(255)
    ,消耗 varchar(255)
    ,曝光量 varchar(255)
    ,点击量 varchar(255)
    ,点击率 varchar(255)
    ,点击成本 varchar(255)
    ,直接加购数 varchar(255)
    ,间接加购数 varchar(255)
    ,总加购量 varchar(255)
    ,收藏量 varchar(255)
    ,加购收藏成本 varchar(255)
    ,预售成交笔数 varchar(255)
    ,预售成交金额 varchar(255)
    ,直接成交笔数 varchar(255)
    ,间接成交笔数 varchar(255)
    ,总成交笔数 varchar(255)
    ,直接成交金额 varchar(255)
    ,间接成交金额 varchar(255)
    ,总成交金额 varchar(255)
    ,成交转化率 varchar(255)
    ,roi varchar(255)
    ,新客成交金额贡献占比 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."万相台_创意_爆发收割" 
(
    日期 varchar(255)
    ,效果模型 varchar(255)
    ,统计口径 varchar(255)
    ,转化周期 varchar(255)
    ,创意id varchar(255)
    ,商品id varchar(255)
    ,店铺url varchar(255)
    ,计划id varchar(255)
    ,计划名称 varchar(255)
    ,消耗 varchar(255)
    ,曝光量 varchar(255)
    ,点击量 varchar(255)
    ,点击率 varchar(255)
    ,点击成本 varchar(255)
    ,直接加购量 varchar(255)
    ,间接加购量 varchar(255)
    ,总加购量 varchar(255)
    ,收藏量 varchar(255)
    ,加购收藏成本 varchar(255)
    ,预售成交笔数 varchar(255)
    ,预售成交金额 varchar(255)
    ,直接成交笔数 varchar(255)
    ,间接成交笔数 varchar(255)
    ,总成交笔数 varchar(255)
    ,直接成交金额 varchar(255)
    ,间接成交金额 varchar(255)
    ,总成交金额 varchar(255)
    ,成交转化率 varchar(255)
    ,roi varchar(255)
    ,千次展现成本 varchar(255)
    ,加购收藏量 varchar(255)
    ,店铺订阅量 varchar(255)
    ,店铺订阅成本 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_品牌专区报表_创意" 
(
    日期 varchar(255)
    ,计划名称 varchar(255)
    ,单元名称 varchar(255)
    ,创意名称 varchar(255)
    ,展现量 varchar(255)
    ,自然流量增量曝光 varchar(255)
    ,触达访客数 varchar(255)
    ,回搜展现量 varchar(255)
    ,回搜触达访客数 varchar(255)
    ,点击量 varchar(255)
    ,跳转点击量 varchar(255)
    ,点击访客数 varchar(255)
    ,点击率 varchar(255)
    ,进店访客数 varchar(255)
    ,互动点击量 varchar(255)
    ,跳转点击率 varchar(255)
    ,回搜点击量 varchar(255)
    ,回搜点击访客数 varchar(255)
    ,店铺收藏数 varchar(255)
    ,宝贝收藏数 varchar(255)
    ,宝贝加购数 varchar(255)
    ,宝贝浏览数 varchar(255)
    ,店铺浏览数 varchar(255)
    ,访问时长 varchar(255)
    ,店铺收藏访客数 varchar(255)
    ,宝贝收藏访客数 varchar(255)
    ,宝贝加购访客数 varchar(255)
    ,宝贝浏览访客数 varchar(255)
    ,成交笔数 varchar(255)
    ,成交金额 varchar(255)
    ,转化率 varchar(255)
    ,预售成交笔数 varchar(255)
    ,成交访客数 varchar(255)
    ,自然流量增量成交 varchar(255)
    ,预售成交金额 varchar(255)
    ,搜索进店率 varchar(255)
    ,进店行动率 varchar(255)
    ,行动成交率 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_品牌专区报表_品牌流量包" 
(
    日期 varchar(255)
    ,计划名称 varchar(255)
    ,单元名称 varchar(255)
    ,词包名称 varchar(255)
    ,展现量 varchar(255)
    ,自然流量增量曝光 varchar(255)
    ,触达访客数 varchar(255)
    ,回搜展现量 varchar(255)
    ,回搜触达访客数 varchar(255)
    ,点击量 varchar(255)
    ,跳转点击量 varchar(255)
    ,点击访客数 varchar(255)
    ,点击率 varchar(255)
    ,进店访客数 varchar(255)
    ,互动点击量 varchar(255)
    ,跳转点击率 varchar(255)
    ,回搜点击量 varchar(255)
    ,回搜点击访客数 varchar(255)
    ,店铺收藏数 varchar(255)
    ,宝贝收藏数 varchar(255)
    ,宝贝加购数 varchar(255)
    ,宝贝浏览数 varchar(255)
    ,店铺浏览数 varchar(255)
    ,访问时长 varchar(255)
    ,店铺收藏访客数 varchar(255)
    ,宝贝收藏访客数 varchar(255)
    ,宝贝加购访客数 varchar(255)
    ,宝贝浏览访客数 varchar(255)
    ,成交笔数 varchar(255)
    ,成交金额 varchar(255)
    ,转化率 varchar(255)
    ,预售成交笔数 varchar(255)
    ,成交访客数 varchar(255)
    ,自然流量增量成交 varchar(255)
    ,预售成交金额 varchar(255)
    ,搜索进店率 varchar(255)
    ,进店行动率 varchar(255)
    ,行动成交率 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_品牌专区报表_定向人群" 
(
    日期 varchar(255)
    ,计划名称 varchar(255)
    ,单元名称 varchar(255)
    ,创意名称 varchar(255)
    ,定向人群名称 varchar(255)
    ,展现量 varchar(255)
    ,自然流量增量曝光 varchar(255)
    ,触达访客数 varchar(255)
    ,回搜展现量 varchar(255)
    ,回搜触达访客数 varchar(255)
    ,点击量 varchar(255)
    ,跳转点击量 varchar(255)
    ,点击访客数 varchar(255)
    ,点击率 varchar(255)
    ,进店访客数 varchar(255)
    ,互动点击量 varchar(255)
    ,跳转点击率 varchar(255)
    ,回搜点击量 varchar(255)
    ,回搜点击访客数 varchar(255)
    ,店铺收藏数 varchar(255)
    ,宝贝收藏数 varchar(255)
    ,宝贝加购数 varchar(255)
    ,宝贝浏览数 varchar(255)
    ,店铺浏览数 varchar(255)
    ,访问时长 varchar(255)
    ,店铺收藏访客数 varchar(255)
    ,宝贝收藏访客数 varchar(255)
    ,宝贝加购访客数 varchar(255)
    ,宝贝浏览访客数 varchar(255)
    ,成交笔数 varchar(255)
    ,成交金额 varchar(255)
    ,转化率 varchar(255)
    ,预售成交笔数 varchar(255)
    ,成交访客数 varchar(255)
    ,自然流量增量成交 varchar(255)
    ,预售成交金额 varchar(255)
    ,搜索进店率 varchar(255)
    ,进店行动率 varchar(255)
    ,行动成交率 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_品牌专区报表_推广单元" 
(
    日期 varchar(255)
    ,计划名称 varchar(255)
    ,单元名称 varchar(255)
    ,展现量 varchar(255)
    ,自然流量增量曝光 varchar(255)
    ,触达访客数 varchar(255)
    ,回搜展现量 varchar(255)
    ,回搜触达访客数 varchar(255)
    ,点击量 varchar(255)
    ,跳转点击量 varchar(255)
    ,点击访客数 varchar(255)
    ,点击率 varchar(255)
    ,进店访客数 varchar(255)
    ,互动点击量 varchar(255)
    ,跳转点击率 varchar(255)
    ,回搜点击量 varchar(255)
    ,回搜点击访客数 varchar(255)
    ,店铺收藏数 varchar(255)
    ,宝贝收藏数 varchar(255)
    ,宝贝加购数 varchar(255)
    ,宝贝浏览数 varchar(255)
    ,店铺浏览数 varchar(255)
    ,访问时长 varchar(255)
    ,店铺收藏访客数 varchar(255)
    ,宝贝收藏访客数 varchar(255)
    ,宝贝加购访客数 varchar(255)
    ,宝贝浏览访客数 varchar(255)
    ,成交笔数 varchar(255)
    ,成交金额 varchar(255)
    ,转化率 varchar(255)
    ,预售成交笔数 varchar(255)
    ,成交访客数 varchar(255)
    ,自然流量增量成交 varchar(255)
    ,预售成交金额 varchar(255)
    ,搜索进店率 varchar(255)
    ,进店行动率 varchar(255)
    ,行动成交率 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_品牌专区报表_推广计划" 
(
    日期 varchar(255)
    ,计划名称 varchar(255)
    ,展现量 varchar(255)
    ,自然流量增量曝光 varchar(255)
    ,触达访客数 varchar(255)
    ,回搜展现量 varchar(255)
    ,回搜触达访客数 varchar(255)
    ,点击量 varchar(255)
    ,跳转点击量 varchar(255)
    ,点击访客数 varchar(255)
    ,点击率 varchar(255)
    ,进店访客数 varchar(255)
    ,互动点击量 varchar(255)
    ,跳转点击率 varchar(255)
    ,回搜点击量 varchar(255)
    ,回搜点击访客数 varchar(255)
    ,店铺收藏数 varchar(255)
    ,宝贝收藏数 varchar(255)
    ,宝贝加购数 varchar(255)
    ,宝贝浏览数 varchar(255)
    ,店铺浏览数 varchar(255)
    ,访问时长 varchar(255)
    ,店铺收藏访客数 varchar(255)
    ,宝贝收藏访客数 varchar(255)
    ,宝贝加购访客数 varchar(255)
    ,宝贝浏览访客数 varchar(255)
    ,成交笔数 varchar(255)
    ,成交金额 varchar(255)
    ,转化率 varchar(255)
    ,预售成交笔数 varchar(255)
    ,成交访客数 varchar(255)
    ,自然流量增量成交 varchar(255)
    ,预售成交金额 varchar(255)
    ,搜索进店率 varchar(255)
    ,进店行动率 varchar(255)
    ,行动成交率 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_品牌特秀_单元" 
(
    单元名称 varchar(255)
    ,定向条件 varchar(255)
    ,所属计划 varchar(255)
    ,曝光量 varchar(255)
    ,点击量 varchar(255)
    ,互动点击量 varchar(255)
    ,曝光点击率 varchar(255)
    ,点击人数 varchar(255)
    ,互动曝光量 varchar(255)
    ,互动完成量 varchar(255)
    ,曝光人数 varchar(255)
    ,UV点击率 varchar(255)
    ,互动点击率 varchar(255)
    ,互动完成率 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_品牌特秀_普通创意" 
(
    创意名称 varchar(255)
    ,创意曝光量 varchar(255)
    ,创意点击量 varchar(255)
    ,创意互动点击量 varchar(255)
    ,创意曝光点击率 varchar(255)
    ,创意点击人数 varchar(255)
    ,创意互动曝光量 varchar(255)
    ,创意互动完成量 varchar(255)
    ,创意曝光人数 varchar(255)
    ,创意uv点击率 varchar(255)
    ,资源名称 varchar(255)
    ,所属计划 varchar(255)
    ,所属单元 varchar(255)
    ,资源曝光量 varchar(255)
    ,资源点击量 varchar(255)
    ,资源互动点击量 varchar(255)
    ,资源曝光点击率 varchar(255)
    ,资源点击人数 varchar(255)
    ,资源互动曝光量 varchar(255)
    ,资源互动完成量 varchar(255)
    ,资源曝光人数 varchar(255)
    ,资源UV点击率 varchar(255)
    ,资源互动点击率 varchar(255)
    ,资源互动完成率 varchar(255)
    ,创意互动点击率 varchar(255)
    ,创意互动完成率 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_品牌首推报表_创意" 
(
    日期 varchar(255)
    ,计划名称 varchar(255)
    ,单元名称 varchar(255)
    ,创意名称 varchar(255)
    ,展现量 varchar(255)
    ,自然流量增量曝光 varchar(255)
    ,消耗 varchar(255)
    ,触达访客数 varchar(255)
    ,千次展现成本 varchar(255)
    ,点击单价 varchar(255)
    ,跳转点击单价 varchar(255)
    ,点击量 varchar(255)
    ,跳转点击量 varchar(255)
    ,点击访客数 varchar(255)
    ,互动点击量 varchar(255)
    ,点击率 varchar(255)
    ,跳转点击率 varchar(255)
    ,进店访客数 varchar(255)
    ,店铺收藏数 varchar(255)
    ,宝贝收藏数 varchar(255)
    ,宝贝加购数 varchar(255)
    ,行动访客数 varchar(255)
    ,店铺收藏访客数 varchar(255)
    ,宝贝收藏访客数 varchar(255)
    ,宝贝加购访客数 varchar(255)
    ,成交笔数 varchar(255)
    ,成交金额 varchar(255)
    ,回报率 varchar(255)
    ,转化率 varchar(255)
    ,自然流量增量成交 varchar(255)
    ,预售成交笔数 varchar(255)
    ,预售成交金额 varchar(255)
    ,成交访客数 varchar(255)
    ,搜索进店率 varchar(255)
    ,进店行动率 varchar(255)
    ,行动成交率 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_品牌首推报表_品牌流量包" 
(
    日期 varchar(255)
    ,计划名称 varchar(255)
    ,单元名称 varchar(255)
    ,词包名称 varchar(255)
    ,展现量 varchar(255)
    ,自然流量增量曝光 varchar(255)
    ,消耗 varchar(255)
    ,触达访客数 varchar(255)
    ,千次展现成本 varchar(255)
    ,点击单价 varchar(255)
    ,跳转点击单价 varchar(255)
    ,点击量 varchar(255)
    ,跳转点击量 varchar(255)
    ,点击访客数 varchar(255)
    ,互动点击量 varchar(255)
    ,点击率 varchar(255)
    ,跳转点击率 varchar(255)
    ,进店访客数 varchar(255)
    ,店铺收藏数 varchar(255)
    ,宝贝收藏数 varchar(255)
    ,宝贝加购数 varchar(255)
    ,行动访客数 varchar(255)
    ,店铺收藏访客数 varchar(255)
    ,宝贝收藏访客数 varchar(255)
    ,宝贝加购访客数 varchar(255)
    ,成交笔数 varchar(255)
    ,成交金额 varchar(255)
    ,回报率 varchar(255)
    ,转化率 varchar(255)
    ,自然流量增量成交 varchar(255)
    ,预售成交笔数 varchar(255)
    ,预售成交金额 varchar(255)
    ,成交访客数 varchar(255)
    ,搜索进店率 varchar(255)
    ,进店行动率 varchar(255)
    ,行动成交率 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_品牌首推报表_定向人群" 
(
    日期 varchar(255)
    ,计划名称 varchar(255)
    ,单元名称 varchar(255)
    ,创意名称 varchar(255)
    ,定向人群名称 varchar(255)
    ,展现量 varchar(255)
    ,自然流量增量曝光 varchar(255)
    ,消耗 varchar(255)
    ,触达访客数 varchar(255)
    ,千次展现成本 varchar(255)
    ,点击单价 varchar(255)
    ,跳转点击单价 varchar(255)
    ,点击量 varchar(255)
    ,跳转点击量 varchar(255)
    ,点击访客数 varchar(255)
    ,互动点击量 varchar(255)
    ,点击率 varchar(255)
    ,跳转点击率 varchar(255)
    ,进店访客数 varchar(255)
    ,店铺收藏数 varchar(255)
    ,宝贝收藏数 varchar(255)
    ,宝贝加购数 varchar(255)
    ,行动访客数 varchar(255)
    ,店铺收藏访客数 varchar(255)
    ,宝贝收藏访客数 varchar(255)
    ,宝贝加购访客数 varchar(255)
    ,成交笔数 varchar(255)
    ,成交金额 varchar(255)
    ,回报率 varchar(255)
    ,转化率 varchar(255)
    ,自然流量增量成交 varchar(255)
    ,预售成交笔数 varchar(255)
    ,预售成交金额 varchar(255)
    ,成交访客数 varchar(255)
    ,搜索进店率 varchar(255)
    ,进店行动率 varchar(255)
    ,行动成交率 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_品牌首推报表_推广单元" 
(
    日期 varchar(255)
    ,计划名称 varchar(255)
    ,单元名称 varchar(255)
    ,展现量 varchar(255)
    ,自然流量增量曝光 varchar(255)
    ,消耗 varchar(255)
    ,触达访客数 varchar(255)
    ,千次展现成本 varchar(255)
    ,点击单价 varchar(255)
    ,跳转点击单价 varchar(255)
    ,点击量 varchar(255)
    ,跳转点击量 varchar(255)
    ,点击访客数 varchar(255)
    ,互动点击量 varchar(255)
    ,点击率 varchar(255)
    ,跳转点击率 varchar(255)
    ,进店访客数 varchar(255)
    ,店铺收藏数 varchar(255)
    ,宝贝收藏数 varchar(255)
    ,宝贝加购数 varchar(255)
    ,行动访客数 varchar(255)
    ,店铺收藏访客数 varchar(255)
    ,宝贝收藏访客数 varchar(255)
    ,宝贝加购访客数 varchar(255)
    ,成交笔数 varchar(255)
    ,成交金额 varchar(255)
    ,回报率 varchar(255)
    ,转化率 varchar(255)
    ,自然流量增量成交 varchar(255)
    ,预售成交笔数 varchar(255)
    ,预售成交金额 varchar(255)
    ,成交访客数 varchar(255)
    ,搜索进店率 varchar(255)
    ,进店行动率 varchar(255)
    ,行动成交率 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_品牌首推报表_推广计划" 
(
    日期 varchar(255)
    ,计划名称 varchar(255)
    ,展现量 varchar(255)
    ,自然流量增量曝光 varchar(255)
    ,消耗 varchar(255)
    ,触达访客数 varchar(255)
    ,千次展现成本 varchar(255)
    ,点击单价 varchar(255)
    ,跳转点击单价 varchar(255)
    ,点击量 varchar(255)
    ,跳转点击量 varchar(255)
    ,点击访客数 varchar(255)
    ,互动点击量 varchar(255)
    ,点击率 varchar(255)
    ,跳转点击率 varchar(255)
    ,进店访客数 varchar(255)
    ,店铺收藏数 varchar(255)
    ,宝贝收藏数 varchar(255)
    ,宝贝加购数 varchar(255)
    ,行动访客数 varchar(255)
    ,店铺收藏访客数 varchar(255)
    ,宝贝收藏访客数 varchar(255)
    ,宝贝加购访客数 varchar(255)
    ,成交笔数 varchar(255)
    ,成交金额 varchar(255)
    ,回报率 varchar(255)
    ,转化率 varchar(255)
    ,自然流量增量成交 varchar(255)
    ,预售成交笔数 varchar(255)
    ,预售成交金额 varchar(255)
    ,成交访客数 varchar(255)
    ,搜索进店率 varchar(255)
    ,进店行动率 varchar(255)
    ,行动成交率 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_品牌首推报表_账户" 
(
    日期 varchar(255)
    ,搜索量 varchar(255)
    ,搜索访客数 varchar(255)
    ,访客触达率 varchar(255)
    ,展现量 varchar(255)
    ,自然流量增量曝光 varchar(255)
    ,消耗 varchar(255)
    ,触达访客数 varchar(255)
    ,千次展现成本 varchar(255)
    ,点击单价 varchar(255)
    ,跳转点击单价 varchar(255)
    ,点击量 varchar(255)
    ,跳转点击量 varchar(255)
    ,点击访客数 varchar(255)
    ,互动点击量 varchar(255)
    ,点击率 varchar(255)
    ,跳转点击率 varchar(255)
    ,进店访客数 varchar(255)
    ,店铺收藏数 varchar(255)
    ,宝贝收藏数 varchar(255)
    ,宝贝加购数 varchar(255)
    ,行动访客数 varchar(255)
    ,店铺收藏访客数 varchar(255)
    ,宝贝收藏访客数 varchar(255)
    ,宝贝加购访客数 varchar(255)
    ,成交笔数 varchar(255)
    ,成交金额 varchar(255)
    ,回报率 varchar(255)
    ,转化率 varchar(255)
    ,自然流量增量成交 varchar(255)
    ,预售成交笔数 varchar(255)
    ,预售成交金额 varchar(255)
    ,成交访客数 varchar(255)
    ,搜索进店率 varchar(255)
    ,进店行动率 varchar(255)
    ,行动成交率 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_品销puSH报表_创意" 
(
    日期 varchar(255)
    ,所属计划 varchar(255)
    ,推广单元名称 varchar(255)
    ,创意名称 varchar(255)
    ,展现量 varchar(255)
    ,点击量 varchar(255)
    ,"点击率(%)" varchar(255)
    ,"花费(元)" varchar(255)
    ,"千次展现成本（元）" varchar(255)
    ,"点击单价（元）" varchar(255)
    ,"跳转点击单价（元）" varchar(255)
    ,展示成交笔数 varchar(255)
    ,"展示成交金额（元）" varchar(255)
    ,宝贝收藏数 varchar(255)
    ,店铺收藏数 varchar(255)
    ,宝贝加购数 varchar(255)
    ,展示回报率 varchar(255)
    ,点击回报率 varchar(255)
    ,点击成交笔数 varchar(255)
    ,"点击成交金额（元）" varchar(255)
    ,触达访客数 varchar(255)
    ,点击访客数 varchar(255)
    ,"展现转化率(%)" varchar(255)
    ,"点击转化率(%)" varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_品销puSH报表_推广单元" 
(
    日期 varchar(255)
    ,所属计划 varchar(255)
    ,推广单元名称 varchar(255)
    ,展现量 varchar(255)
    ,点击量 varchar(255)
    ,"点击率(%)" varchar(255)
    ,"花费(元)" varchar(255)
    ,"千次展现成本（元）" varchar(255)
    ,"点击单价（元）" varchar(255)
    ,"跳转点击单价（元）" varchar(255)
    ,展示成交笔数 varchar(255)
    ,"展示成交金额（元）" varchar(255)
    ,宝贝收藏数 varchar(255)
    ,店铺收藏数 varchar(255)
    ,宝贝加购数 varchar(255)
    ,展示回报率 varchar(255)
    ,点击回报率 varchar(255)
    ,点击成交笔数 varchar(255)
    ,"点击成交金额（元）" varchar(255)
    ,触达访客数 varchar(255)
    ,点击访客数 varchar(255)
    ,"展现转化率(%)" varchar(255)
    ,"点击转化率(%)" varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_品销puSH报表_推广计划" 
(
    日期 varchar(255)
    ,计划名称 varchar(255)
    ,展现量 varchar(255)
    ,点击量 varchar(255)
    ,"点击率(%)" varchar(255)
    ,"花费(元)" varchar(255)
    ,"千次展现成本（元）" varchar(255)
    ,"点击单价（元）" varchar(255)
    ,"跳转点击单价（元）" varchar(255)
    ,展示成交笔数 varchar(255)
    ,"展示成交金额（元）" varchar(255)
    ,宝贝收藏数 varchar(255)
    ,店铺收藏数 varchar(255)
    ,宝贝加购数 varchar(255)
    ,展示回报率 varchar(255)
    ,点击回报率 varchar(255)
    ,点击成交笔数 varchar(255)
    ,"点击成交金额（元）" varchar(255)
    ,触达访客数 varchar(255)
    ,点击访客数 varchar(255)
    ,"展现转化率(%)" varchar(255)
    ,"点击转化率(%)" varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."品销宝_品销puSH报表_账户" 
(
    日期 varchar(255)
    ,展现量 varchar(255)
    ,点击量 varchar(255)
    ,"点击率(%)" varchar(255)
    ,"花费(元)" varchar(255)
    ,"千次展现成本（元）" varchar(255)
    ,"点击单价（元）" varchar(255)
    ,"跳转点击单价（元）" varchar(255)
    ,展示成交笔数 varchar(255)
    ,"展示成交金额（元）" varchar(255)
    ,宝贝收藏数 varchar(255)
    ,店铺收藏数 varchar(255)
    ,宝贝加购数 varchar(255)
    ,展示回报率 varchar(255)
    ,点击回报率 varchar(255)
    ,点击成交笔数 varchar(255)
    ,"点击成交金额（元）" varchar(255)
    ,触达访客数 varchar(255)
    ,点击访客数 varchar(255)
    ,"展现转化率(%)" varchar(255)
    ,"点击转化率(%)" varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."商家中心_商品信息" 
(
    商品id varchar(255)
    ,发布时间 varchar(255)
    ,商品名称 varchar(255)
    ,品牌 varchar(255)
    ,系列 varchar(255)
    ,规格 varchar(255)
    ,商品编码 varchar(255)
    ,条形码 varchar(255)
    ,商品价格 varchar(255)
    ,颜色分类 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."商家中心_订单表商品" 
(
    子订单编号 varchar(255)
    ,主订单编号 varchar(255)
    ,标题 varchar(255)
    ,价格 varchar(255)
    ,购买数量 varchar(255)
    ,外部系统编号 varchar(255)
    ,商品属性 varchar(255)
    ,套餐信息 varchar(255)
    ,备注 varchar(255)
    ,订单状态 varchar(255)
    ,商家编码 varchar(255)
    ,支付单号 varchar(255)
    ,买家应付货款 varchar(255)
    ,买家实际支付金额 varchar(255)
    ,退款状态 varchar(255)
    ,退款金额 varchar(255)
    ,订单创建时间 varchar(255)
    ,订单付款时间 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."商家中心_订单表明细" 
(
    订单编号 text
    ,买家会员名 varchar(255)
    ,买家会员昵称 text
    ,支付单号 text
    ,支付详情 text
    ,买家应付货款 text
    ,买家应付邮费 text
    ,买家支付积分 varchar(255)
    ,总金额 text
    ,返点积分 text
    ,买家实际支付金额 varchar(255)
    ,买家实际支付积分 varchar(255)
    ,订单状态 varchar(255)
    ,买家留言 varchar(255)
    ,收货人姓名 varchar(255)
    ,收货地址 varchar(255)
    ,运送方式 varchar(255)
    ,联系电话 varchar(255)
    ,联系手机 varchar(255)
    ,虚拟号过期时间 varchar(255)
    ,订单创建时间 varchar(255)
    ,订单付款时间 varchar(255)
    ,宝贝标题 varchar(255)
    ,宝贝种类 varchar(255)
    ,物流单号 varchar(255)
    ,物流公司 varchar(255)
    ,订单备注 varchar(255)
    ,宝贝总数量 varchar(255)
    ,店铺Id varchar(255)
    ,店铺名称 varchar(255)
    ,订单关闭原因 varchar(255)
    ,卖家服务费 varchar(255)
    ,买家服务费 varchar(255)
    ,发票抬头 varchar(255)
    ,是否手机订单 varchar(255)
    ,分阶段订单信息 varchar(255)
    ,特权订金订单id varchar(255)
    ,是否上传合同照片 varchar(255)
    ,是否上传小票 varchar(255)
    ,是否代付 varchar(255)
    ,定金排名 varchar(255)
    ,修改后的sku varchar(255)
    ,修改后的收货地址 varchar(255)
    ,异常信息 varchar(255)
    ,天猫卡券抵扣 varchar(255)
    ,集分宝抵扣 varchar(255)
    ,是否是O2O交易 varchar(255)
    ,新零售交易类型 varchar(255)
    ,新零售导购门店名称 varchar(255)
    ,新零售导购门店id varchar(255)
    ,新零售发货门店名称 varchar(255)
    ,新零售发货门店id varchar(255)
    ,退款金额 varchar(255)
    ,预约门店 varchar(255)
    ,确认收货时间 varchar(255)
    ,打款商家金额 varchar(255)
    ,含应开票给个人的个人红包 varchar(255)
    ,是否信用购 varchar(255)
    ,体验期结束时间 varchar(255)
    ,前N有礼 varchar(255)
    ,配送类型 varchar(255)
    ,直播返现状态 varchar(255)
    ,返现金额 varchar(255)
    ,延迟发货自动赔付金额 varchar(255)
    ,新零售成交门店昵称 varchar(255)
    ,新零售成交门店id varchar(255)
    ,新零售成交经销商id varchar(255)
    ,预售订单 varchar(255)
    ,发货时间 varchar(255)
    ,商家备忘 varchar(255)
    ,主订单编号 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_商品销售_商品列表" 
(
    统计日期 varchar(255)
    ,商品id varchar(255)
    ,商品名称 varchar(255)
    ,货号 varchar(255)
    ,商品状态 varchar(255)
    ,商品标签 varchar(255)
    ,商品访客数 varchar(255)
    ,商品浏览量 varchar(255)
    ,平均停留时长 varchar(255)
    ,商品详情页跳出率 varchar(255)
    ,商品收藏人数 varchar(255)
    ,商品加购件数 varchar(255)
    ,商品加购人数 varchar(255)
    ,下单买家数 varchar(255)
    ,下单件数 varchar(255)
    ,下单金额 varchar(255)
    ,下单转化率 varchar(255)
    ,支付买家数 varchar(255)
    ,支付件数 varchar(255)
    ,支付金额 varchar(255)
    ,商品支付转化率 varchar(255)
    ,支付新买家数 varchar(255)
    ,支付老买家数 varchar(255)
    ,老买家支付金额 varchar(255)
    ,聚划算支付金额 varchar(255)
    ,访客平均价值 varchar(255)
    ,成功退款金额 varchar(255)
    ,竞争力评分 varchar(255)
    ,年累计支付金额 varchar(255)
    ,月累计支付金额 varchar(255)
    ,月累计支付件数 varchar(255)
    ,搜索引导支付转化率 varchar(255)
    ,搜索引导访客数 varchar(255)
    ,搜索引导支付买家数 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_活动分析_商品" 
(
    商品id varchar(255)
    ,商品标题 varchar(255)
    ,是否活动商品 varchar(255)
    ,商品链接 varchar(255)
    ,支付金额 varchar(255)
    ,加购件数 varchar(255)
    ,访客数 varchar(255)
    ,"访问-支付转化率" varchar(255)
    ,加购人数 varchar(255)
    ,收藏件数 varchar(255)
    ,收藏人数 varchar(255)
    ,"访问-加购转化率" varchar(255)
    ,"访问-收藏转化率" varchar(255)
    ,支付买家数 varchar(255)
    ,支付件数 varchar(255)
    ,实际折扣率 varchar(255)
    ,爆发系数 varchar(255)
    ,活动id varchar(255)
    ,统计日期 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_活动分析_引流_pc端" 
(
    流量来源 varchar(255)
    ,预热访客数 varchar(255)
    ,预热加购人数 varchar(255)
    ,预热收藏人数 varchar(255)
    ,访客数 varchar(255)
    ,支付买家数 varchar(255)
    ,支付金额 varchar(255)
    ,支付转化率 varchar(255)
    ,加购转化率 varchar(255)
    ,收藏转化率 varchar(255)
    ,活动id varchar(255)
    ,统计日期 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_活动分析_引流_无线端" 
(
    流量来源 varchar(255)
    ,预热访客数 varchar(255)
    ,预热加购人数 varchar(255)
    ,预热收藏人数 varchar(255)
    ,访客数 varchar(255)
    ,支付买家数 varchar(255)
    ,支付金额 varchar(255)
    ,支付转化率 varchar(255)
    ,加购转化率 varchar(255)
    ,收藏转化率 varchar(255)
    ,活动id varchar(255)
    ,统计日期 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_活动分析_正式期kpi" 
(
    支付金额 varchar(255)
    ,支付件数 varchar(255)
    ,客单价 varchar(255)
    ,支付买家数 varchar(255)
    ,访客数 varchar(255)
    ,支付转化率 varchar(255)
    ,活动商品支付金额 varchar(255)
    ,活动商品支付件数 varchar(255)
    ,活动商品支付买家数 varchar(255)
    ,活动商品访客数 varchar(255)
    ,活动申请退款金额 varchar(255)
    ,统计日期 varchar(255)
    ,活动id varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_活动分析_汇总" 
(
    活动名称 varchar(255)
    ,活动类型 varchar(255)
    ,预热期开始时间 varchar(255)
    ,预热期结束时间 varchar(255)
    ,活动期开始时间 varchar(255)
    ,活动期结束时间 varchar(255)
    ,累计支付金额 varchar(255)
    ,累计引流访客数 varchar(255)
    ,活动id varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_活动分析_预热期kpi" 
(
    预热加购人数 varchar(255)
    ,预热加购件数 varchar(255)
    ,预热收藏人数 varchar(255)
    ,预热收藏次数 varchar(255)
    ,预热访客数 varchar(255)
    ,加购转化率 varchar(255)
    ,收藏转化率 varchar(255)
    ,统计日期 varchar(255)
    ,活动id varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_直播_直播分商品效果" 
(
    商品标题 varchar(255)
    ,商品id varchar(255)
    ,直播标题 varchar(255)
    ,场次id varchar(255)
    ,开播时间 varchar(255)
    ,商品点击人数 varchar(255)
    ,商品点击次数 varchar(255)
    ,加购商品件数 varchar(255)
    ,加购人数 varchar(255)
    ,成交人数 varchar(255)
    ,成交件数 varchar(255)
    ,成交笔数 varchar(255)
    ,成交金额 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_直播_直播分场次" 
(
    直播标题 varchar(255)
    ,场次id varchar(255)
    ,开播时间 varchar(255)
    ,直播间观看次数 varchar(255)
    ,直播间观看人数 varchar(255)
    ,新增粉丝数 varchar(255)
    ,转粉率 varchar(255)
    ,商品点击人数 varchar(255)
    ,商品点击次数 varchar(255)
    ,商品点击率 varchar(255)
    ,商品加购人数 varchar(255)
    ,商品加购件数 varchar(255)
    ,成交转化率 varchar(255)
    ,成交人数 varchar(255)
    ,成交件数 varchar(255)
    ,成交笔数 varchar(255)
    ,成交金额 varchar(255)
    ,互助率 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_直播_直播间大盘" 
(
    开播场次 varchar(255)
    ,直播访问人数 varchar(255)
    ,直播间新增粉丝数 varchar(255)
    ,种草成交人数 varchar(255)
    ,种草成交金额 varchar(255)
    ,直播间浏览次数 varchar(255)
    ,粉丝回访率 varchar(255)
    ,商品点击人数 varchar(255)
    ,商品点击次数 varchar(255)
    ,当日确定收货金额 varchar(255)
    ,当日成功退款金额 varchar(255)
    ,开播时长 varchar(255)
    ,新零售种草成交人数 varchar(255)
    ,新零售种草成交金额 varchar(255)
    ,统计日期 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_直播_直播订单明细" 
(
    直播标题 varchar(255) NULL
    ,"开播时间" varchar(255) NULL
    ,"场次id" varchar(255) NULL
    ,"代播主播名称" varchar(255) NULL
    ,"会员名" varchar(255) NULL
    ,"直播间粉丝层级" varchar(255) NULL
    ,"商品id" varchar(255) NULL
    ,"商品标题" varchar(255) NULL
    ,"商品一级类目" varchar(255) NULL
    ,"父订单" varchar(255) NULL
    ,"子订单" varchar(255) NULL
    ,"下单时间" varchar(255) NULL
    ,"支付时间" varchar(255) NULL
    ,"支付金额" varchar(255) NULL
    ,"确认收货时间" varchar(255) NULL
    ,"确认收货金额" varchar(255) NULL
    ,"代播id" varchar(255) NULL
    ,"数据入库日期" varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_营销玩法和工具_新客折扣" 
(
    统计日期 varchar(255)
    ,商品名称 varchar(255)
    ,商品id varchar(255)
    ,新客拉新费率 varchar(255)
    ,商品新访客数 varchar(255)
    ,新客支付人数 varchar(255)
    ,新客支付人数占比 varchar(255)
    ,新客支付金额 varchar(255)
    ,新客支付金额占比 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_营销玩法和工具_购物金" 
(
    统计日期 varchar(255)
    ,商品名称 varchar(255)
    ,充值总金额 varchar(255)
    ,商品访客数 varchar(255)
    ,充值买家数 varchar(255)
    ,充值本金金额 varchar(255)
    ,充值转化率 varchar(255)
    ,人均充值金额 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."生参_v任务_v任务分析_直播任务" 
(
    商品id varchar(255)
    ,商品名称 varchar(255)
    ,商品商品点击次数 varchar(255)
    ,商品商品点击人数 varchar(255)
    ,商品种草成交金额 varchar(255)
    ,商品种草成交人数 varchar(255)
    ,商品种草成交笔数 varchar(255)
    ,商品种草成交件数 varchar(255)
    ,作者类型 varchar(255)
    ,任务id varchar(255)
    ,内容id varchar(255)
    ,主播id varchar(255)
    ,内容类型 varchar(255)
    ,主播 varchar(255)
    ,开播时间 varchar(255)
    ,内容名称 varchar(255)
    ,内容种草成交件数 varchar(255)
    ,内容商品点击次数 varchar(255)
    ,内容种草成交金额 varchar(255)
    ,内容种草成交人数 varchar(255)
    ,内容商品点击人数 varchar(255)
    ,内容种草成交笔数 varchar(255)
    ,交付时间 varchar(255)
    ,任务名称 varchar(255)
    ,任务金额 varchar(255)
    ,下单角色 varchar(255)
    ,服务方 varchar(255)
    ,拍下时间 varchar(255)
    ,统计日期 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."营销活动_88会员_现货" 
(
    活动名称 varchar(255)
    ,预热时间 varchar(255)
    ,活动时间 varchar(255)
    ,商品名称 varchar(255)
    ,商品id varchar(255)
    ,营销id varchar(255)
    ,商品状态 varchar(255)
    ,一口价 varchar(255)
    ,活动价格 varchar(255)
    ,库存 varchar(255)
    ,销量 varchar(255)
    ,报名时间 varchar(255)
    ,活动状态 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."营销活动_天猫行业" 
(
    活动名称 varchar(255)
    ,预热时间 varchar(255)
    ,活动时间 varchar(255)
    ,商品名称 varchar(255)
    ,商品id varchar(255)
    ,营销id varchar(255)
    ,商品状态 varchar(255)
    ,一口价 varchar(255)
    ,活动价格 varchar(255)
    ,库存 varchar(255)
    ,报名时间 varchar(255)
    ,活动状态 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."营销活动_官方大促_现货" 
(
    活动名称 varchar(255)
    ,活动时间 varchar(255)
    ,商品名称 varchar(255)
    ,商品id varchar(255)
    ,营销id varchar(255)
    ,商品状态 varchar(255)
    ,一口价 varchar(255)
    ,活动价格 varchar(255)
    ,报名时间 varchar(255)
    ,预计普惠成交价 varchar(255)
    ,商品价格力 varchar(255)
    ,活动状态 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."营销活动_官方直播间" 
(
    活动名称 varchar(255)
    ,预热时间 varchar(255)
    ,活动时间 varchar(255)
    ,商品名称 varchar(255)
    ,商品id varchar(255)
    ,营销id varchar(255)
    ,商品状态 varchar(255)
    ,一口价 varchar(255)
    ,活动价格 varchar(255)
    ,库存 varchar(255)
    ,报名时间 varchar(255)
    ,活动状态 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."营销活动_淘宝直播" 
(
    活动名称 varchar(255)
    ,预热时间 varchar(255)
    ,活动时间 varchar(255)
    ,商品名称 varchar(255)
    ,商品id varchar(255)
    ,营销id varchar(255)
    ,商品状态 varchar(255)
    ,一口价 varchar(255)
    ,活动价格 varchar(255)
    ,库存 varchar(255)
    ,报名时间 varchar(255)
    ,活动状态 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."营销活动_百亿补贴" 
(
    活动名称 varchar(255)
    ,预热时间l varchar(255)
    ,活动时间 varchar(255)
    ,商品名称 varchar(255)
    ,商品id varchar(255)
    ,营销id varchar(255)
    ,商品状态 varchar(255)
    ,一口价 varchar(255)
    ,活动价格 varchar(255)
    ,成本价 varchar(255)
    ,库存 varchar(255)
    ,销量 varchar(255)
    ,报名时间 varchar(255)
    ,预计普惠成交价 varchar(255)
    ,全网低价 varchar(255)
    ,商品价格力 varchar(255)
    ,活动状态 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."营销活动_聚划算_现货" 
(
    活动名称 varchar(255)
    ,预热时间 varchar(255)
    ,活动时间 varchar(255)
    ,商品名称 varchar(255)
    ,商品id varchar(255)
    ,营销id varchar(255)
    ,商品状态 varchar(255)
    ,一口价 varchar(255)
    ,活动价格 varchar(255)
    ,库存 varchar(255)
    ,销量 varchar(255)
    ,报名时间 varchar(255)
    ,预计普惠成交价 varchar(255)
    ,最低成交价 varchar(255)
    ,商品价格力 varchar(255)
    ,活动状态 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;

CREATE TABLE IF NOT EXISTS bi_master_test.public."营销活动_行业活动" 
(
    活动名称 varchar(255)
    ,预热时间 varchar(255)
    ,活动时间 varchar(255)
    ,商品名称 varchar(255)
    ,商品id varchar(255)
    ,营销id varchar(255)
    ,商品状态 varchar(255)
    ,一口价 varchar(255)
    ,活动价格 varchar(255)
    ,库存 varchar(255)
    ,报名时间 varchar(255)
    ,活动状态 varchar(255)
    ,数据入库日期 varchar(200) NULL
)
;
CREATE TABLE IF NOT EXISTS bi_master_test.public."商家中心_店铺信息" (
  "物流服务" varchar(255),
  "描述相符" varchar(255),
  "服务态度" varchar(255),
  "店铺id" varchar(255),
  "用户id" varchar(255),
  "用户昵称" varchar(255),
  "综合体验星级" varchar(255),
  "主营类目" varchar(255),
  "店铺类型" varchar(255),
  "数据入库日期" varchar(255)
)
;
CREATE TABLE IF NOT EXISTS bi_master_test.public."test" (
  "入库日期" varchar(255),
  "数量" varchar(255)
);

CREATE TABLE  IF NOT EXISTS  bi_master_test.public."触发器事件表"
(
    入库日期 varchar(100) null,
    更新表数量 int          null,
    创建时间 timestamp(10) NOT NULL DEFAULT CURRENT_TIMESTAMP
);
create or replace function public.usp_table_insert_new() 
RETURNS trigger AS 
$BODY$
	 
begin
     insert INTO bi_master_test.public."test" select '2022-03-17','1' ; 
     return NEW; 
end;

$BODY$
LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS  trigger_table_insert_new ON  bi_master_test.public."触发器事件表";
CREATE TRIGGER trigger_table_insert_new
after insert on bi_master_test.public."触发器事件表"
for each row execute procedure public.usp_table_insert_new();

`
