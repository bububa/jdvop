package order

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type GetPromiseTipsRequest struct {
	SkuId    string `json:"skuId"`
	Province string `json:"province"`
	City     string `json:"city"`
	County   string `json:"county"`
	Town     string `json:"town,omitempty"`
}

func (this *GetPromiseTipsRequest) Method() string {
	return "api/order/getPromiseTips"
}

func (this *GetPromiseTipsRequest) Values() url.Values {
	values := url.Values{}
	values.Set("skuId", this.SkuId)
	values.Set("province", this.Province)
	values.Set("city", this.City)
	values.Set("county", this.County)
	if this.Town != "" {
		values.Set("town", this.Town)
	}
	return values
}

func GetPromiseTips(clt *jdvop.Client, skuId string, province string, city string, county string, town string) (string, error) {
	req := &GetPromiseTipsRequest{
		SkuId:    skuId,
		Province: province,
		City:     city,
		County:   county,
		Town:     town,
	}
	resp, err := clt.Do(req)
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
