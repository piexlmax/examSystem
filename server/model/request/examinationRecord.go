// 自动生成模板ExaminationRecord
package request

type ReqExaminationRecord struct {
	SysUserID   string `json:"sysUserId" form:"sysUserId"`
	TestPaperID string `json:"testPaperId" form:"testPaperId"`
}
