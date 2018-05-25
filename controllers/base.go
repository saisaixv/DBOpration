package controllers

import(

	"io/ioutil"
	"encoding/json"
	"strings"
	"strconv"

	"DBOpration/utils"
	"DBOpration/common"

	"github.com/astaxie/beego"
	clog "github.com/cihub/seelog"
)

var (
	ShowReqBody =false
	ShowRspBody=false
)

type BaseController struct {
	beego.Controller
	Token string
	AuthType int
	UserId string
	Language string
	TimeZone int
	RealIP string
	NeedLog bool//是否需要打印调试日志
}

func (self *BaseController)Prepare()  {
	
	//解析请求头
	self.fetchAuthType()
	self.fetchToken()
	self.fetchLanguage()
	self.fetchTimeZone()
	if self.Token != "" {
		self.UserId = GetRootUserIdByToken(self.AuthType, self.Token)
	}
	self.fetchRealIP()
	self.NeedLog = true

}

func (self *BaseController)fetchRealIP()  {
	self.RealIP=self.Ctx.Input.Header("X-Real-IP")
}

func GetRootUserIdByToken(authType int,token string) (rootUserId string) {
	rootUserId=""
	defer func ()  {
		if err:=recover();err!=nil{
			clog.Error(err.(string))
		}
	}()

	if authType==common.AUTH_TYPE_LOGIN{
		keys:=strings.Split(token,common.SPLIT_CHAR)
		if len(keys)!=2{
			return ""
		}

		rootUserId=keys[0]
	}else if authType ==common.AUTH_TYPE_KEY{
		keys:=strings.Split(token,common.SPLIT_CHAR)
		if len(keys)!=4{
			return ""
		}
		rootUserId=keys[0]
	}else{
		clog.Error("出现未定义鉴权类型。auth type："+utils.Itoa(authType))
		return ""
	}

	return rootUserId
}

func (self *BaseController)fetchAuthType()  {
	ul:=self.Ctx.Input.Header("x-us-authtype")
	if ul!=""{
		i,err:=strconv.Atoi(ul)
		if err==nil{
			self.AuthType=i
		}
	}
}

func (self *BaseController)fetchToken()  {
	token:=self.Ctx.Input.Header("x-us-token")
	self.Token=token
}

func (self *BaseController)fetchLanguage()  {
	language:=self.Ctx.Input.Header("accept-language")	
	self.Language=language
}

func (self *BaseController)fetchTimeZone()  {
	timeZone:=self.Ctx.Input.Header("time-zone")
	self.TimeZone=utils.Atoi(timeZone)
}

func (self *BaseController)FetchJsonBody(v interface{}) error {
	r:=self.Ctx.Request

	defer r.Body.Close()
	body,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		clog.Error(err)
		return err
	}

	if len(body)==0{
		return nil
	}

	if ShowReqBody{
		clog.Trace("[http req]:",string(body))
	}

	if err:=json.Unmarshal(body,v);err!=nil{
		clog.Error(err)
		return err
	}
	return nil
}

func (self *BaseController)SetRspCode(code int)  {
	self.Ctx.ResponseWriter.WriteHeader(code)
}

