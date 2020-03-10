package product

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type GetSimilarSkuRequest struct {
	SkuId string `json:"skuId"`
}

func (this *GetSimilarSkuRequest) Method() string {
	return "api/product/getSimilarSku"
}

func (this *GetSimilarSkuRequest) Values() url.Values {
	values := url.Values{}
	values.Set("skuId", this.SkuId)
	return values
}

func GetSimilarSku(clt *jdvop.Client, skuId string) ([]SimilarProduct, error) {
	resp, err := clt.Do(&GetSimilarSkuRequest{SkuId: skuId})
	if err != nil {
		return nil, err
	}
	var ret []SimilarProduct
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
