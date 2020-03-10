package price

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type GetSellPriceRequest struct {
	Sku       string `json:"sku"`
	QueryExts string `json:"queryExts"` // 为英文半角分隔的多个枚举值，枚举值不同，本接口的出参不同。枚举值如下：Price //大客户默认价格(根据合同类型查询价格)，该字段必传。 marketPrice //市场价。containsTax //税率。出参增加tax,taxPrice,nakedPrice 3个字段。 nakedPrice//未税价。出参增加nakedPrice字段
}

func (this *GetSellPriceRequest) Method() string {
	return "api/price/getSellPrice"
}

func (this *GetSellPriceRequest) Values() url.Values {
	values := url.Values{}
	values.Set("sku", this.Sku)
	if this.QueryExts != "" {
		values.Set("queryExts", this.QueryExts)
	}
	return values
}

func GetSellPrice(clt *jdvop.Client, sku string, queryExts string) ([]SellPrice, error) {
	resp, err := clt.Do(&GetSellPriceRequest{Sku: sku, QueryExts: queryExts})
	if err != nil {
		return nil, err
	}
	var ret []SellPrice
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
