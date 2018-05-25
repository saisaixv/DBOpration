package models

import(

	"DBOpration/common"
	"DBOpration/utils"
	"DBOpration/utils/cache"
	"DBOpration/utils/db"
)

func SaveUser(userInfo *common.UserInfoA) bool {
	
	sql := `insert into us_user(user_id,user_name,password,email,create_at,mobile_phone,company,product_code,pic_url,is_complete)` +
		` values(?,?,?,?,?,?,?,?,?,?)`
	ret,_:=db.DoExec(common.DBmysqlUser,sql,
		userInfo.Id,
		userInfo.Name,
		utils.Pbkdf2(userInfo.Password),
		userInfo.Email,
		utils.GetNowUTC2(),
		userInfo.MobilePhone,
		userInfo.Company,
		userInfo.ProductCode,
		userInfo.PicUrl,
		1)

	return ret
}

func CacheUserInfo(userInfo *common.UserInfoA) bool {
	userInfo.Password=""
	redisKey,redisKeyEx:=common.GetKeyUserInfoByEmail(userInfo.Email)
	cache.DoSet(redisKey,userInfo,redisKeyEx)

	redisKey,redisKeyEx=common.GetKeyUserInfoById(userInfo.Id)
	cache.DoSet(redisKey,userInfo,redisKeyEx)

	return true
}

func GetUserInfoEmailNoCache(account string) (bool,*common.UserInfoA) {
	userInfo:=new(common.UserInfoA)
	sql := `select user_id,user_name,password,email,mobile_phone,company,product_code,pic_url,is_complete` +
		` from us_user` +
		` where email = ? `
		results,err:=db.DoQuery(common.DBmysqlUser,sql,account)

		if err==nil && len(results)>0{
			userInfo.Id = results[0][0]
		userInfo.Name = results[0][1]
		userInfo.Password = results[0][2]
		userInfo.Email = results[0][3]
		userInfo.MobilePhone = results[0][4]
		userInfo.Company = results[0][5]
		userInfo.ProductCode = results[0][6]
		userInfo.PicUrl = results[0][7]
		userInfo.IsComplete = utils.Atoi(results[0][8])
		return true, userInfo
	} else {
		return false, userInfo
	}
}

func GetUserInfo(userId string) (bool,*common.UserInfoA) {
	userInfo:=new(common.UserInfoA)
	sql := `select user_id,user_name,password,email,mobile_phone,company,product_code,pic_url,is_complete` +
		` from us_user` +
		` where user_id = ? `
	results,err:=db.DoQuery(common.DBmysqlUser,sql,userId)

	if err==nil && len(results)>0{
		userInfo.Id = results[0][0]
		userInfo.Name = results[0][1]
		userInfo.Password = results[0][2]
		userInfo.Email = results[0][3]
		userInfo.MobilePhone = results[0][4]
		userInfo.Company = results[0][5]
		userInfo.ProductCode = results[0][6]
		userInfo.PicUrl = results[0][7]
		userInfo.IsComplete = utils.Atoi(results[0][8])
		return true, userInfo
	} else {
		return false, userInfo
	}
}