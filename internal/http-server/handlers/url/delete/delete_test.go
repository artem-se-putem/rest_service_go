package delete_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"rest_service_go/internal/http-server/handlers/url/delete"
	"rest_service_go/internal/http-server/handlers/url/delete/mocks"
	"rest_service_go/internal/lib/logger/handlers/slogdiscard"
)

func TestDeleteHandler(t *testing.T) {
	cases := []struct {
		name      string
		url       string
		alias     string
		respError string
		mockError error
	}{
		{
			name:  "Success",
			url:   "https://youtube123.com",
			alias: "test_alias",
		},
		{
			name:      "Empty alias",
			url:       "https://qwe.com",
			alias:     "",
			respError: "field Alias is a required field",
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			// t.Parallel()

			urlDeleterMock := mocks.NewURLDeleter(t) // пакет mocks, метод оттуда

			if tc.respError == "" || tc.mockError != nil {
				urlDeleterMock.On("DeleteURL", tc.alias).
					Return(tc.mockError).
					Once()
			}

			handler := delete.New(slogdiscard.NewDiscardLogger(), urlDeleterMock)

			input := fmt.Sprintf(`{"url": "%s","alias": "%s"}`, tc.url, tc.alias)

			req, err := http.NewRequest(http.MethodPost, "/delete", bytes.NewReader([]byte(input)))
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			require.Equal(t, rr.Code, http.StatusOK)

			body := rr.Body.String()

			var resp delete.DeleteResponse

			require.NoError(t, json.Unmarshal([]byte(body), &resp)) // Валидация парсинга

			require.Equal(t, tc.respError, resp.Error) // Проверка ошидаемой ошибки

			// TODO: add more checks
		})
	}
}
