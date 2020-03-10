package order

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type SelectJdOrderRequest struct {
	OrderId   uint64 `json:"jdOrderId"`           // 京东订单号
	QueryExts string `json:"queryExts,omitempty"` // 扩展参数。支持多个状态组合查询[英文逗号间隔]: orderType, jdOrderState, name, address, mobile, poNo, finishTime, createOrderTime, paymentType
}

func (this *SelectJdOrderRequest) Method() string {
	return "api/order/selectJdOrderIdByThirdOrder"
}

func (this *SelectJdOrderRequest) Values() url.Values {
	values := url.Values{}
	values.Set("jdOrderId", strconv.FormatUint(this.OrderId, 10))
	if this.QueryExts != "" {
		values.Set("queryExts", this.QueryExts)
	}
	return values
}

func SelectJdOrder(clt *jdvop.Client, orderId uint64, queryExts string) (*Order, error) {
	resp, err := clt.Do(&SelectJdOrderRequest{OrderId: orderId, QueryExts: queryExts})
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
