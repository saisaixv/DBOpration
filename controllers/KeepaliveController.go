package controllers

import(

	"fmt"

	"DBOpration/common"
	"DBOpration/models"
)

type KeepaliveController struct {
	BaseController
}

func (self *KeepaliveController)Get()  {
	rsp:=new(common.KeepaliveRsp)

	rsp.Error=0

	defer func ()  {
	
		if err:=recover();err!=nil{
			rsp.Error=-1
			self.SetRspCode(500)
		}

		if rsp.Error!=0{
			rsp.Error=-1
			self.SetRspCode(500)
		}

		self.Data["json"]=rsp

		self.ServeJSON()
	}()

	//检查header
	if self.Language==""||self.Token==""{
		fmt.Println("请求头不能为空")
		rsp.Error=-1
		self.SetRspCode(500)
		return
	}

	ret:=models.RefreshKeepAlive(self.Token,self.UserId)

	if ret!=0{
		fmt.Println("刷新token失败")
		rsp.Error=-1
	}
}