// Code generated by goctl. DO NOT EDIT.
package types

import (
	"github.com/lib/pq"
)


type UserInfo struct {
	Username     string   `json:"username"`
	Email        string   `json:"email"`
	Phone        string   `json:"phone"`
	Avatar       string   `json:"avatar"`
	Role         pq.StringArray `json:"role"`
	IsSuperAdmin bool     `json:"is_super_admin"`
	Age          int      `json:"age"`
	Gender       int8     `json:"gender"`
}


type UserInfoWithToken struct {
	UserInfo
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RegisterReq struct {
	Email    string `form:"email"`
	Password string `form:"password"`
	Username string `form:"username"`
}

type RegisterResp struct {
	BaseResponse
	Data UserInfo `json:"data"`
}

type LoginReq struct {
	Count    string `form:"count" binding:"required"` // tag只能为json or form
	Password string `form:"password" binding:"required"`
}

type LoginResp struct {
	BaseResponse
	Data UserInfoWithToken `json:"data"`
}

type LogoutResp struct {
	BaseResponse
}

type UpdateUserInfoResp struct {
	BaseResponse
	Data UserInfo `json:"data"`
}


type UpdateUserInfoRep struct {
	UserInfo
}

type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	// Data interface{} `json:"data"`
}
