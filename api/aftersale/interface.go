package aftersale

type ComponentExport struct {
	Code string `json:"code,omitempty"` // 服务类型码：退货(10)、换货(20)、维修(30), 上门取件(4)、客户发货(40)、客户送货(7)
	Name string `json:"name,omitempty"` // 服务类型名称：退货、换货、维修, 上门取件、客户发货、客户送货
}

type AfterSaleCustomer struct {
	Pin         string `json:"customerPin,omitempty"`        // 客户京东账号
	ContactName string `json:"customerContactName"`          // 联系人，最多50字符
	Tel         string `json:"customerTel"`                  // 联系电话，最多50字符
	Mobile      string `json:"customerMobilePhone"`          // 手机号，最多50字符
	Email       string `json:"customerEmail,omitempty"`      // Email
	Postcode    string `json:"customerPostercode,omitempty"` // 邮编
}

type AfterSalePickware struct {
	Type     uint   `json:"pickwareType"`     // 取件方式：4上门取件； 7 客户送货；40客户发货
	Province uint64 `json:"pickwareProvince"` // 取件省，取件方式为4（上门取件）时，必填
	City     uint64 `json:"pickwareCity"`     // 取件市，取件方式为4（上门取件）时，必填
	County   uint64 `json:"pickwareCounty"`   // 取件县，取件方式为4（上门取件）时，必填
	Village  uint64 `json:"pickwareVillage"`  // 取件乡镇
	Address  string `json:"pickwareAddress"`  // 取件街道地址，取件方式为4（上门取件）时必填，最多500字符
}

type AfterSaleReturnware struct {
	Type     uint   `json:"returnwareType"`     // 返件方式：自营配送(10),第三方配送(20);
	Province uint64 `json:"returnwareProvince"` // 返件省
	City     uint64 `json:"returnwareCity"`     // 返件市
	County   uint64 `json:"returnwareCounty"`   // 返件县
	Village  uint64 `json:"returnwareVillage"`  // 返件乡镇
	Address  string `json:"returnwareAddress"`  // 返件街道地址，最多500字符
}

type AfterSaleAddressInfo struct {
	PostCode string `json:"postCode,omitempty"` // 售后邮编
	LinkMan  string `json:"linkMan"`            // 售后联系人
	Tel      string `json:"tel"`                // 售后电话
	Address  string `json:"address"`            // 售后地址
}

type AfterSaleDetail struct {
	SkuId  uint64 `json:"skuId"`
	SkuNum uint   `json:"skuNum"` // 商品申请数量
}

type AfsServiceByCustomerPinPage struct {
	ServiceInfoList []AfsServiceByCustomerPin `json:"serviceInfoList,omitempty"` // 售后服务单列表
	Total           int                       `json:"totalNum,omitempty"`        // 总记录数
	PageSize        int                       `json:"pageSize,omitempty"`        // 每页记录数
	PageIndex       int                       `json:"pageCount,omitempty"`       // 当前页数
	PageCount       int                       `json:"pageNum,omitempty"`         // 总页数
}

type AfsServiceByCustomerPin struct {
	Id                 uint64 `json:"afsServiceId,omitempty"`       // 服务单号
	CustomerExpect     uint   `json:"customerExpect,omitempty"`     // 服务类型码：退货(10)、换货(20)、维修(30)
	CustomerExpectName string `json:"customerExpectName,omitempty"` // 服务类型名称
	ApplyTime          string `json:"afsApplyTime,omitempty"`       // 服务单申请时间
	OrderId            uint64 `json:"orderId,omitempty"`            // 订单号
	WareId             uint64 `json:"wareId,omitempty"`             // 商品编号
	WareName           string `json:"wareName,omitempty"`           // 商品名称
	Step               uint   `json:"afsServiceStep,omitempty"`     // 服务单环节：申请阶段(10), 审核不通过(20), 客服审核(21), 商家审核(22), 京东收货(31), 商家收货(32), 京东处理(33), 商家处理(34), 用户确认(40), 完成(50), 取消(60);
	StepName           string `json:"afsServiceStepName,omitempty"` // 服务单环节名称：申请阶段,客服审核,商家审核,京东收货,商家收货, 京东处理,商家处理, 用户确认,完成, 取消;
	Cancel             uint   `json:"cancel,omitempty"`             // 是否可取消: 0代表否，1代表是
}

