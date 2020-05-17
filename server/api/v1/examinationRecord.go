package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags ExaminationRecord
// @Summary 创建ExaminationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaminationRecord true "创建ExaminationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /examinationRecord/createExaminationRecord [post]
func CreateExaminationRecord(c *gin.Context) {
	var examinationRecord model.ExaminationRecord
	_ = c.ShouldBindJSON(&examinationRecord)
	err := service.CreateExaminationRecord(examinationRecord)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ExaminationRecord
// @Summary 删除ExaminationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaminationRecord true "删除ExaminationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /examinationRecord/deleteExaminationRecord [delete]
func DeleteExaminationRecord(c *gin.Context) {
	var examinationRecord model.ExaminationRecord
	_ = c.ShouldBindJSON(&examinationRecord)
	err := service.DeleteExaminationRecord(examinationRecord)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ExaminationRecord
// @Summary 更新ExaminationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaminationRecord  true "更新ExaminationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /examinationRecord/updateExaminationRecord [put]
func UpdateExaminationRecord(c *gin.Context) {
	var examinationRecord model.ExaminationRecord
	_ = c.ShouldBindJSON(&examinationRecord)
	err := service.UpdateExaminationRecord(&examinationRecord)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ExaminationRecord
// @Summary 用id查询ExaminationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaminationRecord true "用id查询ExaminationRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /examinationRecord/findExaminationRecord [get]
func FindExaminationRecord(c *gin.Context) {
	var examinationRecord  request.ReqExaminationRecord
	_ = c.ShouldBindQuery(&examinationRecord)
	err, reexaminationRecord := service.GetExaminationRecord(examinationRecord)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"examinationRecord": reexaminationRecord}, c)
	}
}

// @Tags ExaminationRecord
// @Summary 分页获取ExaminationRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取ExaminationRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /examinationRecord/getExaminationRecordList [get]
func GetExaminationRecordList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetExaminationRecordInfoList(pageInfo)
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
