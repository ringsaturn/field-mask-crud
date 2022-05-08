package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mennanov/fmutils"
	"github.com/ringsaturn/field-mask-crud/api"
)

type GetProfileRequest struct {
	Paths []string `form:"paths"`
}

func main() {

	e := gin.Default()

	e.GET("/foo", func(ctx *gin.Context) {
		req := &GetProfileRequest{}
		if err := ctx.ShouldBindQuery(req); err != nil {
			ctx.String(500, err.Error())
			return
		}
		demoUser := &api.Profile{
			DisplayName: "xxxx",
			Address:     "???",
			URI:         "http://example.com",
		}
		fmutils.Filter(demoUser, req.Paths)
		ctx.JSON(200, demoUser)
	})
	e.Run("localhost:5010")
	// curl "http://localhost:5010/foo?paths=display_name&paths=URI"
	// {"display_name":"xxxx","URI":"http://example.com"}
}
