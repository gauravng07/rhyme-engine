package handler

import (
	"encoding/json"
	"net/http"
	"rhyme-engine/internal/httpext"
	"rhyme-engine/internal/logger"
	"rhyme-engine/internal/rhyme/model"
	"rhyme-engine/internal/rhyme/service"
)

func RhymeHandler(service service.Rhyme) httpext.Handler {
	return func(res httpext.Response, req *http.Request, testHeader string) {
		ctx := req.Context()
		var rhymingWordRequest model.RhymingWordRequest
		err := json.NewDecoder(req.Body).Decode(&rhymingWordRequest)
		if err != nil {
			logger.Errorf(ctx, "error unmarshalling request body: %+v", err)
			res.Send(ctx, httpext.NoResponse, httpext.NewBadRequestError(httpext.InvalidRequestBodyErrorCode, httpext.BadRequestErrorMessage, err, nil))
		}

		response, e := service.FindMatchingWord(ctx, rhymingWordRequest)
		if e != nil {
			res.Send(ctx, httpext.NoResponse, e)
			return
		}
		res.Send(ctx, response, nil)

	}
}

func GetReferenceWordList(service service.ReferenceRhymeWordClient) httpext.Handler {
	return func(res httpext.Response, req *http.Request, testHeader string) {
		res.Send(req.Context(), httpext.SuccessResponse{HttpStatusCode: http.StatusOK, Body: service.GetRhymeWords()}, nil)
	}
}
