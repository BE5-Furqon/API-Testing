package common

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func SuccessResponse(data interface{}) Response {
	return Response{
		Code: 200,
		Message: "success operation",
		Data: data,
	}
}

func ErrorResponse(message string, code int) Response {
	return Response{
		Code: code,
		Message: message,
	}
}