package area

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type GetCityRequest struct {
	ProvinceId uint64 `json:"id"`
}

func (this *GetCityRequest) Method() string {
	return "api/area/getCity"
}

func (this *GetCityRequest) Values() url.Values {
	values := url.Values{}
	values.Set("id", strconv.FormatUint(this.ProvinceId, 10))
	return values
}

func GetCity(clt *jdvop.Client, provinceId uint64) ([]Location, error) {
	resp, err := clt.Do(&GetCityRequest{ProvinceId: provinceId})
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
