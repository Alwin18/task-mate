package handler

type (
	Handler interface {
		LoginHandler
	}

	LoginHandler interface{}
)

func NewHandler() Handler {
	return &handler{}
}

type handler struct{}
