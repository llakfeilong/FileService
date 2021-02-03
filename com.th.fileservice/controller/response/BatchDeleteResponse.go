package response

type BatchDeleteResponse struct {
	Sucesslist Sucesslist `json:"sucesslist"`
	Faillist   Faillist   `json:"faillist"`
}

type Sucesslist struct {
	Filekey []FilekeyInfo `json:"filekeyinfo"`
}

type Faillist struct {
	Filekey []FilekeyInfo `json:"filekeyinfo"`
}

type FilekeyInfo struct {
	Filekey     string `json:"filekey"`
	Filename    string `json:"filename"`
	Filemessage string `json:"filemessage"`
}

func NewBatchDeleteInstance() *BatchDeleteResponse {
	return new(BatchDeleteResponse)
}

func NewSucesslist() *Sucesslist {
	return new(Sucesslist)
}

func NewFaillist() *Faillist {
	return new(Faillist)
}

func NewFilekeyInfo() *FilekeyInfo {
	return new(FilekeyInfo)
}

func (f *FilekeyInfo) SetFilekey(filekey string) {
	f.Filekey = filekey
}

func (f *FilekeyInfo) SetFilename(filename string) {
	f.Filename = filename
}

func (f *FilekeyInfo) SetFilemessage(filemessage string) {
	f.Filemessage = filemessage
}

func (f *Faillist) SetFilekeyInfo(filekeys []FilekeyInfo) {
	f.Filekey = filekeys
}

func (s *Sucesslist) SetFilekeyInfo(filekeys []FilekeyInfo) {
	s.Filekey = filekeys
}

func (s *Sucesslist) Add(f FilekeyInfo) {
	s.Filekey = append(s.Filekey, f)
}

func (f *Faillist) Add(info FilekeyInfo) {
	f.Filekey = append(f.Filekey, info)
}

func (b *BatchDeleteResponse) SetSucessList(s Sucesslist) {
	b.Sucesslist = s
}

func (b *BatchDeleteResponse) SetFailList(f Faillist) {
	b.Faillist = f
}
