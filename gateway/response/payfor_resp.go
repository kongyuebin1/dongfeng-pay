package response

/**
* 返回自动代付结果
 */
type PayForResponse struct {
	ResultCode      string `json:"resultCode,omitempty"`
	ResultMsg       string `json:"resultMsg,omitempty"`
	MerchantOrderId string `json:"merchantOrderId,omitempty"`
	SettAmount      string `json:"settAmount,omitempty"`
	SettFee         string `json:"settFee,omitempty"`
	Sign            string `json:"sign,omitempty"`
}

/**
* 返回商户代付结果查询结果
 */
type PayForQueryResponse struct {
	ResultMsg       string `json:"resultMsg,omitempty"`
	MerchantOrderId string `json:"merchantOrderId,omitempty"`
	SettAmount      string `json:"settAmount,omitempty"`
	SettFee         string `json:"settFee,omitempty"`
	SettStatus      string `json:"settStatus,omitempty"`
	Sign            string `json:"sign,omitempty"`
}

/**
* 返回商户查询余额结果
 */
type BalanceResponse struct {
	ResultCode      string `json:"resultCode,omitempty"`
	Balance         string `json:"balance,omitempty"`
	AvailableAmount string `json:"availableAmount,omitempty"`
	FreezeAmount    string `json:"freezeAmount,omitempty"`
	WaitAmount      string `json:"waitAmount,omitempty"`
	LoanAmount      string `json:"loanAmount,omitempty"`
	PayforAmount    string `json:"payforAmount,omitempty"`
	ResultMsg       string `json:"resultMsg,omitempty"`
	Sign            string `json:"sign,omitempty"`
}
