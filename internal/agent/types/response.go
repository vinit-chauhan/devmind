package types

type Response interface {
	IsDone() bool
	GetResponse() string
}

// EmptyResponse is a placeholder for responses that do not return any data.

type EmptyResponse struct{}

func (r EmptyResponse) IsDone() bool {
	return true
}
func (r EmptyResponse) GetResponse() string {
	return ""
}
