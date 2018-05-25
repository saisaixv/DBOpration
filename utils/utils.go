package utils

import(
	"os"
	"runtime"
	"strconv"
	"fmt"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha1"
	"time"

	clog "github.com/cihub/seelog"
	"gopkg.in/mgo.v2/bson"
	"github.com/dchest/pbkdf2"
	"github.com/satori/go.uuid"
)

//普通常量
const(

	REPORT_ID_CYDEX_DAY  = "001" //CYDEX 按天统计报表
	REPORT_ID_CYDEX_HOUR = "002" //CYDEX 按小时统计报表

	REDIS_DB_STATISTICS_USER_INFO_DT = 1
	REDIS_DB_STATISTICS_CYDEX_DAY    = 2
	REDIS_DB_STATISTICS_CYDEX_HOUR   = 3
	REDIS_DB_DATA_SERVICE            = 4
	REDIS_DB_TRANSFER_NODE           = 5
	REDIS_DB_NOTIFY_NODE             = 6
	REDIS_DB_TRANSFER_SERVICE        = 7
	REDIS_DB_USER_SYSTEM             = 8
	REDIS_DB_CYDEX_MANAGER           = 9
	REDIS_DB_ZONE                    = 10
	REDIS_DB_OPERATION_MAINTENANCE   = 11
	REDIS_DB_CALCULATE               = 12

	REDIS_NM_SERVER_B = 20
	REDIS_NM_CLIENT   = 21

	REDIS_DB_SL_CYDEX_MANAGER = 29

	PRODUCT_CODE_CYDEX     = "001"
	PRODUCT_CODE_WEB       = "002" //官网
	PRODUCT_CODE_CATON_NET = "003"


	//checkErr 函数用到的常量 1表示记录日志并退出 2表示仅记录日志，程序不退出
	CHECK_FLAG_EXIT = 1
	CHECK_FLAG_LOGONLY=2

	//分隔符
	SPLIT_CHAR="_"
	//注释常量
	COMMENT_STR="=========================="


	TIME_MINUTE_ONE=60
	TIME_MINUTE_FIVE=5*60
	TIME_MINUTE_TEN=10*60
	TIME_HOUR_ONE    = 60 * 60
	TIME_HOUR_TWO    = 2 * 60 * 60
	TIME_DAY_ONE     = 24 * 60 * 60
	TIME_MAX         = TIME_DAY_ONE * 365 * 100

	TIME_CACHE = TIME_HOUR_TWO

	//月份 1,01,Jan,January
	//日　 2,02,_2
	//时　 3,03,15,PM,pm,AM,am
	//分　 4,04
	//秒　 5,05
	//年　 06,2006
	//周几 Mon,Monday
	//时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
	//时区字母缩写 MST

	// 日期格式 YYYY-MM-DD
	DATE_FORMAT_SHORT = "2006-01-02"
	// 时间格式 YYYY-MM-DD HH:MM:SS
	DATE_FORMAT_LONG = "2006-01-02 15:04:05"
)

//CheckErr错误处理函数，
//程序中错误分为两种
//一种需要终止程序
//另一种仅仅是记录错误日志
func CheckErr(err error,flag int)  {
	var path string

	if err!=nil{
		_,file,line,_:=runtime.Caller(1)
		path="--"+file+":"+strconv.Itoa(line)

		switch flag {
		case CHECK_FLAG_EXIT:
			clog.Critical(err.Error()+path)
			clog.Critical(StackTrace(false))
			panic(err)
		case CHECK_FLAG_LOGONLY:
			clog.Critical(err.Error()+path)
			clog.Critical(StackTrace(false))
		default:
			clog.Info(err.Error()+path)
		}
	}
}

func StackTrace(all bool) string {
	buf:=make([]byte,10240)

	for{
		size:=runtime.Stack(buf,all)
		//the size of the buffer may be not enough to hold the stacktrace,
		//so double the buffer size
		if size==len(buf){
			buf=make([]byte,len(buf)<<1)
			continue
		}
		break
	}
	return string(buf)
}

//FileIsExist 检查文件或目录是否存在，如果有filename指定的文件或目录存在则返回true
func FileIsExist(filename string) bool {
	_,err:=os.Stat(filename)
	return err==nil || os.IsExist(err)
}

//Args2Str 把不定长参数转成字符串，主要用以打印日志
func Args2Str(args ...interface{}) (ret string) {
	split:=","
	for _,v:=range args{
		switch v.(type) {
		case string:
			ret=ret+v.(string)+split
		case int:
			ret=ret+Itoa(v.(int))+split
		case int64:
			ret=ret+Itoa64(v.(int64))+split
		default:
			ret=ret+fmt.Sprintf("%T",v)+split
		}
	}
	return
}

func Atoi(str string) int {
	if str == "" {
		return 0
	}
	i, err := strconv.Atoi(str)
	CheckErr(err, CHECK_FLAG_LOGONLY)
	return i
}
func Atoi64(str string) int64 {
	if str == "" {
		return 0
	}
	i, err := strconv.ParseInt(str, 10, 64)
	CheckErr(err, CHECK_FLAG_LOGONLY)
	return i
}
func Atof64(str string) float64 {
	if str == "" {
		return 0.0
	}
	i, err := strconv.ParseFloat(str, 64)
	CheckErr(err, CHECK_FLAG_LOGONLY)
	return i
}

