package service
import (
	"Gooddoctor/model"
	"github.com/xormplus/xorm"
)

//存放用户模块基础功能接口定义及实现
func NewNroDoctorService(db *xorm.Engine) NroDoctorService{
	return &nrodoctorSevice{
		engine: db,
	}
}

type NroDoctorService interface {
	GetByDoctorPhoneNumAndPassword(DocPhoneNum,password string) (model.DocInfo,bool)
	GetNroDoctorCount() (int64,error)

}
type  nrodoctorSevice struct{
	engine *xorm.Engine
}

func (dc *nrodoctorSevice) GetNroDoctorCount() (int64,error){
	count,err := dc.engine.Count(new(model.DocInfo))
	if err!=nil{
		panic(err.Error())
		return 0, err
	}
	return count ,nil
}
//用户名和密码查询管理员
func (dc *nrodoctorSevice) GetByDoctorPhoneNumAndPassword(docphonenum,password string)(model.DocInfo,bool){
	var nrodoctor model.DocInfo
	dc.engine.Where("DocPhoneNum = ? and password = ?",docphonenum,password).Get(&nrodoctor)
	return nrodoctor,nrodoctor.DoctorId !=0
}