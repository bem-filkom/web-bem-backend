package response

type ErrorResponse struct {
	HttpStatusCode int    `json:"-"`
	Message        string `json:"message"`
	Detail         any    `json:"detail,omitempty"`
	RefCode        string `json:"ref_code,omitempty"`
	TraceID        uint   `json:"trace_id,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func NewError(httpStatusCode int, refCode, message string) *ErrorResponse {
	return &ErrorResponse{
		HttpStatusCode: httpStatusCode,
		Message:        message,
		RefCode:        refCode,
	}
}

func (e *ErrorResponse) WithDetail(payload any) *ErrorResponse {
	e.Detail = payload
	return e
}

func (e *ErrorResponse) WithRefCode(refCode string) *ErrorResponse {
	e.RefCode = refCode
	return e
}

func (e *ErrorResponse) WithTraceID(traceID uint) *ErrorResponse {
	e.TraceID = traceID
	return e
}
