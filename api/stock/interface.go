package stock

type SkuNum struct {
	SkuId uint64 `json:"skuId,omitempty"`
	Num   int    `json:"num,omitempty"`
}

type SkuStock struct {
	SkuId     uint64 `json:"skuId,omitempty"`
	AreaId    string `json:"areaId,omitempty"`         // 入参时传入的区域码area。因京东目前是3、4级地址均支持，存在areaId在传入的3级地址后填充4级地址“_0“后返回的情况。
	StateId   uint64 `json:"stockStateId,omitempty"`   // 库存状态编号。参考枚举值：33,39,40,36,34,99
	StateDesc string `json:"stockStateDesc,omitempty"` // 库存状态描述。以下为stockStateId不同时，此字段不同的返回值：33 有货 现货-下单立即发货 39 有货 在途-正在内部配货，预计2~6天到达本仓库 40 有货 可配货-下单后从有货仓库配货 36 预订 34 无货 99 无货开预定，此时desc返回的数值代表预计到货天数，并且该功能需要依赖合同上有无货开预定权限的用户，到货周期略长，谨慎采用该功能。
	RemainNum int    `json:"remainNum,omitempty"`      // 剩余数量。当此值为-1时，为未查询到。StockStateDesc为33 或34（有货）：入参的skuNums字段，skuId对应的num<50，此字段为实际库存。入参的skuNums字段，skuId对应的50<=num<100，此字段为-1。入参的skuNums字段，skuId对应的num>100，此字段等于num。(此种情况并未返回真实京东库存)
}
