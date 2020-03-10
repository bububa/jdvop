package area

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type CheckAreaRequest struct {
	ProvinceId uint64 `json:"provinceId"`
	City       uint64 `json:"cityId"`
	CountyId   uint64 `json:"countyId"`
	TownId     uint64 `json:"townId"`
}

func (this *CheckAreaRequest) Method() string {
	return "api/area/checkArea"
}

func (this *CheckAreaRequest) Values() url.Values {
	values := url.Values{}
	values.Set("provinceId", strconv.FormatUint(this.ProvinceId, 10))
	values.Set("cityId", strconv.FormatUint(this.City, 10))
	values.Set("countyId", strconv.FormatUint(this.CountyId, 10))
	values.Set("townId", strconv.FormatUint(this.TownId, 10))
	return values
}

type CheckAreaResponse struct {
	Success    bool   `json:"success"`
	ResultCode int    `json:"resultCode"`
	AddressId  uint64 `json:"addressId"`
	Message    string `json:"message"`
}

func (this *CheckAreaResponse) IsError() bool {
	return this.Success
}

func (this *CheckAreaResponse) Error() string {
	return fmt.Sprintf("Code: %d, Msg: %s", this.ResultCode, this.Message)
}

func (this *CheckAreaRequest) Run(clt *jdvop.Client) (uint64, error) {
	resp, err := clt.Do(this)
	if err != nil {
		return 0, err
	}
	var ret CheckAreaResponse
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return 0, err
	}
	if ret.IsError() {
		return 0, &ret
	}
	return ret.AddressId, nil
}
