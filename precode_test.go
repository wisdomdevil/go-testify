package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenAllRight(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count="+strconv.Itoa(totalCount), nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenCityWrong(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?city=spb&count="+strconv.Itoa(totalCount), nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, "wrong city value", responseRecorder.Body)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count="+strconv.Itoa(totalCount+1), nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
	assert.Equal(t, strings.Join(cafeList["moscow"][:totalCount], ","), responseRecorder.Body)
}
