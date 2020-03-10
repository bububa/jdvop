package invoice

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type GetInvoiceListRequest struct {
	OrderId   uint64 `json:"jdOrderId"`
	IvcType   uint   `json:"ivcType"`             // 发票类型：1 普票，2 专票，3 电子发票
	QueryExts string `json:"queryExts,omitempty"` // 扩展参数：英文逗号间隔输入 prefixZero：专票发票号前面补齐零 electronicVAT：专票电子化，（返回独立的对象：BizIvcChainInvoiceRespVo）
}

func (this *GetInvoiceListRequest) Method() string {
	return "api/invoice/getInvoiceList"
}

func (this *GetInvoiceListRequest) Values() url.Values {
	values := url.Values{}
	values.Set("jdOrderId", strconv.FormatUint(this.OrderId, 10))
	values.Set("ivcType", strconv.FormatUint(uint64(this.IvcType), 10))
	if this.QueryExts != "" {
		values.Set("queryExts", this.QueryExts)
	}
	return values
}

func GetInvoiceList(clt *jdvop.Client, orderId uint64, ivcType uint, queryExts string) ([]BizInvoice, error) {
	resp, err := clt.Do(&GetInvoiceListRequest{OrderId: orderId, IvcType: ivcType, QueryExts: queryExts})
	if err != nil {
		return nil, err
	}
	var ret []BizInvoice
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
