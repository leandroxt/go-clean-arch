package usecase

import (
	"errors"
	"strings"

	"github.com/leandroxt/go-clean-arch/entity"
	"github.com/leandroxt/go-clean-arch/interface-adapters/outbound"
)

const ProhibitedWords = "violence kill"

type GetGiphyByTagUseCase struct {
	gateway outbound.Gateway
}

func (u GetGiphyByTagUseCase) Execute(tag string) (*entity.Giphy, error) {
	lowerTag := strings.ToLower(tag)
	if strings.Contains(ProhibitedWords, lowerTag) {
		return nil, errors.New("tag_not_allowed")
	}

	giphy, err := u.gateway.GetGiphy(tag)
	if err != nil {
		return nil, errors.New("cant_retrieve_giphy")
	}

	return giphy, nil
}
