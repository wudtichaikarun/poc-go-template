package common

type Ctx interface {
	Bind(interface{}) error
	JSON(int, interface{})
	TransactionID() string
	Audience() string
	GetToken() string
	GetParam(string) (int, error)
}