type CompitableServiceDetail struct {
	Id                    uint64                     `json:"afsServiceId,omitempty"`                 // 服务单号
	CustomerExpect        uint                       `json:"customerExpect,omitempty"`               // 服务类型码：退货(10)、换货(20)、维修(30)
	ApplyTime             string                     `json:"afsApplyTime,omitempty"`                 // 服务单申请时间
	OrderId               uint64                     `json:"orderId,omitempty"`                      // 订单号
	IsHasInvoice          uint                       `json:"isHasInvoice,omitempty"`                 // 是不是有发票：0没有 1有
	IsNeedDetectionReport uint                       `json:"isNeedDetectionReport,omitempty"`        // 是不是有检测报告：0没有 1有
	IsHasPackage          uint                       `json:"isHasPackage,omitempty"`                 // 是不是有包装：0没有 1有
	QuestionPic           string                     `json:"questionPic,omitempty"`                  // 问题描述图片.最多2000字符 支持多张图片，用逗号分隔（英文逗号）
	Step                  uint                       `json:"afsServiceStep,omitempty"`               // 服务单环节：申请阶段(10), 审核不通过(20), 客服审核(21), 商家审核(22), 京东收货(31), 商家收货(32), 京东处理(33), 商家处理(34), 用户确认(40), 完成(50), 取消(60);
	StepName              string                     `json:"afsServiceStepName,omitempty"`           // 服务单环节名称：申请阶段,客服审核,商家审核,京东收货,商家收货, 京东处理,商家处理, 用户确认,完成, 取消;
	ApproveNotes          string                     `json:"approveNotes,omitempty"`                 // 审核意见
	QuestionDesc          string                     `json:"questionDesc,omitempty"`                 // 产品问题描述，最多1000字符
	ApprovedResult        uint                       `json:"approvedResult,omitempty"`               // 审核结果：直赔积分 (11),直赔余额 (12),直赔优惠卷 (13),直赔京豆 (14),直赔商品 (21),上门换新 (22),自营取件 (31),客户送货(32),客户发货 (33),闪电退款 (34),虚拟退款 (35),大家电检测 (80),大家电安装 (81),大家电移机 (82),大家电维修 (83),大家电其它(84);
	ApprovedResultName    string                     `json:"approvedResultName,omitempty"`           // 审核结果名称：直赔积分 ,直赔余额 ,直赔优惠卷 ,直赔京豆,直赔商品,上门换新,自营取件 ,客户送货,客户发货,闪电退款,虚拟退款,大家电检测,大家电安装,大家电移机,大家电维修 ,大家电其它;
	ProcessResult         uint                       `json:"processResult,omitempty"`                // 处理结果：返修换新(23),退货(40), 换良(50),原返(60),病单 (71),出检 (72),维修(73),强制关单 (80),线下换新 (90)
	ProcessResultName     string                     `json:"processResultName,omitempty"`            // 处理结果名称：返修换新,退货 , 换良,原返,病单,出检,维修,强制关单,线下换新
	CustomerInfo          *AfterSaleCustomer         `json:"serviceCustomerInfoDTO,omitempty"`       // 客户信息
	AddressInfo           *AfterSaleAddressInfo      `json:"serviceAftersalesAddressInfo,omitempty"` // 售后地址信息
	ExpressInfo           *ServiceExpressInfo        `json:"serviceExpressInfoDTO,omitempty"`        // 客户发货信息
	FinanceDetailInfo     []ServiceFinanceDetailInfo `json:"serviceFinanceDetailInfoDTOs,omitempty"` // 退款明细
	TrackInfo             []ServiceTrackInfo         `json:"serviceTrackInfoDTOs,omitempty"`         // 服务单追踪信息
	DetailInfo            []ServiceDetailInfo        `json:"serviceDetailInfoDTOs,omitempty"`        // 服务单商品明细
	AllowOperations       []uint                     `json:"allowOperations,omitempty"`              // 获取服务单允许的操作列表：列表为空代表不允许操作；列表包含1代表取消；列表包含2代表允许填写或者修改客户发货信息；
}

type ServiceExpressInfo struct {
	ServiceId      uint64 `json:"afsServiceId,omitempty"`   // 服务单号
	FreightMoney   string `json:"freightMoney,omitempty"`   // 运费
	ExpressCompany string `json:"expressCompany,omitempty"` // 快递公司名称
	DeliverDate    string `json:"deliverDate,omitempty"`    // 客户发货日期：格式为yyyy-MM-dd HH:mm:ss
	ExpressCode    string `json:"expressCode,omitempty"`    // 快递单号
}

