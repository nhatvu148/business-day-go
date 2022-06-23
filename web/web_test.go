package web_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	web "github.com/nhatvu148/business-day-go/web"
)

func TestGetCatFact(t *testing.T) {
	cases := []struct {
		description      string
		server           *httptest.Server
		expectedResponse string
		expectedError    error
	}{
		{
			description: "Test Cat fact response",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)

				w.Write([]byte(`{"Fact": "Cats can predict earthquakes. We humans are not 100% sure how they do it. There are several different theories.", "Length": "111"}`))
			})),
			expectedResponse: "Cats can predict earthquakes. We humans are not 100% sure how they do it. There are several different theories.",
			expectedError:    nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			defer tt.server.Close()
			resp, err := web.GetCatFact(tt.server.URL)

			if !reflect.DeepEqual(resp, tt.expectedResponse) {
				t.Errorf("expected (%v), got (%v)", tt.expectedResponse, resp)
			}
			if !errors.Is(err, tt.expectedError) {
				t.Errorf("expected (%v), got (%v)", tt.expectedError, err)
			}
		})
	}
}
