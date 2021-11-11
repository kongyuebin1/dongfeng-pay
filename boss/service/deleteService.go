package service

import (
	"boss/datas"
	"boss/models/accounts"
	"boss/models/agent"
	"boss/models/merchant"
	"boss/models/road"
	"boss/models/system"
	"boss/models/user"
	"boss/utils"
	"github.com/beego/beego/v2/core/logs"
	"sort"
	"strings"
)

type DeleteService struct {
}

func (c *DeleteService) Finish() {
	remainderFirstMenuUid := make([]string, 0)
	remainderFirstMenu := make([]string, 0)
	remainderSecondMenuUid := make([]string, 0)
	remainderSecondMenu := make([]string, 0)
	remainderPowerId := make([]string, 0)
	remainderPower := make([]string, 0)
	allRoleInfo := system.GetRole()
	//如果有删除任何的东西，需要重新赋值权限
	for _, r := range allRoleInfo {
		for _, showFirstUid := range strings.Split(r.ShowFirstUid, "||") {
			if system.FirstMenuUidIsExists(showFirstUid) {
				remainderFirstMenuUid = append(remainderFirstMenuUid, showFirstUid)
				menuInfo := system.GetMenuInfoByMenuUid(showFirstUid)
				remainderFirstMenu = append(remainderFirstMenu, menuInfo.FirstMenu)
			}
		}
		for _, showSecondUid := range strings.Split(r.ShowSecondUid, "||") {
			if system.SecondMenuUidIsExists(showSecondUid) {
				remainderSecondMenuUid = append(remainderSecondMenuUid, showSecondUid)
				secondMenuInfo := system.GetSecondMenuInfoBySecondMenuUid(showSecondUid)
				remainderSecondMenu = append(remainderSecondMenu, secondMenuInfo.SecondMenu)
			}
		}
		for _, showPowerId := range strings.Split(r.ShowPowerUid, "||") {
			if system.PowerUidExists(showPowerId) {
				remainderPowerId = append(remainderPowerId, showPowerId)
				powerInfo := system.GetPowerById(showPowerId)
				remainderPower = append(remainderPower, powerInfo.PowerItem)
			}
		}
		r.ShowFirstUid = strings.Join(remainderFirstMenuUid, "||")
		r.ShowFirstMenu = strings.Join(remainderFirstMenu, "||")
		r.ShowSecondUid = strings.Join(remainderSecondMenuUid, "||")
		r.ShowSecondMenu = strings.Join(remainderSecondMenu, "||")
		r.ShowPowerUid = strings.Join(remainderPowerId, "||")
		r.ShowPower = strings.Join(remainderPower, "||")
		r.UpdateTime = utils.GetBasicDateTime()
		system.UpdateRoleInfo(r)
	}
}

func (c *DeleteService) DeleteMenu(menuUid, userID string) *datas.BaseDataJSON {

	dataJSON := new(datas.BaseDataJSON)
	menuInfo := system.GetMenuInfoByMenuUid(menuUid)
	if menuInfo.MenuUid == "" {
		dataJSON.Msg = "不存在该菜单"
		dataJSON.Code = -1
	} else {
		logs.Info(userID + "，执行了删除一级菜单操作")
		system.DeleteMenuInfo(menuUid)
		//删除该一级目下下的所有二级目录
		system.DeleteSecondMenuByFirstMenuUid(menuUid)
		SortFirstMenuOrder()
		dataJSON.Code = 200
	}

	return dataJSON
}
func (c *DeleteService) DeleteSecondMenu(secondMenuUid string) *datas.BaseDataJSON {

	secondMenuInfo := system.GetSecondMenuInfoBySecondMenuUid(secondMenuUid)
	dataJSON := new(datas.BaseDataJSON)
	if secondMenuUid == "" || secondMenuInfo.SecondMenuUid == "" {
		dataJSON.Code = -1
		dataJSON.Msg = "该二级菜单不存在"
	} else {
		if system.DeleteSecondMenuBySecondMenuUid(secondMenuUid) {
			dataJSON.Code = 200
			ml := system.GetSecondMenuLenByFirstMenuUid(secondMenuInfo.FirstMenuUid)
			//删除该二级页面下的所有权限项
			system.DeletePowerBySecondUid(secondMenuUid)
			if ml == 0 {
				//如果该二级类目已经被全部删除，那么对应的一级类目也应当删除
				system.DeleteMenuInfo(secondMenuInfo.FirstMenuUid)
				SortFirstMenuOrder()
			} else {
				secondMenuInfoList := system.GetSecondMenuListByFirstMenuUid(secondMenuInfo.FirstMenuUid)
				sort.Sort(system.SecondMenuSlice(secondMenuInfoList))
				for i := 0; i < len(secondMenuInfoList); i++ {
					m := secondMenuInfoList[i]
					system.UpdateSecondMenuOrderBySecondUid(m.SecondMenuUid, i+1)
				}
			}
		} else {
			dataJSON.Code = -1
			dataJSON.Msg = "删除失败"
		}
	}
	return dataJSON
}

func (c *DeleteService) DeletePowerItem(powerID string) *datas.BaseDataJSON {
	system.DeletePowerItemByPowerID(powerID)
	dataJSON := new(datas.BaseDataJSON)
	dataJSON.Code = 200
	return dataJSON
}

func (c *DeleteService) DeleteRole(roleUid string) *datas.BaseDataJSON {
	dataJSON := new(datas.BaseDataJSON)

	if system.DeleteRoleByRoleUid(roleUid) {
		dataJSON.Code = 200
	} else {
		dataJSON.Code = -1
	}
	return dataJSON
}

