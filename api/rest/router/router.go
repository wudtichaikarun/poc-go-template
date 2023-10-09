package router

type Context interface {
	Bind(interface{}) error
	JSON(int, interface{}) error
	TransactionID() string
	Audience() string
	GetToken() string
	GetParam(string) (int, error)
}
