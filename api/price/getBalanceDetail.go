package price

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type GetBalanceDetailRequest struct {
	PageNum   int    `json:"pageNum"`
	PageSize  int    `json:"pageSize"`
	OrderId   string `json:"orderId"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

func (this *GetBalanceDetailRequest) Method() string {
	return "api/price/getBalanceDetail"
}

func (this *GetBalanceDetailRequest) Values() url.Values {
	values := url.Values{}
	values.Set("orderId", this.OrderId)
	values.Set("startDate", this.StartDate)
	values.Set("endDate", this.EndDate)
	if this.PageNum > 0 {
		values.Set("pageNum", strconv.FormatInt(int64(this.PageNum), 10))
	}
	if this.PageSize > 0 {
		values.Set("pageSize", strconv.FormatInt(int64(this.PageSize), 10))
	}
	return values
}

func (this *GetBalanceDetailRequest) Get(clt *jdvop.Client) (*BalanceDetails, error) {
	resp, err := clt.Do(this)
	if err != nil {
		return nil, err
	}
	var ret BalanceDetails
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
