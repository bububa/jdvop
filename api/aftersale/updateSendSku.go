package aftersale

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type UpdateSendSkuRequest struct {
	AfsServiceId   uint64  `json:"afsServiceId"`   // 服务单号
	FreightMoney   float64 `json:"freightMoney"`   // 运费
	ExpressCompony string  `json:"expressCompany"` // 发运公司：圆通快递、申通快递、韵达快递、中通快递、宅急送、EMS、顺丰快递
	DeliverDate    string  `json:"deliverDate"`    // 发货日期，格式为yyyy-MM-dd HH:mm:ss
	ExpressCode    string  `json:"expressCode"`    // 货运单号，最大50字符
}

func (this *UpdateSendSkuRequest) Method() string {
	return "api/afterSale/updateSendSku"
}

func (this *UpdateSendSkuRequest) Values() url.Values {
	values := url.Values{}
	js, _ := json.Marshal(this)
	values.Set("param", string(js))
	return values
}

func (this *UpdateSendSkuRequest) Create(clt *jdvop.Client) error {
	_, err := clt.Do(this)
	return err
}
