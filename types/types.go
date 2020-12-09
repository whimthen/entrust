package types

// 账户类型
// 0: 现货账户
// 1: 逐仓杠杆
// 2: 全仓杠杆
type acctType struct {
	MainAccount       int
	LeverAccount      int
	CrossLeverAccount int
}

// 下单类型
// 0: 限价单
// 1: PostOnly
// 2: IOC
type orderType struct {
	Limit    int
	PostOnly int
	Ioc      int
}

// 交易类型
// 0: 限价卖
// 1: 限价买
// 2: PostOnly卖
// 3: PostOnly买
// 4: IOC卖
// 5: IOC买
type tradeType struct {
	Sell         int
	Buy          int
	PostOnlySell int
	PostOnlyBuy  int
	IocSell      int
	IocBuy       int
}

// 根据下单类型判断是否是买单
func (t tradeType) IsBuy(types int) bool {
	return (types & 1) == 1
}

// 根据下单类型判断是否是卖单
func (t tradeType) IsSell(types int) bool {
	return !t.IsBuy(types)
}

// 根据下单账户类型判断是否是杠杆账户(包含逐仓和全仓)
func (a acctType) IsLevel(acctType int) bool {
	return acctType == a.CrossLeverAccount || acctType == a.LeverAccount
}

var (
	AcctType = acctType{
		MainAccount:       mainAccount,
		LeverAccount:      leverAccount,
		CrossLeverAccount: crossLeverAccount,
	}
	OrderType = orderType{
		Limit:    limit,
		PostOnly: postOnly,
		Ioc:      ioc,
	}
	TradeType = tradeType{
		Sell:         sell,
		Buy:          buy,
		PostOnlySell: postOnlySell,
		PostOnlyBuy:  postOnlyBuy,
		IocSell:      iocSell,
		IocBuy:       iocBuy,
	}
)

const (
	sell = iota
	buy
	postOnlySell
	postOnlyBuy
	iocSell
	iocBuy
)

const (
	mainAccount = iota
	leverAccount
	crossLeverAccount
)

const (
	limit = iota
	postOnly
	ioc
)
