package common

import(
	"DBOpration/utils"
	clog "github.com/cihub/seelog"
)


//缓存的key定义
const(
	//user_system
	PREFIX="us_"
	PREFIX_LOG="[cache key] : "


	KEY_ACTIVE_APP=PREFIX+"aa__"//激活assist
	KEY_ACTIVE_APP_REMOVE=PREFIX+"aar_"//assist被强制下线后，暂存下线前的对象


	KEY_USER_INFO              = PREFIX + "ui__" //登陆后缓存用户信息
	KEY_TOKEN                  = PREFIX + "tk__" //token
	KEY_USER_INFO_BY_EMAIL=PREFIX+"uiem_"//缓存用户信息（KEY：EMAIL）
	KEY_USER_INFO_BY_ID=PREFIX+"uiid_"//缓存用户信息（KEY：ID）
	KEY_REGISTER_TOKEN=PREFIX+"rt___"//用户注册时生成的token
	KEY_PASSWORD_RESET_TOKEN=PREFIX+"prt__"//重置密码申请时生成的token
	KEY_PRODUCTS=PREFIX+"pros_"//products
	KEY_USER_LIST=PREFIX+"ul___"//用户列表

)
//缓存超时时间设置
const(
	CYDEX_MANAGER_EX = utils.TIME_HOUR_TWO

	KEY_USER_INFO_EX              = utils.TIME_DAY_ONE * 7
	KEY_ACTIVE_APP_EX             = 20
	KEY_ACTIVE_APP_REMOVE_EX      = utils.TIME_MINUTE_ONE
	KEY_GET_PROJECTS_EX           = CYDEX_MANAGER_EX
	KEY_GET_PROJECT_EX            = utils.TIME_MINUTE_ONE
	KEY_GET_MEMBERS_EX            = CYDEX_MANAGER_EX
	KEY_GET_PERMISSION_OF_ROLE_EX = CYDEX_MANAGER_EX
	KEY_GET_PROJECT_USER_EX       = CYDEX_MANAGER_EX
	KEY_GET_SETTINGS_EVENTS_EX    = CYDEX_MANAGER_EX
	KEY_GET_DEFAULT_RECIPIENTS_EX = CYDEX_MANAGER_EX
	KEY_GET_GROUPS_EX             = CYDEX_MANAGER_EX
	KEY_GET_GROUP_EX              = CYDEX_MANAGER_EX
	KEY_GET_ZONES_EX              = CYDEX_MANAGER_EX
	KEY_INVITE_EX                 = utils.TIME_DAY_ONE
	KEY_GET_USER_INFO_BY_MAIL_EX  = utils.TIME_DAY_ONE * 7
	KEY_TOKEN_EX                  = utils.TIME_HOUR_TWO
	KEY_GET_ZONEA_EX              = CYDEX_MANAGER_EX
	KEY_LOGIN_ERR_CNT_EX          = utils.TIME_MINUTE_FIVE
	KEY_USER_QUOTA_REAL_EX        = utils.TIME_MINUTE_FIVE //缓存时间不能太长，结算程序每小时会重新结算
	KEY_LOCK_EX                   = 3                      //分布式锁的超时时间,单位：秒，默认3秒
	KEY_GET_SERVER_CONFIG_EX      = utils.TIME_DAY_ONE
	KEY_USER_QUOTA_EX             = utils.TIME_MINUTE_FIVE
	KEY_GET_PROJECT_LOG_EX        = utils.TIME_MINUTE_TEN
	KEY_TOKEN_PKG_DOWNLOAD_EX     = utils.TIME_DAY_ONE * 7
	KEY_QUERY_QUOTA_DETAIL_EX     = utils.TIME_MINUTE_FIVE
	KEY_ID_CODE_EX                = utils.TIME_MINUTE_TEN * 3
	KEY_URL_PKG_DOWNLOAD_EX       = utils.TIME_DAY_ONE * 7

	CA_KEY_LOCK_EX = 3 //分布式锁的超时时间,单位：秒，默认3秒

	USER_SYSTEM_EX = utils.TIME_HOUR_TWO

	KEY_USER_INFO_BY_EMAIL_EX   = utils.TIME_DAY_ONE * 7
	KEY_USER_INFO_BY_ID_EX      = utils.TIME_DAY_ONE * 7
	KEY_REGISTER_TOKEN_EX       = utils.TIME_DAY_ONE * 7
	KEY_PASSWORD_RESET_TOKEN_EX = utils.TIME_DAY_ONE
	KEY_PRODUCTS_EX             = utils.TIME_DAY_ONE
	KEY_USER_LIST_EX            = utils.TIME_DAY_ONE
)


func GetKeyUserInfoByEmail(account string) (string,int) {
	key:=KEY_USER_INFO_BY_EMAIL+account
	clog.Trace(PREFIX_LOG+key)
	return key,KEY_USER_INFO_BY_EMAIL_EX
}


func GetKeyUserInfoById(userId string) (string, int) {
	key := KEY_USER_INFO_BY_ID + userId
	clog.Trace(PREFIX_LOG + key)
	return key, KEY_USER_INFO_BY_ID_EX
}

func GetKeyRegisterToken(token string) (string, int) {
	key := KEY_REGISTER_TOKEN + token
	clog.Trace(PREFIX_LOG + key)
	return key, KEY_REGISTER_TOKEN_EX
}

func GetKeyPasswordResetToken(token string) (string, int) {
	key := KEY_PASSWORD_RESET_TOKEN + token
	clog.Trace(PREFIX_LOG + key)
	return key, KEY_PASSWORD_RESET_TOKEN_EX
}

func GetKeyProducts() (string, int) {
	key := KEY_PRODUCTS
	clog.Trace(PREFIX_LOG + key)
	return key, KEY_PRODUCTS_EX
}

func GetKeyUserList() (string, int) {
	key := KEY_USER_LIST
	clog.Trace(PREFIX_LOG + key)
	return key, KEY_USER_LIST_EX
}

func GetKeyToken(uid string) string {
	key := KEY_TOKEN + uid
	clog.Trace(PREFIX_LOG + key)
	return key
}

func GetKeyUserInfo(userId string) string {
	key := KEY_USER_INFO + userId
	clog.Trace(PREFIX_LOG + key)
	return key
}


func GetKeyActiveApp(userId string) string {
	key := KEY_ACTIVE_APP + userId
	//	clog.Trace(PREFIX_LOG + key)
	return key
}