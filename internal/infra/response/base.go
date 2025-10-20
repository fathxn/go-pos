package response

type Response struct {
	HttpStatus     int         `json:"-"`
	StatusCode     string      `json:"status_code"`
	Success        bool        `json:"success"`
	Message        string      `json:"message"`
	Payload        interface{} `json:"payload,omitempty"`
	Error          interface{} `json:"error,omitempty"`
	AdditionalInfo string      `json:"additional_info,omitempty"`
}

type OptResponse func(*Response) *Response

func WithPayload(payload interface{}) OptResponse {
	return func(r *Response) *Response {
		r.Payload = payload
		return r
	}
}

func WithStatusCode(statusCode string) OptResponse {
	return func(r *Response) *Response {
		r.StatusCode = statusCode
		return r
	}
}