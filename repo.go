package model

import "github.com/jinzhu/gorm"

type Repo struct {
	gorm.Model
	RepoID                string   `gorm:"type:varchar(50);unique_index"`
	SearchScore           string   `gorm:"type:varchar(25);"`
	URL                   string   `gorm:"type:varchar(100);"`
	Email                 string   `gorm:"type:varchar(50);"`
	Name                  string   `gorm:"type:varchar(50);"`
	Phone                 string   `gorm:"type:varchar(25);"`
	Created               string   `gorm:"type:varchar(25);"`
	LastModified          string   `gorm:"type:varchar(25);"`
	DownloadURL           string   `gorm:"type:varchar(100);`
	HomepageURL           string   `gorm:"type:varchar(100);`
	DisclaimerURL         string   `gorm:"type:varchar(100);`
	DisclaimerText        string   `gorm:"type:varchar(100);`
	LaborHours            int      `gorm:"type:varchar(100);`
	Languages             []string `gorm:"type:varchar(100);"`
	Name                  string   `gorm:"type:varchar(50);"`
	Version               string   `gorm:"type:varchar(50);"`
	Description           string   `gorm:"type:varchar(100);"`
	Organization          string   `gorm:"type:varchar(100);"`
	LicenseName           string   `gorm:"type:varchar(50);"`
	LicenseURL            string   `gorm:"type:varchar(100);"`
	AgencyName            string   `gorm:"type:varchar(100);"`
	AgencyAcronym         string   `gorm:"type:varchar(25);"`
	Website               string   `gorm:"type:varchar(100);"`
	CodeURL               string   `gorm:"type:varchar(100);"`
	FallbackFile          string   `gorm:"type:varchar(100);"`
	AgencyWidePolicy      int      `gorm:"type:varchar(100);"`
	OpenSourceRequirement int      `gorm:"type:varchar(100);"`
	InventoryRequirement  int      `gorm:"type:varchar(100);"`
	SchemaFormat          float32  `gorm:"type:varchar(100);"`
	OverallCompliance     float32  `gorm:"type:varchar(100);"`
	ComplianceDashboard   bool     `gorm:"type:bool;"`
}
