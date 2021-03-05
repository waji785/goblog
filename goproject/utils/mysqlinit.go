package utils

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
)

func InitMysql() {

	fmt.Println("InitMysql....")
	driverName := beego.AppConfig.String("driverName")

	//注册数据库驱动
	orm.RegisterDriver(driverName, orm.DRMySQL)

	//数据库连接
	user := beego.AppConfig.String("root")
	pwd := beego.AppConfig.String("123456")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	//dbConn := "root:yu271400@tcp(127.0.0.1:3306)/cmsproject?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	db, _ = sql.Open(driverName, dbConn)
}
	//创建用户表
	func CreateTableWithUser() {
		sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`

		ModifyDB(sql)
	}
	func ModifyDB(sql string, args ...interface{}) (int64, error) {
		result, err := db.Exec(sql, args...)
		if err != nil {
			log.Println(err)
			return 0, err
		}
		count, err := result.RowsAffected()
		if err != nil {
			log.Println(err)
			return 0, err
		}
		return count, nil
		}
//查询
func QueryRowDB(sql string) *sql.Row{
	return db.QueryRow(sql)
}