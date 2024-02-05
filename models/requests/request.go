package request

type GetSumInput struct {
	UserName string `form:"user_name"`
	A        int64  `form:"a"`
	B        int64  `form:"b"`
}
