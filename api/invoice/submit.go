package invoice

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

/*
每个批次订单数不能超过500，超过500要拆分多个批次。
总批次超过1时，当所有批次都收到后，融聚才会推送开票。
一个第三方申请单号可以对应多个当前批次，但总批次应该一样。
结算单号主要用于对账，与开票无关。一个结算单号可以对应多个第三方申请单号。
*/
type SubmitRequest struct {
	SupplierOrder                string  `json:"supplierOrder,omitempty"`                 // 子订单号，批量以英文逗号分割
	MarkId                       string  `json:"markId"`                                  // 第三方申请单号：申请发票的唯一id标识 (该标记下可以对应多张发票信息)
	SettlementId                 string  `json:"settlementId"`                            // 结算单号（一个结算单号可对对应多个第三方申请单号）
	SettlementNum                int     `json:"settlementNum,omitempty"`                 // 结算单子订单总数
	SettlementNakedPrice         float64 `json:"settlementNakedPrice,omitempty"`          // 结算单不含税总金额
	SettlementTaxPrice           float64 `json:"settlementTaxPrice,omitempty"`            // 结算单总税额
	InvoiceType                  uint    `json:"invoiceType"`                             // 发票类型 2：增值税专用发票  3 : 电子发票
	InvoiceOrg                   uint64  `json:"invoiceOrg"`                              // 开票机构ID（联系京东业务确定）
	BizInvoiceContent            uint    `json:"bizInvoiceContent"`                       // 开票内容：1, "明细" 100, "大类"
	InvoiceDate                  string  `json:"invoiceDate"`                             // 期望开票时间，格式：2013-11-8
	Title                        string  `json:"title"`                                   // 发票抬头
	BillToParty                  string  `json:"billToParty,omitempty"`                   // 收票单位
	EnterpriseTaxpayer           string  `json:"enterpriseTaxpayer,omitempty"`            // 纳税人识别号（专票必须）
	BillToer                     string  `json:"billToer,omitempty"`                      // 收票人 （专票必须）
	BillToContact                string  `json:"billToContact"`                           // 收票人联系电话
	BillToProvince               uint64  `json:"billToProvince,omitempty"`                // 收票人地址（省）  （专票必须）
	BillToCity                   uint64  `json:"billToCity,omitempty"`                    // 收票人地址（市）  （专票必须）
	BillToCounty                 uint64  `json:"billToCounty,omitempty"`                  // 收票人地址（区）  （专票必须）
	BillToTown                   uint64  `json:"billToTown,omitempty"`                    // 收票人地址（街道）（专票有四级地址则必传，否则传0）
	BillToAddress                string  `json:"billToAddress,omitempty"`                 // 收票人地址（详细地址）（专票必须）
	RepaymentDate                string  `json:"repaymentDate,omitempty"`                 // 预计还款时间2013-11-8（开票方式为预借时必传）
	InvoiceNum                   int     `json:"invoiceNum,omitempty"`                    // 当前批次子订单总数
	InvoiceNakedPrice            float64 `json:"invoiceNakedPrice,omitempty"`             // 当前批次不含税总金额（单位：元）
	InvoiceTaxPrice              float64 `json:"invoiceTaxPrice,omitempty"`               // 当前批次总税额(参考用)
	InvoicePrice                 float64 `json:"invoicePrice"`                            // 当前批次含税总金额
	CurrentBatch                 uint64  `json:"currentBatch"`                            // 当前批次号
	TotalBatch                   uint64  `json:"totalBatch"`                              // 总批次数
	TotalBatchInvoiceNakedAmount float64 `json:"totalBatchInvoiceNackedAmount,omitempty"` // 总批次开发票金额（不含税）
	TotalBatchInvoiceTaxAmount   float64 `json:"totalBatchInvoiceTaxAmount,omitempty"`    // 总批次开发票税额
	TotalBatchInvoiceAmount      float64 `json:"totalBatchInvoiceAmount"`                 // 总批次开发票价税合计
	BillType                     uint    `json:"billType,omitempty"`                      // 1-集中开票，2-分别开票（不传默认为集中开票）
	IsMerge                      uint    `json:"isMerge,omitempty"`                       // 合并开票，（不传默认为合并开票）1-合并SKU，空和其他-分别开票
	PoNo                         string  `json:"poNo,omitempty"`                          // 采购单号，长度范围[1-26]
	EnterpriseRegAddress         string  `json:"enterpriseRegAddress,omitempty"`          // 企业注册地址
	EnterpriseRegPhone           string  `json:"enterpriseRegPhone,omitempty"`            // 企业注册电话
	EnterpriseBankName           string  `json:"enterpriseBankName,omitempty"`            // 企业注册开户行名称
	EnterpriseBankAccount        string  `json:"enterpriseBankAccount,omitempty"`         // 企业注册开户行账号
}

func (this *SubmitRequest) Method() string {
	return "api/invoice/submit"
}

