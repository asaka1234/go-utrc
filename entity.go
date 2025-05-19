package go_utrc

// ----------pre generate-------------------------
type UTRCDepositReq struct {
	// System assigned merchant code
	Code string `json:"code"`
	// User ID
	UID string `json:"uid"`
	// Merchant deposit amount
	OrderAmount string `json:"orderAmount"`
	// Currency type
	CurrencyType string `json:"currencyType"`
	// Order ID
	OrderID string `json:"orderId"`
}

// --------------------------------------------------------
type UTRCDepositRsp struct {
	Msg  string           `json:"msg"`
	Code int              `json:"code"`
	Data *UTRCDepositData `json:"data"`
}

type UTRCDepositData struct {
	URL string `json:"url"`
}

// --------------------------------------------------------

type UTRCDepositCompleteBackReq struct {
	SysNo      string `json:"sys_no"`
	Uid        string `json:"uid"`
	Amount     string `json:"amount"`
	AmountUsdt string `json:"amount_usdt"`
	Sign       string `json:"sign"`
	Currency   string `json:"currency"`
	BillNo     string `json:"bill_no"`
}
