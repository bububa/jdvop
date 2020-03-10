package area

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type GetTownRequest struct {
	CountyId uint64 `json:"id"`
}

func (this *GetTownRequest) Method() string {
	return "api/area/getTown"
}

func (this *GetTownRequest) Values() url.Values {
	values := url.Values{}
	values.Set("id", strconv.FormatUint(this.CountyId, 10))
	return values
}

func GetTown(clt *jdvop.Client, countyId uint64) ([]Location, error) {
	resp, err := clt.Do(&GetTownRequest{CountyId: countyId})
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
