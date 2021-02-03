package controller

import (
	"bytes"
	"com.th.fileservice/config"
	"com.th.fileservice/constant"
	"com.th.fileservice/controller/request"
	"com.th.fileservice/controller/response"
	"com.th.fileservice/dbManger"
	"com.th.fileservice/dbManger/dbStruct"
	"com.th.fileservice/httpMux"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"
)

type Tests struct {
	Id  string `json:"id" @NotNull:"DefaultMsg:id不能为空" @Length:"Value:2;DefaultMsg:id长度必须少于2位"`
	Age int    `json:"age" @NotNull:"DefaultMsg:Age不能为空" @Length:"Value:2;DefaultMsg:Age长度必须少于2位"`
}

func Test(m *httpMux.MuxContext) {
	var test Tests
	//获取post请求JSON 并解析到相应结构体
	json.Unmarshal(m.GetPostJson(), &test)
	log.Println("###id =", test.Id, "age:", test.Age)
	//校验结构体
	validtor := m.Validator.VaildSturct(test)
	//json校验结果
	b, _ := json.Marshal(validtor.ErrorResults)
	//遍历失败结果集取出异常原因
	for _, value := range validtor.ErrorResults {
		log.Println(value)
	}
	//直接获取校验结构体结果
	if m.Validator.ValidResult {
		//成功
	} else {
		//失败
	}
	//输出
	m.WritedString(true, "1000", "sucess", validtor.ErrorResults)
	log.Println("返回结果:", string(b))
}

/**
 *  获取文件返回[]byte
 */
func GetFileByteArray(m *httpMux.MuxContext) {
	downloadResponse := response.NewDownResponseInstance()
	result := httpMux.NewResultInstance()
	fileKey := m.Query("fileKey")
	fmt.Println("文件秘钥filekey", fileKey)
	var fileinfo dbStruct.FileInfo
	dbManger.GetService().Where("file_key = ? AND deleted=?", fileKey, "N").First(&fileinfo)
	if fileinfo == (dbStruct.FileInfo{}) {
		log.Println("###找不到文件秘钥")
		downloadResponse.SetFileSize("")
		downloadResponse.SetFileKey("")
		m.WritedString(false, constant.UNFIND_FILEKEY_ERROR, constant.BussinessErrorType[constant.UNFIND_FILEKEY_ERROR], downloadResponse)
		return
	} else {
		filepath := fileinfo.FilePath + fileinfo.FileName
		fmt.Println("path =" + filepath)
		//判断文件是否存在
		file, err := os.Open(filepath)
		if err != nil {
			//文件打开错误
			fmt.Println(err)
			downloadResponse.SetFileSize("")
			downloadResponse.SetFileKey("")
			m.WritedString(false, constant.UNFIND_FILE_ERROR, constant.BussinessErrorType[constant.UNFIND_FILE_ERROR], downloadResponse)
			return
		}
		defer file.Close()
		info, err := file.Stat()
		if err != nil {
			fmt.Println(err)
			downloadResponse.SetFileSize("")
			downloadResponse.SetFileKey("")
			m.WritedString(false, constant.UNFIND_FILE_ERROR, constant.BussinessErrorType[constant.UNFIND_FILE_ERROR], downloadResponse)
			return
		}

		filesize := info.Size()
		buffer := make([]byte, filesize)
		file.Read(buffer)
		fmt.Println("file key = ", fileinfo.FileKey)
		downloadResponse.SetFileSize(strconv.FormatInt(filesize, 10))
		downloadResponse.SetFileKey(fileinfo.FileKey)
		downloadResponse.SetFileData(buffer)
		downloadResponse.SetFileOriginalName(fileinfo.FileOriginalName)
		fmt.Println("filekey:", downloadResponse.FileKey)
		result.SucessDefault(downloadResponse)
		m.Writer.Header().Set("content-type", "application/json")
		m.WriteString(result)
	}

}

/**
 * 获取文件base64流
 */
