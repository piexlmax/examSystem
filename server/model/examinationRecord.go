// 自动生成模板ExaminationRecord 
package model

import (
    "github.com/jinzhu/gorm"
)

type ExaminationRecord  struct {
      gorm.Model
      SysUserID  uint `json:"sysUserId"`
      SysUser    SysUser `json:"sysUser"`
      TestPaperID uint `json:"testPaperId"`
      TestPaper  TestPaper `json:"testPaper"`
      Agreement  bool `json:"agreement" gorm:"column:agreement"`
      TestPath  string `json:"testPath" gorm:"column:test_path"`
      Score  string `json:"score" gorm:"column:score"`
}

type ReqExaminationRecord struct {
      gorm.Model
      SysUserID  string `json:"sysUserId"`
      TestPaperID string `json:"testPaperId"`
}
