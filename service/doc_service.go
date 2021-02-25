package service

import (
	"Gooddoctor/model"
	"github.com/xormplus/xorm"
)

//存放用户模块基础功能接口定义及实现
func NewDoctorService(db *xorm.Engine) DoctorService{
	return &doctorSevice{
		engine: db,
	}
}

type DoctorService interface {
	GetByDoctorPhoneNumAndPassword(DocPhoneNum,password string) (model.DocInfo,bool)
	GetDoctorCount() (int64,error)

}
type  doctorSevice struct{
	engine *xorm.Engine
}

func (dc *doctorSevice) GetDoctorCount() (int64,error){
	count,err := dc.engine.Count(new(model.DocInfo))
	if err!=nil{
		panic(err.Error())
		return 0, err
	}
	return count ,nil
}
//用户名和密码查询管理员
func (dc *doctorSevice) GetByDoctorPhoneNumAndPassword(docphonenum,password string)(model.DocInfo,bool){
	var doctor model.DocInfo
	dc.engine.Where("DocPhoneNum = ? and password = ?",docphonenum,password).Get(&doctor)
	return doctor,doctor.DoctorId !=0
}