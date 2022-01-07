package inbound

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/leandroxt/go-clean-arch/usecase"
)

type GiphyHandler struct {
	usecase usecase.UseCases
}

func (in *GiphyHandler) GetGiphyByTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tag := ps.ByName("tag")

	giphy, err := in.usecase.GetGiphyByTag.Execute(tag)
	if err != nil {
		serverErrorResponse(w, r, err)
		return
	}

	writeJSON(w, 200, envelope{"data": giphy}, nil)
}