func (c *DeleteService) DeleteOperator(userId string) *datas.BaseDataJSON {

	dataJSON := new(datas.BaseDataJSON)

	if user.DeleteUserByUserId(userId) {
		dataJSON.Code = 200
	} else {
		dataJSON.Code = -1
	}
	return dataJSON
}

func (c *DeleteService) DeleteBankCardRecord(uid string) *datas.BankCardDataJSON {

	dataJSON := new(datas.BankCardDataJSON)
	dataJSON.Code = -1

	if system.DeleteBankCardByUid(uid) {
		dataJSON.Code = 200
	}
	return dataJSON
}

func (c *DeleteService) DeleteRoad(roadUid string) *datas.BaseDataJSON {

	dataJSON := new(datas.BaseDataJSON)
	dataJSON.Code = -1

	if road.DeleteRoadByRoadUid(roadUid) {
		dataJSON.Code = 200
	}
	params := make(map[string]string)
	roadPoolInfoList := road.GetAllRollPool(params)
	//将轮询池中的对应的通道删除
	for _, roadPoolInfo := range roadPoolInfoList {
		var uids []string
		roadInfoList := strings.Split(roadPoolInfo.RoadUidPool, "||")
		for _, uid := range roadInfoList {
			if uid != roadUid {
				uids = append(uids, uid)
			}
		}
		roadPoolInfo.RoadUidPool = strings.Join(uids, "||")
		roadPoolInfo.UpdateTime = utils.GetBasicDateTime()
		road.UpdateRoadPool(roadPoolInfo)
	}
	return dataJSON
}

func (c *DeleteService) DeleteRoadPool(roadPoolCode string) *datas.BaseDataJSON {
	dataJSON := new(datas.BaseDataJSON)
	dataJSON.Code = -1

	if road.DeleteRoadPoolByCode(roadPoolCode) {
		dataJSON.Code = 200
	} else {
		dataJSON.Msg = "删除通道池失败"
	}
	return dataJSON
}

func (c *DeleteService) DeleteMerchant(merchantUid string) *datas.KeyDataJSON {
	keyDataJSON := new(datas.KeyDataJSON)
	if merchantUid == "" {
		keyDataJSON.Code = -1
		return keyDataJSON
	}

	if merchant.DeleteMerchantByUid(merchantUid) {
		keyDataJSON.Code = 200
	} else {
		keyDataJSON.Code = -1
	}
	return keyDataJSON
}

func (c *DeleteService) DeleteAccount(accountUid string) *datas.BaseDataJSON {
	dataJSON := new(datas.BaseDataJSON)
	merchant.IsExistByMerchantUid(accountUid)
	if merchant.IsExistByMerchantUid(accountUid) || agent.IsExistByAgentUid(accountUid) {
		dataJSON.Code = -1
		dataJSON.Msg = "用户还存在，不能删除"
	} else {
		if accounts.DeleteAccountByUid(accountUid) {
			dataJSON.Code = 200
			dataJSON.Msg = "删除账户成功"
		} else {
			dataJSON.Code = -1
			dataJSON.Msg = "删除账户失败"
		}
	}
	return dataJSON
}

func (c *DeleteService) DeleteAgent(agentUid string) *datas.KeyDataJSON {
	keyDataJSON := new(datas.KeyDataJSON)
	//判断是否有商户还绑定了该代理
	if merchant.IsExistMerchantByAgentUid(agentUid) {
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "已有商户绑定改代理，不能删除"
	} else {
		if agent.DeleteAgentByAgentUid(agentUid) {
			keyDataJSON.Code = 200
		} else {
			keyDataJSON.Code = -1
			keyDataJSON.Msg = "删除失败"
		}
	}
	return keyDataJSON
}
func (c *DeleteService) DeleteAgentRelation(merchantUid string) *datas.KeyDataJSON {

	merchantInfo := merchant.GetMerchantByUid(merchantUid)

	keyDataJSON := new(datas.KeyDataJSON)
	keyDataJSON.Code = 200

	if merchantInfo.MerchantUid == "" {
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "不存在这样的商户"
	} else {
		merchantInfo.UpdateTime = utils.GetBasicDateTime()
		merchantInfo.BelongAgentUid = ""
		merchantInfo.BelongAgentName = ""

		if !merchant.UpdateMerchant(merchantInfo) {
			keyDataJSON.Code = -1
			keyDataJSON.Msg = "更新商户失败"
		}
	}
	return keyDataJSON
}

/*
* 对一级菜单重新进行排序
 */
func SortFirstMenuOrder() {
	menuInfoList := system.GetMenuAll()
	sort.Sort(system.MenuInfoSlice(menuInfoList))

	for i := 0; i < len(menuInfoList); i++ {
		m := menuInfoList[i]
		m.UpdateTime = utils.GetBasicDateTime()
		m.MenuOrder = i + 1
		system.UpdateMenuInfo(m)
		//对应的二级菜单也应该重新分配顺序号
		SortSecondMenuOrder(m)
	}
}

/*
* 对二级菜单分配顺序号
 */
func SortSecondMenuOrder(firstMenuInfo system.MenuInfo) {
	secondMenuInfoList := system.GetSecondMenuListByFirstMenuUid(firstMenuInfo.MenuUid)
	for _, sm := range secondMenuInfoList {
		sm.FirstMenuOrder = firstMenuInfo.MenuOrder
		sm.UpdateTime = utils.GetBasicDateTime()
		system.UpdateSecondMenu(sm)
		//删除下下一级的所有权限项
		system.DeletePowerBySecondUid(sm.SecondMenuUid)
	}
}
