package search

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type SearchRequest struct {
	Keyword    string `json:"keyword"` // 搜索关键字，需要编码
	CatId      string `json:"catId"`   // 分类Id,只支持三级类目Id
	PageIndex  int    `json:"pageIndex"`
	PageSize   int    `json:"pageSize"`
	Min        string `json:"min"`
	Max        string `json:"max"`
	Brands     string `json:"brands"`     // 品牌搜索 多个品牌以逗号分隔，需要编码
	Cid1       string `json:"cid1"`       // 一级分类
	Cid2       string `json:"cid2"`       // 二级分类
	SortType   string `json:"sortType"`   // 销量降序="sale_desc"; 价格升序="price_asc"; 价格降序="price_desc"; 上架时间降序="winsdate_desc"; 按销量排序_15天销售额="sort_totalsales15_desc"; 按15日销量排序="sort_days_15_qtty_desc"; 按30日销量排序="sort_days_30_qtty_desc"; 按15日销售额排序="sort_days_15_gmv_desc"; 按30日销售额排序="sort_days_30_gmv_desc";
	PriceCol   string `json:"priceCol"`   // 价格汇总  priceCol=”yes”
	ExtAttrCol string `json:"extAttrCol"` // 扩展属性汇总 extAttrCol=”yes”
}

func (this *SearchRequest) Method() string {
	return "api/search/search"
}

func (this *SearchRequest) Values() url.Values {
	values := url.Values{}
	values.Set("keyword", this.Keyword)
	if this.CatId != "" {
		values.Set("catId", this.CatId)
	}
	values.Set("pageIndex", strconv.FormatInt(int64(this.PageIndex), 10))
	values.Set("pageSize", strconv.FormatInt(int64(this.PageSize), 10))
	if this.Min != "" {
		values.Set("min", this.Min)
	}
	if this.Max != "" {
		values.Set("max", this.Max)
	}
	if this.Max != "" {
		values.Set("max", this.Max)
	}
	if this.Brands != "" {
		values.Set("brands", this.Brands)
	}
	if this.Cid1 != "" {
		values.Set("cid1", this.Cid1)
	}
	if this.Cid2 != "" {
		values.Set("cid2", this.Cid2)
	}
	if this.SortType != "" {
		values.Set("sortType", this.SortType)
	}
	if this.PriceCol != "" {
		values.Set("priceCol", this.PriceCol)
	}
	if this.ExtAttrCol != "" {
		values.Set("extAttrCol", this.ExtAttrCol)
	}
	return values
}

func (this *SearchRequest) Search(clt *jdvop.Client) (*SearchResult, error) {
	resp, err := clt.Do(this)
	if err != nil {
		return nil, err
	}
	var ret SearchResult
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
