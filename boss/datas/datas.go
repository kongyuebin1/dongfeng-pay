package datas

import (
	"boss/models/accounts"
	"boss/models/agent"
	"boss/models/merchant"
	"boss/models/order"
	"boss/models/payfor"
	"boss/models/road"
	"boss/models/system"
	"boss/models/user"
)

type BaseDataJSON struct {
	Msg  string
	Code int
}

type KeyDataJSON struct {
	Msg  string
	Code int
	Key  string
}

type MenuDataJSON struct {
	StartIndex   int
	DisplayCount int
	CurrentPage  int
	TotalPage    int
	Code         int
	Msg          string
	MenuList     []system.MenuInfo
}

type SecondMenuDataJSON struct {
	StartIndex     int
	DisplayCount   int
	CurrentPage    int
	TotalPage      int
	Code           int
	Msg            string
	SecondMenuList []system.SecondMenuInfo
}

type PowerItemDataJSON struct {
	StartIndex    int
	DisplayCount  int
	CurrentPage   int
	TotalPage     int
	Code          int
	Msg           string
	PowerItemList []system.PowerInfo
}

type RoleInfoDataJSON struct {
	StartIndex   int
	DisplayCount int
	CurrentPage  int
	TotalPage    int
	Code         int
	Msg          string
	RoleInfoList []system.RoleInfo
}

type DeployTreeJSON struct {
	Msg               string
	Code              int
	Key               string
	AllFirstMenu      []system.MenuInfo
	ShowFirstMenuUid  map[string]bool
	AllSecondMenu     []system.SecondMenuInfo
	ShowSecondMenuUid map[string]bool
	AllPower          []system.PowerInfo
	ShowPowerUid      map[string]bool
}

type OperatorDataJSON struct {
	Msg          string
	Code         int
	StartIndex   int
	DisplayCount int
	CurrentPage  int
	TotalPage    int
	OperatorList []user.UserInfo
}

type EditOperatorDataJSON struct {
	Code         int
	Msg          string
	OperatorList []user.UserInfo
	RoleList     []system.RoleInfo
}

type BankCardDataJSON struct {
	Msg              string
	Code             int
	StartIndex       int
	DisplayCount     int
	CurrentPage      int
	TotalPage        int
	BankCardInfoList []system.BankCardInfo
}

type RoadDataJSON struct {
	Msg          string
	Code         int
	StartIndex   int
	DisplayCount int
	CurrentPage  int
	TotalPage    int
	RoadInfoList []road.RoadInfo
	RoadPool     road.RoadPoolInfo
}

type RoadPoolDataJSON struct {
	Msg              string
	Code             int
	StartIndex       int
	DisplayCount     int
	CurrentPage      int
	TotalPage        int
	RoadPoolInfoList []road.RoadPoolInfo
}

type MerchantDataJSON struct {
	Msg          string
	Code         int
	StartIndex   int
	DisplayCount int
	CurrentPage  int
	TotalPage    int
	MerchantList []merchant.MerchantInfo
}

type MerchantDeployDataJSON struct {
	Code           int
	Msg            string
	MerchantDeploy merchant.MerchantDeployInfo
}

type AccountDataJSON struct {
	Msg          string
	Code         int
	StartIndex   int
	DisplayCount int
	CurrentPage  int
	TotalPage    int
	AccountList  []accounts.AccountInfo
}

type AccountHistoryDataJSON struct {
	Msg                string
	Code               int
	StartIndex         int
	DisplayCount       int
	CurrentPage        int
	TotalPage          int
	AccountHistoryList []accounts.AccountHistoryInfo
}

type AgentDataJSON struct {
	Msg          string
	Code         int
	StartIndex   int
	DisplayCount int
	CurrentPage  int
	TotalPage    int
	AgentList    []agent.AgentInfo
}

type ProductDataJSON struct {
	Msg        string
	Code       int
	ProductMap map[string]string
}

type OrderDataJSON struct {
	Msg          string
	Code         int
	StartIndex   int
	DisplayCount int
	CurrentPage  int
	TotalPage    int
	OrderList    []order.OrderInfo
	AllAmount    float64
	SuccessRate  string
	NotifyUrl    string
}

type ListDataJSON struct {
	Msg            string
	Code           int
	StartIndex     int
	DisplayCount   int
	CurrentPage    int
	TotalPage      int
	List           []order.OrderProfitInfo
	AllAmount      float64
	SupplierProfit float64
	AgentProfit    float64
	PlatformProfit float64
}

type PayForDataJSON struct {
	Msg          string
	Code         int
	StartIndex   int
	DisplayCount int
	CurrentPage  int
	TotalPage    int
	PayForList   []payfor.PayforInfo
}

type BalanceDataJSON struct {
	Msg     string
	Code    int
	Balance float64
}

type NotifyBankOrderIdListJSON struct {
	Msg          string
	Code         int
	NotifyIdList []string
}

type ProfitListJSON struct {
	TotalAmount         float64
	PlatformTotalProfit float64
	AgentTotalProfit    float64
	Msg                 string
	Code                int
	ProfitList          []order.PlatformProfit
}
