package controllers

import(
	"DBOpration/common"
	"DBOpration/models"

	clog "github.com/cihub/seelog"
)

type GetUserInfoController struct {
	BaseController
}

func (self *GetUserInfoController)Get()  {
	rsp:=new(common.UserInfoRsp)
	ret:=new(common.UserInfoA)

	rsp.Error=0

	defer func ()  {
		if err:=recover();err!=nil{
			rsp.Error=-1
			self.SetRspCode(500)
			clog.Error(err)
		}

		if rsp.Error==0{
			self.Data["json"]=ret
		}else{
			rspErr:=new(common.BaseRsp)
			rspErr.Error=rsp.Error
			self.Data["json"]=rspErr
		}

		self.ServeJSON()
	}()

	//检查header
	if self.Token==""||self.Language==""{
		rsp.Error=-1
		self.SetRspCode(500)
		return
	}

	
	retUserInfo,userInfo:=models.GetUserInfo(self.UserId)
	if retUserInfo==false{
		rsp.Error=-1
		self.SetRspCode(500)
		return
	}

	ret.Id=userInfo.Id
	ret.Name=userInfo.Name
	ret.Email=userInfo.Email
	ret.Password=userInfo.Password
	ret.MobilePhone=userInfo.MobilePhone
	ret.Company=userInfo.Company
	ret.IsComplete=userInfo.IsComplete
	ret.PicUrl=userInfo.PicUrl
	ret.ProductCode=userInfo.ProductCode


}