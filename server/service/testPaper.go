package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"time"
)

// @title    CreateTestPaper
// @description   create a TestPaper
// @param     testPaper               model.TestPaper
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateTestPaper(testPaper model.TestPaper) (err error) {
	err = global.GVA_DB.Create(&testPaper).Error
	return err
}

// @title    DeleteTestPaper
// @description   delete a TestPaper
// @auth                     （2020/04/05  20:22）
// @param     testPaper               model.TestPaper
// @return                    error

func DeleteTestPaper(testPaper model.TestPaper) (err error) {
	if testPaper.TestPaperMould!=""{
		os.Remove(testPaper.TestPaperMould)
	}
	if testPaper.TestPaperSvg !=""{
		os.Remove(testPaper.TestPaperSvg)
	}
	err = global.GVA_DB.Delete(testPaper).Error
	return err
}

// @title    UpdateTestPaper
// @description   update a TestPaper
// @param     testPaper          *model.TestPaper
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateTestPaper(testPaper *model.TestPaper) (err error) {
	err = global.GVA_DB.Save(testPaper).Error
	return err
}

// @title    GetTestPaper
// @description   get the info of a TestPaper
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    TestPaper        TestPaper

func GetTestPaper(id uint) (err error, testPaper model.TestPaper) {
	err = global.GVA_DB.Where("id = ?", id).First(&testPaper).Error
	return
}

// @title    GetTestPaperInfoList
// @description   get TestPaper list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetTestPaperInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var testPapers []model.TestPaper
	err = db.Find(&testPapers).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&testPapers).Error
	return err, testPapers, total
}

func GetTestPaperSvg(ID uint) (err error, svg string) {
	err, testPaper := GetTestPaper(ID)
	if err != nil {
		return err, ""
	}
	fi, err := os.Open(testPaper.TestPaperSvg)
	if err != nil {
		return err, ""
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return nil, string(fd)

}

func GetTestPaperMould(ID uint) (err error, mould string) {
	err, testPaper := GetTestPaper(ID)
	if err != nil {
		return err, ""
	}
	return nil, string(testPaper.TestPaperMould)
}

func RemoveTestPaperFile(testPaper *model.TestPaper) (err error) {
	updateMap := map[string]string{}
	if testPaper.TestPaperSvg != "" {
		updateMap["test_paper_svg"] = ""
	}
	if testPaper.TestPaperMould != "" {
		updateMap["test_paper_mould"] = ""
	}
	db:= global.GVA_DB.Where("id = ?",testPaper.ID).First(testPaper)
	if db.RecordNotFound(){
		return nil
	}
	err = db.Updates(updateMap).Error
	if testPaper.TestPaperSvg != "" {
		err = os.Remove(testPaper.TestPaperSvg)
	}
	if testPaper.TestPaperMould != "" {
		err = os.Remove(testPaper.TestPaperMould)
	}
	return err
}

func FindAndCreateTestPaperSvgNode(node *model.TestPaperSvgNode)(err error,reNode model.TestPaperSvgNode){
	err = global.GVA_DB.Where("node_id = ? AND test_paper_id = ?",node.NodeId,node.TestPaperID).FirstOrCreate(node).Error
	return err,*node
}

func UploadTestPaperSvgNode(filetype string,nodeId string,testPaperID string,file *multipart.FileHeader)(err error){
	var node model.TestPaperSvgNode
	timeString := time.Now().Format("20060102150405")
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out, err := os.Create("./test-resource/node/"+filetype+"/"+ timeString +"-"+file.Filename)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	if err != nil {
		return err
	}
	updateMap:=map[string]string{}
	updateMap[filetype] = "./test-resource/node/"+filetype+"/"+ timeString +"-"+file.Filename
	err = global.GVA_DB.Where("node_id = ? AND test_paper_id = ?",nodeId,testPaperID).First(&node).Updates(updateMap).Error
	return err
}


func ClearTestPaperSvgNode(filetype string,nodeId string,testPaperID uint)(err error) {
	var node model.TestPaperSvgNode
	updateMap := map[string]string{}
	var path string
	db := global.GVA_DB.Where("node_id = ? AND test_paper_id = ?",nodeId,testPaperID).First(&node)
	switch filetype {
	case "flow":
		path = node.Flow
	case "log":
		path = node.Log
	case "configuration":
		path = node.Configuration
	case "sourceCode":
		path = node.SourceCode
	}
	updateMap[filetype] = ""
	err = db.Updates(updateMap).Error
	if err!=nil {
		return err
	}
	err = os.Remove(path)
	return err
}

func PublicTestPaper(testPaper *model.TestPaper)(err error){
	var testPapers []model.TestPaper
	var testPaperNode model.TestPaper
	err = global.GVA_DB.Find(&testPapers).Update("test_paper_status",false).Error
	if err!=nil{
		return err
	}
	err = global.GVA_DB.Where("id = ?",testPaper.ID).First(&testPaperNode).Update("test_paper_status",testPaper.TestPaperStatus).Error
	return err
}

func GetActiveTestPaper()(err error,testPaper model.TestPaper,status int){
	flag := global.GVA_DB.Where("test_paper_status = ?",true).First(&testPaper).RecordNotFound()
	if flag {
		return nil,testPaper,0
	}else{
		if testPaper.TestPaperStartTime.Before(time.Now()) && time.Now().Before(testPaper.TestPaperEndTime) {
			return nil,testPaper,1
		}else{
			return nil,testPaper,2
		}
	}
}