package base

type ResponseCode int

const (
	ResponseCodeOk                 = 1000
	ResponseCodeValidtionerror     = 2000
	ResponseCodeRequestParamsError = 2100
	ResponseCodeInterServerError   = 5000
	ResponseCodeBizError           = 6000
)

type Response struct {
	Code    ResponseCode    `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
