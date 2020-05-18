// 自动生成模板ExaminationRecord 
package model

import (
    "github.com/jinzhu/gorm"
)

type ExaminationRecord  struct {
      gorm.Model
      SysUserID  uint `json:"sysUserId" form:"sysUserId"`
      SysUser    SysUser `json:"sysUser" form:"sysUser"`
      TestPaperID uint `json:"testPaperId" form:"testPaperId"`
      TestPaper  TestPaper `json:"testPaper" form:"testPaper"`
      Agreement  bool `json:"agreement" gorm:"column:agreement" form:"agreement"`
      Appraise  string `json:"appraise" gorm:"column:appraise" form:"appraise"`
      TestPath  string `json:"testPath" gorm:"column:test_path"`
      Score  float64 `json:"score" gorm:"column:score" form:"agreement"`
}

type ReqExaminationRecord struct {
      gorm.Model
      SysUserID  string `json:"sysUserId"`
      TestPaperID string `json:"testPaperId"`
}
