package product

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type SkuImageRequest struct {
	Sku string `json:"sku"` // 商品编号，支持批量，以“,”（半角）分隔  (最高支持100个商品)
}

func (this *SkuImageRequest) Method() string {
	return "api/product/skuImage"
}

func (this *SkuImageRequest) Values() url.Values {
	values := url.Values{}
	values.Set("sku", this.Sku)
	return values
}

func GetSkuImage(clt *jdvop.Client, sku string) (map[string][]*SkuImage, error) {
	resp, err := clt.Do(&SkuImageRequest{Sku: sku})
	if err != nil {
		return nil, err
	}
	ret := make(map[string][]*SkuImage)
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
