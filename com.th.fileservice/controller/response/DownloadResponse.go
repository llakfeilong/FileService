package response

type DownloadResponse struct {
	FileOriginalName string `json:"fileOriginalName"`
	Filesize         string `json:"filesize"`
	FileKey          string `json:"fileKey"`
	FileBase64       string `json:"fileBase64"`
	FileData         []byte `json:"fileData"`
}

//实例化downloadResponse
func NewDownResponseInstance() *DownloadResponse {
	return new(DownloadResponse)
}

//文件大小赋值
func (d *DownloadResponse) SetFileSize(filesize string) {
	d.Filesize = filesize
}

//文件大小取值
func (d *DownloadResponse) GetFileSize() string {
	return d.Filesize
}

//文件key赋值
func (d *DownloadResponse) SetFileKey(fileKey string) {
	d.FileKey = fileKey
}

//文件key取值
func (d *DownloadResponse) GetFileKey() string {
	return d.FileKey
}

func (d *DownloadResponse) SetFileOriginalName(fileOriginalName string) {
	d.FileOriginalName = fileOriginalName
}

func (d *DownloadResponse) GetFileOriginalName() string {
	return d.FileOriginalName
}

func (d *DownloadResponse) SetFileBase64(fileBase64 string) {
	d.FileBase64 = fileBase64
}

func (d *DownloadResponse) GetFileBase64() string {
	return d.FileBase64
}

func (d *DownloadResponse) SetFileData(fileData []byte) {
	d.FileData = fileData
}

func (d *DownloadResponse) GetFileData() []byte {
	return d.FileData
}
