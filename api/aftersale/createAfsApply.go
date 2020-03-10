package aftersale

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type CreateAfsApplyRequest struct {
	OrderId               uint64              `json:"jdOrderId"`
	CustomerExpect        uint                `json:"customerExpect"`         // 售后类型：退货(10)、换货(20)、维修(30)
	QuestionDesc          string              `json:"questionDesc,omitempty"` // 产品问题描述，最多1000字符
	IsNeedDetectionReport bool                `json:"isNeedDetectionReport"`  // 是否需要检测报告
	QuestionPic           string              `json:"questionPic,omitempty"`  // 问题描述图片.最多2000字符 支持多张图片，用逗号分隔（英文逗号）
	IsHasPackage          bool                `json:"isHasPackage"`           // 是否有包装
	PackageDesc           uint                `json:"packageDesc"`            // 包装描述：0 无包装 10 包装完整 20 包装破损
	AsCustomerDto         AfterSaleCustomer   `json:"asCustomerDto"`          // 客户信息实体
	AsPickwareDto         AfterSalePickware   `json:"asPickwareDto"`          // 取件信息实体
	AsReturnwareDto       AfterSaleReturnware `json:"asReturnwareDto"`        // 返件信息实体
	AsDetailDto           AfterSaleDetail     `json:"asDetailDto"`            // 申请单明细
}

func (this *CreateAfsApplyRequest) Method() string {
	return "api/afterSale/createAfsApply"
}

func (this *CreateAfsApplyRequest) Values() url.Values {
	values := url.Values{}
	js, _ := json.Marshal(this)
	values.Set("param", string(js))
	return values
}

func (this *CreateAfsApplyRequest) Create(clt *jdvop.Client) error {
	_, err := clt.Do(this)
	return err
}
