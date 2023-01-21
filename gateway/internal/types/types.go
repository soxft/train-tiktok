// Code generated by goctl. DO NOT EDIT.
package types

type Resp struct {
	Code int32  `json:"status_code"`
	Msg  string `json:"status_msg"`
}

type PingResp struct {
	Resp
	Msg string `json:"msg"`
}

type LoginReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type RegisterReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type PublishReq struct {
	Token string `form:"token"`
	Title string `form:"Title"`
}
