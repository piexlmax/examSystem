import service from '@/utils/request'

// @Tags ExaminationRecord 
// @Summary 创建ExaminationRecord 
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaminationRecord  true "创建ExaminationRecord "
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /examinationRecord /createExaminationRecord  [post]
export const createExaminationRecord = (data) => {
    return service({
        url: "/examinationRecord/createExaminationRecord",
        method: 'post',
        data
    })
}


// @Tags ExaminationRecord 
// @Summary 删除ExaminationRecord 
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaminationRecord  true "删除ExaminationRecord "
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /examinationRecord /deleteExaminationRecord  [delete]
export const deleteExaminationRecord = (data) => {
    return service({
        url: "/examinationRecord/deleteExaminationRecord",
        method: 'delete',
        data
    })
}

// @Tags ExaminationRecord 
// @Summary 更新ExaminationRecord 
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaminationRecord  true "更新ExaminationRecord "
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /examinationRecord /updateExaminationRecord  [put]
export const updateExaminationRecord = (data) => {
    return service({
        url: "/examinationRecord/updateExaminationRecord",
        method: 'put',
        data
    })
}


// @Tags ExaminationRecord 
// @Summary 用id查询ExaminationRecord 
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaminationRecord  true "用id查询ExaminationRecord "
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /examinationRecord /findExaminationRecord  [get]
export const findExaminationRecord = (params) => {
    return service({
        url: "/examinationRecord/findExaminationRecord",
        method: 'get',
        params
    })
}


// @Tags ExaminationRecord 
// @Summary 分页获取ExaminationRecord 列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取ExaminationRecord 列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /examinationRecord /getExaminationRecord List [get]
export const getExaminationRecordList = (params) => {
    return service({
        url: "/examinationRecord/getExaminationRecordList",
        method: 'get',
        params
    })
}

export const offlineAppraise = (data) => {
    return service({
        url: "/examinationRecord/offlineAppraise",
        method: 'post',
        data,
        responseType: 'blob'
    })
}