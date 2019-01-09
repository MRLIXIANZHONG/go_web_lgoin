package model
import (
	"log"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"web_test/pkg/setting"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"created_at"`
}

func init() {
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")//读取配置文件
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return tablePrefix + defaultTableName+"s";//前缀+表名+后缀
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)//设置闲置的连接数。
	db.DB().SetMaxOpenConns(100)//设置最大打开的连接数，默认值为0表示不限制
}
//其他包调用db
func DbInstance() *gorm.DB {
	return db
}

//关闭db
func CloseDB() {
	defer db.Close()
}
