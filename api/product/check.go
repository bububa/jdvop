package product

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type CheckRequest struct {
	SkuIds    string `json:"skuIds"`
	QueryExts string `json:"queryExts"` // 扩展参数：英文逗号间隔输入, noReasonToReturn //无理由退货类型, thwa //无理由退货文案类型, isSelf // 是否自营, isJDLogistics // 是否京东配送, taxInfo //商品税率
}

func (this *CheckRequest) Method() string {
	return "api/product/check"
}

func (this *CheckRequest) Values() url.Values {
	values := url.Values{}
	values.Set("skuIds", this.SkuIds)
	values.Set("queryExts", this.QueryExts)
	return values
}

func Check(clt *jdvop.Client, skuIds string, queryExts string) ([]SkuSaleState, error) {
	resp, err := clt.Do(&CheckRequest{SkuIds: skuIds, QueryExts: queryExts})
	if err != nil {
		return nil, err
	}
	var ret []SkuSaleState
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
