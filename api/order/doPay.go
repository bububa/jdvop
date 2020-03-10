package order

import (
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type DoPayRequest struct {
	OrderId uint64 `json:"jdOrderId"` // 京东的订单单号(下单返回的父订单号)
}

func (this *DoPayRequest) Method() string {
	return "api/order/doPay"
}

func (this *DoPayRequest) Values() url.Values {
	values := url.Values{}
	values.Set("jdOrderId", strconv.FormatUint(this.OrderId, 10))
	return values
}

func DoPay(clt *jdvop.Client, orderId uint64) error {
	_, err := clt.Do(&DoPayRequest{OrderId: orderId})
	return err
}
