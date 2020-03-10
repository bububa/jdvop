package message

type Message interface {
	MsgType() uint
}

// 京东订单可能会被多次拆单； 例如：订单1 首先被拆成订单2、订单3；然后订单2有继续被拆成订单4、订单5；最终订单1的子单是订单3、订单4、订单5；每拆一次单我们都会发送一次拆单消息，但父订单号只会传递订单1（原始单），需要通过查询接口获取到最新所有子单，进行相关更新；
type OrderSplitMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		ParentOrderId uint64 `json:"pOrder,omitempty"`
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *OrderSplitMsg) MsgType() uint {
	return 1
}

// 此消息数据量大且频繁，考虑到客户侧不一定有技术实力对接此消息，此消息开通需要走线下邮件开通申请。
type SkuPriceChangeMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		SkuId uint64 `json:"pOrder,omitempty"`
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *SkuPriceChangeMsg) MsgType() uint {
	return 2
}

// 指的是在商品池里的商品在京东主站上发生了“上下架”
type SkuOnlineChangeMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		SkuId uint64 `json:"skuId,omitempty"`
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *SkuOnlineChangeMsg) MsgType() uint {
	return 4
}

// 该订单已妥投（买断模式代表外单已妥投或外单已拒收）。
type OrderShippingChangeMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		OrderId uint64 `json:"orderId,omitempty"`
		State   uint   `json:"state,omitempty"` // 1是妥投，2是拒收
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *OrderShippingChangeMsg) MsgType() uint {
	return 5
}

// 商品池内商品添加、删除消息
type ProductPoolProductChangeMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		SkuId   uint64 `json:"skuId,omitempty"`
		PageNum string `json:"page_num,omitempty"` // 商品池编号
		State   uint   `json:"state,omitempty"`    // 1添加，2删除
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *ProductPoolProductChangeMsg) MsgType() uint {
	return 6
}

// 订单取消消息
type OrderCancelMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		OrderId uint64 `json:"orderId,omitempty"`
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *OrderCancelMsg) MsgType() uint {
	return 10
}

// 发票开票进度消息
type InvoiceProcessMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		MarkId string `json:"markId,omitempty"`
		State  uint   `json:"state,omitempty"` // 状态：1：代表全部开票；2：代表驳回；3：部分开票；4：发票寄出
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *InvoiceProcessMsg) MsgType() uint {
	return 11
}

// 配送单生成成功消息
type OrderShipMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		OrderId uint64 `json:"orderId,omitempty"`
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *OrderShipMsg) MsgType() uint {
	return 12
}

// 换新订单生成成功消息
type AfterSaleNewOrderMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		OrderId   uint64 `json:"orderId,omitempty"`
		ServiceId uint64 `json:"afsServiceId,omitempty"`
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *AfterSaleNewOrderMsg) MsgType() uint {
	return 13
}

// 支付失败消息
type PaymentFailedMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		OrderId uint64 `json:"orderId,omitempty"`
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *PaymentFailedMsg) MsgType() uint {
	return 14
}

// 7天未支付取消消息/未确认取消消息
type OrderAutoCancelMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		OrderId    uint64 `json:"orderId,omitempty"`
		CancelType uint   `json:"cancelType,omitempty"` // 1: 7天未支付取消消息; 2: 未确认取消
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *OrderAutoCancelMsg) MsgType() uint {
	return 15
}

// 商品信息变更
// 包含：
// 商品名称，介绍，规格参数和商品图变更
type SkuChangeMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		SkuId uint64 `json:"skuId,omitempty"`
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *SkuChangeMsg) MsgType() uint {
	return 16
}

// 赠品促销变更消息
type PromotionChangeMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		SkuId uint64 `json:"skuId,omitempty"`
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *PromotionChangeMsg) MsgType() uint {
	return 17
}

// 订单等待确认收货消息
type OrderWaitConfirmMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		OrderId uint64 `json:"orderId,omitempty"`
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *OrderWaitConfirmMsg) MsgType() uint {
	return 18
}

// 订单配送退货消息
type OrderShippingReturnMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		OrderId uint64 `json:"orderId,omitempty"`
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *OrderShippingReturnMsg) MsgType() uint {
	return 23
}

// 新订单消息
type NewOrderMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		OrderId uint64 `json:"orderId,omitempty"`
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *NewOrderMsg) MsgType() uint {
	return 25
}

// 预定订单消息
type OrderReserveMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		OrderId uint64 `json:"orderId,omitempty"`
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *OrderReserveMsg) MsgType() uint {
	return 26
}

