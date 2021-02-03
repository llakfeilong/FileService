package response

type ExceptionResult struct {
	FiledName      string `json:"filedName"`
	FailureMessage string `json:"failureMessage"`
}

func NewExceptionResult() *ExceptionResult {
	return new(ExceptionResult)
}

func (e *ExceptionResult) SetFiledName(filedname string) {
	e.FiledName = filedname
}

func (e *ExceptionResult) SetFailureMessage(failureMessage string) {
	e.FailureMessage = failureMessage
}
