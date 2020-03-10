package order

type SkuNum struct {
	SkuId      uint64  `json:"skuId,omitempty"`
	Num        int     `json:"num,omitempty"`
	BNeedAnnex bool    `json:"bNeedAnnex,omitempty"` // 表示附件，必须传true，附件是必不可少的，当出现下单报附件无货只能等附件到货后再下单。
	BNeedGift  bool    `json:"bNeedGift,omitempty"`  // 表示赠品，true要赠品，false不要赠品;（传true是当订单主商品存在赠品活动且赠品还有库存时生成订单会有对应的赠品）
	Yanbao     *Yanbao `json:"yanbao,omitempty"`     // 延保这个字段可以直接不传
}

type Yanbao struct {
	SkuId uint64 `json:"skuId"`
}

type Freight struct {
	Freight             uint   `json:"freight,omitempty"`             // 总运费
	BaseFreight         uint   `json:"baseFreight,omitempty"`         // 基础运费
	RemoteRegionFreight uint   `json:"remoteRegionFreight,omitempty"` // 偏远地区加收运费
	RemoteSku           uint64 `json:"remoteSku,omitempty"`           // 需收取偏远运费的sku
}

type PromiseCalendarResult struct {
	SkuClassifyResult    *SkuClassifyResult    `json:"skuClassifyResult,omitempty"`    // 入参sku的物流的分类。（大家电、中小件）
	CalendarListResult   *CalendarListResult   `json:"calendarListResult,omitempty"`   // 中小件的预约日历结果。
	LaCalendarListResult *LaCalendarListResult `json:"laCalendarListResult,omitempty"` // 大家电的预约日历结果。
}

type SkuClassifyResult struct {
	ResultCode     int             `json:"resultCode,omitempty"`      // 状态，只有1表示处理成功
	ResultMessage  string          `json:"resultMessage,omitempty"`   // 对resultCode的简要说明
	SkuClassifyMap map[string]uint `json:"skuClassifyMaps,omitempty"` // key表示sku编码; value表示sku类型，1：表示中小件，2：表示大家电
}

type CalendarListResult struct {
	ResultCode    int               `json:"resultCode,omitempty"`    // 状态，只有1表示处理成功
	ResultMessage string            `json:"resultMessage,omitempty"` // 对resultCode的简要说明
	PromiseTime   map[string]string `json:"proimseTime,omitempty"`   // 履约时间 1、只工作日送货(双休日、假日不用送); 2、只双休日、假日送货(工作日不用送) ;3、工作日、双休日与假日均可送货;Sting里面放置返回的promise字符串（包括html语句）；若没有，则返回空字符串有预约日历时，才有值
	TipMsg        string            `json:"tipMsg,omitempty"`        // 提示信息 311区域：“配送服务升级，晚间也可以送货上门啦！” 211和次日达区域：“配送服务升级，可以指定上午和下午送货上门啦！”
	CalendarList  []CalendarDay     `json:"calendarList,omiempty"`   // 中小件预约日历的详细信息
}

type CalendarDay struct {
	Date     string         `json:"dateStr,omitempty"`  // 预约日期，格式：yyyy-MM-dd
	Week     uint           `json:"week,omitempty"`     // dateStr是周几，1：表示周1，7:表示周日
	TimeList []CalendarTime `json:"timeList,omitempty"` // dateStr这天的预约时间段
	Today    bool           `json:"today,omitempty"`    // 今天是否就是dateStr
}

type CalendarTime struct {
	TimeRange     string `json:"timeRange,omitempty"`     // 可预约的时间段
	Enable        bool   `json:"enable,omitempty"`        // 这个时间段是否可以选择预约
	Selected      bool   `json:"selected,omitempty"`      // 是否为默认系统选中的时间段
	TimeRangeCode int    `json:"timeRangeCode,omitempty"` // 时间段timeRange的编码
	BatchId       int    `json:"batchId,omitempty"`       // 波次id
	ReservingDate int    `json:"reservingDate,omitempty"` // 天的偏移量
}

