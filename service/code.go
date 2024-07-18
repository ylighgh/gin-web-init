package service

import (
	"gin-web-init/dto/request"
)

type TestService interface {
	TestCheck(req *request.TestRequest) error
}

type testServiceImpl struct{}

func NewTestService() TestService {
	return new(testServiceImpl)
}

func (*testServiceImpl) TestCheck(req *request.TestRequest) error {
	return nil
}
