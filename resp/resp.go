package resp

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(data any) Result {
	return Result{
		Code:    0,
		Message: "success",
		Data:    data,
	}
}

func Failure(msg string) Result {
	return Result{
		Code:    500,
		Message: msg,
		Data:    nil,
	}
}
