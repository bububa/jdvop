package aftersale

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type GetServiceListPageRequest struct {
	OrderId   uint64 `json:"jdOrderId"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
}

func (this *GetServiceListPageRequest) Method() string {
	return "api/afterSale/getServiceListPage"
}

func (this *GetServiceListPageRequest) Values() url.Values {
	values := url.Values{}
	js, _ := json.Marshal(this)
	values.Set("param", string(js))
	return values
}

func GetServiceListPage(clt *jdvop.Client, orderId uint64, pageIndex int, pageSize int) (*AfsServiceByCustomerPinPage, error) {
	resp, err := clt.Do(&GetServiceListPageRequest{OrderId: orderId, PageIndex: pageIndex, PageSize: pageSize})
	if err != nil {
		return nil, err
	}
	var ret AfsServiceByCustomerPinPage
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
