package router

import (
	"gin-vue-admin/api/v1"
    "gin-vue-admin/middleware"
    "github.com/gin-gonic/gin"
)

func InitTestPaperRouter(Router *gin.RouterGroup) {
	TestPaperRouter := Router.Group("testPaper").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		TestPaperRouter.GET("getActiveTestPaper",v1.GetActiveTestPaper)   // 发布考题
		TestPaperRouter.PUT("publicTestPaper",v1.PublicTestPaper)   // 发布考题
		TestPaperRouter.DELETE("clearTestPaperSvgNode",v1.ClearTestPaperSvgNode)   // 删除节点下文件
		TestPaperRouter.POST("downloadTestPaperSvgNode",v1.DownloadTestPaperSvgNode) // 下载节点文件
		TestPaperRouter.POST("uploadTestPaperSvgNode",v1.UploadTestPaperSvgNode) // 右键获取节点信息
		TestPaperRouter.POST("findAndCreateTestPaperSvgNode",v1.FindAndCreateTestPaperSvgNode) // 右键获取节点信息
		TestPaperRouter.DELETE("removeTestPaperFile",v1.RemoveTestPaperFile)    // 删除文件
		TestPaperRouter.GET("getTestPaperMould", v1.GetTestPaperMould)        // 根据ID获取Mould
		TestPaperRouter.POST("uploadMould", v1.UploadMould)   // 新建TestPaperMould
		TestPaperRouter.GET("getTestPaperSvg", v1.GetTestPaperSvg)        // 根据ID获取SVG
		TestPaperRouter.POST("uploadSvg", v1.UploadSVG)   // 新建TestPaperSvg
		TestPaperRouter.POST("createTestPaper", v1.CreateTestPaper)   // 新建TestPaper
		TestPaperRouter.DELETE("deleteTestPaper", v1.DeleteTestPaper) //删除TestPaper
		TestPaperRouter.PUT("updateTestPaper", v1.UpdateTestPaper)    //更新TestPaper
		TestPaperRouter.GET("findTestPaper", v1.FindTestPaper)        // 根据ID获取TestPaper
		TestPaperRouter.GET("getTestPaperList", v1.GetTestPaperList)  //获取TestPaper列表
	}
}