type LaCalendarListResult struct {
	ResultCode              int              `json:"resultCode,omitempty"`           // 状态，只有1表示处理成功
	ResultMessage           string           `json:"resultMessage,omitempty"`        // 对resultCode的简要说明
	SupportShip             bool             `json:"supportShip,omitempty"`          // 是否支持配送
	SupportInstall          bool             `json:"supportInstall,omitempty"`       // 是否支持安装
	SupportNightShip        bool             `json:"supportNightShip,omitempty"`     // 是否支持夜间配送
	ReservingDateList       []int            `json:"reservingDateList,omitempty`     // 可选的预约配送时间,只有supportShip为true时才会有值。
	ReservingInstallDateMap map[string][]int `json:"reservingInstallDate,omitempty"` // 可选的预约安装时间，只有supportInstall为true时才会有值;
	SkuInfoList             []LaSku          `json:"skuInfoList,omitempty"`          // Sku列表   List<LaSku> 不支持配送 或支持安装时，此字段有值，其他情况下，此字段为null
	ReservedCalendarList    []CalendarDay    `json:"reservedCalendarList,omiempty"`  // 预约日期列表
	IsSetTimeArrive         bool             `json:"isSetTimeArrive,omitempty"`      // 是否支持定时达
}

type LaSku struct{}

type OrderPriceSnap struct {
	Price float64 `json:"price"` // 商品价格
	SkuId uint64  `json:"skuId"`
}

type BizSku struct {
	Id           uint64  `json:"skuId,omitempty"`
	Num          uint    `json:"num,omitempty"`          // 购买商品数量
	Category     uint64  `json:"category,omitempty"`     // 商品分类编号
	Price        float64 `json:"price,omitempty"`        // 商品单价
	Name         string  `json:"name,omitempty"`         // 商品名称
	Tax          float64 `json:"tax,omitempty"`          // 商品税率
	TaxPrice     float64 `json:"taxPrice,omitempty"`     // 商品税额
	NakedPrice   float64 `json:"nakedPrice,omitempty"`   // 商品未税价
	Type         uint    `json:"type,omitempty"`         // 商品类型：0普通、1附件、2赠品
	Oid          uint64  `json:"oid,omitempty"`          // 主商品skuid，如果本身是主商品，则oid为0
	SplitFreight float64 `json:"splitFreight,omitempty"` // 运费拆分价格
}

type ParentOrder interface {
	IsNumber() bool
}

type ParentOrderId uint64

func (this ParentOrderId) IsNumber() bool {
	return true
}

