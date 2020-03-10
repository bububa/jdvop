package product

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type GetSkuGiftRequest struct {
	SkuId    string `json:"skuId"`
	Province string `json:"province"`
	City     string `json:"city"`
	County   string `json:"county"`
	Town     string `json:"town,omitempty"`
}

func (this *GetSkuGiftRequest) Method() string {
	return "api/product/getSkuGift"
}

func (this *GetSkuGiftRequest) Values() url.Values {
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

func GetSkuGift(clt *jdvop.Client, skuId string, province string, city string, county string, town string) (*SkuGifts, error) {
	req := &GetSkuGiftRequest{
		SkuId:    skuId,
		Province: province,
		City:     city,
		County:   county,
		Town:     town,
	}
	resp, err := clt.Do(req)
	if err != nil {
		return nil, err
	}
	var ret SkuGifts
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
