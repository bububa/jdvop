package order

import (
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type SaveOrUpdatePoNoRequest struct {
	OrderId uint64 `json:"jdOrderId"` // 京东订单号
	PoNo    string `json:"poNo"`      // 采购单号，长度范围[1-26]
}

func (this *SaveOrUpdatePoNoRequest) Method() string {
	return "api/order/saveOrUpdatePoNo"
}

func (this *SaveOrUpdatePoNoRequest) Values() url.Values {
	values := url.Values{}
	values.Set("jdOrderId", strconv.FormatUint(this.OrderId, 10))
	values.Set("poNo", this.PoNo)
	return values
}

func SaveOrUpdatePoNo(clt *jdvop.Client, orderId uint64, poNo string) error {
	_, err := clt.Do(&SaveOrUpdatePoNoRequest{OrderId: orderId, PoNo: poNo})
	return err
}
