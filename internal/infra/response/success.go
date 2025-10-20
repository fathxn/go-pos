package response

import "net/http"

func newSuccess(httpStatus int, msg string, opts ...OptResponse) Response {
	var resp = Response{
		HttpStatus: httpStatus,
		Success:    true,
		Message:    msg,
	}

	for _, opt := range opts {
		opt(&resp)
	}

	return resp
}

func NewSuccessCreated(msg string, opts ...OptResponse) Response {
	return newSuccess(http.StatusCreated, msg, opts...)
}

func NewSuccessOk(msg string, opts ...OptResponse) Response {
	return newSuccess(http.StatusOK, msg, opts...)
}
