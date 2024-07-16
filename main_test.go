package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=7&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	// Проверка кол-ва count
	cafes := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, cafes, totalCount, "Не соответсвует кол-ву")
	// Возрашаем все доступные кафе из полученных
	expectedCafes := strings.Join(cafeList["moscow"], ",")
	assert.Equal(t, expectedCafes, responseRecorder.Body.String(), "Кафе не соответсвуют ожидаемым кафе")
}

func TestMainHandlerWhenStatusOkAndBodyNotEmpty(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=7&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	// Проверяем код 200
	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Ожидался код 200")
	// Проверка на пустой ответ в теле
	assert.NotEmpty(t, responseRecorder.Body.String(), "Тело ответа пришло пустым")
}

func TestMainHandlerWhenStatusBadRequestAndWrongErrorMessage(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=7&city", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	// Проверяем код 400
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Ожидался код 400")
	// Проверка сообщения об ошибке
	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "Получено другое сообщение об ошибке")
}
