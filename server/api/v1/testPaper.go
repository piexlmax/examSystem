package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"time"
)


// @Tags TestPaper
// @Summary 创建TestPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestPaper true "创建TestPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testPaper/createTestPaper [post]
func CreateTestPaper(c *gin.Context) {
	var testPaper model.TestPaper
	_ = c.ShouldBindJSON(&testPaper)
	err := service.CreateTestPaper(testPaper)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}


// @Tags TestPaper
// @Summary 删除TestPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestPaper true "删除TestPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testPaper/deleteTestPaper [delete]
func DeleteTestPaper(c *gin.Context) {
	var testPaper model.TestPaper
	_ = c.ShouldBindJSON(&testPaper)
	err := service.DeleteTestPaper(testPaper)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}


// @Tags TestPaper
// @Summary 更新TestPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestPaper true "更新TestPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testPaper/updateTestPaper [put]
func UpdateTestPaper(c *gin.Context) {
	var testPaper model.TestPaper
	_ = c.ShouldBindJSON(&testPaper)
	err := service.UpdateTestPaper(&testPaper)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}


// @Tags TestPaper
// @Summary 用id查询TestPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestPaper true "用id查询TestPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testPaper/findTestPaper [get]
func FindTestPaper(c *gin.Context) {
	var testPaper model.TestPaper
	_ = c.ShouldBindQuery(&testPaper)
	err,retestPaper := service.GetTestPaper(testPaper.ID)
	if err != nil {
	response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData( gin.H{"retestPaper":retestPaper,}, c)
	}
}


// @Tags TestPaper
// @Summary 分页获取TestPaper列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TestPaper列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testPaper/getTestPaperList [get]
func GetTestPaperList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetTestPaperInfoList(pageInfo)
	if err != nil {
	    response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
        response.OkWithData(resp.PageResult{
                    List:     list,
                    Total:    total,
                    Page:     pageInfo.Page,
                    PageSize: pageInfo.PageSize,
                }, c)
	}
}


// 上传svg
func UploadSVG(c *gin.Context) {
	file,err := c.FormFile("svg")
	timeString := time.Now().Format("20060102150405")
	if err!=nil {
		response.FailWithMessage(fmt.Sprintf("上传文件失败，%v", err), c)
		return
	}
	fileErr := c.SaveUploadedFile(file,"./test-resource/svg/"+timeString+"-"+file.Filename);
	if fileErr!=nil{
		response.FailWithMessage(fmt.Sprintf("保存文件失败，%v", fileErr), c)
		return
	}
	response.OkWithData(gin.H{
		"path":"./test-resource/svg/"+ timeString +"-"+file.Filename,
	}, c)
}

//获取svg
func GetTestPaperSvg(c *gin.Context) {
	var testPaper model.TestPaper
	_ = c.ShouldBindQuery(&testPaper)
	err,retestPaper := service.GetTestPaperSvg(testPaper.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData( gin.H{"retestPaper":retestPaper,}, c)
	}
}

// 上传Mould
func UploadMould(c *gin.Context) {
	file,err := c.FormFile("mould")
	timeString := time.Now().Format("20060102150405")
	if err!=nil {
		response.FailWithMessage(fmt.Sprintf("上传文件失败，%v", err), c)
		return
	}
	fileErr := c.SaveUploadedFile(file,"./test-resource/mould/"+timeString+"-"+file.Filename);
	if fileErr!=nil{
		response.FailWithMessage(fmt.Sprintf("保存文件失败，%v", fileErr), c)
		return
	}
	response.OkWithData(gin.H{
		"path":"./test-resource/mould/"+ timeString +"-"+file.Filename,
	}, c)
}

//获取Mould
func GetTestPaperMould(c *gin.Context) {
	var testPaper model.TestPaper
	_ = c.ShouldBindQuery(&testPaper)
	err,testPaperMouldPath := service.GetTestPaperMould(testPaper.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("下载失败，%v", err), c)
	} else {
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "考试模板.zip")) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.Header().Add("success", "true")
		c.File(testPaperMouldPath)
	}
}

//清除Mould或svg
func RemoveTestPaperFile(c *gin.Context) {
	var testPaper model.TestPaper
	_ = c.ShouldBindJSON(&testPaper)
	err := service.RemoveTestPaperFile(&testPaper)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.Ok(c)
	}
}

//创建或者获取右键对象
func FindAndCreateTestPaperSvgNode(c *gin.Context){
	var svgNode model.TestPaperSvgNode
	c.ShouldBindJSON(&svgNode)
	err,node := service.FindAndCreateTestPaperSvgNode(&svgNode)
	if err!=nil{
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	}else{
		response.OkWithData( gin.H{"node":node,}, c)
	}
}

//上传节点文件
func UploadTestPaperSvgNode(c *gin.Context){
	filetype := c.PostForm("type")
	nodeId := c.PostForm("nodeId")
	testPaperID := c.PostForm("testPaperID")
	file,_ := c.FormFile("file")
	err := service.UploadTestPaperSvgNode(filetype,nodeId,testPaperID,file)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("添加失败，%v", err), c)
	} else {
		response.Ok(c)
	}
}

//下载节点文件
func DownloadTestPaperSvgNode(c *gin.Context){
	var reqPath request.ReqPath
	c.ShouldBindJSON(&reqPath)
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "考试模板.zip")) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.Header().Add("success", "true")
	c.File(reqPath.Path)
}

//清除节点文件
func ClearTestPaperSvgNode(c *gin.Context){
	var reqClearNode request.ReqClearNode
	c.ShouldBindJSON(&reqClearNode)
	err := service.ClearTestPaperSvgNode(reqClearNode.Type,reqClearNode.NodeId,reqClearNode.TestPaperID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.Ok(c)
	}
}

//发布考题
func PublicTestPaper(c *gin.Context){
	var testPaper model.TestPaper
	_ = c.ShouldBindJSON(&testPaper)
	err := service.PublicTestPaper(&testPaper)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("发布失败，%v", err), c)
	} else {
		response.Ok(c)
	}
}

// 获取进行中考题

func GetActiveTestPaper(c *gin.Context){
	err,testPaper,status := service.GetActiveTestPaper()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败,或没有进行中的考试，%v", err), c)
	} else {
		response.OkWithData( gin.H{"testPaper":testPaper,"status":status}, c)
	}
}