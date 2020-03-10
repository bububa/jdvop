package checkorder

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type CheckNewOrderRequest struct {
	Date         string `json:"date"`              // 查询日期，格式2018-11-7（不包含当天）
	PageNo       int    `json:"pageNo"`            //页码，默认1
	PageSize     int    `json:"pageSize"`          // 页大小取值范围[1,100]，默认20
	OrderIdIndex uint64 `json:"jdOrderIdIndex"`    // 最小订单号索引游标，为解决大于1W条订单无法查询问题。注意事项：该字段和pageNo互斥，订单数小于1W可以用pageNo分页的方式来查询，订单数大于1W则需要使用索引游标的方式来读取数据。使用方式：第一次查询无需传入该字段，返回订单信息后（第一次记录订单总条数）；第二次查询将第一次查询结果中最小的订单号传入，查询返回结果中不包含传入的订单号；递归这个流程，直到接口返回无数据为止，订单查询完毕，核对本地订单数和第一次接口返回的订单数目是否一致。如果使用本字段：订单号必须大于1
	EndDate      string `json:"endDate,omitempty"` // 结束日期，格式2018-11-7。主要用于查询时间段内，跟date配合使用。
}

func (this *CheckNewOrderRequest) Method() string {
	return "api/checkOrder/checkNewOrder"
}

func (this *CheckNewOrderRequest) Values() url.Values {
	values := url.Values{}
	values.Set("date", this.Date)
	if this.OrderIdIndex > 0 {
		values.Set("jdOrderIdIndex", strconv.FormatUint(this.OrderIdIndex, 10))
	} else {
		if this.PageNo < 1 {
			this.PageNo = 1
		}
		values.Set("pageNo", strconv.FormatInt(int64(this.PageNo), 10))
	}
	values.Set("pageSize", strconv.FormatInt(int64(this.PageSize), 10))
	if this.EndDate != "" {
		values.Set("endDate", this.EndDate)
	}
	return values
}

func (this *CheckNewOrderRequest) Check(clt *jdvop.Client) (*CheckOrderResult, error) {
	resp, err := clt.Do(this)
	if err != nil {
		return nil, err
	}
	var ret CheckOrderResult
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
