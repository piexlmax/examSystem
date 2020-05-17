package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateExaminationRecord
// @description   create a ExaminationRecord
// @param     examinationRecord                model.ExaminationRecord
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateExaminationRecord(examinationRecord model.ExaminationRecord) (err error) {
	err = global.GVA_DB.Create(&examinationRecord).Error
	return err
}

// @title    DeleteExaminationRecord
// @description   delete a ExaminationRecord
// @auth                     （2020/04/05  20:22）
// @param     examinationRecord                model.ExaminationRecord
// @return                    error

func DeleteExaminationRecord(examinationRecord model.ExaminationRecord) (err error) {
	err = global.GVA_DB.Delete(examinationRecord).Error
	return err
}

// @title    UpdateExaminationRecord
// @description   update a ExaminationRecord
// @param     examinationRecord           *model.ExaminationRecord
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateExaminationRecord(examinationRecord *model.ExaminationRecord) (err error) {
	err = global.GVA_DB.Save(examinationRecord).Error
	return err
}

// @title    GetExaminationRecord
// @description   get the info of a ExaminationRecord
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    ExaminationRecord         ExaminationRecord

func GetExaminationRecord(ResExaminationRecord request.ReqExaminationRecord) (err error, examinationRecord model.ExaminationRecord) {
	flag := global.GVA_DB.Where("sys_user_id = ? AND test_paper_id = ?", ResExaminationRecord.SysUserID,ResExaminationRecord.TestPaperID).First(&examinationRecord).RecordNotFound()
	if flag {
		return nil,examinationRecord
	} else {
		return err,examinationRecord
	}
}

// @title    GetExaminationRecord InfoList
// @description   get ExaminationRecord  list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetExaminationRecordInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var examinationRecords []model.ExaminationRecord
	err = db.Find(&examinationRecords).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&examinationRecords).Error
	return err, examinationRecords, total
}
