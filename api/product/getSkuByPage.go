package product

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type GetSkuByPageRequest struct {
	PageNum string `json:"pageNum"`
	PageNo  int    `json:"pageNo"`
}

func (this *GetSkuByPageRequest) Method() string {
	return "api/product/getSkuByPage"
}

func (this *GetSkuByPageRequest) Values() url.Values {
	values := url.Values{}
	values.Set("pageNum", this.PageNum)
	values.Set("pageNo", strconv.FormatInt(int64(this.PageNo), 10))
	return values
}

type GetSkuByPageResponse struct {
	PageCount int      `json:"pageCount"`
	SkuIds    []uint64 `json:"skuIds"`
}

func GetSkuByPage(clt *jdvop.Client, pageNum string, pageNo int) (*GetSkuByPageResponse, error) {
	resp, err := clt.Do(&GetSkuByPageRequest{PageNum: pageNum, PageNo: pageNo})
	if err != nil {
		return nil, err
	}
	var ret GetSkuByPageResponse
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
