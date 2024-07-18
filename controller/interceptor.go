package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin-web-init/dto/response"
	"gin-web-init/utils"
	"sync"
)

var log = utils.NewLogger(utils.NewConcreteLogConfigBuilder())

type RequestInterceptor func(request *http.Request) error

type requestInterceptorChain struct {
	mutex        *sync.RWMutex
	interceptors []RequestInterceptor
	whitelist    map[string]struct{}
}

func (c *requestInterceptorChain) Do(request *http.Request) error {
	if _, ok := c.whitelist[request.URL.EscapedPath()]; ok {
		return nil
	}
	return nil
}

var defaultRequestInterceptorChain = &requestInterceptorChain{
	mutex:        new(sync.RWMutex),
	interceptors: make([]RequestInterceptor, 0),
	whitelist:    make(map[string]struct{}),
}

func PermitRequest(paths ...string) {
	mutex := defaultRequestInterceptorChain.mutex
	mutex.Lock()
	defer mutex.Unlock()
	for _, path := range paths {
		defaultRequestInterceptorChain.whitelist[path] = struct{}{}
	}
}

func AddRequestInterceptors(interceptors ...RequestInterceptor) {
	mutex := defaultRequestInterceptorChain.mutex
	mutex.Lock()
	defer mutex.Unlock()

	for _, interceptor := range interceptors {
		defaultRequestInterceptorChain.interceptors = append(defaultRequestInterceptorChain.interceptors, interceptor)
	}
}

var _ = AddRequestInterceptors

func Interceptor(ctx *gin.Context) {
	if err := defaultRequestInterceptorChain.Do(ctx.Request); err != nil {
		log.Errorf("do interceptor error = %s", err.Error())
		ctx.JSON(http.StatusBadRequest, response.BadWithCode(400))
	} else {
		// hand over request to another HttpHandler once the interceptor has done
		ctx.Next()
	}
}