func (this *SubmitRequest) Values() url.Values {
	values := url.Values{}
	values.Set("supplierOrder", this.SupplierOrder)
	values.Set("markId", this.MarkId)
	values.Set("settlementId", this.SettlementId)
	if this.SettlementNum > 0 {
		values.Set("settlementNum", strconv.FormatInt(int64(this.SettlementNum), 10))
	}
	if this.SettlementNakedPrice > 0 {
		values.Set("settlementNakedPrice", strconv.FormatFloat(this.SettlementNakedPrice, 'f', -1, 64))
	}
	if this.SettlementTaxPrice > 0 {
		values.Set("settlementTaxPrice", strconv.FormatFloat(this.SettlementTaxPrice, 'f', -1, 64))
	}
	values.Set("invoiceType", strconv.FormatUint(uint64(this.InvoiceType), 10))
	values.Set("invoiceOrg", strconv.FormatUint(this.InvoiceOrg, 10))
	values.Set("bizInvoiceContent", strconv.FormatUint(uint64(this.BizInvoiceContent), 10))
	values.Set("invoiceDate", this.InvoiceDate)
	values.Set("title", this.Title)
	if this.BillToParty != "" {
		values.Set("billToParty", this.BillToParty)
	}
	if this.EnterpriseTaxpayer != "" {
		values.Set("enterpriseTaxpayer", this.EnterpriseTaxpayer)
	}
	if this.BillToer != "" {
		values.Set("billToer", this.BillToer)
	}
	values.Set("billToContact", this.BillToContact)
	if this.BillToProvince > 0 {
		values.Set("billToProvince", strconv.FormatUint(this.BillToProvince, 10))
	}
	if this.BillToCity > 0 {
		values.Set("billToCity", strconv.FormatUint(this.BillToCity, 10))
	}
	if this.BillToCounty > 0 {
		values.Set("billToCounty", strconv.FormatUint(this.BillToCounty, 10))
	}
	if this.BillToTown > 0 {
		values.Set("billToTown", strconv.FormatUint(this.BillToTown, 10))
	}
	if this.BillToAddress != "" {
		values.Set("billToAddress", this.BillToAddress)
	}
	if this.RepaymentDate != "" {
		values.Set("repaymentDate", this.RepaymentDate)
	}
	values.Set("invoiceNum", strconv.FormatInt(int64(this.InvoiceNum), 10))
	if this.InvoiceNakedPrice > 0 {
		values.Set("invoiceNakedPrice", strconv.FormatFloat(this.InvoiceNakedPrice, 'f', -1, 64))
	}
	if this.InvoiceTaxPrice > 0 {
		values.Set("invoiceTaxPrice", strconv.FormatFloat(this.InvoiceTaxPrice, 'f', -1, 64))
	}
	values.Set("invoicePrice", strconv.FormatFloat(this.InvoicePrice, 'f', -1, 64))
	values.Set("currentBatch", strconv.FormatUint(this.CurrentBatch, 10))
	values.Set("totalBatch", strconv.FormatUint(this.TotalBatch, 10))
	if this.TotalBatchInvoiceNakedAmount > 0 {
		values.Set("totalBatchInvoiceNakedAmount", strconv.FormatFloat(this.TotalBatchInvoiceNakedAmount, 'f', -1, 64))
	}
	if this.TotalBatchInvoiceTaxAmount > 0 {
		values.Set("totalBatchInvoiceTaxAmount", strconv.FormatFloat(this.TotalBatchInvoiceTaxAmount, 'f', -1, 64))
	}
	values.Set("totalBatchInvoiceAmount", strconv.FormatFloat(this.TotalBatchInvoiceAmount, 'f', -1, 64))
	if this.BillType == 1 || this.BillType == 2 {
		values.Set("billType", strconv.FormatUint(uint64(this.BillType), 10))
	}
	if this.IsMerge == 1 {
		values.Set("isMerge", strconv.FormatUint(uint64(this.IsMerge), 10))
	}
	if this.PoNo != "" {
		values.Set("poNo", this.PoNo)
	}
	if this.EnterpriseRegAddress != "" {
		values.Set("enterpriseRegAddress", this.EnterpriseRegAddress)
	}
	if this.EnterpriseRegPhone != "" {
		values.Set("enterpriseRegPhone", this.EnterpriseRegPhone)
	}
	if this.EnterpriseBankName != "" {
		values.Set("enterpriseBankName", this.EnterpriseBankName)
	}
	if this.EnterpriseBankAccount != "" {
		values.Set("enterpriseBankAccount", this.EnterpriseBankAccount)
	}
	return values
}

func (this *SubmitRequest) Submit(clt *jdvop.Client) (bool, error) {
	resp, err := clt.Do(this)
	if err != nil {
		return false, err
	}
	var ret bool
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return false, err
	}
	return ret, nil
}
