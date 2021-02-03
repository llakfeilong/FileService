package response

type PageResult struct {
	PageNum  int         `json:"pageNum"`
	PageSize int         `json:"pageSize"`
	Pages    int         `json:"pages"`
	Total    int64       `json:"total"`
	List     interface{} `json:"list"`
}

func NewPageResult() *PageResult {
	return new(PageResult)
}

func (p *PageResult) SetPageNum(pageNum int) {
	p.PageNum = pageNum
}

func (p *PageResult) SetPageSize(pageSize int) {
	p.PageSize = pageSize
}

func (p *PageResult) Setlist(list interface{}) {
	p.List = list
}
func (p *PageResult) SetPages(Pages int) {
	p.Pages = Pages
}

func (p *PageResult) SetTotal(total int64) {
	p.Total = total
}
