package types

type Readable interface {
	IsDone() bool
	GetResponse() string
}

// EmptyResponse is a placeholder for responses that do not return any data.
type emptyResponse struct{}

func (r emptyResponse) IsDone() bool {
	return true
}
func (r emptyResponse) GetResponse() string {
	return ""
}

var EmptyResponse Readable = emptyResponse{}
