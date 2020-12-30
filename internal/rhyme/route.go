package rhyme

import (
	"net/http"
	"rhyme-engine/internal/rhyme/handler"
	"rhyme-engine/internal/rhyme/service"

	"github.com/gorilla/mux"
)

func Configure(mainRouter *mux.Router, refWordClient service.ReferenceRhymeWordClient)  {
	r := mainRouter.PathPrefix("/rhymes").Subrouter()
	r.Handle("/match-word", handler.RhymeHandler(service.NewRhymeSvcImpl(refWordClient.GetRhymeWords()))).Methods(http.MethodPost)
	r.Handle("/reference-list", handler.GetReferenceWordList(refWordClient))
}