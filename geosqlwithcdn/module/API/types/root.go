package types

type RegisterUserReq struct {
	UserName    string   `json:"username" binding:"required"`
	Description string   `json:"description"`
	Hobby       []string `json:"hobby"`
	Latitude    float64  `json:"latitude" binding:"required,min=-90,max=90"`
	Longitude   float64  `json:"longitude" binding:"required,min=-180,max=180"`
	//min max can also be used in string type's length
}

type AroundUsersReq struct {
	UserName string `form:"username" binding:"required"`
	Range    int64  `form:"range" binding:"required"`
	Limit    int64  `form:"limit" binding:"required"`
}
