package search

type SearchResult struct {
	Total      int                `json:"resultCount,omitempty"`
	PageCount  int                `json:"pageCount,omitempty"`
	PageSize   int                `json:"pageSize,omitempty"`
	PageIndex  int                `json:"pageIndex,omitempty"`
	Brands     *BrandAggregate    `json:"brandAggregate,omitempty"`
	Categories *CategoryAggregate `json:"categoryAggregate,omitempty"`
	Prices     []PriceInterval    `json:"priceIntervalAggregate,omitempty"`
	Hits       []HitResult        `json:"hitResult,omitempty"`
}

type BrandAggregate struct {
	Pinyin []string `json:"pinyinAggr,omitempty"`
	Brands []Brand  `json:"brandList,omitempty"`
}

type Brand struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Pinyin string `json:"pinyin,omitempty"`
}

type PriceInterval struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type CategoryAggregate struct {
	First  []Category `json:"firstCategory,omitempty"`
	Second []Category `json:"secondaryCategory,omitempty"`
	Third  []Category `json:"thirdCategory,omitempty"`
}

type Category struct {
	Id     uint64 `json:"catId,omitempty"`
	Count  int    `json:"count,omitempty"`
	Name   string `json:"name,omitempty"`
	Weight int    `json:"weight,omitempty"`
}

type HitResult struct {
	Brand    string `json:"brand,omitempty"`
	ImageUrl string `json:"imageUrl,omitempty"`
	WareName string `json:"wareName,omitempty"`
	WareId   string `json:"wareId,omitempty"`
	WarePid  string `json:"warePId,omitempty"`
	BrandId  string `json:"brandId,omitempty"`
	Cid1     string `json:"cid1,omitempty"`
	Cid2     string `json:"cid2,omitempty"`
	CatId    string `json:"catId,omitempty"`
	WState   string `json:"wstate,omitempty"` // 上柜状态，1.有效
	Wyn      string `json:"wyn,omitempty"`    // 商品状态，1.有效
	Cid1Name string `json:"cid1Name,omitempty"`
	Cid2Name string `json:"cid2Name,omitempty"`
	CatName  string `json:"catName,omitempty"`
	Synonyms string `json:"synonyms,omitempty"` // 商品同义词
}
