// 自动生成模板TestPaper
package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TestPaper struct {
	gorm.Model
	TestPaperSubmitTimes bool               `json:"testPaperSubmitTimes" gorm:"column:test_paper_submit_times"`
	TestPaperStatus      bool               `json:"testPaperStatus" gorm:"column:test_paper_status"`
	TestPaperStartTime   time.Time          `json:"testPaperStartTime" gorm:"column:test_paper_start_time"`
	TestPaperEndTime     time.Time          `json:"testPaperEndTime" gorm:"column:test_paper_end_time"`
	TestPaperNote        string             `json:"testPaperNote" gorm:"column:test_paper_note"`
	TestPaperName        string             `json:"testPaperName" gorm:"column:test_paper_name"`
	TestPaperAuthor      string             `json:"testPaperAuthor" gorm:"column:test_paper_author"`
	TestPaperSvg         string             `json:"testPaperSvg" gorm:"column:test_paper_svg"`
	TestPaperMould       string             `json:"testPaperMould" gorm:"column:test_paper_mould"`
	TestPaperSvgNodes    []TestPaperSvgNode `json:"testPaperSvgNodes"`
}

type TestPaperSvgNode struct {
	gorm.Model
	NodeId        string `json:"nodeId"`
	TestPaperID   uint   `json:"testPaperID"`
	Flow          string `json:"flow"`
	Log           string `json:"log"`
	Configuration string `json:"configuration"`
	SourceCode    string `json:"sourceCode"`
}
