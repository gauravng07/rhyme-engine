package service

import (
	"context"
	"net/http"
	"rhyme-engine/internal/httpext"
	"rhyme-engine/internal/logger"
	"rhyme-engine/internal/rhyme/model"
)

type Rhyme interface {
	FindMatchingWord(ctx context.Context, words model.RhymingWordRequest) (httpext.SuccessResponse, error)
}

type rhymeImpl struct {
	refWords []string
}

func NewRhymeSvcImpl(refWords []string) *rhymeImpl {
	return &rhymeImpl{refWords: refWords}
}

func (r *rhymeImpl) FindMatchingWord(ctx context.Context, wordReq model.RhymingWordRequest) (httpext.SuccessResponse, error) {
	words := wordReq.Words
	logger.Infof(ctx, "words: %v", words)

	phonemeModifiers := FindPhoneticWithModifiers()
	refWordPhonetics := SplitWordIntoPhonetics(r.refWords, phonemeModifiers)
	inputWordPhonetics := SplitWordIntoPhonetics(words, phonemeModifiers)

	result := make(map[string][]string)
	for _, word := range words {
		var bswords []string
		bestMatch := FindBestMatches(refWordPhonetics, inputWordPhonetics, word)
		for i := 0; i < len(bestMatch); i++ {
			k := bestMatch[i]
			for k, _ := range k {
				bswords = append(bswords, k)
			}
		}
		result[word] = bswords
	}
	return httpext.SuccessResponse{HttpStatusCode: http.StatusOK, Body: result}, nil
}
