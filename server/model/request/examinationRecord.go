// 自动生成模板ExaminationRecord
package request

import "gin-vue-admin/model"

type ReqExaminationRecord struct {
	SysUserID   string `json:"sysUserId" form:"sysUserId"`
	TestPaperID string `json:"testPaperId" form:"testPaperId"`
}
type SearchRecordParams struct {
	model.ExaminationRecord
	PageInfo
	OrderKey string `json:"orderKey"`
	Desc     bool   `json:"desc"`
}