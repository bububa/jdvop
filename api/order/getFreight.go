package order

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type GetFreightRequest struct {
	Skus        []SkuNum `json:"sku"` //  [{“skuId”:商品编号1,”num”:商品数量1},{“skuId”:商品编号2,”num”:商品数量2}]（最多支持50种商品）
	Province    string   `json:"province"`
	City        string   `json:"city"`
	County      string   `json:"county"`
	Town        string   `json:"town,omitempty"`
	PaymentType int      `json:"paymentType,omitempty"` // 京东支付方式
}

func (this *GetFreightRequest) Method() string {
	return "api/order/getFreight"
}

func (this *GetFreightRequest) Values() url.Values {
	values := url.Values{}
	js, _ := json.Marshal(this.Skus)
	values.Set("sku", string(js))
	values.Set("province", this.Province)
	values.Set("city", this.City)
	values.Set("county", this.County)
	if this.Town != "" {
		values.Set("town", this.Town)
	}
	values.Set("paymentType", strconv.FormatInt(int64(this.PaymentType), 10))
	return values
}

func GetFreight(clt *jdvop.Client, skus []SkuNum, province string, city string, county string, town string, paymentType int) ([]Freight, error) {
	req := &GetFreightRequest{
		Skus:        skus,
		Province:    province,
		City:        city,
		County:      county,
		Town:        town,
		PaymentType: paymentType,
	}
	resp, err := clt.Do(req)
	if err != nil {
		return nil, err
	}
	var ret []Freight
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
