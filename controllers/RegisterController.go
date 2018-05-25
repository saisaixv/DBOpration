package controllers

import(
	"strings"

	"DBOpration/common"
	"DBOpration/utils"
	"DBOpration/models"

	clog "github.com/cihub/seelog"
)

type RegisterController struct {
	BaseController
}


func (self *RegisterController)Post()  {
	req:=new(common.RegisterReq)
	rsp:=new(common.RegisterRsp)
	ret:=new(common.UserInfo)

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

	//检查请求头
	if self.Language==""{
		rsp.Error=-1
		return
	}

	//获取post参数
	if err:=self.FetchJsonBody(req);err!=nil{
		rsp.Error=-1
		self.SetRspCode(400)
		return
	}

	userId:=utils.GetMongoObjectId()
	userInfo:=new(common.UserInfoA)
	userInfo.Id=userId
	userInfo.Email=req.Account
	userInfo.Name=req.UserName
	userInfo.Password=req.Password
	userInfo.ProductCode=req.ProductCode

	//save
	retSave:=models.SaveUser(userInfo)
	if retSave==false{
		rsp.Error=-1
		self.SetRspCode(500)
		return
	}

	//缓存
	models.CacheUserInfo(userInfo)

	ret.Id=userInfo.Id
	ret.Name=userInfo.Name

}


func checkHttpHead(language string,timeZone int) bool {
	if language==""||timeZone!=-8{
		return false
	}

	if strings.ToUpper(language)!="ZH-CN" && 
		strings.ToUpper(language)!="ZH-TW" &&
		strings.ToUpper(language) !="EN"{
			
		return false
	}
	return true
}