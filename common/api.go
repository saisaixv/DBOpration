package common


type BaseRsp struct {
	Error int `json:"errno,omitempty"`
}


type UserInfo struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type UserInfoA struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"passworfd"`
	MobilePhone string `json:"mobile_phone"`
	Company     string `json:"company"`
	ProductCode string `json:"product_code"`
	PicUrl      string `json:"pic_url"`
	IsComplete  int    `json:"is_complete,omitempty"`
}


type RegisterRsp struct {
	BaseRsp
}

type RegisterReq struct {
	Account string `json:"account"`
	Password string `json:"password"`
	UserName string `json:"user_name"`
	ProductCode string `json:"product_code"`
	CaptchaId string `json:"captcha_id"`
	InputValue string `json:"input_value"`
	MobilePhone string `json:"mobile_phone"`
}

type LoginRsp struct {
	BaseRsp
	UserId string `json:"user_id"`
	Token string `json:"token"`
}

type LoginReq struct {
	Account string `json:"account"`
	Password string `json:"password"`
	CaptchaId string `json:"captcha_id"`
	Inputvalue string `json:"input_value"`
}

type UserInfoRsp struct {
	BaseRsp
}

type KeepaliveRsp struct {
	BaseRsp
}

type ActiveApp struct {
	UserId string `json:"user_id"`
	Token string `json:"token"`
}
