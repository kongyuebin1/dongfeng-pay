package fast

type OrderInfo struct {
	Id              string `orm:"pk;column(id)"`
	MerchantOrderId string
	ShopName        string
	OrderPeriod     string
	BankOrderId     string
	BankTransId     string
	OrderAmount     float64
	ShowAmount      float64
	FactAmount      float64
	RollPoolCode    string
	RollPoolName    string
	RoadUid         string
	RoadName        string
	PayProductCode  string
	PayProductName  string
	PayTypeCode     string
	PayTypeName     string
	OsType          string
	Status          string
	Refund          string
	RefundTime      string
	Freeze          string
	FreezeTime      string
	Unfreeze        string
	UnfreezeTime    string
	ReturnUrl       string
	NotifyUrl       string
	MerchantUid     string
	MerchantName    string
	AgentUid        string
	AgentName       string
	Response        string
	UpdateTime      string
	CreateTime      string
}
