package order

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type SubmitOrderRequest struct {
	ThirdOrder           string            `json:"thirdOrder"`                     // 第三方的订单单号，必须在100字符以内
	Skus                 []SkuNum          `json:"sku"`                            // 下单商品信息
	Name                 string            `json:"name"`                           // 收货人姓名，最多20个字符
	Province             string            `json:"province"`                       // 一级地址编码：收货人省份地址编码
	City                 string            `json:"city"`                           // 二级地址编码：收货人市级地址编码
	County               string            `json:"county"`                         // 三级地址编码：收货人县（区）级地址编码
	Town                 string            `json:"town,omitempty"`                 // 四级地址编码：收货人乡镇地址编码(如果该地区有四级地址，则必须传递四级地址，没有四级地址则传0)
	Address              string            `json:"address"`                        // 收货人详细地址，最多100个字符
	Zip                  string            `json:"zip,omitempty"`                  // 邮编，最多20个字符
	Phone                string            `json:"phone,omitempty"`                // 座机号，最多20个字符
	Mobile               string            `json:"mobile"`                         // 手机号，最多20个字符
	Email                string            `json:"email"`                          // 邮箱
	Remark               string            `json:"remark,omitempty"`               // 备注（少于100字）
	InvoiceState         uint              `json:"invoiceState"`                   // 开票方式(2为集中开票，4 订单完成后开票)
	InvoiceType          uint              `json:"invoiceType"`                    // 发票类型（2增值税专用发票；3 电子票）当发票类型为2时，开票方式只支持2集中开票
	SelectedInvoiceTitle uint              `json:"selectedInvoiceTitle"`           // 发票类型：4：个人，5：单位
	CompanyName          string            `json:"companyName"`                    // 发票抬头  (如果selectedInvoiceTitle=5则此字段必须)
	InvoiceContent       uint              `json:"invoiceContent"`                 // 1:明细，100：大类 备注:若增值税专用发票则只能选1 明细
	DoOrderPriceMode     uint              `json:"doOrderPriceMode"`               // 1价格校验（建议传1）；0 不需要
	OrderPriceSnap       []OrderPriceSnap  `json:"orderPriceSnap"`                 // doOrderPriceMode传1的时候必传
	PaymentType          int               `json:"paymentType"`                    // 支付方式 (1：货到付款， 4：在线支付，即账户余额，此支付方式需先充值，5：公司转账， 101：金采支付，20：混合支付)
	IsUseBalance         uint              `json:"isUseBalance"`                   // 使用余额paymentType=4时，此值固定是1 其他支付方式0
	SubmitState          uint              `json:"submitState"`                    // 是否预占库存，0是预占库存（需要调用确认订单接口），1是不预占库存
	InvoiceName          string            `json:"invoiceName,omitempty"`          // 增专票收票人姓名
	InvoicePhone         string            `json:"invoicePhone"`                   // 收票人电话
	InvoiceProvice       string            `json:"invoiceProvice,omitempty"`       // 增专票收票人所在省(京东地址编码)当发票了类型为2增值税专用发票时，该字段必填。
	InvoiceCity          string            `json:"invoiceCity,omitempty"`          // 增专票收票人所在市(京东地址编码) 当invoiceType =2时，该字段必填。
	InvoiceCounty        string            `json:"invoiceCounty,omitempty"`        // 增专票收票人所在区/县(京东地址编码) 当invoiceType =2时，该字段必填。
	InvoiceAddress       string            `json:"invoiceAddress,omitempty"`       // 增专票收票人所在地址当invoiceType =2时，该字段必填。
	RegCompanyName       string            `json:"RegCompanyName,omitempty"`       // 专票资质公司名称当invoiceType =2时，该字段必填。
	RegCode              string            `json:"regCode,omitempty"`              // 专票资质纳税人识别号 当invoiceType =2时，该字段必填。
	RegAddr              string            `json:"regAddr,omitempty"`              // 专票资质注册地址 当invoiceType =2时，该字段必填。
	RegPhone             string            `json:"regPhone,omitempty"`             // 专票资质注册电话 当invoiceType =2时，该字段必填。
	RegBank              string            `json:"regBank,omitempty"`              // 专票资质注册银行 当invoiceType =2时，该字段必填。
	RegBankAccount       string            `json:"regBankAccount,omitempty"`       // 专票资质银行账号 当invoiceType =2时，该字段必填。
	ReservingDate        int               `json:"reservingDate,omitempty"`        // 大家电配送日期：默认值为-1，0表示当天，1表示明天，2：表示后天; 如果为-1表示不使用大家电预约日历
	InstallDate          int               `json:"installDate,omitempty"`          // 大家电安装日期：默认按-1处理，0表示当天，1表示明天，2：表示后天
	NeedInstall          bool              `json:"needInstall,omitempty"`          // 是否选择了安装，默认为true，选择了“暂缓安装”，此为必填项，必填值为false。
	PromiseDate          string            `json:"promiseDate,omitempty"`          // 中小件配送预约日期，格式：yyyy-MM-dd
	PromiseTimeRange     string            `json:"promiseTimeRange,omitempty"`     // 中小件配送预约时间段，时间段如： 9:00-15:00
	PromiseTimeRangeCode int               `json:"promiseTimeRangeCode,omitempty"` // 中小件预约时间段的标记
	ReservedDate         string            `json:"reservedDateStr,omitempty"`      // 家电配送预约日期，格式：yyyy-MM-dd
	ReservedTimeRange    string            `json:"reservedTimeRange,omitempty"`    // 大家电配送预约时间段，如果：9:00-15:00
	CycleCalendar        map[string]string `json:"cycleCalendar,omitempty"`        // 循环日历, 客户传入最近一周可配送的时间段,客户入参:{"3": "09:00-10:00,12:00-19:00","4": "09:00-15:00"}
	PoNo                 string            `json:"poNo,omitempty"`                 // 采购单号，长度范围[1-26]
	ValidHolidayVocation bool              `json:"validHolidayVocation,omitempty"` // 节假日不可配送，默认值为false，表示节假日可以配送，为true表示节假日不配送
}

