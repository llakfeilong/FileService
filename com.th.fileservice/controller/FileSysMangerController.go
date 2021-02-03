package controller

import (
	"com.th.fileservice/constant"
	"com.th.fileservice/controller/request"
	"com.th.fileservice/controller/response"
	"com.th.fileservice/dbManger"
	"com.th.fileservice/dbManger/dbStruct"
	"com.th.fileservice/httpMux"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"time"
)

/**
 * 新增文件类型
 */
func AddFileType(m *httpMux.MuxContext) {
	m.Writer.Header().Set("content-type", "application/json")
	result := httpMux.NewResultInstance()
	var fileTypeRequest request.AddFileTypeRequest
	jsonerr := json.Unmarshal(m.GetPostJson(), &fileTypeRequest)
	if jsonerr != nil {
		result.Fail(constant.BussinessErrorType[constant.JSON_PARSE_ERROR], constant.JSON_PARSE_ERROR, "")
		m.WriteString(result)
		return
	}
	//校验结构体参数
	valid := m.Validator.VaildSturct(fileTypeRequest)
	if !valid.ValidResult {
		//校验不通过
		result.Fail(constant.BussinessErrorType[constant.NOTVALID_ERROR], constant.NOTVALID_ERROR, valid.ErrorResults)
		m.WriteString(result)
		return
	}
	var fileTypeManger dbStruct.FileTypeManager
	//校验文件名是否重复
	dbManger.GetService().Where("file_type_name=? AND deleted=?", fileTypeRequest.FileTypeName, "N").First(&fileTypeManger)
	//判断结构体是否为空
	if fileTypeManger != (dbStruct.FileTypeManager{}) {
		result.Fail(constant.BussinessErrorType[constant.DUPLICATENAME], constant.DUPLICATENAME, valid.ErrorResults)
		m.WriteString(result)
		return
	}
	//把结构体置空
	t := time.Now()
	fileTypeManger = dbStruct.FileTypeManager{}
	fileTypeManger.Create_name = "system"
	fileTypeManger.CreateDate = &t
	fileTypeManger.Update_name = ""
	fileTypeManger.Enable = fileTypeRequest.Enable
	fileTypeManger.File_expiry_date = fileTypeRequest.FileExpiryDate
	fileTypeManger.File_format = fileTypeRequest.FileFormat
	fileTypeManger.File_max_size = fileTypeRequest.FileMaxSize
	fileTypeManger.File_type_name = fileTypeRequest.FileTypeName
	fileTypeManger.Remark = fileTypeRequest.Remark
	fileTypeManger.Deleted = "N"
	//存入数据
	err := dbManger.CreateStruct(fileTypeManger)
	if err != nil {
		result.Fail(constant.BussinessErrorType[constant.SYSTEM_ERR], constant.SYSTEM_ERR, valid.ErrorResults)
		m.WriteString(result)
		return
	}
	//再次查询数据返回id
	fileTypeManger = dbStruct.FileTypeManager{}
	dbManger.GetService().Where("file_type_name=? AND deleted=?", fileTypeRequest.FileTypeName, "N").First(&fileTypeManger)
	result.SucessDefault(fileTypeManger.File_Type_Id)
	m.WriteString(result)
}

//更新文件类型controller
func UpdateFileType(m *httpMux.MuxContext) {
	m.Writer.Header().Set("content-type", "application/json")
	result := httpMux.NewResultInstance()
	var updateTypeRequest request.UpdateFileTypeRequest
	jsonerr := json.Unmarshal(m.GetPostJson(), &updateTypeRequest)
	if jsonerr != nil {
		result.Fail(constant.BussinessErrorType[constant.JSON_PARSE_ERROR], constant.JSON_PARSE_ERROR, "")
		m.WriteString(result)
		return
	}
	valid := m.Validator.VaildSturct(updateTypeRequest)
	if !valid.ValidResult {
		//校验不通过
		result.Fail(constant.BussinessErrorType[constant.NOTVALID_ERROR], constant.NOTVALID_ERROR, valid.ErrorResults)
		m.WriteString(result)
		return
	}
	var fileTypeManger dbStruct.FileTypeManager
	//校验文件名是否重复
	dbManger.GetService().Where("file_type_name=? AND file_type_id !=? AND deleted=?", updateTypeRequest.FileTypeName, updateTypeRequest.FileTypeId, "N").First(&fileTypeManger)
	//判断结构体是否为空
	if fileTypeManger != (dbStruct.FileTypeManager{}) {
		result.Fail(constant.BussinessErrorType[constant.DUPLICATENAME], constant.DUPLICATENAME, valid.ErrorResults)
		m.WriteString(result)
		return
	}
	//清空数据
	fileTypeManger = dbStruct.FileTypeManager{}
	dbManger.GetService().Where("file_type_id = ? AND deleted=?", updateTypeRequest.FileTypeId, "N").First(&fileTypeManger)
	if fileTypeManger == (dbStruct.FileTypeManager{}) {
		result.Fail(constant.BussinessErrorType[constant.DUPLICATENAME], constant.DUPLICATENAME, valid.ErrorResults)
		m.WriteString(result)
		return
	}
	t := time.Now()
	fileTypeManger.Remark = updateTypeRequest.Remark
	fileTypeManger.File_type_name = updateTypeRequest.FileTypeName
	fileTypeManger.File_max_size = updateTypeRequest.FileMaxSize
	fileTypeManger.File_format = updateTypeRequest.FileFormat
	fileTypeManger.File_expiry_date = updateTypeRequest.FileExpiryDate
	fileTypeManger.Enable = updateTypeRequest.Enable
	fileTypeManger.Update_name = "system"
	fileTypeManger.UpdateDate = &t
	err := dbManger.UpdateStruct(fileTypeManger)
	if err != nil {
		result.Fail(constant.BussinessErrorType[constant.SYSTEM_ERR], constant.SYSTEM_ERR, valid.ErrorResults)
		m.WriteString(result)
		return
	}
	result.SucessDefault(fileTypeManger.File_Type_Id)
	m.WriteString(result)
}

