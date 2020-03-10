package area

type Location struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type Address struct {
	NationId   string `json:"nationId,omitempty"`
	Nation     string `json:"nation,omitempty"`
	ProvinceId uint64 `json:"provinceId,omitempty"`
	Province   string `json:"province,omitempty"`
	CityId     uint64 `json:"cityId,omitempty"`
	City       string `json:"city,omitempty"`
	CountyId   uint64 `json:"countyId,omitempty"`
	County     string `json:"county,omitempty"`
	TownId     uint64 `json:"townId,omitempty"`
	Town       string `json:"town,omitempty"`
}
