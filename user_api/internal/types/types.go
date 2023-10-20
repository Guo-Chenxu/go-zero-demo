// Code generated by goctl. DO NOT EDIT.
package types

type GetUserReq struct {
	Id int64 `path:"id" validate:"required,gte=0,lte=10"`
}

type GetUserRep struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`     // The username
	Password string `json:"password"` // The user password
	Mobile   string `json:"mobile"`   // The mobile phone number
	Type     int64  `json:"type"`     // The user type, 0:normal,1:vip, for test golang keyword
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

type SetUserReq struct {
	Name     string `json:"name"`     // The username
	Password string `json:"password"` // The user password
	Mobile   string `json:"mobile"`   // The mobile phone number
	Type     int64  `json:"type"`
}
