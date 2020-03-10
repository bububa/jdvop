package order

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type OrderTrackRequest struct {
	OrderId     uint64 `json:"jdOrderId"`             // 京东订单号
	WaybillCode uint   `json:"waybillCode,omitempty"` //是否返回订单的配送信息。0不返回配送信息。1，返回配送信息。只支持最近2个月的配送信息查询。
}

func (this *OrderTrackRequest) Method() string {
	return "api/order/orderTrack"
}

func (this *OrderTrackRequest) Values() url.Values {
	values := url.Values{}
	values.Set("jdOrderId", strconv.FormatUint(this.OrderId, 10))
	values.Set("waybillCode", strconv.FormatUint(uint64(this.WaybillCode), 10))
	return values
}

func GetOrderTrack(clt *jdvop.Client, orderId uint64, waybillCode uint) (*OrderTrackResult, error) {
	resp, err := clt.Do(&OrderTrackRequest{OrderId: orderId, WaybillCode: waybillCode})
	if err != nil {
		return nil, err
	}
	var ret OrderTrackResult
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
