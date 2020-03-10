package price

type SellPrice struct {
	SkuId       uint64  `json:"skuId,omitempty"`
	Price       float64 `json:"price,omitempty"`       // 京东销售价。当此值为-1时，为未查询到。
	JdPrice     float64 `json:"jdPrice,omitempty"`     // 京东价。当此值为-1时，为未查询到。
	MarketPrice float64 `json:"marketPrice,omitempty"` // 入参中的queryExts中包含marketPrice时，输出此字段。京东的前台划线价。现在只有图书频道能露出，其他的因政策原因已不允许展示。此值为-1时，为未查询到。
	Tax         float64 `json:"tax,omitempty"`         // 税率。当queryExts中包含containsTax时，出参中有此字段。例如：此值为16时，代表税率为“16%”。
	TaxPrice    float64 `json:"taxPrice,omitempty"`    // 税额。当queryExts中包含containsTax时，出参中有此字段。
	NakedPrice  float64 `json:"nakedPrice,omitempty"`  // 未税价。（当queryExts中包含nakedPrice或containsTax时，出参中有此字段）当此值为-1时，为未查询到。
}

type UnionBalance struct {
	Balance *Balance `json:"balance,omitempty"` // 当type入参中包含1时，此对象出现
	Geious  *Geious  `json:"geious,omitempty"`  // 当type入参中包含2时，此对象出现
}

type Balance struct {
	Pin         string  `json:"pin,omitempty"`
	RemainLimit float64 `json:"remainLimit,omitempty"`
}

type Geious struct {
	Pin           string  `json:"pin,omitempty"`
	PenaltySumAmt float64 `json:"penaltySumAmt,omitempty"` // 金采的违约金金额。
	CreditLimit   float64 `json:"creditLimit,omitempty"`   // 金采的总授信额度。
	DebtSumAmt    float64 `json:"debtSumAmt,omitempty"`    // 金采的待还款额度。
	RemainLimit   float64 `json:"remainLimit,omitempty"`   // 金采账户余额。
}

type BalanceDetails struct {
	Total     int             `json:"total,omitempty"`
	PageSize  int             `json:"pageSize,omitempty"`
	PageNo    int             `json:"pageNo,omitempty"`
	PageCount int             `json:"pageCount,omitempty"`
	Data      []BalanceDetail `json:"data,omitempty"`
}

type BalanceDetail struct {
	Id            uint64  `json:"id,omitempty"`          // 余额明细ID
	AccountType   uint    `json:"accountType,omitempty"` // 账户类型  1：可用余额 2：锁定余额
	Amount        float64 `json:"amount,omitempty"`      // 金额（元），有正负，可以是零，表示订单流程变化，如退款时会先有一条退款申请的记录，金额为0
	Pin           string  `json:"pin,omitempty"`
	OrderId       string  `json:"orderId,omitempty"`
	TradeType     int     `json:"tradeType,omitempty"`     // 业务类型
	TradeTypeName string  `json:"tradeTypeName,omitempty"` // 业务类型名称
	CreatedDate   string  `json:"createdDate,omitempty"`   // 余额变动日期
	NotePub       string  `json:"notePub,omitempty"`       // 备注信息
	TradeNo       uint64  `json:"tradeNo,omitempty"`       // 业务号，一般由余额系统，在每一次操作成功后自动生成，也可以由前端业务系统传入
}
