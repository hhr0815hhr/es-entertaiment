package mysql

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlStruct struct {
	Host         string
	Port         int
	User         string
	Pass         string
	Db           string
	MaxIdleConns int
	MaxOpenConns int
}

var (
	dbM, dbS *gorm.DB
	pool     map[string]*gorm.DB
	once     sync.Once
	err      error
)

func InitDB(master, slaver interface{}) {
	pool = make(map[string]*gorm.DB, 2)
	m := master.(MysqlStruct)
	s := slaver.(MysqlStruct)
	once.Do(func() {
		if master != (MysqlStruct{}) {
			connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", m.User, m.Pass, m.Host, m.Port, m.Db)
			dbM, err = gorm.Open("mysql", connStr)
			if err != nil {
				panic(err)
			}
			dbM.DB().SetMaxIdleConns(m.MaxIdleConns)
			dbM.DB().SetMaxOpenConns(m.MaxOpenConns)
			pool["master"] = dbM
		}
		if slaver != (MysqlStruct{}) {
			connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", s.User, s.Pass, s.Host, s.Port, s.Db)
			dbS, err = gorm.Open("mysql", connStr)
			if err != nil {
				panic(err)
			}
			dbS.DB().SetMaxIdleConns(s.MaxIdleConns)
			dbS.DB().SetMaxOpenConns(s.MaxOpenConns)
			pool["master"] = dbS
		}
	})
	fmt.Println("init db success...")
}

func GetDB(db string) *gorm.DB {
	return pool[db]
}
