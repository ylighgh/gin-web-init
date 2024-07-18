package response

type BaseResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data *DataRef `json:"data"`
}

func OK() *BaseResponse {
	return &BaseResponse{
		Code: 200,
		Msg:  "success",
		Data: nil,
	}
}

func OKWithData(data *DataRef) *BaseResponse {
	return &BaseResponse{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
}

func BadWithReason(reason Reason) *BaseResponse {
	return &BaseResponse{
		Code: reason.Code(),
		Msg:  reason.Message(),
		Data: nil,
	}
}

func BadWithBindReason(err error) *BaseResponse {
	return &BaseResponse{
		Code: 1001,
		Msg:  err.Error(),
		Data: nil,
	}
}

func BadWithCode(code int) *BaseResponse {
	return &BaseResponse{
		Code: code,
		Msg:  "internal error",
		Data: nil,
	}
}

type DataRef struct {
	Items any `json:"items"`
	Total int `json:"total"`
}

type Reason interface {
	Code() int
	Message() string
}
