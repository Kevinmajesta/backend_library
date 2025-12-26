package response

type CustomErrorResponse struct {
	Message        string `json:"message"`
	ZiyadErrorCode string `json:"ziyad_error_code"`
	TraceID        string `json:"trace_id"`
}
