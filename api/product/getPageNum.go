package product

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type GetPageNumRequest struct {
	QueryExts string `json:"queryExts"`
}

func (this *GetPageNumRequest) Method() string {
	return "api/product/getPageNum"
}

func (this *GetPageNumRequest) Values() url.Values {
	values := url.Values{}
	if this.QueryExts != "" {
		values.Set("queryExts", this.QueryExts)
	}
	return values
}

func GetPageNum(clt *jdvop.Client, queryExts string) ([]PageNum, error) {
	resp, err := clt.Do(&GetPageNumRequest{QueryExts: queryExts})
	if err != nil {
		return nil, err
	}
	var ret []PageNum
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
