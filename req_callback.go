package go_utrc

// 充值回调
func (cli *Client) CashierCallback(req UTRCDepositCompleteBackReq, processor func(UTRCDepositCompleteBackReq) error) error {
	//验证签名
	//TODO

	//开始处理
	return processor(req)
}
