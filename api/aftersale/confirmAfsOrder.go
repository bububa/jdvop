package aftersale

import (
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type ConfirmAfsOrderRequest struct {
	CustomerName string `json:"customerName"`
	Username     string `json:"username"`     // 服务单对应的用户username
	ServiceId    uint64 `json:"afsServiceId"` // 售后服务单号
}

func (this *ConfirmAfsOrderRequest) Method() string {
	return "api/afterSale/confirmAfsOrder"
}

func (this *ConfirmAfsOrderRequest) Values() url.Values {
	values := url.Values{}
	values.Set("customerName", this.CustomerName)
	values.Set("username", this.Username)
	values.Set("afsServiceId", strconv.FormatUint(this.ServiceId, 10))
	return values
}

func ConfirmAfsOrder(clt *jdvop.Client, customerName string, username string, serviceId uint64) error {
	_, err := clt.Do(&ConfirmAfsOrderRequest{CustomerName: customerName, Username: username, ServiceId: serviceId})
	return err
}
