package checkorder

type Order struct {
	Id           uint64 `json:"jdOrderId,omitempty"`
	State        uint   `json:"state,omitempty"`        // 订单状态 0 是新建  1是妥投   2是拒收
	HangUpState  uint   `json:"hangUpState,omitempty"`  // 是否挂起   0为为挂起    1为挂起
	InvoiceState uint   `json:"invoiceState,omitempty"` // 开票方式(1为随货开票，0为订单预借，2为集中开票 )
	Time         string `json:"time,omitempty"`         // 订单创建时间，格式：yyyy-MM-dd HH:mm:ss
}

type CheckOrderResult struct {
	Total     int     `json:"total,omitempty"`
	TotalPage int     `json:"totalPage,omitempty"`
	CurPage   int     `json:"curPage,omitempty"`
	Orders    []Order `json:"order,omitempty"`
}
