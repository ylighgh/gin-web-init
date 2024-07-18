package request

type TestRequest struct {
	Name string `form:"name" json:"name" xml:"name" binding:"required"`
}
