package common

import(
	"database/sql"
)

var(
	DBmysqlUser *sql.DB
)

const(
	SPLIT_CHAR="_"

	//鉴权类型 当auth_type=1时，token=uid+"_"+唯一随机码
	//当auth——type=2时，token=access key +"_"+cur_time+"_"+摘要码
	AUTH_TYPE_LOGIN=1
	AUTH_TYPE_KEY=2
)