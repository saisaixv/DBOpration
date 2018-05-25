package models

import(

	"fmt"

	"DBOpration/common"
	"DBOpration/utils/cache"
)

func DoKeepAlive(token string,userId string) int  {
	redisKey:=common.GetKeyActiveApp(userId)
	redisKeyEx:=common.KEY_ACTIVE_APP_EX

	activeAppTmp:=new(common.ActiveApp)
	activeAppTmp.Token=token
	activeAppTmp.UserId=userId

	fmt.Println("key = "+redisKey)
	ret:=cache.DoSet(redisKey,activeAppTmp,redisKeyEx)

	if ret==false{
		fmt.Println("缓存token失败")
		return -1
	}
	fmt.Println("缓存token成功")
	return 0
}

func RefreshKeepAlive(token string,userId string) int {
	redisKey:=common.GetKeyActiveApp(userId)
	redisKeyEx:=common.KEY_ACTIVE_APP_EX

	activeAppTmp:=new(common.ActiveApp)

	fmt.Println("key = "+redisKey)
	ret:=cache.DoGet(redisKey,activeAppTmp)

	if ret==false{
		fmt.Println("获取缓存token失败")
		return -1
	}

	if activeAppTmp.Token==token{
		ret:=cache.DoExpire(redisKey,redisKeyEx)
		if ret==false{
			fmt.Println("刷新token过期时间失败")
			return -2
		}
		return 0
	}
	return -1
}