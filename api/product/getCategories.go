package product

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type GetCategoriesRequest struct {
	PageNo   int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
	ParentId uint64 `json:"parentId"`
	Class    uint   `json:"catClass"`
}

func (this *GetCategoriesRequest) Method() string {
	return "api/product/getCategories"
}

func (this *GetCategoriesRequest) Values() url.Values {
	values := url.Values{}
	values.Set("pageNo", strconv.FormatInt(int64(this.PageNo), 10))
	values.Set("pageSize", strconv.FormatInt(int64(this.PageSize), 10))
	if this.ParentId > 0 {
		values.Set("parentId", strconv.FormatUint(this.ParentId, 10))
	}
	values.Set("catClass", strconv.FormatInt(int64(this.Class), 10))
	return values
}

type GetCategoriesResponse struct {
	Total      int        `json:"totalRows,omitempty"`
	Categories []Category `json:"categorys,omitempty"`
	Page       int        `json:"pageNo,omitempty"`
	Limit      int        `json:"pageSize,omitempty"`
}

func (this *GetCategoriesRequest) Get(clt *jdvop.Client) (*GetCategoriesResponse, error) {
	resp, err := clt.Do(this)
	if err != nil {
		return nil, err
	}
	var ret GetCategoriesResponse
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
