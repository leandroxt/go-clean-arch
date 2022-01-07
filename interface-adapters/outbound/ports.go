package outbound

import "github.com/leandroxt/go-clean-arch/entity"

type Gateway interface {
	GetGiphy(tag string) (*entity.Giphy, error)
}

type Gatewayimpl struct {
	giphyKey string
}

func NewGatewayImpl(giphyKey string) Gateway {
	return Gatewayimpl{giphyKey: giphyKey}
}
