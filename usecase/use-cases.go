package usecase

import "github.com/leandroxt/go-clean-arch/interface-adapters/outbound"

type UseCases struct {
	GetGiphyByTag GetGiphyByTagUseCase
}

func NewUseCases(giphyKey string) UseCases {
	gateway := outbound.NewGatewayImpl(giphyKey)

	return UseCases{
		GetGiphyByTag: GetGiphyByTagUseCase{gateway: gateway},
	}
}
