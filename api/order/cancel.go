package order

import (
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type CancelRequest struct {
	OrderId uint64 `json:"jdOrderId"` // 京东的订单单号(下单返回的父订单号)
}

func (this *CancelRequest) Method() string {
	return "api/order/cancel"
}

func (this *CancelRequest) Values() url.Values {
	values := url.Values{}
	values.Set("jdOrderId", strconv.FormatUint(this.OrderId, 10))
	return values
}

func Cancel(clt *jdvop.Client, orderId uint64) error {
	_, err := clt.Do(&CancelRequest{OrderId: orderId})
	return err
}
