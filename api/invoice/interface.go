package invoice

type BizInvoiceDetail struct {
	Id          string  `json:"invoiceId,omitempty"`          // 发票号码
	Code        string  `json:"invoiceCode,omitempty"`        // 发票代码
	Date        string  `json:"invoiceDate,omitempty"`        // 发票日期
	NakedAmount float64 `json:"invoiceNakedAmount,omitempty"` // 发票金额（裸价）（专票有值）
	TaxRate     float64 `json:"invoiceTaxRate,omitempty"`     // 发票税额（专票有值）
	Amount      float64 `json:"invoiceAmount,omitempty"`      // 价税合计
	Type        uint    `json:"invoiceType,omitempty"`        // 发票类型 2：专票 3：电子普通发票
	Success     bool    `json:"success,omitempty"`            // 开票状态
}

type InvoiceItem struct {
	Id           string             `json:"invoiceId,omitempty"`           // 发票号码
	Code         string             `json:"invoiceCode,omitempty"`         // 发票代码
	State        uint               `json:"state,omitempty"`               // 发票状态【1正常2作废3冲红】
	OriginalId   string             `json:"originalInvoiceId,omitempty"`   // 冲红的原始发票号
	OriginalCode string             `json:"originalInvoiceCode,omitempty"` // 冲红的原始发票代码
	Date         string             `json:"invoiceDate,omitempty"`         // 发票日期
	NakedAmount  float64            `json:"invoiceNakedAmount,omitempty"`  // 发票金额（裸价）（专票有值）
	TaxRate      float64            `json:"invoiceTaxRate,omitempty"`      // 发票税额（专票有值）
	TaxAmount    float64            `json:"invoiceTaxAmount,omitempty"`    // 发票税额
	Amount       float64            `json:"invoiceAmount,omitempty"`       // 价税合计
	Type         uint               `json:"invoiceType,omitempty"`         // 发票类型 2：专票 3：电子普通发票
	Title        string             `json:"title,omitempty"`               // 发票抬头
	Taxpayer     string             `json:"taxpayer,omitempty"`            // 纳税人识别号（专票有值）
	Address      string             `json:"address,omitempty"`             // 地址（专票有值）
	Tel          string             `json:"tel,omitempty"`                 // 电话（专票有值）
	Bank         string             `json:"bank,omitempty"`                // 开户行（专票有值）
	Account      string             `json:"account,omitempty"`             // 账号（专票有值）
	Remark       string             `json:"remark,omitempty"`              // 备注
	SkuDetails   []InvoiceSkuDetail `json:"skuDetails,omitempty"`          // 商品明细
}

type InvoiceSkuDetail struct {
	OrderId       string  `json:"jdOrderId,omitempty"`
	SkuId         string  `json:"skuId,omitempty"`
	SkuName       string  `json:"skuName,omitempty"`
	Price         float64 `json:"price,omitempty"`
	Num           float64 `json:"num,omitempty"`
	TaxRate       float64 `json:"taxRate,omitempty"`
	Amount        float64 `json:"amount,omitempty"`        // 总金额（含税）
	AmountUnTax   float64 `json:"amountUnTax,omitempty"`   // 总金额（不含税）
	TaxAmount     float64 `json:"taxAmount,omitempty"`     // 税额
	Specification string  `json:"specification,omitempty"` // 规格型号
	SettleUnit    string  `json:"settleUnit,omitempty"`    // 结算单位
}

type BizInvoice interface {
	Type() uint
}

type BizInvoiceActualBillCResp struct {
	IvcCode        string  `json:"ivcCode,omitempty"`        // 发票代码
	IvcNo          string  `json:"ivcNo,omitempty"`          // 发票号
	IvcType        uint    `json:"ivcType,omitempty"`        // 发票类型
	IvcContentType uint    `json:"ivcContentType,omitempty"` // 开票内容
	IvcContentName string  `json:"ivcContentName,omitempty"` // 开票内容名称
	IvcTitle       string  `json:"ivcTitle,omitempty"`       // 发票抬头
	IFlag          uint    `json:"iFlag,omitempty"`          // 发票类别（1.正数；2.负数）1代表蓝票 2代表红票
	BussinessId    string  `json:"bussinessId,omitempty"`    // 订单号
	InvoiceTime    string  `json:"invoiceTime,omitempty"`    // 开票日期
	TotalPrice     float64 `json:"totalPrice,omitempty"`     // 开票金额,订单在本发票中的总金额
	TotalTaxPrice  float64 `json:"totalTaxPrice,omitempty"`  // 税额
	BlueIsn        string  `json:"blueIsn,omitempty"`        // 红票对应蓝票序列号
	Isn            string  `json:"isn,omitempty"`            // 发票ISN
	Remark         string  `json:"remark,omitempty"`         // 备注
	TaxRate        float64 `json:"taxRate,omitempty"`        // 税率
	FileUrl        string  `json:"fileUrl,omitempty"`        // 电子票下载地址
	Yn             uint    `json:"yn,omitempty"`             // 作废标识(1,有效；0.作废)--不作为判断是否是冲红的标志
}

