package router

import (
	"gin-vue-admin/api/v1"
    "gin-vue-admin/middleware"
    "github.com/gin-gonic/gin"
)

func InitExaminationRecordRouter(Router *gin.RouterGroup) {
	ExaminationRecordRouter := Router.Group("examinationRecord").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ExaminationRecordRouter.POST("createExaminationRecord", v1.CreateExaminationRecord)     // 新建ExaminationRecord
		ExaminationRecordRouter.DELETE("deleteExaminationRecord", v1.DeleteExaminationRecord)   //删除ExaminationRecord
		ExaminationRecordRouter.PUT("updateExaminationRecord", v1.UpdateExaminationRecord)   //更新ExaminationRecord
		ExaminationRecordRouter.GET("findExaminationRecord", v1.FindExaminationRecord)           // 根据ID获取ExaminationRecord
		ExaminationRecordRouter.GET("getExaminationRecordList", v1.GetExaminationRecordList) //获取ExaminationRecord 列表
	}
}
