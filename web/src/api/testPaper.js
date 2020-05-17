import service from '@/utils/request'

// @Tags TestPaper
// @Summary 创建TestPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestPaper true "创建TestPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testPaper/createTestPaper [post]
export const createTestPaper = (data) => {
     return service({
         url: "/testPaper/createTestPaper",
         method: 'post',
         data
     })
 }


// @Tags TestPaper
// @Summary 删除TestPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestPaper true "删除TestPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /testPaper/deleteTestPaper [delete]
 export const deleteTestPaper = (data) => {
     return service({
         url: "/testPaper/deleteTestPaper",
         method: 'delete',
         data
     })
 }

// @Tags TestPaper
// @Summary 更新TestPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestPaper true "更新TestPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /testPaper/updateTestPaper [put]
 export const updateTestPaper = (data) => {
     return service({
         url: "/testPaper/updateTestPaper",
         method: 'put',
         data
     })
 }


// @Tags TestPaper
// @Summary 用id查询TestPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestPaper true "用id查询TestPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /testPaper/findTestPaper [get]
 export const findTestPaper = (params) => {
     return service({
         url: "/testPaper/findTestPaper",
         method: 'get',
         params
     })
 }


// @Tags TestPaper
// @Summary 分页获取TestPaper列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TestPaper列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /testPaper/getTestPaperList [get]
 export const getTestPaperList = (params) => {
     return service({
         url: "/testPaper/getTestPaperList",
         method: 'get',
         params
     })
 }

 export const getTestPaperSvg = (params) =>{
    return service({
        url: "/testPaper/getTestPaperSvg",
        method: 'get',
        params
    })
 }


 
 export const getTestPaperMould = (params) =>{
    return service({
        url: "/testPaper/getTestPaperMould",
        method: 'get',
        params,
        responseType: 'blob'
    })
 }


 export const removeTestPaperFile = (data) =>{
    return service({
        url: "/testPaper/removeTestPaperFile",
        method: 'delete',
        data
    })
 }

 

 export const findAndCreateTestPaperSvgNode = (data) =>{
    return service({
        url: "/testPaper/findAndCreateTestPaperSvgNode",
        method: 'post',
        data
    })
 }


 export const downloadTestPaperSvgNode = (data) =>{
    return service({
        url: "/testPaper/downloadTestPaperSvgNode",
        method: 'post',
        data,
        responseType: 'blob'
    })
 }

 export const clearTestPaperSvgNode = (data) =>{
    return service({
        url: "/testPaper/clearTestPaperSvgNode",
        method: 'delete',
        data,
    })
 }
 
 export const publicTestPaper = (data) =>{
    return service({
        url: "/testPaper/publicTestPaper",
        method: 'put',
        data,
    })
 }
 
 export const getActiveTestPaper = () =>{
    return service({
        url: "/testPaper/getActiveTestPaper",
        method: 'get'
    })
 }
 

 