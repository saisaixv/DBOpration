package db

import(
	"database/sql"
	"time"

	"DBOpration/utils"

	clog "github.com/cihub/seelog"
	_ "github.com/go-sql-driver/mysql"
)

func OpenDB(driver string,url string,maxLT int,maxOC int,maxIC int) (DBmysql *sql.DB) {
	var err error

	DBmysql,err=sql.Open(driver,url)
	if err!=nil{
		clog.Critical(err)
		panic(err)
	}

	utils.CheckErr(err,utils.CHECK_FLAG_EXIT)

	DBmysql.SetConnMaxLifetime(time.Duration(maxLT)*time.Second)
	DBmysql.SetMaxOpenConns(maxOC)
	DBmysql.SetMaxIdleConns(maxIC)

	clog.Info("[db opened]")
	
	return DBmysql

}

func CloseDB(DBmysql *sql.DB)  {
	DBmysql.Close()
	clog.Info("[db closed] mysql")
}

func DoQuery(DBmysql *sql.DB,sql string,args ...interface{}) (results [][]string,err error) {
	
	clog.Trace("[sql]: ",sql+" args:"+utils.Args2Str(args))

	rows,err:=DBmysql.Query(sql,args...)
	utils.CheckErr(err,utils.CHECK_FLAG_LOGONLY)
	if err!=nil{
		return nil,err
	}

	cols,_:=rows.Columns()
	values:=make([][]byte,len(cols))
	scans:=make([]interface{},len(cols))

	for i:=range values{
		scans[i]=&values[i]
	}
	results=make([][]string,0)

	i:=0
	for rows.Next(){
		err=rows.Scan(scans...)
		utils.CheckErr(err,utils.CHECK_FLAG_LOGONLY)
		row:=make([]string,0)
		for _,v:=range values{//每行数据是放在values里面，现在把它挪到row里
			row=append(row,string(v))
		}
		results=append(results,row)
		i++
	}

	rows.Close()
	return results,nil
}

func DoExec(DBmysql *sql.DB,sql string,args ...interface{}) (bool,error) {
	clog.Trace("[sql]: ",sql+" args:"+utils.Args2Str(args...))

	_,err:=DBmysql.Exec(sql,args...)
	utils.CheckErr(err,utils.CHECK_FLAG_LOGONLY)
	if err==nil{
		return true,err
	}else{
		return false,err
	}
}