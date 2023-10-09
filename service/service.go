package service

type Service struct {
	CommonService
}

func New() *Service {
	return &Service{
		CommonService: NewCommonService(),
	}
}
