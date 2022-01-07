package outbound

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/leandroxt/go-clean-arch/entity"
)

func (g Gatewayimpl) GetGiphy(tag string) (*entity.Giphy, error) {
	url := "https://api.giphy.com/v1/gifs/random?api_key=" + g.giphyKey + "&tag=" + tag

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= 300 {
		return nil, errors.New("giphy_response_bad_statuscode")
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(data)
	decoder := json.NewDecoder(r)

	giphy := &GiphyResponse{}
	if err = decoder.Decode(giphy); err != nil {
		return nil, err
	}

	return &entity.Giphy{
		ID:  giphy.Data.ID,
		URL: giphy.Data.Images.Original.URL,
	}, nil
}
