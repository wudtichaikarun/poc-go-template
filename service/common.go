package service

type CommonService interface {
	HealthCheck() bool
}

type CommonServiceContext struct{}

func (c *CommonServiceContext) HealthCheck() bool {
	return true
}

func NewCommonService() CommonService {
	return &CommonServiceContext{}
}
