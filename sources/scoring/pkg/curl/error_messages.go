package curl

// Error messages for curl package
const (
	ErrMsgServiceUnavailable     = "external service is unavailable"
	ErrMsgInvalidResponse        = "invalid response format from service"
	ErrMsgResourceNotFound       = "resource not found in external service"
	ErrMsgBadRequest             = "bad request to external service"
	ErrMsgRequestCreationFailed  = "failed to create HTTP request"
	ErrMsgRequestExecutionFailed = "failed to execute HTTP request"
	ErrMsgResponseReadFailed     = "failed to read response body"
	ErrMsgResponseParseFailed    = "failed to parse response JSON"
	ErrMsgUnexpectedStatusCode   = "unexpected status code from service"
	ErrMsgContextTimeout         = "request context timeout"
	ErrMsgContextCanceled        = "request context canceled"
)

// Error context keys for detailed logging
const (
	ErrCtxURL          = "url"
	ErrCtxMethod       = "method"
	ErrCtxStatusCode   = "status_code"
	ErrCtxResponseBody = "response_body"
	ErrCtxServiceName  = "service_name"
	ErrCtxQuestionID   = "question_id"
)
