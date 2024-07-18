package v1

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"gin-web-init/dto/request"
	"gin-web-init/dto/response"
	"gin-web-init/service"
	"gin-web-init/utils"
)

var log = utils.NewLogger(utils.NewConcreteLogConfigBuilder())

var testService = service.NewTestService()

func Test(ctx *gin.Context) {
	var _ = &request.TestRequest{}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read request body"})
		return
	}

	log.Info(string(body))

	//if err := ctx.MustBindWith(req, binding.JSON); err != nil {
	//	ctx.JSON(http.StatusOK, response.BadWithReason(error.NewMustBindError(err)))
	//	return
	//}
	//err = testService.TestCheck(req)
	//
	//if err != nil {
	//	return
	//}

	ctx.JSON(http.StatusOK, response.OK())
}
