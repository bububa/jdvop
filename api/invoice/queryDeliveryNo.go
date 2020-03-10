package invoice

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type QueryDeliveryNoRequest struct {
	OrderId uint64 `json:"jdOrderId"`
}

func (this *QueryDeliveryNoRequest) Method() string {
	return "api/invoice/queryDeliveryNo"
}

func (this *QueryDeliveryNoRequest) Values() url.Values {
	values := url.Values{}
	values.Set("jdOrderId", strconv.FormatUint(this.OrderId, 10))
	return values
}

func QueryDeliveryNo(clt *jdvop.Client, orderId uint64) ([]BizInvoiceTrace, error) {
	resp, err := clt.Do(&QueryDeliveryNoRequest{OrderId: orderId})
	if err != nil {
		return nil, err
	}
	var ret []BizInvoiceTrace
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
