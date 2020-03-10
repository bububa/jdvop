package aftersale

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

// 如果不传jdOrderId、sku则查询username下某时间段内所有售后服务单列表情况；
type GetAfsServiceListPageRequest struct {
	Username  string `json:"username,omitempty"` // 用户名称，若为空，以生成token用户查询，非空以username为准
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	OrderId   uint64 `json:"jdOrderId,omitempty"`
	SkuId     string `json:"sku,omitempty"`
}

func (this *GetAfsServiceListPageRequest) Method() string {
	return "api/afterSale/getAfsServiceListPage"
}

func (this *GetAfsServiceListPageRequest) Values() url.Values {
	values := url.Values{}
	js, _ := json.Marshal(this)
	values.Set("param", string(js))
	return values
}

func (this *GetAfsServiceListPageRequest) Get(clt *jdvop.Client) (*AfsServiceByCustomerPinPage, error) {
	resp, err := clt.Do(this)
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
