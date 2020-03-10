package product

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type GetIsCodRequest struct {
	SkuIds    string `json:"skuIds"`
	Province  string `json:"province"`
	City      string `json:"city"`
	County    string `json:"county"`
	Town      string `json:"town,omitempty"`
	QueryExts string `json:"queryExts,omitempty"` // skuIds //返回具体的skuId明细，例102194,13781
}

func (this *GetIsCodRequest) Method() string {
	return "api/product/getIsCod"
}

func (this *GetIsCodRequest) Values() url.Values {
	values := url.Values{}
	values.Set("skuIds", this.SkuIds)
	values.Set("province", this.Province)
	values.Set("city", this.City)
	values.Set("county", this.County)
	if this.Town != "" {
		values.Set("town", this.Town)
	}
	if this.QueryExts != "" {
		values.Set("queryExts", this.QueryExts)
	}
	return values
}

func GetIsCod(clt *jdvop.Client, skuIds string, province string, city string, county string, town string, queryExts string) (bool, error) {
	req := &GetIsCodRequest{
		SkuIds:    skuIds,
		Province:  province,
		City:      city,
		County:    county,
		Town:      town,
		QueryExts: queryExts,
	}
	resp, err := clt.Do(req)
	if err != nil {
		return false, err
	}
	var ret bool
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return false, err
	}
	return ret, nil
}
