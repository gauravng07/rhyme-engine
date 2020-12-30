package httpext

import (
	"context"
	"encoding/json"
	"net/http"
	"rhyme-engine/internal/logger"
)

type Response struct {
	http.ResponseWriter
}

type SuccessResponse struct {
	HttpStatusCode int
	Body           interface{}
}

var NoResponse SuccessResponse

type Handler func(res Response, req *http.Request, testHeader string)

func (h Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	h(Response{rw}, req, "")
}

func (r Response) Send(ctx context.Context, res SuccessResponse, err error) {
	if err != nil {
		customErr, ok := err.(ErrorResponse)
		if !ok {
			customErr = InternalServerError
		}

		logger.Errorf(ctx, "%+v", err)
		r.WriteHeader(customErr.httpStatusCode)
		_ = json.NewEncoder(r).Encode(customErr)
		return
	}

	r.WriteHeader(res.HttpStatusCode)
	if res.Body != nil {
		_ = json.NewEncoder(r).Encode(res.Body)
	}
}

func (r Response) SendHttpErrorResponse(ctx context.Context, err error) {
	if err != nil {
		customErr, ok := err.(ErrorResponse)
		if !ok {
			customErr = InternalServerError
		}
		logger.Errorf(ctx, "%+v", err)
		r.WriteHeader(customErr.httpStatusCode)
		_, _ = r.Write([]byte(err.Error()))
		return
	}
	r.WriteHeader(http.StatusInternalServerError)
}
