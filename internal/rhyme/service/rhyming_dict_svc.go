package service

import (
	"context"
	"encoding/json"
	"fmt"
	"rhyme-engine/internal/config"
	"rhyme-engine/internal/logger"
	"rhyme-engine/internal/wrapper"
	"sync"

	"github.com/spf13/viper"
)

type ReferenceRhymeWordClient interface {
	GetRhymeWords() []string
	FetchReferenceRhymeWords(ctx context.Context, reader wrapper.Reader) []string
}

type referenceRhymeWordSvcImpl struct {
	word []string
	mux sync.Mutex
}

func NewReferenceRhymeSvcImpl(ctx context.Context, reader wrapper.Reader) *referenceRhymeWordSvcImpl {
	impl := referenceRhymeWordSvcImpl{}
	word := impl.FetchReferenceRhymeWords(ctx, reader)
	return &referenceRhymeWordSvcImpl{
		word: word,
	}
}

func (r *referenceRhymeWordSvcImpl) GetRhymeWords() []string {
	r.mux.Lock()
	defer r.mux.Unlock()
	return r.word
}

func (r *referenceRhymeWordSvcImpl) FetchReferenceRhymeWords(ctx context.Context, reader wrapper.Reader) []string  {
	var result map[string][]string
	refRhymeWords, err := reader.Read(fmt.Sprintf("%s", viper.GetString(config.RefFile)))
	if err != nil {
		logger.Errorf(ctx, "error getting reference file: %v", err)
	}

	e := json.Unmarshal(refRhymeWords, &result)
	if e != nil {
		logger.Errorf(ctx, "error getting rhyme words: %v", e)
	}

	if v, ok := result["sample_rhyme_words"]; ok {
		return v
	} else {
		logger.Errorf(ctx, "error getting reference rhyme words: %+v", v)
		return nil
	}
}


