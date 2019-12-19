/***************************************************
 ** @Desc : c file for ...
 ** @Time : 2019/8/21 16:51
 ** @Author : yuebin
 ** @File : delete
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/8/21 16:51
 ** @Software: GoLand
****************************************************/
package controllers

import (
	"github.com/astaxie/beego/logs"
	"juhe/service/models"
	"juhe/service/utils"
	"sort"
	"strings"
)

type Deletecontroller struct {
	BaseController
}

func (c *Deletecontroller) Finish() {
	remainderFirstMenuUid := make([]string, 0)
	remainderFirstMenu := make([]string, 0)
	remainderSecondMenuUid := make([]string, 0)
	remainderSecondMenu := make([]string, 0)
	remainderPowerId := make([]string, 0)
	remainderPower := make([]string, 0)
	allRoleInfo := models.GetRole()
	//如果有删除任何的东西，需要重新赋值权限
	for _, r := range allRoleInfo {
		for _, showFirstUid := range strings.Split(r.ShowFirstUid, "||") {
			if models.FirstMenuUidIsExists(showFirstUid) {
				remainderFirstMenuUid = append(remainderFirstMenuUid, showFirstUid)
				menuInfo := models.GetMenuInfoByMenuUid(showFirstUid)
				remainderFirstMenu = append(remainderFirstMenu, menuInfo.FirstMenu)
			}
		}
		for _, showSecondUid := range strings.Split(r.ShowSecondUid, "||") {
			if models.SecondMenuUidIsExists(showSecondUid) {
				remainderSecondMenuUid = append(remainderSecondMenuUid, showSecondUid)
				secondMenuInfo := models.GetSecondMenuInfoBySecondMenuUid(showSecondUid)
				remainderSecondMenu = append(remainderSecondMenu, secondMenuInfo.SecondMenu)
			}
		}
		for _, showPowerId := range strings.Split(r.ShowPowerUid, "||") {
			if models.PowerUidExists(showPowerId) {
				remainderPowerId = append(remainderPowerId, showPowerId)
				powerInfo := models.GetPowerById(showPowerId)
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
		models.UpdateRoleInfo(r)
	}
}

func (c *Deletecontroller) DeleteMenu() {
	menuUid := c.GetString("menuUid")
	menuInfo := models.GetMenuInfoByMenuUid(menuUid)
	dataJSON := new(BaseDataJSON)
	if menuInfo.MenuUid == "" {
		dataJSON.Msg = "不存在该菜单"
		dataJSON.Code = -1
	} else {
		logs.Info(c.GetSession("userID").(string) + "，执行了删除一级菜单操作")
		models.DeleteMenuInfo(menuUid)
		//删除该一级目下下的所有二级目录
		models.DeleteSecondMenuByFirstMenuUid(menuUid)
		SortFirstMenuOrder()
		dataJSON.Code = 200
	}
	c.Data["json"] = dataJSON
	c.ServeJSONP()
}

/*
* 对一级菜单重新进行排序
 */
func SortFirstMenuOrder() {
	menuInfoList := models.GetMenuAll()
	sort.Sort(models.MenuInfoSlice(menuInfoList))

	for i := 0; i < len(menuInfoList); i++ {
		m := menuInfoList[i]
		m.UpdateTime = utils.GetBasicDateTime()
		m.MenuOrder = i + 1
		models.UpdateMenuInfo(m)
		//对应的二级菜单也应该重新分配顺序号
		SortSecondMenuOrder(m)
	}
}

/*
* 对二级菜单分配顺序号
 */
func SortSecondMenuOrder(firstMenuInfo models.MenuInfo) {
	secondMenuInfoList := models.GetSecondMenuListByFirstMenuUid(firstMenuInfo.MenuUid)
	for _, sm := range secondMenuInfoList {
		sm.FirstMenuOrder = firstMenuInfo.MenuOrder
		sm.UpdateTime = utils.GetBasicDateTime()
		models.UpdateSecondMenu(sm)
		//删除下下一级的所有权限项
		models.DeletePowerBySecondUid(sm.SecondMenuUid)
	}
}

func (c *Deletecontroller) DeleteSecondMenu() {
	secondMenuUid := strings.TrimSpace(c.GetString("secondMenuUid"))
	secondMenuInfo := models.GetSecondMenuInfoBySecondMenuUid(secondMenuUid)
	dataJSON := new(BaseDataJSON)
	if secondMenuUid == "" || secondMenuInfo.SecondMenuUid == "" {
		dataJSON.Code = -1
		dataJSON.Msg = "该二级菜单不存在"
	} else {
		if models.DeleteSecondMenuBySecondMenuUid(secondMenuUid) {
			dataJSON.Code = 200
			ml := models.GetSecondMenuLenByFirstMenuUid(secondMenuInfo.FirstMenuUid)
			//删除该二级页面下的所有权限项
			models.DeletePowerBySecondUid(secondMenuUid)
			if ml == 0 {
				//如果该二级类目已经被全部删除，那么对应的一级类目也应当删除
				models.DeleteMenuInfo(secondMenuInfo.FirstMenuUid)
				SortFirstMenuOrder()
			} else {
				secondMenuInfoList := models.GetSecondMenuListByFirstMenuUid(secondMenuInfo.FirstMenuUid)
				sort.Sort(models.SecondMenuSlice(secondMenuInfoList))
				for i := 0; i < len(secondMenuInfoList); i++ {
					m := secondMenuInfoList[i]
					models.UpdateSecondMenuOrderBySecondUid(m.SecondMenuUid, i+1)
				}
			}
		} else {
			dataJSON.Code = -1
			dataJSON.Msg = "删除失败"
		}
	}
	c.Data["json"] = dataJSON
	c.ServeJSON()
}

/*
* 删除权限项
 */
func (c *Deletecontroller) DeletePowerItem() {
	powerID := strings.TrimSpace(c.GetString("powerID"))
	models.DeletePowerItemByPowerID(powerID)
	dataJSON := new(BaseDataJSON)
	dataJSON.Code = 200
	c.GenerateJSON(dataJSON)
}

/*
* 删除角色
 */
func (c *Deletecontroller) DeleteRole() {
	roleUid := strings.TrimSpace(c.GetString("roleUid"))
	dataJSON := new(BaseDataJSON)

	if models.DeleteRoleByRoleUid(roleUid) {
		dataJSON.Code = 200
	} else {
		dataJSON.Code = -1
	}
	c.GenerateJSON(dataJSON)
}

/*
* 删除操作员
 */
func (c *Deletecontroller) DeleteOperator() {
	userId := strings.TrimSpace(c.GetString("userId"))

	dataJSON := new(BaseDataJSON)

	if models.DeleteUserByUserId(userId) {
		dataJSON.Code = 200
	} else {
		dataJSON.Code = -1
	}

	c.GenerateJSON(dataJSON)
}

func (c *Deletecontroller) DeleteBankCardRecord() {
	uid := strings.TrimSpace(c.GetString("uid"))

	dataJSON := new(BankCardDataJSON)
	dataJSON.Code = -1

	if models.DeleteBankCardByUid(uid) {
		dataJSON.Code = 200
	}

	c.GenerateJSON(dataJSON)
}

/*
* 删除通道操作
 */
func (c *Deletecontroller) DeleteRoad() {
	roadUid := strings.TrimSpace(c.GetString("roadUid"))

	dataJSON := new(BaseDataJSON)
	dataJSON.Code = -1

	if models.DeleteRoadByRoadUid(roadUid) {
		dataJSON.Code = 200
	}
	params := make(map[string]string)
	roadPoolInfoList := models.GetAllRollPool(params)
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
		models.UpdateRoadPool(roadPoolInfo)
	}
	c.GenerateJSON(dataJSON)
}

/*
* 删除通道池
 */
func (c *Deletecontroller) DeleteRoadPool() {
	roadPoolCode := strings.TrimSpace(c.GetString("roadPoolCode"))

	dataJSON := new(BaseDataJSON)
	dataJSON.Code = -1

	if models.DeleteRoadPoolByCode(roadPoolCode) {
		dataJSON.Code = 200
	} else {
		dataJSON.Msg = "删除通道池失败"
	}
	c.GenerateJSON(dataJSON)
}

/*
* 删除商户
 */
func (c *Deletecontroller) DeleteMerchant() {
	merchantUid := strings.TrimSpace(c.GetString("merchantUid"))
	keyDataJSON := new(KeyDataJSON)
	if merchantUid == "" {
		keyDataJSON.Code = -1
		c.GenerateJSON(keyDataJSON)
		return
	}

	if models.DeleteMerchantByUid(merchantUid) {
		keyDataJSON.Code = 200
	} else {
		keyDataJSON.Code = -1
	}
	c.GenerateJSON(keyDataJSON)
}

/*
* 删除账户
 */
func (c *Deletecontroller) DeleteAccount() {
	accountUid := strings.TrimSpace(c.GetString("accountUid"))

	dataJSON := new(BaseDataJSON)
	models.IsExistByMerchantUid(accountUid)
	if models.IsExistByMerchantUid(accountUid) || models.IsExistByAgentUid(accountUid) {
		dataJSON.Code = -1
		dataJSON.Msg = "用户还存在，不能删除"
	} else {
		if models.DeleteAccountByUid(accountUid) {
			dataJSON.Code = 200
			dataJSON.Msg = "删除账户成功"
		} else {
			dataJSON.Code = -1
			dataJSON.Msg = "删除账户失败"
		}
	}

	c.GenerateJSON(dataJSON)
}

func (c *Deletecontroller) DeleteAgent() {
	agentUid := strings.TrimSpace(c.GetString("agentUid"))

	keyDataJSON := new(KeyDataJSON)
	//判断是否有商户还绑定了该代理
	if models.IsExistMerchantByAgentUid(agentUid) {
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "已有商户绑定改代理，不能删除"
	} else {
		if models.DeleteAgentByAgentUid(agentUid) {
			keyDataJSON.Code = 200
		} else {
			keyDataJSON.Code = -1
			keyDataJSON.Msg = "删除失败"
		}
	}

	c.GenerateJSON(keyDataJSON)
}

func (c *Deletecontroller) DeleteAgentRelation() {
	merchantUid := strings.TrimSpace(c.GetString("merchantUid"))

	merchantInfo := models.GetMerchantByUid(merchantUid)

	keyDataJSON := new(KeyDataJSON)
	keyDataJSON.Code = 200

	if merchantInfo.MerchantUid == "" {
		keyDataJSON.Code = -1
		keyDataJSON.Msg = "不存在这样的商户"
	} else {
		merchantInfo.UpdateTime = utils.GetBasicDateTime()
		merchantInfo.BelongAgentUid = ""
		merchantInfo.BelongAgentName = ""

		if !models.UpdateMerchant(merchantInfo) {
			keyDataJSON.Code = -1
			keyDataJSON.Msg = "更新商户失败"
		}
	}

	c.GenerateJSON(merchantInfo)
}
