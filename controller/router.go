package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	apiv1 "gin-web-init/controller/v1"
	"gin-web-init/dto/response"
	my "gin-web-init/validator"
)

var R *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for name, fn := range my.RegistryList() {
			_ = v.RegisterValidation(name, fn)
		}
	}

	// permit `healthz` to be allowed to access
	PermitRequest("/healthz")
	// instantiate a gin.Engine instance
	R = gin.New()
	// Interceptor middleware will be included in the handlers chain for every single request
	R.Use(Interceptor)
	R.Use(gin.Recovery())
	R.Use(gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{"/healthz"}}))
	// for liveness/readiness/startup check
	R.Any("/healthz", Health)

	v1 := R.Group("/api/v1")

	v1.POST("/test", apiv1.Test)
}

func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.OK())
}