//批量查询条件-文件类型ID
func equalToFileTypeId(fileTypeId string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("File_Type_Id = ?", fileTypeId)
	}
}

//批量查询条件-文件类型名
func likeFileTypeName(fileTypeName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("file_type_name LIKE ?", "%"+fileTypeName+"%")
	}
}

//批量查询条件-文件类型支持的格式
func likeFileFormat(fileFormat string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("file_format LIKE ?", "%"+fileFormat+"%")
	}
}

//文件类型是否启用
func equalToEnable(enable string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("enable = ?", enable)
	}
}

//分页查询文件类型
func PageList(m *httpMux.MuxContext) {
	m.Writer.Header().Set("content-type", "application/json")
	result := httpMux.NewResultInstance()
	var fileTypePageRquest request.FileTypePageRquest
	jsonerr := json.Unmarshal(m.GetPostJson(), &fileTypePageRquest)
	if jsonerr != nil {
		result.Fail(constant.BussinessErrorType[constant.JSON_PARSE_ERROR], constant.JSON_PARSE_ERROR, "")
		m.WriteString(result)
		return
	}
	valid := m.Validator.VaildSturct(fileTypePageRquest)
	if !valid.ValidResult {
		//校验不通过
		result.Fail(constant.BussinessErrorType[constant.NOTVALID_ERROR], constant.NOTVALID_ERROR, valid.ErrorResults)
		m.WriteString(result)
		return
	}
	var scopes []func(*gorm.DB) *gorm.DB

	if fileTypePageRquest.FileTypeId != "" {
		scopes = append(scopes, equalToFileTypeId(fileTypePageRquest.FileTypeId))
	}
	if fileTypePageRquest.FileTypeName != "" {
		scopes = append(scopes, likeFileTypeName(fileTypePageRquest.FileTypeName))
	}
	if fileTypePageRquest.FileFormat != "" {
		scopes = append(scopes, likeFileFormat(fileTypePageRquest.FileFormat))
	}
	if fileTypePageRquest.Enable != "" {
		scopes = append(scopes, equalToEnable(fileTypePageRquest.Enable))
	}
	var pageresult response.PageResult
	var pages int
	var fileTypeMangers []dbStruct.FileTypeManager
	var fileTypeResponse []dbStruct.FileTypeResponse

	//开始批量查询
	dbManger.PageQuery(scopes, fileTypePageRquest.PageSize, fileTypePageRquest.PageNum, "").Where("deleted=?", "N").Find(&fileTypeMangers).Count(&pages)
	//开始格式化时间
	for _, filetype := range fileTypeMangers {
		var typeResponse dbStruct.FileTypeResponse
		if !filetype.CreateDate.IsZero() {
			typeResponse.CreateDate = filetype.CreateDate.Format("2006.01.02 15:04:05")
		} else {
			typeResponse.CreateDate = ""
		}
		if !filetype.UpdateDate.IsZero() {
			typeResponse.UpdateDate = filetype.UpdateDate.Format("2006.01.02 15:04:05")
		} else {
			typeResponse.UpdateDate = ""
		}
		typeResponse.Create_name = filetype.Create_name
		typeResponse.Update_name = filetype.Update_name
		typeResponse.File_type_name = filetype.File_type_name
		typeResponse.Enable = filetype.Enable
		typeResponse.File_expiry_date = filetype.File_expiry_date
		typeResponse.File_format = filetype.File_format
		typeResponse.File_max_size = filetype.File_max_size
		typeResponse.Remark = filetype.Remark
		typeResponse.File_Type_Id = filetype.File_Type_Id
		typeResponse.File_type_code = filetype.File_type_code
		fileTypeResponse = append(fileTypeResponse, typeResponse)
	}
	pageresult.PageNum = fileTypePageRquest.PageNum
	pageresult.PageSize = fileTypePageRquest.PageSize
	pageresult.Pages = pages
	pageresult.List = fileTypeResponse
	result.SucessDefault(pageresult)
	m.WriteString(result)
}
