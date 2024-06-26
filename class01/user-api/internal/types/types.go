// Code generated by goctl. DO NOT EDIT.
package types

type UserInfoReq struct {
	UserId int64 `form:"userId"` //query传参就把json修改为form
}

type UserInfoResp struct {
	UserId   int64  `json:"userIdgo"`
	NickName string `json:"nickname"`
}

type UserUpdateReq struct {
	UserId   int64  `json:"userId"`
	NickName string `json:"nickName"`
}

type UserUpdateResp struct {
	Flag bool `json:"flag"`
}