// 订单完成消息
type OrderFinishMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		OrderId uint64 `json:"orderId,omitempty"`
		State   uint   `json:"jdOrderState,omitempty"`
		Pin     string `json:"pin,omitempty"`
		Time    string `json:"completeTime,omitempty"`
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *OrderFinishMsg) MsgType() uint {
	return 31
}

// 商品池添加、删除消息
type ProductPoolChangeMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		PoolType string `json:"poolType,omitempty"` // p_skupool 用户的私有商品池；cate_pool 分类商品池; recommend 主推商品池；hot_sale 热销商品池；p_custom_skupool 用户的私有定制商品池
		PageNum  string `json:"page_num,omitempty"` // 商品池编号
		State    uint   `json:"state,omitempty"`    // 1添加，2删除
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *ProductPoolChangeMsg) MsgType() uint {
	return 28
}

// 折扣率变更消息
type DiscountRateChangeMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		OldRate  float64 `json:"oldRate,omitempty"`
		NewRate  float64 `json:"newRate,omitempty"`
		Category string  `json:"category,omitempty"`
		Type     uint    `json:"type,omitempty"`
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *DiscountRateChangeMsg) MsgType() uint {
	return 49
}

// 京东地址变更消息
type JdAddressChangeMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		AreaId      uint64 `json:"areaId,omitempty"`      // 京东地址编码
		AreaName    string `json:"areaName,omitempty"`    // 京东地址名称
		ParentId    uint64 `json:"parentId,omitempty"`    // 父京东ID编码
		AreaLevel   uint   `json:"areaLevel,omitempty"`   // 地址等级(行政级别：国家(1)、省(2)、市(3)、县(4)、镇(5))
		OperateType uint   `json:"operateType,omitempty"` // 操作类型(插入数据为1，更新时为2，删除时为3)}
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *JdAddressChangeMsg) MsgType() uint {
	return 50
}

// 商品税率变更消息（目前未涵盖全部商品）
// 客户侧取outputVAT作为采购税率
type SkuTaxRateChangeMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		Time     int64  `json:"timestamp,omitempty"`
		Features string `json:"features,omitempty"`
		SkuId    uint64 `json:"skuId,omitempty"`
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *SkuTaxRateChangeMsg) MsgType() uint {
	return 100
}

// 专票资质审核进度消息
// pins有可能多个用逗号隔开
type InvoiceAuditProcessMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		CompleteDate string `json:"completeDate,omitempty"`
		PushDate     string `json:"pushDate,omitempty"`
		Pins         string `json:"pins,omitempty"`
		Reason       string `json:"reason,omitempty"`
		Status       uint   `json:"status,omitempty"`
		UnitName     string `json:"unitName,omitempty"`
		SubmitDate   string `json:"submitDate,omitempty"`
		TaxpayerId   string `json:"texpayerId,omitempty"`
		VatId        uint64 `json:"vatId,omitempty"`
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *InvoiceAuditProcessMsg) MsgType() uint {
	return 102
}

// 混合支付微信个人支付成功消息
type PaymentSuccessMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		OrderId uint64 `json:"orderId,omitempty"`
		Pin     string `json:"pin,omitempty"`
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *PaymentSuccessMsg) MsgType() uint {
	return 108
}

// 退款消息
type RefundMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		OrderId       uint64 `json:"orderId,omitempty"`
		RefundType    uint   `json:"refundType,omitempty"`    // 1  售后退款; 2   整单退款; 5   外单退款; 6   移动转售退款; 11  订单多支付退款; 12  余额提现; 13  整单二次退款; 14  APP余额转小金库; 15  PC余额转小金库
		ServiceNumber uint64 `json:"serviceNumber,omitempty"` // 涉及售后退款 serviceNumber 有值
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *RefundMsg) MsgType() uint {
	return 109
}

// 新售后服务单状态变更
type AfterSaleServiceChangeMsg struct {
	Id     uint64 `json:"id,omitempty"` // 推送id
	Result *struct {
		OrderId   uint64 `json:"orderId,omitempty"`
		Pin       string `json:"pin,omitempty"`
		SkuId     uint64 `json:"skuId,omitempty"`
		ServiceId uint64 `json:"afsServiceId,omitempty"`
		State     uint   `json:"state,omitempty"`     // 1：创建；2：审核不通过；3：审核取消；4：审核通过; 5：用户确认; 6：完成
		IsOffline uint   `json:"isOffline,omitempty"` // 是否线上申请
	} `json:"result,omitempty"`
	Type uint   `json:"type,omitempty"`
	Time string `time,omitempty"`
}

func (this *AfterSaleServiceChangeMsg) MsgType() uint {
	return 110
}
