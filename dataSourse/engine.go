package dataSourse

//数据库引擎
import (
	"Gooddoctor/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

func NewMySqlEngine() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(localhost:3306)/goodoctor?charset=utf8")
	err = engine.Sync2(new(
		model.DocInfo),
		new(model.DocInfo))
	if err != nil {
		panic(err.Error())
	}
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(10)
	return engine
}