type ServiceFinanceDetailInfo struct {
	RefundWay     uint    `json:"refundWay,omitempty"`     // 退款方式
	RefundWayName string  `json:"refundWayName,omitempty"` // 退款方式名称
	Status        int     `json:"status,omitempty"`        // 状态
	StatusName    string  `json:"statusName,omitempty"`    // 状态名称
	RefundPrice   float64 `json:"refundPrice,omitempty"`   // 退款金额
	WareName      string  `json:"wareName,omitempty"`      // 商品名称
	WareId        uint64  `json:"wareId,omitempty"`        // 商品编号
	B2bPayType    int     `json:"b2bPayType,omitempty"`    // B2B支付方式
	B2bPayFlag    uint    `json:"b2bPayFlag,omitempty"`    // 个人/企业
}

type ServiceTrackInfo struct {
	ServiceId  uint64 `json:"afsServiceId,omitempty"` // 服务单号
	Title      string `json:"title,omitempty"`        // 追踪标题
	Context    string `json:"context,omitempty"`      // 追踪内容
	CreateDate string `json:"createDate,omitempty"`   // 提交时间：格式为yyyy-MM-dd HH:mm:ss
	CreateName string `json:"createName,omitempty"`   // 操作人昵称
	CreatePin  string `json:"createPin,omitempty"`    // 操作人账号
}

type ServiceDetailInfo struct {
	WareId        uint64 `json:"wareId,omitempty"`
	WareName      string `json:"wareName,omitempty"`
	wareBrand     string `json:"wareBrand,omitempty"`
	AfsDetailType uint   `json:"afsDetailType,omitempty"` //  明细类型：主商品(10), 赠品(20), 附件(30)，拍拍取主商品就可以
	WareDescribe  string `json:"wareDescribe,omitempty"`  // 附件描述
}

type BizRefundDetail struct {
	RequestId          string              `json:"requestId,omitempty"`        // 售后服务单号
	OrderId            string              `json:"orderId,omitempty"`          // 京东订单号
	RefundDetailType   uint                `json:"refundDetailType,omitempty"` // 退款明细类型
	SourceStatus       int                 `json:"sourceStatus,omitempty"`
	RefundAmount       string              `json:"refundAmount,omitempty"`        // 退款金额
	RefundBankName     string              `json:"refundBankName,omitempty"`      // 退款银行名称
	RefundBankProvince string              `json:"refundBankProvince,omitempty"`  // 退款银行省份
	RefundBankCity     string              `json:"refundBankCity,omitempty"`      // 退款银行城市
	RefundSubBankName  string              `json:"refundSubBankName,omitempty"`   // 退款支行名称
	RefundCardNumber   string              `json:"refundCardNumber,omitempty"`    // 退款卡号
	CouponNumber       uint64              `json:"couponNumber,omitempty"`        // 优惠券号码
	CouponPrice        float64             `json:"couponPrice,omitempty"`         // 优惠券金额
	RefundCardUser     string              `json:"refundCardUser,omitempty"`      // 退款卡拥有人
	CommitBankId       string              `json:"commitBankId,omitempty"`        // 银行ID
	FinanceNo          string              `json:"financeNo,omitempty"`           // 财务号
	ConfirmDate        string              `json:"confirmDate,omitempty"`         // 确认时间，格式：2019-03-01 09:03:13
	ActualRefundType   uint                `json:"actualRefundType,omitempty"`    // 实际退款类型
	RefundJdBankId     uint64              `json:"refundJdBankId,omitempty"`      // 京东退款银行Id
	RefundPaymentInfos []RefundPaymentInfo `json:"refundPaymentInfoVo,omitempty"` // 退款支付信息列表
}

type RefundPaymentInfo struct {
	PayId            string  `json:"payId,omitempty"`            // 支付单号
	PayType          uint    `json:"payType,omitempty"`          // 主站支付方式：6：在线支付
	PayEnum          uint    `json:"payEnum,omitempty"`          // 支付枚举值
	PaymentType      int     `json:"paymentType,omitempty"`      // B2B支付方式
	B2BPayFlag       uint    `json:"b2bPayFlag,omitempty"`       // 个人/企业
	RefundableAmount float64 `json:"refundableAmount,omitempty"` // 退款金额
	PayTime          string  `json:"payTime,omitempty"`          // 支付时间
}
