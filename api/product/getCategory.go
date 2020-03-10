package product

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type GetCategoryRequest struct {
	Cid uint64 `json:"skuId"`
}

func (this *GetCategoryRequest) Method() string {
	return "api/product/getCategory"
}

func (this *GetCategoryRequest) Values() url.Values {
	values := url.Values{}
	values.Set("cid", strconv.FormatUint(this.Cid, 10))
	return values
}

func GetCategory(clt *jdvop.Client, cid uint64) (*Category, error) {
	resp, err := clt.Do(&GetCategoryRequest{Cid: cid})
	if err != nil {
		return nil, err
	}
	var ret Category
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
