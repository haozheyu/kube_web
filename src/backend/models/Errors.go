package models

import "errors"

var (
	ErrMemberNoExist             = errors.New("用户不存在")
	ErrMemberExist               = errors.New("用户已存在")
	ErrMemberDisabled            = errors.New("用户被禁用")
	ErrMemberEmailEmpty          = errors.New("用户邮箱不能为空")
	ErrMemberEmailExist          = errors.New("用户邮箱已被使用")
	ErrMemberDescriptionTooLong  = errors.New("用户描述必须小于500字")
	ErrMemberEmailFormatError    = errors.New("邮箱格式不正确")
	ErrMemberPasswordFormatError = errors.New("密码必须在6-50个字符之间")
	ErrMemberAccountFormatError  = errors.New("账号只能由英文字母数字组成，且在3-50个字符")
	ErrMemberRoleError           = errors.New("用户权限不正确")
	// ErrorMemberPasswordError 密码错误.
	ErrorMemberPasswordError = errors.New("用户密码错误")
	// ErrDataNotExist 指定的服务已存在.
	ErrDataNotExist = errors.New("数据不存在")
	// ErrInvalidParameter 参数错误.
	ErrInvalidParameter = errors.New("Invalid parameter")
	ErrPermissionDenied = errors.New("Permission denied")
	ErrGetPOD = errors.New("get pod fail")
	ErrCreatePOD = errors.New("create pod fail")
	ErrDelPOD = errors.New("del pod fail")
	ErrListDeployment = errors.New("ListDeployment fail")
	ErrGetDeployment = errors.New("GetDeployment fail")
	ErrDeleteDeployment = errors.New("DeleteDeployment fail")
	ErrDeploymentParams = errors.New("DeploymentParams fail")
	ErrApplyDeployment = errors.New("ApplyDeployment fail")
	ErrGetNamespace = errors.New("GetNamespace fail")
	ErrCreateNamespace = errors.New("CreateNamespace fail")
	ErrDeleteNamespace = errors.New("DeleteNamespace fail")
	ErrListService = errors.New("ListService fail")
	ErrGetService = errors.New("GetService fail")
	ErrParseFormSVC = errors.New("ParseFormSVC fail")
	ErrCreateService = errors.New("CreateService fail")
	ErrDeleteService = errors.New("DeleteService fail")
	ErrParseFormPVC = errors.New("ParseFormPVC fail")
	ErrCreatePVC = errors.New("CreatePVC fail")
)

type Error struct {
	code    int
	message string
}

func (e Error) Error() string {
	return e.message
}

func (e Error) Code() int {
	return e.code
}

func NewError(code int, message string) Error {
	return Error{code: code, message: message}
}