func Itof64(i int) float64 {
	a := Itoa(i)
	f := Atof64(a)
	return f
}

// Ftoa64 转字符串 prec：保留几位小数
func F64toa(f float64, prec int) string {
	str := strconv.FormatFloat(f, 'f', prec, 64)
	return str
}

// Ftoa64 转int，如有小数，则抛弃
func F64toi(f float64) int {
	y := int(f)
	return y
}

func Itoa(i int) string {
	str := strconv.Itoa(i)
	return str
}
func Itoa64(i int64) string {
	str := strconv.FormatInt(i, 10)
	return str
}

// ItoaZero 数字转字符串，前面补零
func ItoaZero(i int, lenth int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(lenth)+"d", i)
}

// ItoaZero64 数字转字符串，前面补零
func ItoaZero64(i int64, lenth int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(lenth)+"d", i)
}
func Btoa(b bool) string {
	str := strconv.FormatBool(b)
	return str
}

// Itob int -> bool
func Itob(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

// Itob int -> bool
func Atob(i string) bool {
	if i == "1" {
		return true
	} else {
		return false
	}
}

//GetMongoObjectId 取得mongo objectid，24位
func GetMongoObjectId() string {
	ret:=bson.NewObjectId()
	return ret.Hex()
}

func Pbkdf2(str string) string {
	password := str
	//Get random salt
	salt:=make([]byte,32)
	if _,err:=rand.Reader.Read(salt);err!=nil{
		panic("random reader failed")
	}

	salt=[]byte("ae260d66ed648a7ffb6d9286b080ee91")
	//Derive key
	key:=pbkdf2.WithHMAC(sha256.New,[]byte(password),salt,9999,16)
	return fmt.Sprintf("%x",key)
}


// GetNow 取得当前日期时间
func GetNow(format string) string {
	return time.Now().Format(format)
}
func GetNowUTC(format string) string {
	return time.Now().UTC().Format(format)
}
func GetNowUTC2() string {
	return time.Now().UTC().Format(DATE_FORMAT_LONG)
}
func GetNowUTC2Num() int64 {
	return Str2TimeStampL(GetNowUTC2())
}


// Time2StrL long time -> yyyy-mm-dd hh:mm:ss
func Time2StrL(t time.Time) string {
	return t.Format(DATE_FORMAT_LONG)
}

// Time2StrS short time -> yyyy-mm-dd
func Time2StrS(t time.Time) string {
	return t.Format(DATE_FORMAT_SHORT)
}

// TimeStamp2StrL long time -> yyyy-mm-dd hh:mm:ss
func TimeStamp2StrL(ts int64) string {
	t := time.Unix(ts, 0).UTC()
	return t.Format(DATE_FORMAT_LONG)
}

// TimeStamp2StrS short time -> yyyy-mm-dd
func TimeStamp2StrS(ts int64) string {
	t := time.Unix(ts, 0).UTC()
	return t.Format(DATE_FORMAT_SHORT)
}

// Str2TimeL yyyy-mm-dd hh:mm:ss -> time
func Str2TimeL(s string) time.Time {
	loc, _ := time.LoadLocation("UTC")
	t, err := time.ParseInLocation(DATE_FORMAT_LONG, s, loc)
	CheckErr(err, CHECK_FLAG_LOGONLY)
	return t
}

// Str2TimeS yyyy-mm-dd -> time
func Str2TimeS(s string) time.Time {
	loc, _ := time.LoadLocation("UTC")
	t, err := time.ParseInLocation(DATE_FORMAT_SHORT, s, loc)
	CheckErr(err, CHECK_FLAG_LOGONLY)
	return t
}

// Str2TimeStampL yyyy-mm-dd hh:mm:ss -> timeStamp
func Str2TimeStampL(s string) int64 {
	loc, _ := time.LoadLocation("UTC")
	t, err := time.ParseInLocation(DATE_FORMAT_LONG, s, loc)
	CheckErr(err, CHECK_FLAG_LOGONLY)
	return t.Unix()
}

// Str2TimeStampS yyyy-mm-dd -> timeStamp
func Str2TimeStampS(s string) int64 {
	loc, _ := time.LoadLocation("UTC")
	t, err := time.ParseInLocation(DATE_FORMAT_SHORT, s, loc)
	CheckErr(err, CHECK_FLAG_LOGONLY)
	return t.Unix()
}


// DtDiff 计算两个时间差 ret = timeB - timeA ,参数：yyyy-mm-dd hh:mm:ss 返回值：time.Duration
func DtDiff(timeA string, timeB string) (ret time.Duration) {
	dtA, _ := time.Parse(DATE_FORMAT_LONG, timeA)
	dtB, _ := time.Parse(DATE_FORMAT_LONG, timeB)
	d := dtB.Sub(dtA)
	return d
}


// GetToken 生成token
func GetToken() string {
	// Creating UUID Version 4
	u1,_ := uuid.NewV4()

	return Sha1(u1.String())
}

func Sha1(str string) string {
	sum := sha1.Sum([]byte(str))
	return fmt.Sprintf("%x", sum)
}
