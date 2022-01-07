package inbound

import "github.com/leandroxt/go-clean-arch/usecase"

type Handler struct {
	Giphy GiphyHandler
}

func NewHandler(giphyKey string) Handler {
	usecases := usecase.NewUseCases(giphyKey)

	return Handler{
		Giphy: GiphyHandler{usecase: usecases},
	}
}
