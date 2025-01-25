package services

type homeService struct{}

func NewHomeService() HomeServiceInterface {
	return &homeService{}
}

func (service *homeService) GetWelcomeMessage() string {
	return "Hello, World!"
}