func (this *SubmitOrderRequest) Method() string {
	return "api/order/submitOrder"
}

func (this *SubmitOrderRequest) Values() url.Values {
	values := url.Values{}
	values.Set("thirdOrder", this.ThirdOrder)
	{
		js, _ := json.Marshal(this.Skus)
		values.Set("sku", string(js))
	}
	values.Set("name", this.Name)
	values.Set("province", this.Province)
	values.Set("city", this.City)
	values.Set("county", this.County)
	if this.Town != "" {
		values.Set("town", this.Town)
	}
	values.Set("address", this.Address)
	if this.Zip != "" {
		values.Set("zip", this.Zip)
	}
	if this.Phone != "" {
		values.Set("phone", this.Phone)
	}
	values.Set("mobile", this.Mobile)
	values.Set("email", this.Email)
	if this.Remark != "" {
		values.Set("remark", this.Remark)
	}
	values.Set("invoiceState", strconv.FormatUint(uint64(this.InvoiceState), 10))
	values.Set("invoiceType", strconv.FormatUint(uint64(this.InvoiceType), 10))
	values.Set("selectedInvoiceTitle", strconv.FormatUint(uint64(this.SelectedInvoiceTitle), 10))
	if this.SelectedInvoiceTitle == 5 {
		values.Set("companyName", this.CompanyName)
	}
	values.Set("invoiceContent", strconv.FormatUint(uint64(this.InvoiceContent), 10))
	values.Set("doOrderPriceMode", strconv.FormatUint(uint64(this.DoOrderPriceMode), 10))
	if this.OrderPriceSnap != nil {
		js, _ := json.Marshal(this.OrderPriceSnap)
		values.Set("orderPriceSnap", string(js))
	}
	values.Set("paymentType", strconv.FormatInt(int64(this.PaymentType), 10))
	if this.PaymentType == 4 {
		values.Set("isUseBalance", "1")
	} else {
		values.Set("isUseBalance", "0")
	}
	values.Set("submitState", strconv.FormatUint(uint64(this.SubmitState), 10))
	if this.InvoiceName != "" {
		values.Set("invoiceName", this.InvoiceName)
	}
	values.Set("invoicePhone", this.InvoicePhone)
	if this.InvoiceProvice != "" {
		values.Set("invoiceProvice", this.InvoiceProvice)
	}
	if this.InvoiceCity != "" {
		values.Set("invoiceCity", this.InvoiceCity)
	}
	if this.InvoiceCounty != "" {
		values.Set("invoiceCounty", this.InvoiceCounty)
	}
	if this.InvoiceAddress != "" {
		values.Set("invoiceAddress", this.InvoiceAddress)
	}
	if this.RegCompanyName != "" || this.InvoiceType == 2 {
		values.Set("regCompanyName", this.RegCompanyName)
	}
	if this.RegCode != "" || this.InvoiceType == 2 {
		values.Set("regCode", this.RegCode)
	}
	if this.RegAddr != "" || this.InvoiceType == 2 {
		values.Set("regAddr", this.RegAddr)
	}
	if this.RegPhone != "" || this.InvoiceType == 2 {
		values.Set("regPhone", this.RegPhone)
	}
	if this.RegBank != "" || this.InvoiceType == 2 {
		values.Set("regBank", this.RegBank)
	}
	if this.RegBankAccount != "" || this.InvoiceType == 2 {
		values.Set("regBankAccount", this.RegBankAccount)
	}
	if this.ReservingDate >= 0 {
		values.Set("reservingDate", strconv.FormatInt(int64(this.ReservingDate), 10))
	}
	if this.InstallDate >= 0 {
		values.Set("installDate", strconv.FormatInt(int64(this.InstallDate), 10))
	}
	if this.NeedInstall {
		values.Set("needInstall", "true")
	} else {
		values.Set("needInstall", "false")
	}
	if this.PromiseDate != "" {
		values.Set("promiseDate", this.PromiseDate)
	}
	if this.PromiseTimeRange != "" {
		values.Set("promiseTimeRange", this.PromiseTimeRange)
	}
	if this.PromiseTimeRangeCode > 0 {
		values.Set("promiseTimeRangeCode", strconv.FormatInt(int64(this.PromiseTimeRangeCode), 10))
	}
	if this.ReservedDate != "" {
		values.Set("reservedDateStr", this.ReservedDate)
	}
	if this.ReservedTimeRange != "" {
		values.Set("reservedTimeRange", this.ReservedTimeRange)
	}
	if this.CycleCalendar != nil {
		js, _ := json.Marshal(this.CycleCalendar)
		values.Set("cycleCalendar", string(js))
	}
	if this.PoNo != "" {
		values.Set("poNo", this.PoNo)
	}
	if this.ValidHolidayVocation {
		values.Set("validHolidayVocation", "true")
	}
	return values
}

func (this *SubmitOrderRequest) Submit(clt *jdvop.Client) (*Order, error) {
	resp, err := clt.Do(this)
	if err != nil {
		return nil, err
	}
	var ret Order
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
