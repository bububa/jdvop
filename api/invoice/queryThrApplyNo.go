package invoice

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type QueryThrApplyNoRequest struct {
	OrderId uint64 `json:"jdOrderId"`
}

func (this *QueryThrApplyNoRequest) Method() string {
	return "api/invoice/queryThrApplyNo"
}

func (this *QueryThrApplyNoRequest) Values() url.Values {
	values := url.Values{}
	values.Set("jdOrderId", strconv.FormatUint(this.OrderId, 10))
	return values
}

func QueryThrApplyNo(clt *jdvop.Client, orderId uint64) (string, error) {
	resp, err := clt.Do(&QueryThrApplyNoRequest{OrderId: orderId})
	if err != nil {
		return "", err
	}
	var ret string
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return "", err
	}
	return ret, nil
}