func GetFileBase64(m *httpMux.MuxContext) {
	log.Println("##获取文件base64")
	m.Writer.Header().Set("content-type", "application/json")
	downloadResponse := response.NewDownResponseInstance()
	result := httpMux.NewResultInstance()
	fileKey := m.Query("fileKey")
	if fileKey == "" {
		downloadResponse.SetFileSize("")
		downloadResponse.SetFileKey("")
		m.WritedString(false, constant.NOTVALID_ERROR, constant.BussinessErrorType[constant.NOTVALID_ERROR], downloadResponse)
		return
	}
	var fileinfo dbStruct.FileInfo
	dbManger.GetService().Where("file_key = ? AND deleted=?", fileKey, "N").First(&fileinfo)
	filepath := fileinfo.FilePath + fileinfo.FileName
	//判断文件是否存在
	file, err := os.Open(filepath)
	if err != nil {
		//文件打开错误
		log.Println(err)
		downloadResponse.SetFileSize("")
		downloadResponse.SetFileKey("")
		m.WritedString(false, constant.UNFIND_FILE_ERROR, constant.BussinessErrorType[constant.UNFIND_FILE_ERROR], downloadResponse)
		return
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		log.Println(err)
		downloadResponse.SetFileSize("")
		downloadResponse.SetFileKey("")
		m.WritedString(false, constant.UNFIND_FILE_ERROR, constant.BussinessErrorType[constant.UNFIND_FILE_ERROR], downloadResponse)
		return
	}

	filesize := info.Size()
	buffer := make([]byte, filesize)
	file.Read(buffer)
	filebase64 := base64.StdEncoding.EncodeToString(buffer)
	downloadResponse.SetFileSize(strconv.FormatInt(filesize, 10))
	downloadResponse.SetFileKey(fileinfo.FileKey)
	downloadResponse.SetFileBase64(filebase64)
	downloadResponse.SetFileOriginalName(fileinfo.FileOriginalName)
	result.SucessDefault(downloadResponse)
	m.WriteString(result)
}

/**
 * 旧文件上传接口没有FileType
 */
func UploadFile(m *httpMux.MuxContext) {
	m.Writer.Header().Set("content-type", "application/json")
	uploadResponse := response.NewUploadResultInstance()
	result := httpMux.NewResultInstance()
	err := m.ParseMultipartForm(30)
	if err != nil {
		log.Println(err)
		return
	}
	srcfile, srcfileheader, err := m.FormFile("files") // file 是上传表单域的名字
	if err != nil {
		uploadResponse.SetFileKey("")
		uploadResponse.SetFlag("")
		uploadResponse.SetTaskId("")
		result.Fail(constant.BussinessErrorType[constant.NOTVALID_ERROR], constant.NOTVALID_ERROR, uploadResponse)
		m.WriteString(result)
		return
	}
	flag := m.FormValue("flag")
	taskId := m.FormValue("taskId")
	if srcfile == nil || srcfileheader.Size == 0 {
		uploadResponse.SetFileKey("")
		uploadResponse.SetFlag("")
		uploadResponse.SetTaskId("")
		result.Fail(constant.BussinessErrorType[constant.NOTVALID_ERROR], constant.NOTVALID_ERROR, uploadResponse)
		m.WriteString(result)
		return
	}
	defer srcfile.Close()
	var filemaxsize = 30720
	if filemaxsize < int(srcfileheader.Size) {
		uploadResponse.SetFileKey("")
		uploadResponse.SetFlag("")
		uploadResponse.SetTaskId("")
		result.Fail(constant.BussinessErrorType[constant.UPLOADFILE_SIZE_ERROR], constant.UPLOADFILE_SIZE_ERROR, uploadResponse)
		m.WriteString(result)
		return
	}
	//获取文件后缀名位置
	index := strings.LastIndex(srcfileheader.Filename, ".")
	//获取文件后缀名
	subfix := srcfileheader.Filename[index:]
	//生成本地存储的文件名
	localFileName := fmt.Sprintf("%d", time.Now().Unix()) + subfix
	storagePath := config.GetStoragePath()

	//保存文件内容
	dstfile, err := os.OpenFile(storagePath+localFileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("保存文件内容错误:", err)
		uploadResponse.SetFileKey("")
		uploadResponse.SetFlag("")
		uploadResponse.SetTaskId("")
		result.Fail(constant.BussinessErrorType[constant.SAVEFILEFAIL], constant.SAVEFILEFAIL, uploadResponse)
		m.WriteString(result)
		return
	}
	_, err = io.Copy(dstfile, srcfile)
	if err != nil {
		log.Println(err)
		uploadResponse.SetFileKey("")
		uploadResponse.SetFlag("")
		uploadResponse.SetTaskId("")
		result.Fail(constant.BussinessErrorType[constant.SAVEFILEFAIL], constant.SAVEFILEFAIL, uploadResponse)
		m.WriteString(result)
		return
	}
	//使用真随机生成字符串秘钥
	randomkey := CreateRandomString(64)
	//使用sha256加密
	filekeysum := sha256.Sum256([]byte(randomkey))
	filekey := hex.EncodeToString(filekeysum[:])
	//数据组装
	var fileInfo dbStruct.FileInfo
	t := time.Now()
	fileInfo.FileKey = filekey
	fileInfo.FileOriginalName = srcfileheader.Filename
	fileInfo.FileName = localFileName
	fileInfo.FilePath = storagePath
	fileInfo.FileSize = strconv.FormatInt(srcfileheader.Size, 10)
	fileInfo.FileSuffixName = subfix
	fileInfo.FileTypeCode = "tenfineFiles"
	fileInfo.CreateDate = &t
	fileInfo.UpdateName = ""
	fileInfo.CreateName = "system"
	fileInfo.Deleted = "N"
	//存入数据库
	dberr := dbManger.GetService().Create(&fileInfo).Error
	if dberr != nil {
		//存入数据失败
		log.Println(dberr)
		uploadResponse.SetFileKey("")
		uploadResponse.SetFlag("")
		uploadResponse.SetTaskId("")
		result.Fail(constant.BussinessErrorType[constant.SAVEFILEFAIL], constant.SAVEFILEFAIL, uploadResponse)
		m.WriteString(result)
		return
	}
	uploadResponse.SetFileKey(fileInfo.FileKey)
	uploadResponse.SetFlag(flag)
	uploadResponse.SetTaskId(taskId)
	result.SucessDefault(uploadResponse)
	m.WriteString(result)
	return
}

