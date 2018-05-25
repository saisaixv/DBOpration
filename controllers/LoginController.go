package controllers

import(
	"DBOpration/common"
	"DBOpration/models"
	"DBOpration/utils"
	"DBOpration/utils/cache"

	clog "github.com/cihub/seelog"
)

type LoginController struct {
	BaseController
}


func (self *LoginController)Post()  {
	req:=new(common.LoginReq)
	rsp:=new(common.LoginRsp)


	rsp.Error=0

	defer func ()  {
		if err:=recover();err!=nil{
			rsp.Error=-1
			self.SetRspCode(500)
			clog.Error(err)
		}

		if rsp.Error==0{
			self.Data["json"]=rsp
		}else{
			rsp.Error=-1
			self.Data["json"]=rsp
		}

		self.ServeJSON()
	}()

	//检查header
	if self.Language==""{
		rsp.Error=-1
		return
	}

	//解析body
	if err:=self.FetchJsonBody(req);err!=nil{
		rsp.Error=-1
		return
	}

	if req.Account==""||req.Password==""{
		rsp.Error=-1
		return
	}

	retUserInfo,userInfo:=models.GetUserInfoEmailNoCache(req.Account)
	if retUserInfo==false{
		rsp.Error=-1
		self.SetRspCode(500)
		return
	}

	if userInfo.Password!=utils.Pbkdf2(req.Password){
		rsp.Error=-1
		self.SetRspCode(500)
		return
	}

	//生成token
	rsp.Token=userInfo.Id+utils.SPLIT_CHAR+utils.GetToken()
	rsp.UserId=userInfo.Id

	keyToken:=common.GetKeyToken(rsp.Token)
	// keyUserInfo:=common.GetKeyUserInfo(userInfo.Id)

	//缓存token，用以后续鉴权
	retToken:=cache.DoStrSet(keyToken,rsp.Token,common.KEY_TOKEN_EX)
	if retToken==false{
		rsp.Error=-1
		self.SetRspCode(500)
		return
	}

	//缓存用户信息

	//缓存token
	models.DoKeepAlive(rsp.Token,userInfo.Id)


}
