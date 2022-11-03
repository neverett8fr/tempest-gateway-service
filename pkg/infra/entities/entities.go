package entities

type GetRequest struct {
	Host        string
	Port        int
	Route       string
	ContentType string
}
