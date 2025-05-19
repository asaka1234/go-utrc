package go_utrc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 下单
func (cli *Client) Deposit(req UTRCDepositReq) (*UTRCDepositRsp, error) {

	//TODO 貌似没有任何鉴权
	
	rawURL := cli.DepositURL

	// Build URL with query parameters
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %v", err)
	}

	q := u.Query()
	q.Add("code", req.Code)
	q.Add("uid", req.UID)
	q.Add("currency_type", req.CurrencyType)
	q.Add("order_amount", req.OrderAmount)
	q.Add("order_id", req.OrderID)
	u.RawQuery = q.Encode()

	fullURL := u.String()

	// Make GET request
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse JSON response
	var depositRsp UTRCDepositRsp
	if err := json.Unmarshal(body, &depositRsp); err != nil {
		return nil, fmt.Errorf("failed to parse response JSON: %v", err)
	}

	return &depositRsp, nil
}
