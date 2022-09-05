// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AddUserRequest struct {
	Number   string `json:"number"`
	Username string `json:"username"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

type GetUserByNameRequest struct {
	Username string `form:"username"`
}

type LoginResponse struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Gender       string `json:"gender"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}
type UserDetailResponse struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}
type UsersResponse struct {
	UserList []*UserDetailResponse `json:"userDetailResponse"`
}
