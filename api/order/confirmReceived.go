package order

import (
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type ConfirmReceivedRequest struct {
	OrderId uint64 `json:"jdOrderId"` // 京东订单号
}

func (this *ConfirmReceivedRequest) Method() string {
	return "api/order/confirmReceived"
}

func (this *ConfirmReceivedRequest) Values() url.Values {
	values := url.Values{}
	values.Set("jdOrderId", strconv.FormatUint(this.OrderId, 10))
	return values
}

func ConfirmReceived(clt *jdvop.Client, orderId uint64) error {
	_, err := clt.Do(&ConfirmReceivedRequest{OrderId: orderId})
	return err
}
