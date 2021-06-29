package consts

// 这里定义的常量，一般是具有错误代码+错误说明组成，一般用于接口返回
const (

	// 表单验证器前缀
	ValidatorPrefix              string = "Form_Validator_"
	ValidatorParamsCheckFailCode int    = 400
	ValidatorParamsCheckFailMsg  string = "参数校验失败"

	//服务器代码发生错误
	ServerOccurredErrorCode int    = 500
	ServerOccurredErrorMsg  string = "服务器内部发生代码执行错误, "

	// CURD 常用业务状态码
	CurdStatusOkCode     int    = 200
	CurdStatusOkMsg      string = "Success"
	CurdCreatFailCode    int    = 400200
	CurdCreatFailMsg     string = "新增失败"
	CurdUpdateFailCode   int    = 400201
	CurdUpdateFailMsg    string = "更新失败"
	CurdDeleteFailCode   int    = 400202
	CurdDeleteFailMsg    string = "删除失败"
	CurdSelectFailCode   int    = 400203
	CurdSelectFailMsg    string = "查询无数据"
	CurdRegisterFailCode int    = 400204
	CurdRegisterFailMsg  string = "注册失败"
	CurdLoginFailCode    int    = 400205
	CurdLoginFailMsg     string = "登录失败"
)
