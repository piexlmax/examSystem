package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"time"
)

// @title    CreateExaminationRecord
// @description   create a ExaminationRecord
// @param     examinationRecord                model.ExaminationRecord
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateExaminationRecord(examinationRecord model.ExaminationRecord) (err error) {
	var inExaminationRecord model.ExaminationRecord
	db := global.GVA_DB.Where("sys_user_id = ? AND test_paper_id = ?", examinationRecord.SysUserID, examinationRecord.TestPaperID).First(&inExaminationRecord)
	flag := db.RecordNotFound()
	if flag {
		err = global.GVA_DB.Create(&examinationRecord).Error
	} else {
		err = db.Update("agreement", examinationRecord.Agreement).Error
	}
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
	var inExaminationRecord model.ExaminationRecord
	err = global.GVA_DB.Where("id = ?", examinationRecord.ID).First(&inExaminationRecord).Update("score", examinationRecord.Score).Update("appraise", examinationRecord.Appraise).Error
	return err
}

// @title    GetExaminationRecord
// @description   get the info of a ExaminationRecord
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    ExaminationRecord         ExaminationRecord

func GetExaminationRecord(ResExaminationRecord request.ReqExaminationRecord) (err error, examinationRecord model.ExaminationRecord) {
	flag := global.GVA_DB.Where("sys_user_id = ? AND test_paper_id = ?", ResExaminationRecord.SysUserID, ResExaminationRecord.TestPaperID).First(&examinationRecord).RecordNotFound()
	if flag {
		return nil, examinationRecord
	} else {
		return err, examinationRecord
	}
}

// @title    GetExaminationRecord InfoList
// @description   get ExaminationRecord  list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetExaminationRecordInfoList(exa model.ExaminationRecord, info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	if exa.TestPaperID != 0 {
		db = db.Where("test_paper_id = ?", exa.TestPaperID)
	}
	if exa.SysUserID != 0 {
		db = db.Where("sys_user_id = ?", exa.SysUserID)
	}
	db = db.Order("test_paper_id desc")
	var examinationRecords []model.ExaminationRecord
	err = db.Find(&examinationRecords).Count(&total).Error
	err = db.Preload("SysUser").Preload("TestPaper").Limit(limit).Offset(offset).Find(&examinationRecords).Error
	return err, examinationRecords, total
}

func SubmitTest(sysUserId string, testPaperId string, file *multipart.FileHeader) (err error) {
	var testPaper model.TestPaper
	err = global.GVA_DB.Where("id = ?", testPaperId).First(&testPaper).Error
	if err != nil {
		return err
	}
	if !testPaper.TestPaperStatus {
		return errors.New("考官已关闭考试")
	}
	if time.Now().After(testPaper.TestPaperEndTime) {
		return errors.New("已不在考试时间之内")
	}
	var examinationRecord model.ExaminationRecord
	db := global.GVA_DB.Where("sys_user_id = ? AND test_paper_id = ?", sysUserId, testPaperId).First(&examinationRecord)
	if examinationRecord.TestPath != "" && !testPaper.TestPaperSubmitTimes {
		return errors.New("不允许多次提交")
	}
	if examinationRecord.TestPath != "" && testPaper.TestPaperSubmitTimes {
		os.Remove(examinationRecord.TestPath)
	}
	timeString := time.Now().Format("20060102150405")
	utils.CreateDir("./test-resource/record/" + testPaperId)
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	testPath := "./test-resource/record/" + testPaperId + "/" + timeString + "-" + file.Filename
	out, err := os.Create(testPath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)

	err = db.Update("test_path", testPath).Error
	return err
}

func OfflineAppraise(testPaperId string) (err error, path string) {
	files, _ := ioutil.ReadDir("./test-resource/record/" + testPaperId)
	var fileList []string
	for _, f := range files {
		fileList = append(fileList, "./test-resource/record/"+testPaperId+"/"+f.Name())
	}
	err = utils.ZipFiles("./kt.zip", fileList, ".", ".")
	if err != nil {
		return err, ""
	}
	path = "./kt.zip"
	return err, path
}
