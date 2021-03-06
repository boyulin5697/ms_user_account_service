// Code generated by goctl. DO NOT EDIT.
package types

type RegisterReq struct {
	Name        string `json:"name"`
	Password    string `json:"password"`
	Telephone   string `json:"telephone"`
	Email       string `json:"email"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
	Nationality string `json:"nationaility"`
}

type LoginReq struct {
	LoginPayLoad string `json:"loginPayLoad"`
	Password     string `json:"password"`
}

type GeneralResp struct {
	Code        int8   `json:"code"`
	Message     string `json:"message"`
	Information string `json:"information"`
}
