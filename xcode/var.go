package xcode

import "net/http"

const (
	MinReservedCode = -1000 // 系统保留错误码最低值
	MaxReservedCode = 10000 // 系统保留错误码最大值
	ForbiddenCode   = 0     // 禁止的错误码
)

var (
	OK                   = &Code{http.StatusOK, 0, "请求成功"}
	ErrCommon            = &Code{http.StatusBadRequest, 4000, "请求错误"}
	ParamInvalid         = &Code{http.StatusBadRequest, 4001, "参数错误"}
	UserHasNoPhone       = &Code{http.StatusBadRequest, 4002, "当前用户未绑定手机号"}
	UserNotLogin         = &Code{http.StatusUnauthorized, 4011, "用户未登录"}
	UserSessionExpired   = &Code{http.StatusUnauthorized, 4012, "用户SESSION已过期"}
	UserNotAuthorized    = &Code{http.StatusUnauthorized, 4013, "鉴权失败，当前用户没有操作权限"}
	NeedWechatAuthorized = &Code{http.StatusUnauthorized, 4014, "需要额外的微信授权"}
	ErrServCommon        = &Code{http.StatusInternalServerError, 5000, "服务错误"}
	OpeDbErr             = &Code{http.StatusInternalServerError, 5001, "操作db错误"}
	SrvParamInvalid      = &Code{http.StatusInternalServerError, 5002, "服务调用参数出错"}
	SrvOperateDbErr      = &Code{http.StatusInternalServerError, 5003, "服务操作数据库出错"}
	ErrCodeFailed        = &Code{http.StatusInternalServerError, 5100, "错误码定义错误"}
)

var allCode = []*Code{
	OK,
	ErrCommon,
	ParamInvalid,
	UserHasNoPhone,
	UserNotLogin,
	UserSessionExpired,
	UserNotAuthorized,
	NeedWechatAuthorized,
	ErrServCommon,
	OpeDbErr,
	SrvParamInvalid,
	SrvOperateDbErr,
	ErrCodeFailed,
}
