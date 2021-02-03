package response

type UploadBaseResponse struct {
	FileKey string `json:"fileKey"`
	TaskId  string `json:"taskId"`
	Flag    string `json:"flag"`
}

func NewUploadResultInstance() *UploadBaseResponse {
	return new(UploadBaseResponse)
}

func (u *UploadBaseResponse) SetFileKey(filekey string) {
	u.FileKey = filekey
}

func (u *UploadBaseResponse) GetFileKey() string {
	return u.FileKey
}

func (u *UploadBaseResponse) SetTaskId(taskId string) {
	u.TaskId = taskId
}

func (u *UploadBaseResponse) GetTaskId() string {
	return u.TaskId
}

func (u *UploadBaseResponse) SetFlag(flag string) {
	u.Flag = flag
}

func (u *UploadBaseResponse) GetFlag() string {
	return u.Flag
}
