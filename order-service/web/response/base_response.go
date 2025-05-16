package response

type Response struct {
	Message string      `json:"message"`
	Meta    interface{} `json:"meta"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(message string) *Response {
	return &Response{
		Message: message,
	}
}

func DataResponse(message string, data interface{}, meta interface{}) *Response {
	return &Response{
		Message: message,
		Meta:    meta,
		Data:    data,
	}
}