type Order struct {
	ParentOrder   ParentOrder `json:"pOrder,omitempty"`     // 父订单号。为0时，此订单为父单, 父订单响应为父订单详情
	ChildOrders   []Order     `json:"cOrder,omitempty"`     // 子订单详情
	State         uint        `json:"orderState,omitempty"` // 订单状态。0为取消订单  1为有效。
	Id            uint64      `json:"jdOrderId,omitempty"`
	LogisticState uint        `json:"state,omitempty"`           // 物流状态。0 是新建  1是妥投   2是拒收
	SubmitState   uint        `json:"submitState,omitempty"`     // 预占确认状态。0没确认预占。   1为确认预占。
	Type          uint        `json:"type,omitempty"`            // 订单类型。1是父订单   2是子订单。
	Freight       float64     `json:"freight,omitempty"`         // 运费
	Skus          []BizSku    `json:"skus,omitempty"`            // 商品列表
	Price         float64     `json:"orderPrice,omitempty"`      // 订单总金额。orderPrice(订单总金额)=商品总金额+运费
	NakedPrice    float64     `json:"orderNakedPrice,omitempty"` // 订单未含税金额。
	TaxPrice      float64     `json:"orderTaxPrice,omitempty"`   // 订单税额。
	OrderType     uint        `json:"orderType,omitempty"`       // 订单类别。查询参数queryExts中包含orderType。参考枚举值如下：1.普通商品 2.大家电 3.实物礼品卡 4.售后换新单 5.厂家直送订单 6.FBP订单 7.生鲜 20.电子卡 21.机票 22.酒店 23.合约机号卡 24.火车票[@文祥：更新新订单类型；父单子单的订单类型形成规则。特殊说明虚拟订单，虚拟订单通常有专门的查询接口]
	CreateTime    string      `json:"createOrderTime,omitempty"` // 订单创建时间。查询参数queryExts中包含createOrderTime。 输出格式为“yyyy-MM-dd hh:mm:ss”
	FinishTime    string      `json:"finishTime,omitempty"`      // 订单创建时间。 查询参数queryExts中包含createOrderTime。 输出格式为“yyyy-MM-dd hh:mm:ss” 未完成时，此参数返回null。
	OrderState    uint        `json:"jdOrderState,omitempty"`    // 京东状态。查询参数中包含queryExts=jdOrderState。参考枚举值如下：1.新单 2.等待支付 3.等待支付确认 4.延迟付款确认 5.订单暂停 6.店长最终审核 7.等待打印 8.等待出库 9.等待打包 10.等待发货 11.自提途中 12.上门提货 13.自提退货 14.确认自提 16.等待确认收货 17.配送退货 18.货到付款确认 19.已完成 21.收款确认 22.锁定 29.等待三方出库 30.等待三方发货 31.等待三方发货完成
	Address       string      `json:"address,omitempty"`         // 加密后的收货地址。查询参数queryExts中包含address。 解密方式：AD+client_id前6位 解密规则：DES/CBC/PKCS5Padding，Hex
	Name          string      `json:"name,omitempty"`            // 加密后的姓名。查询参数queryExts中包含name。解密方式：NA+client_id前6位，解密规则：DES/CBC/PKCS5Padding，Hex
	Mobile        string      `json:"mobile,omitempty"`          // 加密后的联系方式。查询参数queryExts中包含mobile。解密方式：MO+client_id前6位 解密规则：DES/CBC/PKCS5Padding，Hex
	PaymentType   int         `json:"paymentType,omitempty"`     // 支付方式 (1：货到付款， 4：在线支付，5：公司转账， 101：金采支付，20为混合支付)
	PayDetails    []PayDetail `json:"payDeatails,omitempty"`     // 混合支付明细，当paymentType为20混合支付，返回值
}

func (this *Order) IsNumber() bool {
	return false
}

type PayDetail struct {
	Money       float64 `json:"payMoney,omitempty"`    // 支付金额
	PaymentType string  `json:"paymentType,omitempty"` // 支付类型 枚举定义 17, "微信支付" 101, "金采支付"
	Flag        string  `json:"flag,omitempty"`        //支付标记 枚举定义 1  个人
}

type OrderTrackResult struct {
	OrderTracks  []OrderTrack  `json:"orderTrack,omitempty"`  // 配送的信息
	WaybillCodes []WaybillCode `json:"waybillCode,omitempty"` // 返回订单的运单信息。当入参的waybillCode=1时，返回此字段，
}

type OrderTrack struct {
	Content  string `json:"content,omitempty"`  // 操作内容明细
	MsgTime  string `json:"msgTime,omitempty"`  // 操作时间。日期格式为“yyyy-MM-dd hh:mm:ss”
	Operator string `json:"operator,omitempty"` // 操作员名称。
}

type WaybillCode struct {
	OrderId         uint64 `json:"orderId,omitempty"`
	ParentId        uint64 `json:"parentId,omitempty"`
	Carrier         string `json:"carrier,omitempty"`         // 承运商。可以为“京东快递”或者商家自行录入的承运商名称。
	DeliveryOrderId string `json:"deliveryOrderId,omitempty"` // 运单号。
}
