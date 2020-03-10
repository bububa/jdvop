package area

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type GetJDAddressFromAddressRequest struct {
	Address string `json:"address"`
}

func (this *GetJDAddressFromAddressRequest) Method() string {
	return "api/area/getJDAddressFromAddress"
}

func (this *GetJDAddressFromAddressRequest) Values() url.Values {
	values := url.Values{}
	values.Set("address", this.Address)
	return values
}

func GetJDAddressFromAddress(clt *jdvop.Client, address string) (*Address, error) {
	resp, err := clt.Do(&GetJDAddressFromAddressRequest{Address: address})
	if err != nil {
		return nil, err
	}
	var ret Address
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
