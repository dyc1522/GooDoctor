package controller

import (
	"Gooddoctor/service"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type DoctorController struct {
	Ctx iris.Context
	//doctor功能实体
	Service service.DoctorService

}
func (dc *DoctorController) PostLogin(context iris.Context) mvc.Result{
	iris.New().Logger().Info("doctor login")
	var doctorLogin DocLogin
	dc.Ctx.ReadJSON(&doctorLogin)

	if doctorLogin.DocPhoneNum == "" || doctorLogin.Password == ""{
		return 	mvc.Response{
			Object: map[string]interface{}{
				"status": "0",
				"success":"登陆失败",
				"message":"用户名或密码为空，请重新填写后尝试登录",
			},
		}
	}
	doctor,exist := dc.Service.GetByDoctorPhoneNumAndPassword(doctorLogin.DocPhone,doctorLogin,Password)
	if !exist{
		return mvc.Response{
			Object: map[string]interface{}{
			"status": "0",
			"success":"登陆失败",
			"message":"用户名或密码错误，请重新填写后登录",
			},
		}
	}
	docByte,_:=json.Marshal(doctor)
	dc.Session.Set(DOCTOR,docByte)
	return mvc.Response{
		Object: map[string]interface{}{
			"status": "1",
			"success":"登陆成功",
			"message":"医生登陆成功，欢迎您",
		},
	}



}