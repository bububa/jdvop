package order

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type SelectJdOrderIdByThirdOrderRequest struct {
	ThirdOrder string `json:"thirdOrder"` // 第三方的订单单号，必须在100字符以内
}

func (this *SelectJdOrderIdByThirdOrderRequest) Method() string {
	return "api/order/selectJdOrderIdByThirdOrder"
}

func (this *SelectJdOrderIdByThirdOrderRequest) Values() url.Values {
	values := url.Values{}
	values.Set("thirdOrder", this.ThirdOrder)
	return values
}

func SelectJdOrderIdByThirdOrder(clt *jdvop.Client, thirdOrder string) (string, error) {
	resp, err := clt.Do(&SelectJdOrderIdByThirdOrderRequest{ThirdOrder: thirdOrder})
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
