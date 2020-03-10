package area

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type GetCountyRequest struct {
	CityId uint64 `json:"id"`
}

func (this *GetCountyRequest) Method() string {
	return "api/area/getCounty"
}

func (this *GetCountyRequest) Values() url.Values {
	values := url.Values{}
	values.Set("id", strconv.FormatUint(this.CityId, 10))
	return values
}

func GetCounty(clt *jdvop.Client, cityId uint64) ([]Location, error) {
	resp, err := clt.Do(&GetCountyRequest{CityId: cityId})
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
