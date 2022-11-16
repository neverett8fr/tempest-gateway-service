package entities

type Request struct {
	Host        string
	Port        int
	Route       string
	ContentType string
	Method      string
	Auth        string
	Body        interface{}
}
