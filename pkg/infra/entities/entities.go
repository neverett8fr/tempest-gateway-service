package entities

type Request struct {
	Host        string
	Port        int
	Route       string
	Accept      string
	ContentType string
	Transfer    string
	Method      string
	Auth        string
	Body        interface{}
}
