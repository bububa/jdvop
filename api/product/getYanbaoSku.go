package product

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type GetYanbaoSkuRequest struct {
	SkuIds   string `json:"skuIds"`
	Province string `json:"province"`
	City     string `json:"city"`
	County   string `json:"county"`
	Town     string `json:"town,omitempty"`
}

func (this *GetYanbaoSkuRequest) Method() string {
	return "api/product/getYanbaoSku"
}

func (this *GetYanbaoSkuRequest) Values() url.Values {
	values := url.Values{}
	values.Set("skuIds", this.SkuIds)
	values.Set("province", this.Province)
	values.Set("city", this.City)
	values.Set("county", this.County)
	if this.Town != "" {
		values.Set("town", this.Town)
	}
	return values
}

func GetYanbaoSku(clt *jdvop.Client, skuIds string, province string, city string, county string, town string) (map[uint64]Yanbao, error) {
	req := &GetYanbaoSkuRequest{
		SkuIds:   skuIds,
		Province: province,
		City:     city,
		County:   county,
		Town:     town,
	}
	resp, err := clt.Do(req)
	if err != nil {
		return nil, err
	}
	ret := make(map[uint64]Yanbao)
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
