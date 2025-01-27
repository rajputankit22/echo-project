package services

// struct homeService
type homeService struct{}

// Constructor function for homeService
func NewHomeService() HomeServiceInterface {
	return &homeService{}
}

// GetWelcomeMessage - returns a welcome message
func (service *homeService) GetWelcomeMessage() string {
	return "Hello, World!"
}
