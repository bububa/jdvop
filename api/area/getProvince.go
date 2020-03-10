package area

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type GetProvinceRequest struct {
}

func (this *GetProvinceRequest) Method() string {
	return "api/area/getProvince"
}

func (this *GetProvinceRequest) Values() url.Values {
	return url.Values{}
}

func GetProvince(clt *jdvop.Client) ([]Location, error) {
	resp, err := clt.Do(&GetProvinceRequest{})
	if err != nil {
		return nil, err
	}
	ret := make(map[string]uint64)
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	var locs []Location
	for name, id := range ret {
		locs = append(locs, Location{
			Id:   id,
			Name: name,
		})
	}
	return locs, nil
}
