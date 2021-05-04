package fast

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
)

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

const ORDERINFO = "order_info"

func (c *OrderInfo) TableName() string {
	return ORDERINFO
}

/**
** 获取短时间内的充值金额
 */
func GetRangeDateIncome(startTime, endTime string) float64 {
	o := orm.NewOrm()
	sum := 0.00
	err := o.Raw("select sum(order_amount) from order_info where status = ? and create_time >= ? and create_time <= ?", "success", startTime, endTime).QueryRow(&sum)
	if err != nil {
		logs.Error("获取短时间内金额失败，err：", err)
	}

	return sum
}