/**
 * 文件上传接口带文件类型
 */
func UploadFileToFileType(m *httpMux.MuxContext) {
	m.Writer.Header().Set("content-type", "application/json")
	uploadResponse := response.NewUploadResultInstance()
	result := httpMux.NewResultInstance()
	err := m.ParseMultipartForm(30)
	if err != nil {
		log.Println(err)
		return
	}
	// 写明缓冲的大小。如果超过缓冲，文件内容会被放在临时目录中，而不是内存。过大可能较多占用内存，过小可能增加硬盘 I/O
	// FormFile() 时调用 ParseMultipartForm() 使用的大小是 32 << 20，32MB
	srcfile, srcfileheader, err := m.FormFile("files") // file 是上传表单域的名字
	if err != nil {
		uploadResponse.SetFileKey("")
		uploadResponse.SetFlag("")
		uploadResponse.SetTaskId("")
		result.Fail(constant.BussinessErrorType[constant.NOTVALID_ERROR], constant.NOTVALID_ERROR, uploadResponse)
		m.WriteString(result)
		return
	}
	//校验其他参数
	fileTypeId := m.FormValue("fileTypeId")
	flag := m.FormValue("flag")
	taskId := m.FormValue("taskId")
	if srcfile == nil || srcfileheader.Size == 0 || fileTypeId == "" {
		uploadResponse.SetFileKey("")
		uploadResponse.SetFlag("")
		uploadResponse.SetTaskId("")
		result.Fail(constant.BussinessErrorType[constant.NOTVALID_ERROR], constant.NOTVALID_ERROR, uploadResponse)
		m.WriteString(result)
		return
	}
	defer srcfile.Close() // 此时上传内容的 IO 已经打开，需要手动关闭！！
	var filetype dbStruct.FileTypeManager
	//查数据库
	notfound := dbManger.GetService().Where("file_type_id = ? AND deleted=?", fileTypeId, "N").First(&filetype).RecordNotFound()
	if notfound {
		//查不到数据
		uploadResponse.SetFileKey("")
		uploadResponse.SetFlag("")
		uploadResponse.SetTaskId("")
		result.Fail(constant.BussinessErrorType[constant.NOTVALID_ERROR], constant.NOTVALID_ERROR, uploadResponse)
		m.WriteString(result)
		return
	} else {
		fmt.Println(srcfileheader.Filename, "分割", srcfileheader.Size, "typeid:", fileTypeId)
		//开始校验文件大小
		if filetype.File_max_size < int(srcfileheader.Size) {
			uploadResponse.SetFileKey("")
			uploadResponse.SetFlag("")
			uploadResponse.SetTaskId("")
			result.Fail(constant.BussinessErrorType[constant.UPLOADFILE_SIZE_ERROR], constant.UPLOADFILE_SIZE_ERROR, uploadResponse)
			m.WriteString(result)
			return
		}
		//获取文件后缀名位置
		index := strings.LastIndex(srcfileheader.Filename, ".")
		//获取文件后缀名
		subfix := srcfileheader.Filename[index:]
		//校验格式
		fileformats := strings.Split(filetype.File_format, "、")
		var check bool
		for _, format := range fileformats {
			if subfix == "."+format {
				check = true
			}
		}
		if !check {
			//不支持的格式
			uploadResponse.SetFileKey("")
			uploadResponse.SetFlag("")
			uploadResponse.SetTaskId("")
			result.Fail(constant.BussinessErrorType[constant.UNSUPPORTED_FILETYPE_ERROR], constant.UNSUPPORTED_FILETYPE_ERROR, uploadResponse)
			m.WriteString(result)
			return
		}

		//生成本地存储的文件名
		localFileName := fmt.Sprintf("%d", time.Now().Unix()) + subfix
		storagePath := config.GetStoragePath() + filetype.File_type_code
		//保存文件内容
		dstfile, err := os.OpenFile(storagePath+localFileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
		if err != nil {
			uploadResponse.SetFileKey("")
			uploadResponse.SetFlag("")
			uploadResponse.SetTaskId("")
			result.Fail(constant.BussinessErrorType[constant.SAVEFILEFAIL], constant.SAVEFILEFAIL, uploadResponse)
			m.WriteString(result)
			return
		}
		_, err = io.Copy(dstfile, srcfile)
		if err != nil {
			log.Println(err)
			uploadResponse.SetFileKey("")
			uploadResponse.SetFlag("")
			uploadResponse.SetTaskId("")
			result.Fail(constant.BussinessErrorType[constant.SAVEFILEFAIL], constant.SAVEFILEFAIL, uploadResponse)
			m.WriteString(result)
			return
		}
		//使用真随机生成字符串秘钥
		randomkey := CreateRandomString(64)
		//使用sha256加密
		filekeysum := sha256.Sum256([]byte(randomkey))
		filekey := hex.EncodeToString(filekeysum[:])
		//数据组装
		var fileInfo dbStruct.FileInfo
		t := time.Now()
		fileInfo.FileKey = filekey
		fileInfo.FileOriginalName = srcfileheader.Filename
		fileInfo.FileName = localFileName
		fileInfo.FilePath = storagePath
		fileInfo.FileSize = strconv.FormatInt(srcfileheader.Size, 10)
		fileInfo.FileSuffixName = subfix
		fileInfo.FileTypeCode = filetype.File_type_code
		fileInfo.CreateDate = &t
		fileInfo.UpdateName = ""
		fileInfo.CreateName = "fileservice"
		fileInfo.Deleted = "N"
		//存入数据库
		dberr := dbManger.CreateStruct(fileInfo)
		if dberr != nil {
			//存入数据失败
			log.Println(err)
			uploadResponse.SetFileKey("")
			uploadResponse.SetFlag("")
			uploadResponse.SetTaskId("")
			result.Fail(constant.BussinessErrorType[constant.SAVEFILEFAIL], constant.SAVEFILEFAIL, uploadResponse)
			m.WriteString(result)
			return
		}
		uploadResponse.SetFileKey(fileInfo.FileKey)
		uploadResponse.SetFlag(flag)
		uploadResponse.SetTaskId(taskId)
		result.SucessDefault(uploadResponse)
		m.WriteString(result)
		return
	}
}

/**
 * 批量删除文件接口
 */
func BatchDeleteFile(m *httpMux.MuxContext) {
	m.Writer.Header().Set("content-type", "application/json")
	result := httpMux.NewResultInstance()
	batchresponse := response.NewBatchDeleteInstance()
	var request request.BatchDeleteRequest
	var fileinfos []dbStruct.FileInfo
	jsonerr := json.Unmarshal(m.GetPostJson(), &request)
	if jsonerr != nil {
		result.Fail(constant.BussinessErrorType[constant.JSON_PARSE_ERROR], constant.JSON_PARSE_ERROR, batchresponse)
		m.WriteString(result)
		return
	}
	//查数据
	dbManger.GetService().Where("file_key in (?) AND deleted=?", request.Filekeylist, "N").Find(&fileinfos)
	if fileinfos == nil {
		result.Fail(constant.BussinessErrorType[constant.NOTVALID_ERROR], constant.NOTVALID_ERROR, batchresponse)
		m.WriteString(result)
		return
	}
	var failist response.Faillist
	var sucesslist response.Sucesslist
	//遍历文件信息
	for _, fileinfo := range fileinfos {
		filepath := fileinfo.FilePath + fileinfo.FileName
		err := os.Remove(filepath)
		if err != nil {
			log.Println("批量删除失败", "key =", fileinfo.FileKey)
			//删除失败 在返回里面添加到失败列表
			var filekeyinfo response.FilekeyInfo
			filekeyinfo.SetFilekey(fileinfo.FileKey)
			filekeyinfo.SetFilename(fileinfo.FileName)
			filekeyinfo.SetFilemessage("删除失败")
			failist.Add(filekeyinfo)
		} else {
			log.Println("批量删除成功", "key =", fileinfo.FileKey)
			//删除成功更新数据库状态
			var filekeyinfo response.FilekeyInfo
			t := time.Now()
			filekeyinfo.SetFilekey(fileinfo.FileKey)
			filekeyinfo.SetFilename(fileinfo.FileName)
			filekeyinfo.SetFilemessage("删除成功")
			fileinfo.Deleted = "Y"
			fileinfo.Delete_date = &t
			dbManger.GetService().Model(&fileinfo).Update(fileinfo)
			sucesslist.Add(filekeyinfo)
		}
	}
	batchresponse.SetSucessList(sucesslist)
	batchresponse.SetFailList(failist)
	result.SucessDefault(batchresponse)
	m.WriteString(result)
}

func CreateRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}
