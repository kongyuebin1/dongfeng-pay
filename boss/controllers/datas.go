package controllers

import "boss/models"

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
	MenuList     []models.MenuInfo
}

type SecondMenuDataJSON struct {
	StartIndex     int
	DisplayCount   int
	CurrentPage    int
	TotalPage      int
	Code           int
	Msg            string
	SecondMenuList []models.SecondMenuInfo
}

type PowerItemDataJSON struct {
	StartIndex    int
	DisplayCount  int
	CurrentPage   int
	TotalPage     int
	Code          int
	Msg           string
	PowerItemList []models.PowerInfo
}

type RoleInfoDataJSON struct {
	StartIndex   int
	DisplayCount int
	CurrentPage  int
	TotalPage    int
	Code         int
	Msg          string
	RoleInfoList []models.RoleInfo
}

type DeployTreeJSON struct {
	Msg               string
	Code              int
	Key               string
	AllFirstMenu      []models.MenuInfo
	ShowFirstMenuUid  map[string]bool
	AllSecondMenu     []models.SecondMenuInfo
	ShowSecondMenuUid map[string]bool
	AllPower          []models.PowerInfo
	ShowPowerUid      map[string]bool
}

type OperatorDataJSON struct {
	Msg          string
	Code         int
	StartIndex   int
	DisplayCount int
	CurrentPage  int
	TotalPage    int
	OperatorList []models.UserInfo
}

type EditOperatorDataJSON struct {
	Code         int
	Msg          string
	OperatorList []models.UserInfo
	RoleList     []models.RoleInfo
}

type BankCardDataJSON struct {
	Msg              string
	Code             int
	StartIndex       int
	DisplayCount     int
	CurrentPage      int
	TotalPage        int
	BankCardInfoList []models.BankCardInfo
}

type RoadDataJSON struct {
	Msg          string
	Code         int
	StartIndex   int
	DisplayCount int
	CurrentPage  int
	TotalPage    int
	RoadInfoList []models.RoadInfo
	RoadPool     models.RoadPoolInfo
}

type RoadPoolDataJSON struct {
	Msg              string
	Code             int
	StartIndex       int
	DisplayCount     int
	CurrentPage      int
	TotalPage        int
	RoadPoolInfoList []models.RoadPoolInfo
}

type MerchantDataJSON struct {
	Msg          string
	Code         int
	StartIndex   int
	DisplayCount int
	CurrentPage  int
	TotalPage    int
	MerchantList []models.MerchantInfo
}

type MerchantDeployDataJSON struct {
	Code           int
	Msg            string
	MerchantDeploy models.MerchantDeployInfo
}

type AccountDataJSON struct {
	Msg          string
	Code         int
	StartIndex   int
	DisplayCount int
	CurrentPage  int
	TotalPage    int
	AccountList  []models.AccountInfo
}

type AccountHistoryDataJSON struct {
	Msg                string
	Code               int
	StartIndex         int
	DisplayCount       int
	CurrentPage        int
	TotalPage          int
	AccountHistoryList []models.AccountHistoryInfo
}

type AgentDataJSON struct {
	Msg          string
	Code         int
	StartIndex   int
	DisplayCount int
	CurrentPage  int
	TotalPage    int
	AgentList    []models.AgentInfo
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
	OrderList    []models.OrderInfo
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
	List           []models.OrderProfitInfo
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
	PayForList   []models.PayforInfo
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
	ProfitList          []models.PlatformProfit
}
