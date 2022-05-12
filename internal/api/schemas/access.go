package schemas

type AccessMember struct {
	UserID int64  `json:"userId" example:"1"`
	Role   string `json:"role" enums:"admin,member"`
}