func (this *BizInvoiceActualBillCResp) Type() uint {
	return 3
}

type BizIvcChainInvoiceResp struct {
	Uuid           string                  `json:"uuid,omitempty"`
	InvoiceCode    string                  `json:"invoiceCode,omitempty"`
	InvoiceNo      string                  `json:"invoiceNo,omitempty"`
	InvoiceDate    string                  `json:"invoiceDate,omitempty"`
	InvoiceType    uint                    `json:"invoiceType,omitempty"`    // 发票类型 2：增值税电子发票 1：增值税专用发票
	CheckCode      string                  `json:"checkCode,omitempty"`      // 校验码
	PrinterNo      string                  `json:"printerNo,omitempty"`      // 机器编号
	BuyerName      string                  `json:"buyerName,omitempty"`      // 购货方发票抬头
	BuyerTaxNo     string                  `json:"buyerTaxNo,omitempty"`     // 购货方纳税人识别号
	BuyerAddress   string                  `json:"buyerAddress,omitempty"`   // 购货方纳税人地址
	BuyerPhone     string                  `json:"buyerPhone,omitempty"`     // 购货方电话
	BuyerBank      string                  `json:"buyerBank,omitempty"`      // 购货方开户行
	BuyerAccount   string                  `json:"buyerAccount,omitempty"`   // 购货方银行账号
	Password       string                  `json:"password,omitempty"`       // 密码区
	TotalAmount    float64                 `json:"totalAmount,omitempty"`    // 金额合计
	TotalTaxAmount float64                 `json:"totalTaxAmount,omitempty"` // 税额合计
	AllAmount      float64                 `json:"allAmount,omitempty"`      // 价税合计
	SellerName     string                  `json:"sellerName,omitempty"`     // 销货方名称
	SellerTaxNo    string                  `json:"sellerTaxNo,omitempty"`    // 销货方税号
	SellerAddress  string                  `json:"sellerAddress,omitempty"`  // 销货方地址
	SellerPhone    string                  `json:"sellerPhone,omitempty"`    // 销货方电话
	SellerBank     string                  `json:"sellerBank,omitempty"`     // 销货方银行
	SellerAccount  string                  `json:"sellerAccount,omitempty"`  // 销货方银行账号
	Payee          string                  `json:"payee,omitempty"`          // 收款人
	Reviewer       string                  `json:"reviewer,omitempty"`       // 复核人
	Printor        string                  `json:"printor,omitempty"`        // 开票人
	Remark         string                  `json:"remark,omitempty"`
	IsSaleList     string                  `json:"isSaleList,omitempty"`   // 是否有销货清单
	Pdf            string                  `json:"pdf,omitempty"`          // pdf下载链接
	InvoiceState   uint                    `json:"invoiceState,omitempty"` // 发票状态
	IsRed          string                  `json:"isRed,omitempty"`        // 是否红票 true/false
	BlueCode       string                  `json:"blueCode,omitempty"`     // 红票对应的蓝票代码
	BlueNo         string                  `json:"blueNo,omitempty"`       // 红票对应的蓝票号码
	CreateTime     string                  `json:"createTime,omitempty"`   // 创建时间
	DetailList     []BizIvcChainDetailResp `json:"detailList,omitempty"`   // 明细列表
}

func (this *BizIvcChainInvoiceResp) Type() uint {
	return 2
}

type BizIvcChainDetailResp struct {
	IsDiscount string  `json:"isDiscount,omitempty"` // 是否折扣行 true/false
	IvwName    string  `json:"ivwName,omitempty"`    // 货物名称
	IvSpec     string  `json:"ivSpec,omitempty"`     // 规格型号
	Unit       string  `json:"unit,omitempty"`       // 单位
	Wnumb      float64 `json:"wnumb,omitempty"`      // 数量
	Wprice     float64 `json:"wprice,omitempty"`     // 单价
	Amount     float64 `json:"amount,omitempty"`     // 金额
	TaxRate    float64 `json:"taxRate,omitempty"`    // 税率
	TaxAmount  float64 `json:"taxAmount,omitempty"`  // 税额
	CreateTime string  `json:"createTime,omitempty"` // 创建时间
}

type BizInvoiceDelivery struct {
	PostId      uint64 `json:"ivcPostId,omitempty"`   // 邮寄ID
	DeliveryId  uint64 `json:"deliveryId,omitempty"`  // 运单号
	PostCompany string `json:"postCompany,omitempty"` // 配送公司
	PostTime    string `json:"postTime,omitempty"`    // 配送时间
	State       uint   `json:"state,omitempty"`       // 配送状态（1 已邮寄; 2 已取消;）
}

type BizInvoiceTrace struct {
	Title       string `json:"opeTitle,omitempty"`    // 操作主题
	Remark      string `json:"opeRemark,omitempty"`   // 操作备注
	Name        string `json:"opeName,omitempty"`     // 操作用户
	Time        string `json:"opeTime,omitempty"`     // 操作时间
	WaybillCode string `json:"waybillCode,omitempty"` // 物流单号
}
