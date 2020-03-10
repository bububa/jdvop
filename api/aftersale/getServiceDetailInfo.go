package aftersale

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type GetServiceDetailInfoRequest struct {
	AfsServiceId    uint64 `json:"afsServiceId"`              // 京东售后服务单号
	AppendInfoSteps []uint `json:"appendInfoSteps,omitempty"` // 获取信息模块：不设置数据表示只获取服务单主信息、商品明细以及客户信息；1、代表增加获取售后地址信息；2、代表增加获取客户发货信息；3、代表增加获取退款明细；4、代表增加获取跟踪信息；5、代表增加获取允许的操作信息
}

func (this *GetServiceDetailInfoRequest) Method() string {
	return "api/afterSale/getServiceDetailInfo"
}

func (this *GetServiceDetailInfoRequest) Values() url.Values {
	values := url.Values{}
	js, _ := json.Marshal(this)
	values.Set("param", string(js))
	return values
}

func GetServiceDetailInfo(clt *jdvop.Client, serviceId uint64, steps []uint) (*CompitableServiceDetail, error) {
	resp, err := clt.Do(&GetServiceDetailInfoRequest{AfsServiceId: serviceId, AppendInfoSteps: steps})
	if err != nil {
		return nil, err
	}
	var ret CompitableServiceDetail
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
