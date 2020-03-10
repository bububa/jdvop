package stock

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type GetNewStockByIdRequest struct {
	SkuNums []SkuNum `json:"skuNums"` // 商品和数量  [{skuId: 569172,num:101}]。“{skuId: 569172,num:101}”为1条记录，此参数最多传入100条记录。
	Area    string   `json:"area"`    // 格式：1_0_0 (分别代表1、2、3级地址)
}

func (this *GetNewStockByIdRequest) Method() string {
	return "api/stock/getNewStockById"
}

func (this *GetNewStockByIdRequest) Values() url.Values {
	values := url.Values{}
	js, _ := json.Marshal(this.SkuNums)
	values.Set("skuNums", string(js))
	values.Set("area", this.Area)
	return values
}

func GetNewStockById(clt *jdvop.Client, skuNums []SkuNum, area string) ([]SkuStock, error) {
	resp, err := clt.Do(&GetNewStockByIdRequest{SkuNums: skuNums, Area: area})
	if err != nil {
		return nil, err
	}
	var ret string
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	var stocks []SkuStock
	err = json.Unmarshal([]byte(ret), &stocks)
	if err != nil {
		return nil, err
	}
	return stocks, nil
}
