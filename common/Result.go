package common

// Result 返回的对象
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

const (
	ERROR   = 50000
	SUCCESS = 20000
)

// Success 封装成功返回的对象
func Success(data interface{}) Result {
	return Result{
		Code: SUCCESS,
		Msg:  "success",
		Data: data,
	}
}

// Fail 封装失败返回的对象
func Fail(msg string) Result {
	return Result{
		Code: ERROR,
		Msg:  msg,
		Data: nil,
	}
}
