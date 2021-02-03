package dbManger

import (
	"com.th.fileservice/config"
	"com.th.fileservice/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
	"strings"
)

var db *gorm.DB
var err error

//初始化数据库
func InitDB() {
	log.Println("数据库初始化")
	dbConfig := config.GetDBConfig()
	db, err = gorm.Open("mysql", dbConfig.Username+":"+dbConfig.Password+"@tcp("+dbConfig.Address+")/"+dbConfig.Dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	db.SingularTable(true)
	db.LogMode(true)
	if db.HasTable("file_info") {
		log.Println("执行增量脚本")
		excuteSqlScript(dbConfig.Sqlincrementpath)
	} else {
		log.Println("执行全量脚本")
		excuteSqlScript(dbConfig.Sqlinitpath)
	}
}

//加载SQL脚本
func excuteSqlScript(sqlpath string) {
	files := utils.GetFiles(sqlpath)
	for _, path := range files {
		file, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		str := string(file)
		str = strings.TrimSpace(str) //去除首尾空格
		sqlstr := strings.Split(str, ";")
		for _, sql := range sqlstr {
			if !strings.EqualFold(sql, "") {
				//不为空
				db.Exec(sql)
			}
		}

	}
}

func GetService() *gorm.DB {
	return db
}
