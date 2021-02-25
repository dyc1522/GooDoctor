package model

type DocInfo struct {
	DoctorId       int64  `xorm:"pk autoincr" 	json:"doctor_id"`
	DoctorName     string `xorm:"varchar(32)" 	json:"doctor_name"`
	DoctorAge      int64  `xorm:"default 1" 	json:"doctor_age"`
	DoctorSex      int64  `xorm:"default 0" 	json:"doctor_sex"`
	DocEduBacG     string `xorm:"varchar(16)" 	json:"doc_edu_bac_g"`
	DocPhoneNum    string `xorm:"varchar(11)" 	json:"doc_phone_num"`
	DocLecturer    string `xorm:"varchar(32)" 	json:"doc_lecturer"`
	DocAddress     string `xorm:"varchar(64) 	json:"doc_address""`
	IsStoreManager bool   `xorm:"bool" 		json:"is_store_manager"`
	Password       string `xorm:"varchar(16) json:"password""`
}
