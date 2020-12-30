package service

import (
	"context"
	"net/http"
	"rhyme-engine/internal/httpext"
	"rhyme-engine/internal/rhyme/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRhymeSvc(t *testing.T)  {
	refWords := []string{"Diluting", "Computing", "Polluting", "Commuting", "Recruiting", "Drooping"}
	rhymeImpl := NewRhymeSvcImpl(refWords)
	wordReq := model.RhymingWordRequest{
		Words: []string{"Shooting","Disputing"},
	}
	response, e := rhymeImpl.FindMatchingWord(context.Background(), wordReq)
	require.Nil(t, e)

	expectedResponse := map[string][]string {
			"Disputing": []string{"Computing"},
			"Shooting": []string{
				"Diluting",
				"Commuting",
				"Computing",
				"Polluting",
				"Recruiting",
			},
	}
	successResponse := httpext.SuccessResponse{HttpStatusCode: http.StatusOK, Body: expectedResponse}
	assert.Equal(t, response.HttpStatusCode, successResponse.HttpStatusCode)
	assert.Equal(t, response.Body, expectedResponse)
}

func TestRhymeSvc_1(t *testing.T)  {
	refWords := []string{"Diluting", "Computing", "Polluting", "Commuting", "Recruiting", "Drooping"}
	rhymeImpl := NewRhymeSvcImpl(refWords)
	wordReq := model.RhymingWordRequest{
		Words: []string{"Convoluting"},
	}
	response, e := rhymeImpl.FindMatchingWord(context.Background(), wordReq)
	require.Nil(t, e)

	expectedResponse := map[string][]string {
		"Convoluting": []string{"Diluting"},
	}
	successResponse := httpext.SuccessResponse{HttpStatusCode: http.StatusOK, Body: expectedResponse}
	assert.Equal(t, response.HttpStatusCode, successResponse.HttpStatusCode)
	assert.Equal(t, response.Body, expectedResponse)
}

func TestRhymeSvc_2(t *testing.T)  {
	refWords := []string{"Diluting", "Computing", "Polluting", "Commuting", "Recruiting", "Drooping"}
	rhymeImpl := NewRhymeSvcImpl(refWords)
	wordReq := model.RhymingWordRequest{
		Words: []string{"Orange"},
	}
	response, e := rhymeImpl.FindMatchingWord(context.Background(), wordReq)
	require.Nil(t, e)

	expectedResponse := map[string][]string {
		"Orange": []string{},
	}
	successResponse := httpext.SuccessResponse{HttpStatusCode: http.StatusOK, Body: expectedResponse}
	assert.Equal(t, response.HttpStatusCode, successResponse.HttpStatusCode)
}