package order

import (
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type ConfirmOrderRequest struct {
	OrderId uint64 `json:"jdOrderId"` // 京东的订单单号(下单返回的父订单号)
}

func (this *ConfirmOrderRequest) Method() string {
	return "api/order/confirmOrder"
}

func (this *ConfirmOrderRequest) Values() url.Values {
	values := url.Values{}
	values.Set("jdOrderId", strconv.FormatUint(this.OrderId, 10))
	return values
}

func ConfirmOrder(clt *jdvop.Client, orderId uint64) error {
	_, err := clt.Do(&ConfirmOrderRequest{OrderId: orderId})
	return err
}
