package main

import (
	"Gooddoctor/controller"
	"Gooddoctor/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"time"

)

func main() {

}
func newApp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")
	//注册静态资源
	app.StaticWeb("/static", "./static")
	app.StaticWeb("/manage/static", "./static")
	app.StaticWeb("/img", "./static/img")
	//注册视图文件
	app.RegisterView(iris.HTML("./static", ".html"))
	app.Get("/", func(context context.Context) {
		context.View("index.html")
	})
	return app
}

//mvc 架构模式处理
func mvcHandle(app *iris.Application{
	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessioncookie",
		Expires: 24 * time.Hour,
	})
	engine := datasource.NewMySqlEngine()
	//店主医生模块
	doctorService := service.NewDoctorService(engine)
	doctor := mvc.New(app.Party("/doctor"))
	doctor.Register(
		doctorService,
		sessManager.Start)
	doctor.Handle(new(controller.DoctorController))
	//店员医生模块
	nrodocService := service.NewNroDoctorService(engine)
	nrodoc := mvc.New(app.Party("/nrodocs"))
	nrodoc.Register(
		nrodocService,
		sessManager.Start,
		)



})
