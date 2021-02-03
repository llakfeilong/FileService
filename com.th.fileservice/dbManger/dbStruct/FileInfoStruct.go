package dbStruct

import (
	"time"
)

//文件信息
type FileInfo struct {
	FileID           string     `gorm:"column:file_id;type:varchar(255)"`
	FileKey          string     `gorm:"column:file_key;type:varchar(255)"`
	FileName         string     `gorm:"column:file_name;type:varchar(255)"`
	FileOriginalName string     `gorm:"column:file_original_name;type:varchar(255)"`
	FileSuffixName   string     `gorm:"column:file_suffix_name;type:varchar(255)"`
	FilePath         string     `gorm:"column:file_path;type:varchar(255)"`
	FileSize         string     `gorm:"column:file_size;type:varchar(32)"`
	FileTypeCode     string     `gorm:"column:file_type_code;type:varchar(32)"`
	FileStatus       string     `gorm:"column:file_status;type:varchar(255)"`
	UpdateName       string     `gorm:"column:update_name;type:varchar(40)"`
	CreateName       string     `gorm:"column:create_name;type:varchar(40)"`
	UpdateDate       *time.Time `gorm:"column:update_date"`
	CreateDate       *time.Time `gorm:"column:create_date"`
	Id               int64      `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Deleted          string     `gorm:"column:deleted;type:varchar(1)"`
	Delete_date      *time.Time `gorm:"column:delete_date"`
	Delete_name      string     `gorm:"column:delete_name;type:varchar(40)"`
}

//设置表名接口
func (FileInfo) TableName() string {
	return "file_info"
}

//设置表名
func (FileTypeManager) TableName() string {
	return "file_type_manager"
}

//文件策略
type FileTypeManager struct {
	File_Type_Id     int64      `json:"fileTypeId" gorm:"column:file_type_id;primary_key"`
	File_type_name   string     `json:"fileTypeName" gorm:"column:file_type_name;type:varchar(30)"`
	File_type_code   string     `json:"fileTypeCode" gorm:"column:file_type_code;type:varchar(30)"`
	File_format      string     `json:"fileFormat" gorm:"column:file_format;type:varchar(100)"`
	File_max_size    int        `json:"fileMaxSize" gorm:"column:file_max_size;type:int"`
	File_expiry_date int        `json:"fileExpiryDate" gorm:"column:file_expiry_date;type:int"`
	Enable           string     `json:"enable" gorm:"column:enable;type:varchar(10)"`
	Update_name      string     `json:"updateName" gorm:"column:update_name;type:varchar(40)"`
	Create_name      string     `json:"createName" gorm:"column:create_name;type:varchar(40)"`
	UpdateDate       *time.Time `json:"updateDate" gorm:"column:update_date"`
	CreateDate       *time.Time `json:"createDate" gorm:"column:create_date"`
	Remark           string     `json:"remark" gorm:"column:remark;type:varchar(255)"`
	Deleted          string     `gorm:"column:deleted;type:varchar(1)"`
	Delete_date      *time.Time `gorm:"column:delete_date"`
	Delete_name      string     `gorm:"column:delete_name;type:varchar(40)"`
}

type FileTypeResponse struct {
	File_Type_Id     int64  `json:"fileTypeId" gorm:"column:file_type_id;primary_key"`
	File_type_name   string `json:"fileTypeName" gorm:"column:file_type_name;type:varchar(30)"`
	File_type_code   string `json:"fileTypeCode" gorm:"column:file_type_code;type:varchar(30)"`
	File_format      string `json:"fileFormat" gorm:"column:file_format;type:varchar(100)"`
	File_max_size    int    `json:"fileMaxSize" gorm:"column:file_max_size;type:int"`
	File_expiry_date int    `json:"fileExpiryDate" gorm:"column:file_expiry_date;type:int"`
	Enable           string `json:"enable" gorm:"column:enable;type:varchar(10)"`
	Update_name      string `json:"updateName" gorm:"column:update_name;type:varchar(40)"`
	Create_name      string `json:"createName" gorm:"column:create_name;type:varchar(40)"`
	UpdateDate       string `json:"updateDate" gorm:"column:update_date"`
	CreateDate       string `json:"createDate" gorm:"column:create_date"`
	Remark           string `json:"remark" gorm:"column:remark;type:varchar(255)"`
}
