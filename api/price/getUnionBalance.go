package price

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type GetUnionBalanceRequest struct {
	Pin  string `json:"pin"`  // 京东PIN。必须是相同合同下的pin
	Type string `json:"type"` // 账户余额类型。多选，可用英文逗号拼接。参考枚举值如下：1：账户余额。2：金采余额。
}

func (this *GetUnionBalanceRequest) Method() string {
	return "api/price/getUnionBalance"
}

func (this *GetUnionBalanceRequest) Values() url.Values {
	values := url.Values{}
	values.Set("pin", this.Pin)
	values.Set("type", this.Type)
	return values
}

func GetUnionBalance(clt *jdvop.Client, pin string, balanceType string) (*UnionBalance, error) {
	resp, err := clt.Do(&GetUnionBalanceRequest{Pin: pin, Type: balanceType})
	if err != nil {
		return nil, err
	}
	var ret UnionBalance
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
