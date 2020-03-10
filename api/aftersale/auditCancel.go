package aftersale

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type AuditCancelRequest struct {
	ApproveNotes string   `json:"approveNotes"`  // 审核意见
	ServiceIds   []uint64 `json:"serviceIdList"` // 京东售后服务单集合
}

func (this *AuditCancelRequest) Method() string {
	return "api/afterSale/auditCancel"
}

func (this *AuditCancelRequest) Values() url.Values {
	values := url.Values{}
	js, _ := json.Marshal(this)
	values.Set("param", string(js))
	return values
}

func AuditCancel(clt *jdvop.Client, serviceIds []uint64, notes string) (bool, error) {
	resp, err := clt.Do(&AuditCancelRequest{ServiceIds: serviceIds, ApproveNotes: notes})
	if err != nil {
		return false, err
	}
	var ret bool
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return false, err
	}
	return ret, nil
}
