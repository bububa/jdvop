package aftersale

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type GetOrderPayByOrderIdRequest struct {
	OrderId string `json:"orderId"`
	RefId   uint64 `json:"refId,omitempty"` // 售后服务单号
}

func (this *GetOrderPayByOrderIdRequest) Method() string {
	return "api/afterSale/getOrderPayByOrderId"
}

func (this *GetOrderPayByOrderIdRequest) Values() url.Values {
	values := url.Values{}
	values.Set("orderId", this.OrderId)
	if this.RefId > 0 {
		values.Set("refId", strconv.FormatUint(this.RefId, 10))
	}
	return values
}

func GetOrderPayByOrderId(clt *jdvop.Client, orderId string, refId uint64) (*BizRefundDetail, error) {
	resp, err := clt.Do(&GetOrderPayByOrderIdRequest{OrderId: orderId, RefId: refId})
	if err != nil {
		return nil, err
	}
	var ret BizRefundDetail
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
