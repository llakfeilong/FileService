package request

//批量删除入参
type BatchDeleteRequest struct {
	Filekeylist []string `json:"filekeylist"`
}

//添加文件类型入参
type AddFileTypeRequest struct {
	FileTypeName   string `json:"fileTypeName" @NotNull:"DefaultMsg:fileTypeName不能为空" @Length:"Value:30;DefaultMsg:fileTypeName长度必须少于30位"`
	FileFormat     string `json:"fileFormat" @NotNull:"DefaultMsg:fileFormat不能为空"`
	FileMaxSize    int    `json:"fileMaxSize" @NotNull:"DefaultMsg:FileMaxSize不能为空"`
	FileExpiryDate int    `json:"fileExpiryDate"`
	Enable         string `json:"enable" @NotNull:"DefaultMsg:enable不能为空" @Length:"Value:10;DefaultMsg:enable长度必须少于10个字符"`
	Remark         string `json:"remark" @Length:"Value:255;DefaultMsg:remark长度必须少于255个字符"`
}

//更新文件类型入场
type UpdateFileTypeRequest struct {
	FileTypeId     string `json:"fileTypeId" @NotNull:"DefaultMsg:fileTypeId不能为空"`
	FileTypeName   string `json:"fileTypeName" @NotNull:"DefaultMsg:fileTypeName不能为空" @Length:"Value:30;DefaultMsg:fileTypeName长度必须少于30位"`
	FileFormat     string `json:"fileFormat" @NotNull:"DefaultMsg:fileFormat不能为空"`
	FileMaxSize    int    `json:"fileMaxSize" @NotNull:"DefaultMsg:FileMaxSize不能为空"`
	FileExpiryDate int    `json:"fileExpiryDate"`
	Enable         string `json:"enable" @NotNull:"DefaultMsg:enable不能为空" @Length:"Value:10;DefaultMsg:enable长度必须少于10个字符"`
	Remark         string `json:"remark" @Length:"Value:255;DefaultMsg:remark长度必须少于255个字符"`
}

//批量查询文件类型入参
type FileTypePageRquest struct {
	FileTypeId   string `json:"fileTypeId" @NotNull:"DefaultMsg:fileTypeId不能为空"`
	FileTypeName string `json:"fileTypeName" @NotNull:"DefaultMsg:fileTypeName不能为空" @Length:"Value:30;DefaultMsg:fileTypeName长度必须少于30位"`
	FileFormat   string `json:"fileFormat" @NotNull:"DefaultMsg:fileFormat不能为空"`
	Enable       string `json:"enable" @NotNull:"DefaultMsg:enable不能为空" @Length:"Value:10;DefaultMsg:enable长度必须少于10个字符"`
	PageNum      int    `json:"pageNum" @NotNull:"DefaultMsg:PageNum不能为空"`
	PageSize     int    `json:"pageSize" @NotNull:"DefaultMsg:PageSize不能为空"`
}
