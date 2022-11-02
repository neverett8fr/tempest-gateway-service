package entities

type GetRequest struct {
	Protocol    string
	Host        string
	Port        int
	Route       string
	ContentType string
}
